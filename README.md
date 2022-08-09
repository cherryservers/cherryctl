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
  * [Install binary from Source](#install-binary-from-source)
  * [Install binary from Release Download](#install-binary-from-release-download)
* [Shell Completion](#shell-completion)
* [Authentication](#authentication)
* [Configuring Default Values](#configuring-default-values)
* [Documentation](#documentation)
* [Examples](#examples)

## Introduction

The Cherry Servers CLI wraps the [Cherry Servers Go SDK](https://github.com/cherryservers/cherrygo) allowing interaction with Cherry Servers platform from a command-line interface.

## Requirements

* Cherry Servers authentication token.
* Cherry Servers CLI [binaries](https://github.com/cherryservers/cherryctl/releases).

## Supported Platforms

The [Cherry Servers CLI binaries](https://github.com/cherryservers/cherryctl/releases) are available for Linux, Windows, and Mac OS X for various architectures including ARM.

## Installation

### Install binary from Source

If you have `go` 1.17 or later installed, you can build and install the latest development version with:

```sh
go install github.com/cherryservers/cherryctl@latest
```

You can find the installed executable/binary in either `$GOPATH/bin` or `$HOME/go/bin` folder.

### Install binary from Release Download

Visit the [Releases page](https://github.com/cherryservers/cherryctl/releases) for the
[`cherryctl` GitHub project](https://github.com/cherryservers/cherryctl/doctl), and find the
appropriate binary for your operating system and architecture. Download the appropriate binaries for your platform to the desired location, `chmod +x` it and rename it to `cherryctl`.

## Shell Completion

Once installed, shell completion can be enabled (in Bash) with `source <(cherryctl completion bash)`.

Check `cherryctl completion -h` for instructions to use in other shells.

## Authentication

After installing Cherry Servers CLI, configure your account using `cherryctl init`:

```bash
$ cherryctl init
Cherry Servers API Tokens can be obtained through the portal at https://portal.cherryservers.com/.

Token (hidden): 
Team ID []: 12345
Project ID []: 123456

Writing /Users/username/.config/cherry/cherry.yaml
```

The Cherry Servers authentication token can be stored in the `$CHERRY_AUTH_TOKEN` environment variable or in JSON or YAML configuration files. The configuration file path can be overridden with the `--config` flag.  The default configuration path is "$HOME/config/cherry/cherry.*" (any supported filetype).

## Configuring Default Values

The `cherryctl` configuration file is used to store your API authentication token as well as the defaults for command flags. If you find yourself using certain flags frequently, you can set their default values to avoid typing them every time. This can be useful when, for example, you want to deploy all infrastructure in same region.

`cherryctl` saves its configuration as `${HOME}/cherry/config.yaml`. The `${HOME}/cherry/` directory will be created once you run `cherryctl init`.

To change the default value for `--region` flag, open `.config.yaml` file and add the value at the end of file. In this example, we changed default region to eu_nord_1.
```
...
region: eu_nord_1
```

## Documentation

The full CLI documentation can be found [here](docs/cherryctl.md).

## Examples

### List plans

```sh
cherryctl plan list
```

### List plan images

```sh
cherryctl image list --plan [plna_slug]
```

### List regions

```sh
cherryctl region list
```

### Create a device

```sh
cherryctl server create --plan [plan_slug] -h [hostname] --image [os_slug] --region [region_slug]
```

### Get a device

```sh
cherryctl server get -i [server_ID]
```