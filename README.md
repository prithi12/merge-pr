# merge-pr

Merge your GitHub pull requests from the command line.

## Motivation

Merging pull requests in the browser is nice, sure, but you lose a lot of
clarity into what changed & when at a higher level. When did you add that
feature? Oh, you're releasing a new patch? What changes did you make? No, I
won't read through your commit history.

This tool aims to make it easy to merge PR's and add a line to the
CHANGELOG (`History.markdown` by default). All with one command.

## Installation

You need [Go](https://golang.org) and you need your `$GOPATH` set &
`$GOPATH/bin` in your `$PATH`. Then:

```bash
$ go get github.com/parkr/merge-pr
```

Throw your credentials in `$HOME/.netrc`, like this:

```text
machine api.github.com
  login yourusername
  password mycoolpasswordnooneknows
```

## Usage

```bash
$ merge-pr 7
```

This will go to GitHub, merge the PR, delete the branch if it's on the same
repo, will pull down those changes, open up your editor (`$EDITOR`), then
commit that change.

## Credits / License

MIT License, copyright Parker Moore. Details in the `LICENSE` file.
