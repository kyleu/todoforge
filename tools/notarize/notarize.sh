#!/bin/bash

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

if [ "$PUBLISH_TEST" != "true" ]
then
  xcrun notarytool submit --apple-id "$APPLE_EMAIL" --team-id "$APPLE_TEAM_ID" --password "$APPLE_PASSWORD" ./build/dist/todoforge_0.2.19_darwin_amd64_desktop.dmg
  xcrun notarytool submit --apple-id "$APPLE_EMAIL" --team-id "$APPLE_TEAM_ID" --password "$APPLE_PASSWORD" ./build/dist/todoforge_0.2.19_darwin_arm64_desktop.dmg
  xcrun notarytool submit --apple-id "$APPLE_EMAIL" --team-id "$APPLE_TEAM_ID" --password "$APPLE_PASSWORD" ./build/dist/todoforge_0.2.19_darwin_all_desktop.dmg
fi
