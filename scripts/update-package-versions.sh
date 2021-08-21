#!/usr/bin/env bash
# 2021-07-08 WATERMARK, DO NOT REMOVE - This script was generated from the Kurtosis Bash script template

set -euo pipefail   # Bash "strict mode"
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root_dirpath="$(dirname "${script_dirpath}")"


# ==================================================================================================
#                                             Constants
# ==================================================================================================
SUPPORTED_LANGS_FILENAME="supported-languages.txt"
UPDATE_VERSION_IN_FILE_SCRIPT_FILENAME="update-version-in-file.sh"  # From devtools; expected to be on PATH
SKIP_KEYWORD="SKIP" # Certain languages don't need any package replacement; they will use this keyword for their package file and replacement pattern

# Per-language package filepaths that need version, RELATIVE TO THE LANG ROOT, that need version replacement
declare -A PACKAGE_FILE_RELATIVE_FILEPATHS


declare -A PACKAGE_VERSION_PRINTF_PATTERNS

# Golang -- empty because Go doesn't need any package version replacing
PACKAGE_FILE_RELATIVE_FILEPATHS["golang"]="${SKIP_KEYWORD}"
PACKAGE_VERSION_PRINTF_PATTERNS["golang"]="${SKIP_KEYWORD}"

# Typescript
PACKAGE_FILE_RELATIVE_FILEPATHS["typescript"]="package.json"
PACKAGE_VERSION_PRINTF_PATTERNS["typescript"]="\"version\": \"%s\""

# ==================================================================================================
#                                       Arg Parsing & Validation
# ==================================================================================================
show_helptext_and_exit() {
    echo "Usage: $(basename "${0}") new_version"
    echo ""
    echo "  new_version   The version of this repo that is about to released"
    echo ""
    exit 1  # Exit with an error so that if this is accidentally called by CI, the script will fail
}

new_version="${1:-}"

if [ -z "${new_version}" ]; then
    echo "Error: No new version provided" >&2
    show_helptext_and_exit
fi



# ==================================================================================================
#                                             Main Logic
# ==================================================================================================

echo "Updating versions in package files for all supported langauges..."
supported_langs_filepath="${root_dirpath}/${SUPPORTED_LANGS_FILENAME}"
for lang in $(cat "${supported_langs_filepath}"); do
    to_update_relative_filepath="${PACKAGE_FILE_RELATIVE_FILEPATHS["${lang}"]}"
    if [ -z "${to_update_relative_filepath}" ]; then
        echo "Error: No relative filepath to a package file that needs updating was found for language '${lang}'; this script needs to be updated with this information" >&2
        exit 1
    fi
    if [ "${to_update_relative_filepath}" == "${SKIP_KEYWORD}" ]; then
        continue
    fi

    to_update_abs_filepath="${root_dirpath}/${lang}/${to_update_relative_filepath}"

    replacement_pattern="${PACKAGE_VERSION_PRINTF_PATTERNS["${lang}"]}"
    if [ -z "${replacement_pattern}" ]; then
        echo "Error: No replacement pattern was found for language '${lang}'; this script needs to be updated with this information" >&2
        exit 1
    fi
    if [ "${replacement_pattern}" == "${SKIP_KEYWORD}" ]; then
        continue
    fi

    if ! "${UPDATE_VERSION_IN_FILE_SCRIPT_FILENAME}" "${to_update_abs_filepath}" "${replacement_pattern}" "${new_version}"; then
        echo "Error: An error occurred setting new version '${new_version}' in '${lang}' package file '${to_update_abs_filepath}' using pattern '${replacement_pattern}'" >&2
        exit 1
    fi
done
echo "Successfully updated the versions in the package files for all supported languages"
