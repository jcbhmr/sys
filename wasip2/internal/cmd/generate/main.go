package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var merr error
	defer func() {
		if merr != nil {
			log.Fatal(merr)
		}
	}()

	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		merr = err
		return
	}
	defer os.RemoveAll(tempDir)
	log.Printf("tempDir=%q", tempDir)

	cmd := exec.Command("wkg", "fetch")
	log.Printf("Running %q...", cmd)
	err = cmd.Run()
	if err != nil {
		merr = err
		return
	}

	cmd = exec.Command("go", "tool", "wit-bindgen-go", "generate", "wit", "--out", tempDir, "--package-root", "example.internal")
	log.Printf("Running %q...", cmd)
	err = cmd.Run()
	if err != nil {
		merr = err
		return
	}

	entries, err := os.ReadDir(filepath.Join(tempDir, "wasi"))
	if err != nil {
		merr = err
		return
	}
	for _, e := range entries {
		if !e.IsDir() {
			merr = fmt.Errorf("%v not dir", e)
			return
		}
	}

	filepath.WalkDir(filepath.Join(tempDir, "wasi"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Base(path) == "empty.s" {
			return os.Remove(path)
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		rel, err := filepath.Rel(tempDir, path)
		if err != nil {
			panic(err)
		}

		log.Printf("Post-processing %q...", rel)
		f, err := os.OpenFile(path, os.O_RDWR, 0o666)
		if err != nil {
			return err
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}

		replaced := false
		data = regexp.MustCompile(`^// Code generated .* DO NOT EDIT\.$`).ReplaceAllFunc(data, func(b []byte) []byte {
			if replaced {
				return b
			}
			replaced = true
			return append(b, "\n\n//go:build wasip2"...)
		})
		if !replaced {
			data = append([]byte("//go:build wasip2\n\n"), data...)
		}

		data = regexp.MustCompile(`"example.internal/wasi/(.+?)"`).ReplaceAll(data, []byte(`"go.jcbhmr.com/sys/wasip2/$1"`))

		_, err = f.Write(data)
		return err
	})

	for _, e := range entries {
		log.Printf("Replacing %q with %q...", e.Name(), filepath.Join(tempDir, "wasi", e.Name()))
		err := os.RemoveAll(e.Name())
		if err != nil {
			merr = err
			return
		}
		err = os.CopyFS(e.Name(), os.DirFS(filepath.Join(tempDir, "wasi", e.Name())))
		if err != nil {
			merr = err
			return
		}
	}
}
