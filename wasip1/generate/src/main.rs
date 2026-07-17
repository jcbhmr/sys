use heck::{ToLowerCamelCase, ToShoutySnakeCase, ToSnakeCase, ToUpperCamelCase};
use std::error::Error;
use witx::{BuiltinType, IntRepr, NamedType, RecordDatatype, Type, TypeRef};

fn main() -> Result<(), Box<dyn Error>> {
    let doc = witx::load(&["witx/wasi_snapshot_preview1.witx"])?;

    for m in doc.modules() {
        use std::fmt::Write;
        let mut out = String::new();
        writeln!(out, "{}", "//go:build wasip1")?;
        writeln!(out)?;
        writeln!(out, "package wasip1")?;
        writeln!(out)?;
        for i in m.imports() {
            // TODO?
        }
        for f in m.funcs() {
            writeln!(out, "{}", to_go_doc_comment(&f.docs))?;
            let go_name = f.name.as_str().to_upper_camel_case();
            write!(out, "func {}(", go_name)?;
            let mut first = true;
            for p in &f.params {
                if first {
                    first = false;
                } else {
                    write!(out, ", ")?;
                }
                let go_param_name = p.name.as_str().to_lower_camel_case();
                let go_type_name = type_ref_to_go_type_name(&p.tref);
                write!(out, "{} {}", go_param_name, go_type_name)?;
            }
            write!(out, ")")?;
            match f.results.len() {
                0 => {}
                _ => {
                    write!(out, " (")?;
                    let mut first = true;
                    for r in &f.results {
                        if first {
                            first = false;
                        } else {
                            write!(out, ", ")?;
                        }
                        let go_type_name = type_ref_to_go_type_name(&r.tref);
                        write!(out, "{}", go_type_name)?;
                    }
                    write!(out, ")")?;
                }
            }
            writeln!(out, " {{")?;
            // TODO
            writeln!(out, "}}")?;
            writeln!(out)?;
        }

        let name_flat_case = m.name.as_str().to_snake_case().replace("_", "");
        let path = format!("{}.go", name_flat_case);
        fs_err::write(path, out)?;
    }

    Ok(())
}

fn to_go_doc_comment(input: &str) -> String {
    input
        .lines()
        .map(|l| format!("// {l}"))
        .collect::<Vec<_>>()
        .join("\n")
}

fn type_ref_to_go_type_name(input: &TypeRef) -> String {
    match input {
        TypeRef::Name(n) => n.name.as_str().to_upper_camel_case(),
        TypeRef::Value(v) => match &**v {
            Type::Builtin(t) => builtin_type_to_go_type_name(t),
            Type::Handle(t) => todo!(),
            Type::List(t) => match &**t.type_() {
                Type::Builtin(BuiltinType::Char) => "string".to_owned(),
                _ => {
                    let go_type_name = type_ref_to_go_type_name(t);
                    format!("[]{}", go_type_name)
                }
            },
            Type::Pointer(t) | Type::ConstPointer(t) => {
                let go_type_name = type_ref_to_go_type_name(t);
                format!("*{}", go_type_name)
            }
            Type::Record(t) => {
                println!("Type::Record {:?}", t);
                format!("RECORD")
            }
            Type::Variant(v) if v.is_bool() => "bool".to_owned(),
            Type::Variant(v) => match v.as_expected() {
                Some((ok, err)) => {
                    let err = err.unwrap();
                    match ok {
                        Some(t) => {
                            let a = type_ref_to_go_type_name(t);
                            let b = type_ref_to_go_type_name(err);
                            format!("{}, {}", a, b)
                        },
                        None => {
                            let b = type_ref_to_go_type_name(err);
                            b
                        },
                    }
                }
                None => todo!(),
            },
        },
    }
}

fn builtin_type_to_go_type_name(input: &BuiltinType) -> String {
    match input {
        BuiltinType::Char => "rune".to_owned(),
        BuiltinType::F32 => "float32".to_owned(),
        BuiltinType::F64 => "float64".to_owned(),
        BuiltinType::S16 => "int16".to_owned(),
        BuiltinType::S32 => "int32".to_owned(),
        BuiltinType::S64 => "int64".to_owned(),
        BuiltinType::S8 => "int8".to_owned(),
        BuiltinType::U16 => "uint16".to_owned(),
        BuiltinType::U32 { lang_ptr_size: _ } => "uint32".to_owned(),
        BuiltinType::U64 => "uint64".to_owned(),
        BuiltinType::U8 { lang_c_char: _ } => "uint8".to_owned(),
    }
}

// trait ToGoSource {
//     fn to_go_source(&self) -> String;
// }

// impl ToGoSource for NamedType {
//     fn to_go_source(&self) -> String {
//         println!("{:?}", self);
//         match &self.tref {
//             TypeRef::Name(nt) => todo!(),
//             TypeRef::Value(ty) => match &**ty {
//                 Type::Record(s) => {
//                     let mut out = String::new();
//                     let name = self.name.as_str().to_upper_camel_case();
//                     if let Some(repr) = s.bitflags_repr() {
//                         out.push_str(&format!("type {} = {}\n", &name, repr.to_go_source()));
//                         out.push_str("\n");
//                         out.push_str("const (\n");
//                         for (i, member) in s.members.iter().enumerate() {
//                             let docs = doc_comment(&member.docs);
//                             let member_name = member.name.as_str().to_upper_camel_case();
//                             out.push_str(&format!(
//                                 "{}\n{}{} {} = {}\n",
//                                 &docs, &name, &member_name, &name, i
//                             ));
//                         }
//                         out.push_str(")");
//                     } else {
//                         out.push_str(&format!("type {} struct {{\n", name));
//                         for member in s.members.iter() {
//                             let docs = doc_comment(&member.docs);
//                             out.push_str(&format!(
//                                 "{}\n{} {}",
//                                 &docs,
//                                 member.name.to_go_source(),
//                                 member
//                             ));
//                         }
//                     }
//                     out
//                 }
//                 Type::Handle(h) => {
//                     format!("HANDLE")
//                 }
//                 _ => format!("TODO"),
//             },
//         }
//     }
// }

// impl ToGoSource for IntRepr {
//     fn to_go_source(&self) -> String {
//         match self {
//             IntRepr::U8 => "uint8".to_owned(),
//             IntRepr::U16 => "uint16".to_owned(),
//             IntRepr::U32 => "uint32".to_owned(),
//             IntRepr::U64 => "uint64".to_owned(),
//         }
//     }
// }
