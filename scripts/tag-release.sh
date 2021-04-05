#!/usr/bin/env bash

# This script creates a new annotated tag for release.
# 
# Usage: `hack/tag-release.sh`.

set -o errexit
set -o nounset
set -o pipefail
# set -o xtrace

# Set variables
GIT_REMOTE="${GIT_REMOTE:-origin}"

# Set usage instructions
usage() {
    echo "Usage: ${0} VERSION" 1>&2
    exit 1
}

# Validate input arguments
if [[ ${#} -ne 1 ]]; then
    usage
fi

# Pull all tags from origin
git fetch ${GIT_REMOTE} --tags

# Tag release
git tag \
    --sign \
    --annotate ${1} \
    --message ${1} \

# Push tags
git push ${GIT_REMOTE} ${1}
