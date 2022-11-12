# Groom

`groom` is to Go what `crate` is to Rust.

It provides quality of life features missing currenly in `go` binary.

## Features

- Generate project structure (`groom-create`)
- Install Golang binary (`groom-install`)

## Prequisites
- Git 
- Bash
- Go

`groom` generates a Makefile. To use it install `make`.

## Usage

`groom` has many subcommands, each handles a part of golang project management

Currently 1 subcommand is supported

### Create

This subcommand creates a project out of a predetermined templates.
For the list of templates and get more information [groom-create](https://github.com/golang-groom/groom-create).

## Roadmap

- [ ] Detect available subcommands list them.

## Possible subcommands
- [ ] Groom Release (`groom-release`). Compile and compress for all platforms
- [ ] Groom Deb (`groom-deb`). Generate a deb file for Debian platforms.
- [ ] Groom Install/Remove (`groom-install`/`groom-remove`). Manage golang binaries and their depedencies.
- [ ] Groom Documentation (`groom-doc`). Generate and serve documenation.

## Contributing
Contribution is gladly welcome. Read the [Contributing Guidlines]() before contributing.
