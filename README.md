---
title: Go Basics Workshop
author: Jan Waś
...

# Agenda

Estimated time: 4x3 hours over 2 days

* quick poll
* intro
* installing go, directories overview

Basics

* packaging
* go get, go build, go install
* fmt, imports, linter
* vars, consts
* functions
* data types
* for, if, switch
* pointers
* structs
* arrays and slices
* maps
* methods

# Mid level

* interfaces
* goroutines
* channels
* error handling
* (file) io
* unit tests

# Advanced (not really)

* stdlib
* ide: AC, def/usage, debug
* package/dependency management
* marshall/unmarshall (sed/deser)
* docs
* benchmarking
* calling remote apis
* parsing cli args
* db access
* logging

# Requirements

Bring a laptop!

Go through:
* https://tour.golang.org/welcome/1
* https://gobyexample.com

Setup Go env to be able to compile and run programs. See the [Install](#Install Go) section.

Prepare your favourite text editor/IDE.

# Poll

* Go knowledge
* Other programming languages experience
* Fav environment

# Intro

Go:

* is expressive, concise, clean, and efficient
* has easy concurrency constructs for multicore and networking programs
* its type system enables flexible and modular program construction
* compiles quickly to machine code yet has garbage collection and run-time reflection
* is opinionated

<!--
go is opinionated
-->

# Install Go

## Download

* https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
* https://dl.google.com/go/go1.11.2.darwin-amd64.pkg
* https://dl.google.com/go/go1.11.2.windows-amd64.msi

## Install

```bash
VERSION=1.11.2
OS=linux
ARCH=amd64
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

## Setup

In `~/.bashrc`:

```bash
export PATH=$PATH:~/bin:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)

go env
```

# Packaging

```bash
go get github.com/nineinchnick/go-workshop
cd $GOPATH/src/github.com/nineinchnick/go-workshop

go run 00-pkg/main.go
```

# go get, go build, go install

```bash
go help
go help get
tree -L 2 ~/go
```

```
/home/jwas/go
├── bin
│   ├── dlv
│   ├── errcheck
│   ├── gocode
│   ├── godef
│   ├── goimports
│   ├── golint
│   ├── gometalinter
│   ├── gorename
│   ├── gotags
│   └── tour
├── pkg
│   ├── dep
│   └── linux_amd64
└── src
    ├── github.com
    ├── gopkg.in
    └── sourcegraph.com
```

# fmt, imports, linter

```
% go help fmt
usage: go fmt [-n] [-x] [packages]

Fmt runs the command 'gofmt -l -w' on the packages named
by the import paths. It prints the names of the files that are modified.

For more about gofmt, see 'go doc cmd/gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

To run gofmt with specific options, run gofmt itself.

See also: go fix, go vet.
```

# gometalinter

```bash
gometalinter \
    --skip=vendor \
    --disable-all \
    --enable=golint \
    --enable=misspell \
    --enable=vetshadow \
    --enable=gotype \
    --enable=megacheck \
    --enable=vet \
    --enable=goconst \
    --enable=ineffassign \
    --enable=staticcheck \
    --enable=gocyclo \
    --enable=gosec \
    --deadline=300s \
    ./...;
```

# vars, consts

The following examples contains errors on purpose.
Try to fix them and make everything run.

```bash
cat 01-vars/main.go
```

# functions

# data types

# for, if, switch

# pointers

# structs

# arrays and slices

# maps

# methods

# Links

From https://golang.org/doc/:
* https://tour.golang.org/welcome/1
* https://gobyexample.com
* https://golang.org/doc/effective_go.html
* https://golang.org/doc/faq

Others:
* https://blog.golang.org/package-names
* https://github.com/kelseyhightower/intro-to-go-workshop
* https://www.youtube.com/watch?v=ndmB0bj7eyw

# Exit poll

* how was the instructor prepared?
* was the instructor helpful and making sure everyone keeps on track?
* pace
* scope
* duration
* rate this examples
* open questions: what did you like, what to improve
* rate the venue
