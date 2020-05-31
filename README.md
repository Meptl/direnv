direnv -- unclutter your .profile
=================================

This is a fork of [direnv](https://github.com/direnv/direnv). See their
documentation for basic usage. This fork adds support for running commands in
the interactive shell after loading the environment and before unloading the
environment.

Current functionality only supports zsh, but accomodating additional shells
should be trivial.

# Usage
Add the following to .envrc
```
$ cd ~/my-project
.envrc is blocked.
$ cat .envrc
postload_eval "alias foo='echo 4'"
preunload_eval "unalias foo

# Show that the FOO environment variable is not loaded.
$ echo ${FOO-nope}
nope

# Create a new .envrc. This file is bash code that is going to be loaded by
# direnv.
$ echo export FOO=foo > .envrc
.envrc is not allowed

# The security mechanism didn't allow to load the .envrc. Since we trust it,
# let's allow its execution.
$ direnv allow .
direnv: loading .envrc

$ foo
4

$ cd ..
direnv: unloading

$ foo
foo: command not found
```

You may want your shell native modifiers to be part of ~/.direnvrc
```
alias_add() {
  postload_eval "alias $1='$2'"
  preunload_eval "unalias $1"
}
```

### The stdlib

Exporting variables by hand is a bit repetitive so direnv provides a set of
utility functions that are made available in the context of the `.envrc` file.

As an example, the `PATH_add` function is used to expand and prepend a path to
the $PATH environment variable. Instead of `export PATH=$PWD/bin:$PATH` you
can write `PATH_add bin`. It's shorter and avoids a common mistake where
`$PATH=bin`.

To find the documentation for all available functions check the
[direnv-stdlib(1) man page](man/direnv-stdlib.1.md).

It's also possible to create your own extensions by creating a bash file at
`~/.config/direnv/direnvrc` or `~/.config/direnv/lib/*.sh`. This file is
loaded before your `.envrc` and thus allows you to make your own extensions to
direnv.

## Docs

* [Install direnv](docs/installation.md)
* [Hook into your shell](docs/hook.md)
* [Develop for direnv](docs/development.md)
* [Manage your rubies with direnv and ruby-install](docs/ruby.md)
* [Community Wiki](https://github.com/direnv/direnv/wiki)

Make sure to take a look at the wiki! It contains all sorts of useful
information such as common recipes, editor integration, tips-and-tricks.

### Man pages

* [direnv(1) man page](man/direnv.1.md)
* [direnv-stdlib(1) man page](man/direnv-stdlib.1.md)
* [direnv.toml(1) man page](man/direnv.toml.1.md)

### FAQ

Based on GitHub issues interactions, here are the top things that have been
confusing for users:

1. direnv has a standard library of functions, a collection of utilities that
   I found useful to have and accumulated over the years. You can find it
   here: https://github.com/direnv/direnv/blob/master/stdlib.sh

2. It's possible to override the stdlib with your own set of function by
   adding a bash file to `~/.config/direnv/direnvrc`. This file is loaded and
   it's content made available to any `.envrc` file.

3. direnv is not loading the `.envrc` into the current shell. It's creating a
   new bash sub-process to load the stdlib, direnvrc and `.envrc`, and only
   exports the environment diff back to the original shell. This allows direnv
   to record the environment changes accurately and also work with all sorts
   of shells. It also means that aliases and functions are not exportable
   right now.

## Contributing

Bug reports, contributions and forks are welcome. All bugs or other forms of
discussion happen on http://github.com/direnv/direnv/issues .

Or drop by on [IRC (#direnv on freenode)](irc://irc.freenode.net/#direnv) to
have a chat. If you ask a question make sure to stay around as not everyone is
active all day.

## Complementary projects

Here is a list of projects you might want to look into if you are using direnv.

* [starship](https://starship.rs/) - A cross-shell prompt.
* [nix-direnv](https://github.com/nix-community/nix-direnv) - A fast, persistent use_nix implementation for direnv.

## Related projects

Here is a list of other projects found in the same design space. Feel free to
submit new ones.

* [Environment Modules](http://modules.sourceforge.net/) - one of the oldest (in a good way) environment-loading systems
* [autoenv](https://github.com/kennethreitz/autoenv) - lightweight; doesn't support unloads
* [zsh-autoenv](https://github.com/Tarrasch/zsh-autoenv) - a feature-rich mixture of autoenv and [smartcd](https://github.com/cxreg/smartcd): enter/leave events, nesting, stashing (Zsh-only).
* [asdf](https://github.com/asdf-vm/asdf) - a pure bash solution that has a plugin system. The [asdf-direnv](https://github.com/asdf-community/asdf-direnv) plugin allows using asdf managed tools with direnv.
* [ondir](https://github.com/alecthomas/ondir) - OnDir is a small program to automate tasks specific to certain directories
* [shadowenv](https://shopify.github.io/shadowenv/) - uses an s-expression format to define environment changes that should be executed

## COPYRIGHT

[MIT licence](LICENSE) - Copyright (C) 2019 @zimbatm and [contributors](https://github.com/direnv/direnv/graphs/contributors)
