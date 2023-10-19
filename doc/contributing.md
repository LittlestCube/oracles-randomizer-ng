# Building and contributing to oracles-randomizer-ng

If you plan to build the randomizer from source or contribute to development,
here are some things to know:


## Building

An environment with Git and Go is required to build the randomizer. The
following instructions are for building the dev branch.

First, install git and golang, and add `$HOME/go/bin` to your `$PATH` (for `esc`):

```
sudo apt install git golang
echo "PATH=$HOME/go/bin:$PATH" >> $HOME/.bashrc
```

Alternately, you can copy `$HOME/go/bin/esc` to somewhere that's already
in your `$PATH`.

Then, clone and cd into the repository:

```
mkdir -p $HOME/go/src/github.com/LittlestCube/oracles-randomizer-ng
git clone https://github.com/LittlestCube/oracles-randomizer-ng.git $HOME/go/src/github.com/LittlestCube/oracles-randomizer-ng
cd $HOME/go/src/github.com/LittlestCube/oracles-randomizer-ng
```

Now you can run the makefile:

```
make
```

This runs some first-time setup, as well as runs `go generate` and `go build`
which compiles the code. You will need to run `make` again anytime changes
are made. Now there should be a binary `oracles-randomizer-ng`.


## Organization

Go code is in `randomizer/`, but some types of changes don't even need to touch
the Go code. Logic is in `logic/`, GBC assembly code is in `asm/`, owl hint
names are in `hints/`, and various ROM addresses and values are in `romdata/`.
All the non-Go directories use YAML, although sometimes the contents of the
YAML amount to something more like a domain-specific language.


## Code style

Always run `go fmt ./randomizer/` before commits (note that `go fmt` coerces
all line-initial indentation to tabs). Wrap lines (in Go and YAML) at 80
characters when possible, assuming 8-space tabs for Go. YAML is necessarily
indented using spaces; the standard varies between 2 and 4 spaces, depending on
the conventions of the directory.
