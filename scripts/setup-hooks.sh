#!/usr/bin/env bash

# Execute this script from repo root to configure git to use .githooks/pre-commit as the precommit hook for this repository.

set -euo pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"

git -C "$REPO_ROOT" config --local core.hooksPath .githooks

echo "Git hooks enabled."
