![README](data/gfx/README.gif)

**This is an unofficial fork not related to the original project**

[Toolbox](https://containertoolbx.org/) is a tool for Linux operating systems,
which allows the use of containerized command line environments. It is built
on top of [Podman](https://podman.io/) and other standard container
technologies from [OCI](https://opencontainers.org/).

### Why this fork?

The original upstream project seems stalled - they have over 68 PRs some of which are years old. There are long standing issues
that affect quality of life and also missing features when comparing to the more actively developed solution `distrobox`.

As I don't have a lot of hopes in PRs being merged in a timely manner by the upstream project I have taken upon myself to develop
`toolbox` in a way that makes it comparable to `distrobox`. I believe the static image approach is prone to less user-facing issues and
I also think it is good to have multiple `toolbox` implementations that are somewhat compatible.

### CI Status

| Description | Status |
|-|-|
|Development Build|[![dev-build](https://github.com/akdev1l/toolbox/actions/workflows/build-dev-release.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/build-dev-release.yml)|
|Code Linting|[![code-lint](https://github.com/akdev1l/toolbox/actions/workflows/linting.yaml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/linting.yaml)|


### Cross-Distro Support

|Distribution Name|Dockerhub Link|Build Status|
|-|-|-|
|Alpine|[https://hub.docker.com/repository/docker/akdev1l/alpine-toolbox](https://hub.docker.com/repository/docker/akdev1l/alpine-toolbox)|[![build alpine toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/alpine.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/alpine.yml)|
|Arch Linux|[https://hub.docker.com/repository/docker/akdev1l/archlinux-toolbox](https://hub.docker.com/repository/docker/akdev1l/archlinux-toolbox)|[![build archlinux toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/archlinux.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/archlinux.yml)|
|CentOS Stream|[https://hub.docker.com/repository/docker/akdev1l/centos-toolbox](https://hub.docker.com/repository/docker/akdev1l/centos-toolbox)|[![build centos toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/centos.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/centos.yml)|
|Kali Linux|[https://hub.docker.com/repository/docker/akdev1l/kalilinux-toolbox](https://hub.docker.com/repository/docker/akdev1l/kalilinux-toolbox)|[![build kalilinux toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/kalilinux.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/kalilinux.yml)|
|OpenSuSE Tumbleweed|[https://hub.docker.com/repository/docker/akdev1l/tumbleweed-toolbox](https://hub.docker.com/repository/docker/akdev1l/tumbleweed-toolbox)|[![build tumbleweed toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/tumbleweed.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/tumbleweed.yml)|
|Ubuntu|[https://hub.docker.com/repository/docker/akdev1l/ubuntu-toolbox](https://hub.docker.com/repository/docker/akdev1l/ubuntu-toolbox)|[![build ubuntu toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/ubuntu.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/ubuntu.yml)|


## Installation & Use

#### Using wrapper script

You can deploy this version of toolbox quickly by using the power of podman. We support a custom
toolbox container image and will automatically pull from the dockerhub when the wrapper is run.

To quickly test the version stored in Dockerub follow these steps:

```
$ curl https://raw.githubusercontent.com/akdev1l/toolbox/akdev/toolbox -o toolbox
$ chmod +x toolbox
$ ./toolbox list
```

The wrapper will deploy a temporary version of toolbox in `/run/user/$(id -u)/toolbox/bin/toolbox`.
This is meant to be used as a quick way of testing new versions - it doesn't do any permanent changes to the
system.

#### Getting a development build

You can fetch the latest successful build by going to [dev-build](https://github.com/akdev1l/toolbox/actions/workflows/build-dev-release.yml)
Github Action and click on the latest successful build.

There are two artifacts - `toolbox` and `toolbox-bin`. The first contains the entire binary release including man pages, the latter 
only contains the `toolbox` which is all that is required to run it.

## Usage


```
$ toolbox --help
Tool for containerized command line environments on Linux

Usage:
  toolbox [command]

Available Commands:
  completion     Generate the autocompletion script for the specified shell
  create         Create a new toolbox container
  enter          Enter a toolbox container for interactive use
  export         Exports an application, binary or service to the host
  help           Help about any command
  list           List existing toolbox containers and images
  rm             Remove one or more toolbox containers
  rmi            Remove one or more toolbox images
  run            Run a command in an existing toolbox container

Flags:
  -y, --assumeyes          Automatically answer yes for all questions
  -h, --help               help for toolbox
      --log-level string   Log messages at the specified level: trace, debug, info, warn, error, fatal or panic (default "info")
      --log-podman         Show the log output of Podman. The log level is handled by the log-level option
  -v, --verbose count      Set log-level to 'debug'
      --version            version for toolbox

Use "toolbox [command] --help" for more information about a command.
```

#### Features

1. Experimental isolated home support
2. Ephemeral toolboxes support
3. Containerized released for easier testing
4. Bug fix addressing long standing issues with the upstream implementation
5. No binary patching, statically linked binary will work across different systems
