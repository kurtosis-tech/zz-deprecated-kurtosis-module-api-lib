#!/usr/bin/env bash
# ^^^^^^^^^^^^^^^^^ this is the most platform-agnostic way to guarantee this script runs with Bash
# 2021-07-08 WATERMARK, DO NOT REMOVE - This script was generated from the Kurtosis Bash script template

set -euo pipefail   # Bash "strict mode"
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root_dirpath="$(dirname "${script_dirpath}")"



# ==================================================================================================
#                                             Constants
# ==================================================================================================
GENERATOR_SCRIPT_FILENAME="generate-protobuf-bindings.sh"  # Must be on the PATH
PROTOBUF_DIRNAME="protobuf-api"
BINDINGS_DIRNAME="kurtosis_module_rpc_api_bindings"
GOLANG_DIRNAME="golang"
TYPESCRIPT_DIRNAME="typescript"



# ==================================================================================================
#                                             Main Logic
# ==================================================================================================
input_dirpath="${root_dirpath}/${PROTOBUF_DIRNAME}"

# Golang
go_output_dirpath="${root_dirpath}/${GOLANG_DIRNAME}/${BINDINGS_DIRNAME}"
if ! GO_MOD_FILEPATH="${root_dirpath}/${GOLANG_DIRNAME}/go.mod" "${GENERATOR_SCRIPT_FILENAME}" "${input_dirpath}" "${go_output_dirpath}" golang; then
    echo "Error: An error occurred generating Go bindings in directory '${go_output_dirpath}'" >&2
    exit 1
fi
echo "Successfully generated Go bindings in directory '${go_output_dirpath}'"

# TypeScript
typescript_output_dirpath="${root_dirpath}/${TYPESCRIPT_DIRNAME}/src/${BINDINGS_DIRNAME}"
if ! "${GENERATOR_SCRIPT_FILENAME}" "${input_dirpath}" "${typescript_output_dirpath}" typescript; then
    echo "Error: An error occurred generating TypeScript bindings in directory '${typescript_output_dirpath}'" >&2
    exit 1
fi
echo "Successfully generated TypeScript bindings in directory '${typescript_output_dirpath}'"
