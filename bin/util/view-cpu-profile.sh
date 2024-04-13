#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Starts a pprof server using the (previously-recorded) CPU profile at ./tmp/cpu.pprof

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

echo "=== launching profiler for cpu.pprof ==="
go tool pprof -http=":8000" ./build/debug/todoforge ./tmp/cpu.pprof

