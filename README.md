# Groom

`groom` is a project management tool for Golang.
It tries to merge multiple tools and provide quality of life features missing in the Golang binary.

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


### Create

This subcommand creates a project out of a predetermined templates.
For the list of templates and get more information [groom-create](https://github.com/golang-groom/groom-create).

## Roadmap

- [ ] Detect available subcommands list them.

## Possible ideas.
- [ ] Groom task (`groom-task`). A system-agnostic task runner.
- [ ] Groom Release (`groom-release`). Compile and compress for all platforms
- [ ] Groom Deb (`groom-deb`). Generate a deb file for Debian platforms.
- [ ] Groom Install/Remove (`groom-install`/`groom-remove`). Manage golang binaries and their depedencies.
- [ ] Groom Documentation (`groom-doc`). Generate and serve documenation.

## Contributing
Contribution is gladly welcome. Read the [Contributing Guidlines]() before contributing.
