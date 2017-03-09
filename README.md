# git-pair
[![Build Status](http://52.170.46.90/api/badges/seanknox/git-pair/status.svg)](http://52.170.46.90/seanknox/git-pair)

Commit code as multiple people!

## Developers Guide

This guide explains how to set up your environment for developing on
git-pair.

## Prerequisites

- Go 1.6.0 or later
- Glide 0.12.0 or later

## Building git-pair

We use Make to build our programs. The simplest way to get started is:

```console
$ make bootstrap build
```

This will build git-pair. `make bootstrap` will attempt to
install certain tools if they are missing.

To run all of the tests (without running the tests for `vendor/`), run
`make test`.

## Contribution Guidelines

We welcome contributions. This project has set up some guidelines in
order to ensure that (a) code quality remains high, (b) the project
remains consistent, and (c) contributions follow the open source legal
requirements. Our intent is not to burden contributors, but to build
elegant and high-quality open source code so that our users will benefit.

### Structure of the Code

The code is organized as follows:

- The individual programs are located in `cmd/`. Code inside of `cmd/`
  is not designed for library re-use.
- Shared libraries are stored in `pkg/`.

Go dependencies are managed with
[Glide](https://github.com/Masterminds/glide) and stored in the
`vendor/` directory.

### Git Conventions

We use Git for our version control system. The `master` branch is the
home of the current development candidate.

We accept changes to the code via GitHub Pull Requests (PRs). One
workflow for doing this is as follows:

1. Use `go get` go clone the repo into your $GOPATH: `go get github.com/seanknox/git-pair`
2. Fork that repository into your GitHub account
3. Add your repository as a remote for `$GOPATH/src/github.com/seanknox/git-pair`
4. Create a new working branch (`git checkout -b feat/my-feature`) and
   do your work on that branch.
5. When you are ready for us to review, push your branch to GitHub, and
   then open a new pull request with us.

### Go Conventions

We follow the Go coding style standards very closely. Typically, running
`go fmt` will make your code beautiful for you.

We also typically follow the conventions recommended by `go lint` and
`gometalinter`.

Read more:

- Effective Go [introduces formatting](https://golang.org/doc/effective_go.html#formatting).
- The Go Wiki has a great article on [formatting](https://github.com/golang/go/wiki/CodeReviewComments).
