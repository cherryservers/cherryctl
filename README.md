cherryctl
================
<p align="left">
    <img width="100" height="100" src="https://pbs.twimg.com/profile_images/900630217630285824/p46dA56X_400x400.jpg" alt="Cherry Servers ctl." />
</p>
<p align="left">
  <a href="https://goreportcard.com/report/github.com/cherryservers/cherryctl">
    <img src="https://goreportcard.com/badge/github.com/cherryservers/cherryctl" alt="Go Report Card" />
  </a>
</p>

## Table of Contents

* [Introduction](#introduction)
* [Requirements](#requirements)
* [Supported Platforms](#supported-platforms)
* [Installation](#installation)
    * [Install via Homebrew](#use-homebrew-to-install-cherryctl)
    * [Install binary from Source](#install-binary-from-source)
    * [Install binary from Release Download](#install-binary-from-release-download)
* [Shell Completion](#shell-completion)
* [Authentication](#authentication)
    * [Working With Multiple Contexts](#working-with-multiple-contexts)
* [Configuring Default Values](#configuring-default-values)
* [Documentation](#documentation)
* [Examples](#examples)

## Introduction

The Cherry Servers CLI wraps the [Cherry Servers Go SDK](https://github.com/cherryservers/cherrygo) allowing interaction
with Cherry Servers platform from a command-line interface.

## Requirements

* Cherry Servers authentication token.
* Cherry Servers CLI [binaries](https://github.com/cherryservers/cherryctl/releases).

## Supported Platforms

The [Cherry Servers CLI binaries](https://github.com/cherryservers/cherryctl/releases) are available for Linux, Windows,
and Mac OS X for various architectures including ARM.

## Installation

### Install `cherryctl` Using [Homebrew](https://brew.sh/) Package Manager

```sh
brew tap cherryservers/cherryctl
brew install cherryctl
```

### Install `cherryctl` from the [AUR](https://aur.archlinux.org/packages/cherryctl)

```sh
paru -S cherryctl
```

### Install `cherryctl` from Source

If you have `go` 1.17 or later installed, you can build and install the latest development version with:

```sh
go install github.com/cherryservers/cherryctl@latest
```

You can find the installed executable/binary in either `$GOPATH/bin` or `$HOME/go/bin` folder.

### Install a Specific `cherryctl` Release from Source

Visit the [Releases page](https://github.com/cherryservers/cherryctl/releases) for the
[`cherryctl` GitHub project](https://github.com/cherryservers/cherryctl), and find the
appropriate binary for your operating system and architecture. Download the appropriate binaries for your platform to
the desired location, `chmod +x` it and rename it to `cherryctl`.

## Shell Auto-completion

Once `cherryctl` is installed, you may generate an auto-completion script and load it to your current shell session to
use `cherryctl` more conveniently. For instance, to enable auto-completion for bash shell use the following command:

```sh
source <(cherryctl completion bash)
```

If you want to make the auto-completion script load every time you initiate a bash session, place a new shell script in
the bash completion directory:

```sh
cherryctl completion bash > /etc/bash_completion.d/cherry-autocomplete.sh
```

Check `cherryctl completion -h` for instructions, if you are using other shells.

## Authentication

After installing Cherry Servers CLI, configure your account using `cherryctl init`:

```sh
$ cherryctl init
Cherry Servers API Tokens can be obtained through the portal at https://portal.cherryservers.com/.

Token (hidden): 
Team ID []: 12345
Project ID []: 123456

Writing configuration to: /Users/username/.config/cherry/default.yaml
```

This will create the `cherryctl` directory with a default context in your default OS user configuration directory. If
you wish to store this context in a
different location, you can pass a custom path with the `--config` option or by setting the `CHERRY_CONFIG` environment
variable. You can also override the API token on a case-by-case basis by setting the `CHERRY_AUTH_TOKEN` environment
variable, but a context is still required.

### Working With Multiple Contexts

A context is a collection of settings specific to a certain user that is stored in a configuration file. It
consists of at least an API token, a Team ID and a Project ID, but you may add many additional configuration options.

You may work with multiple contexts at the same time, since `cherryctl` allows you to switch between them by using
the `--context` option. You can also switch between context directories, with the `--config` option.

By default, the `--context` option has a value `default`. To create a new context, run
`cherryctl init --context <new_context_name>`. You will be prompted for a Token, a Team ID and a Project ID which will
be associated with the new context. You will be able to add any other options by editing the newly generated
context file.

To use a non-default context name to any `cherryctl` command:

```sh

cherryctl servers list --context <new_context_name>

```

## Configuring Default Values

If you find yourself using certain flags frequently, you can set their default values to avoid typing them
every time. This can be useful when, for example, you want to deploy all infrastructure in the same region.

If you want to change the default value for a `--region` flag, open your context file and add the corresponding
key-value pair at the end of the file. For instance, in the following example we have changed the default region to
LT-Siauliai:

```yaml
[ ... ]
region: LT-Siauliai
```

## Documentation

The full CLI documentation can be found [here](docs/cherryctl.md).

## Examples

### List Plans

You may list all the plans that are available on Cherry Servers stock by using the following command:

```sh
cherryctl plan list
```

`cherryctl` only allows you to order services with hourly and spot billing cycles. If you wish to get a fixed term
plan, use the Client Portal instead.

### List Plan Images

Every plan may have a different set of available images. Use a selected plan slug obtained from the
`cherryctl plan list`
command to get a list of available images for that plan:

```sh
cherryctl image list --plan [plan_slug]
```

### List Regions

You may find relevant information about available regions by using the following command:

```sh
cherryctl region list
```

### Deploy a New Server

Provision a new server:

```sh
cherryctl server create --plan [plan_slug] --image [os_slug] --region [region_slug] --hostname [hostname]
```

### List Your Servers

You may check the full list of your active servers:

```sh
cherryctl server list
```

### Get Information About Existing Server

If you want to inspect a single server, you may use the following command. You may specify a `-o json` or `-o yaml` flag
if you want to change the output format.

```sh
cherryctl server get -o json [server_ID]
```
