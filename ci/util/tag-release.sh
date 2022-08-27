#!/bin/bash


version="$(awk '/ version:/{print $2}' FS="'" meson.build)"

git tag -a "${version}"

