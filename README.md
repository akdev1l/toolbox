![README](data/gfx/README.gif)

**This is an unofficial fork not related to the original project**

[Toolbox](https://containertoolbx.org/) is a tool for Linux operating systems,
which allows the use of containerized command line environments. It is built
on top of [Podman](https://podman.io/) and other standard container
technologies from [OCI](https://opencontainers.org/).

### CI Status

| Description | Status |
|-|-|
|Development Build|[![dev-build](https://github.com/akdev1l/toolbox/actions/workflows/build-dev-release.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/build-dev-release.yml)|
|Code Linting|[![code-lint](https://github.com/akdev1l/toolbox/actions/workflows/linting.yaml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/linting.yaml)|


### Cross-Distro Support

|Distribution Name|Dockerhub Link|Build Status|
|-|-|-|
|Alpine|[docker.io/akdev1l/alpine-toolbox](docker.io/akdev1l/alpine-toolbox)|[![build alpine toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/alpine.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/alpine.yml)|
|Arch Linux|[docker.io/akdev1l/archlinux-toolbox](docker.io/akdev1l/archlinux-toolbox)|[![build archlinux toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/archlinux.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/archlinux.yml)|
|CentOS Stream|[docker.io/akdev1l/centos-toolbox](docker.io/akdev1l/centos-toolbox)|[![build centos toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/centos.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/centos.yml)|
|Kali Linux|[docker.io/akdev1l/kalilinux-toolbox](docker.io/akdev1l/kalilinux-toolbox)|[![build kalilinux toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/kalilinux.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/kalilinux.yml)|
|OpenSuSE Tumbleweed|[docker.io/akdev1l/tumbleweed-toolbox](docker.io/akdev1l/tumbleweed-toolbox)|[![build tumbleweed toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/tumbleweed.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/tumbleweed.yml)|
|Ubuntu|[docker.io/akdev1l/ubuntu-toolbox](docker.io/akdev1l/ubuntu-toolbox)|[![build ubuntu toolbox container images](https://github.com/akdev1l/toolbox/actions/workflows/ubuntu.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/ubuntu.yml)|


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
