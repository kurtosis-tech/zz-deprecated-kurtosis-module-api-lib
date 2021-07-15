#!/usr/bin/env bash
set -euo pipefail
script_dirpath="$(cd "$(dirname "${0}")" && pwd)"
lang_root_dirpath="$(dirname "${script_dirpath}")"

cd "${lang_root_dirpath}"
go test ./...
