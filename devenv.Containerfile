FROM docker.io/akdev1l/devenv:latest

RUN dnf install -y go meson go-md2man ShellCheck
