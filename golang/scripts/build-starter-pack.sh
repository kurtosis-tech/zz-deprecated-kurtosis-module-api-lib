#!/usr/bin/env bash
# ^^^^^^^^^^^^^^^^^ this is the most platform-agnostic way to guarantee this script runs with Bash
# 2021-07-08 WATERMARK, DO NOT REMOVE - This script was generated from the Kurtosis Bash script template

set -euo pipefail   # Bash "strict mode"
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root_dirpath="$(dirname "${script_dirpath}")"


# ==================================================================================================
#                                             Constants
# ==================================================================================================
KURTOSIS_DOCKERHUB_ORG="kurtosistech"
IMAGE_NAME="leandro-lambda-starter-pack"


# ==================================================================================================
#                                             Main Logic
# ==================================================================================================
docker build -t "${KURTOSIS_DOCKERHUB_ORG}/${IMAGE_NAME}" --progress=plain -f "${root_dirpath}/Dockerfile" "${root_dirpath}"
