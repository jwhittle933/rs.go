# rs.go
When you want to write Rust, but you're stuck writing Go.

## Motivation
TODO

## Limitations
Rust is semantically magnificent, and that's a compiler matter. Go's compiler is fine and fast, but it cannot be coerced into operating like the Rust compiler (and you shouldn't expect it to). While some of the semantics, types, and behavior can be accomplished through ordinary Go means, many cannot.

- Go's Generics (1.18) are immature and dissapointing if Rust is your standard. Interfaces cannot define generic methods, which is severe limitation if you're trying to implement something like a `Result`.
- The semantics of type instantiation in Rust (`Thing::new()`) obviously cannot be recreated in Go. The closest option would be `(Thing{}).new()`, which already exists and don't do that.
- Go's packages differ from Rust's modules. In Rust a file can have many modules, but many files can never be a single module.
- Import semantics: they're different and nothing can be done about it. Go's imports must be at the top of the file, and they follow the file system/module path. And there is no re-exporting without some type aliasing hack. Don't do that.
- Rust's `match` can only be accomplished in Go with a `switch`. Close enough.
- Enums are amazing, and Go doesn't have them. Go figure. Hacks will follow.