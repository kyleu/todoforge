#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Uses `tools/desktop` to build a desktop application, intended to be run from Docker

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/.."

echo "updating dependencies..."
go mod tidy
