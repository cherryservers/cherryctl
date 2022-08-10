#!/usr/bin/env bash

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/../"
OUT_D=${OUT_D:-${DIR}/builds}
mkdir -p "$OUT_D"

base="-X github.com/cherryservers/cherryctl"
version="$(git rev-parse --short HEAD)"
ldflags="${base}/cmd.Version=${version}"

(
  export GOOS=${GOOS:-darwin}
  export GOARCH=${GOARCH:-arm64}
  go build -ldflags "$ldflags" -o "${OUT_D}/cherryctl_${GOOS}_${GOARCH}"
)