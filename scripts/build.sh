#!/usr/bin/env bash
# 2021-07-08 WATERMARK, DO NOT REMOVE - This script was generated from the Kurtosis Bash script template

set -euo pipefail   # Bash "strict mode"
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root_dirpath="$(dirname "${script_dirpath}")"


# ==================================================================================================
#                                             Constants
# ==================================================================================================
SUPPORTED_LANGS_FILENAME="supported-languages.txt"
LANG_SCRIPTS_DIRNAME="scripts"
LANG_BUILDSCRIPT_FILENAME="build.sh"



# ==================================================================================================
#                                             Main Logic
# ==================================================================================================
supported_langs_filepath="${root_dirpath}/${SUPPORTED_LANGS_FILENAME}"
for lang in $(cat "${supported_langs_filepath}"); do
    lang_buildscript_filepath="${root_dirpath}/${lang}/${LANG_SCRIPTS_DIRNAME}/${LANG_BUILDSCRIPT_FILENAME}"
    if ! "${lang_buildscript_filepath}"; then
        echo "Error: Build of lang '${lang}' failed" >&2
        exit 1
    fi
done
