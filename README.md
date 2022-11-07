# Groom

`groom` is to Go what `crate` is to Rust.

It provides quality of life features missing currenly in `go` binary.

## Features

- Create multiple projects.
- Automatic Makefile generation
- Automatic Extract documentation
- Setup tests
- And Many More...

## Prequisites
- make
- Git 
- Bash
- Go

`groom` generates a Makefile. To use it install `make`.

## Usage

`groom` has many subcommands, each handles a part of golang project management

Currently 1 subcommand is supported

### Create

This subcommand creates a project out of a predetermined templates.
For the list of templates see [here]().


## Roadmap

- [ ] Integrate subcommands like `cargo` does. Example `groom-deb` should have support for making deb files.
- [ ] Provide `release` subcommand. Thus `groom release` should compile for multiple architectures.
- [ ] Handle global package management. Provide `groom install/remove` commands.
- [ ] Generate `documentation`.

