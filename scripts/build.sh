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
  
  # Enable shadow stack (SHSTK) support for x86-64 Linux builds
  # This adds Intel CET (Control-flow Enforcement Technology) protection
  if [[ "$GOOS" == "linux" && "$GOARCH" == "amd64" ]]; then
    export CGO_ENABLED=1
    export CGO_CFLAGS="-fcf-protection=full"
    export CGO_LDFLAGS="-Wl,-z,shstk -Wl,-z,ibt -Wl,-z,cet-report=error"
  fi
  
  go build -ldflags "$ldflags" -o "${OUT_D}/cherryctl_${GOOS}_${GOARCH}"
)