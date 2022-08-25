![README](data/gfx/README.gif)

[![dev-build](https://github.com/akdev1l/toolbox/actions/workflows/build-dev-release.yml/badge.svg)](https://github.com/akdev1l/toolbox/actions/workflows/build-dev-release.yml)

**This is an unofficial fork not related to the original project**

[Toolbox](https://containertoolbx.org/) is a tool for Linux operating systems,
which allows the use of containerized command line environments. It is built
on top of [Podman](https://podman.io/) and other standard container
technologies from [OCI](https://opencontainers.org/).

This is particularly useful on
[OSTree](https://ostree.readthedocs.io/en/latest/) based operating systems like
[Fedora CoreOS](https://coreos.fedoraproject.org/) and
[Silverblue](https://silverblue.fedoraproject.org/). The intention of these
systems is to discourage installation of software on the host, and instead
install software as (or in) containers — they mostly don't even have package
managers like DNF or YUM. This makes it difficult to set up a development
environment or install tools for debugging in the usual way.

Toolbox solves this problem by providing a fully mutable container within
which one can install their favourite development and debugging tools, editors
and SDKs. For example, it's possible to do `yum install ansible` without
affecting the base operating system.

However, this tool doesn't *require* using an OSTree based system. It works
equally well on Fedora Workstation and Server, and that's a useful way to
incrementally adopt containerization.

The toolbox environment is based on an [OCI](https://www.opencontainers.org/)
image. On Fedora this is the `fedora-toolbox` image. This image is used to
create a toolbox container that seamlessly integrates with the rest of the
operating system by providing access to the user's home directory, the Wayland
and X11 sockets, networking (including Avahi), removable devices (like USB
sticks), systemd journal, SSH agent, D-Bus, ulimits, /dev and the udev
database, etc..


## Installation & Use

See our guides on
[installing & getting started](https://containertoolbx.org/install/) with
Toolbox and [Linux distro support](https://containertoolbx.org/distros/).

