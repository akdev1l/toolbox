FROM fedora:latest

RUN dnf install -y \
      ShellCheck \
      bash-completion \
      bats \
      flatpak-session-helper \
      golang \
      golang-github-cpuguy83-md2man \
      httpd-tools \
      meson \
      ninja-build \
      openssl \
      podman \
      skopeo \
      systemd \
      udisks2 && \
    mkdir -p /build && \
    meson setup /source /build/ \
      --prefix /usr/local \
      -Dinstall_completions=true && \
    TOOLBOX_PATH=/source/build/install/bin/toolbox ninja -C /build -j12 install

FROM fedora:latest

ENV CONTAINER_HOST=unix:/run/user/1000/podman/podman.sock 

RUN dnf install -y podman-remote man

COPY --from=0 /usr/local /usr/local

ENTRYPOINT ["/usr/local/bin/toolbox"]
