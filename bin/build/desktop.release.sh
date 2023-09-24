#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Meant to be run as part of the release process, builds desktop apps

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

docker build -f tools/desktop/Dockerfile.desktop -t todoforge .

rm -rf tmp/release
mkdir -p tmp/release

cd "tmp/release"

id=$(docker create todoforge)
docker cp $id:/dist - > ./desktop.tar
docker rm -v $id
tar -xvf "desktop.tar"
rm "desktop.tar"

mv dist/darwin_amd64/todoforge "todoforge.darwin"
mv dist/darwin_arm64/todoforge "todoforge.darwin.arm64"
mv dist/linux_amd64/todoforge "todoforge"
mv dist/windows_amd64/todoforge "todoforge.exe"
rm -rf "dist"

# darwin amd64
cp -R "../../tools/desktop/template" .

mkdir -p "./TODO Forge.app/Contents/Resources"
mkdir -p "./TODO Forge.app/Contents/MacOS"

cp -R "./template/darwin/Info.plist" "./TODO Forge.app/Contents/Info.plist"
cp -R "./template/darwin/icons.icns" "./TODO Forge.app/Contents/Resources/icons.icns"

cp "todoforge.darwin" "./TODO Forge.app/Contents/MacOS/todoforge"

echo "signing amd64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./TODO Forge.app/Contents/MacOS/todoforge"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./TODO Forge.app"

cp "./template/darwin/appdmg.config.json" "./appdmg.config.json"

echo "building macOS amd64 DMG..."
appdmg "appdmg.config.json" "./todoforge_${TGT}_darwin_amd64_desktop.dmg"
zip -r "todoforge_${TGT}_darwin_amd64_desktop.zip" "./TODO Forge.app"

# darwin arm64
cp "todoforge.darwin.arm64" "./TODO Forge.app/Contents/MacOS/todoforge"

echo "signing arm64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./TODO Forge.app/Contents/MacOS/todoforge"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./TODO Forge.app"

echo "building macOS arm64 DMG..."
appdmg "appdmg.config.json" "./todoforge_${TGT}_darwin_arm64_desktop.dmg"
zip -r "todoforge_${TGT}_darwin_arm64_desktop.zip" "./TODO Forge.app"

# macOS universal
rm "./TODO Forge.app/Contents/MacOS/todoforge"
lipo -create -output "./TODO Forge.app/Contents/MacOS/todoforge" todoforge.darwin todoforge.darwin.arm64

echo "signing universal desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./TODO Forge.app/Contents/MacOS/todoforge"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./TODO Forge.app"

echo "building macOS universal DMG..."
appdmg "appdmg.config.json" "./todoforge_${TGT}_darwin_all_desktop.dmg"
zip -r "todoforge_${TGT}_darwin_all_desktop.zip" "./TODO Forge.app"

# linux
echo "building Linux zip..."
zip "todoforge_${TGT}_linux_amd64_desktop.zip" "./todoforge"

#windows
echo "building Windows zip..."
curl -L -o webview.dll https://github.com/webview/webview/raw/master/dll/x64/webview.dll
curl -L -o WebView2Loader.dll https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll
zip "todoforge_${TGT}_windows_amd64_desktop.zip" "./todoforge.exe" "./webview.dll" "./WebView2Loader.dll"

mkdir -p "../../build/dist"
mv "./todoforge_${TGT}_darwin_amd64_desktop.dmg" "../../build/dist"
mv "./todoforge_${TGT}_darwin_amd64_desktop.zip" "../../build/dist"
mv "./todoforge_${TGT}_darwin_arm64_desktop.dmg" "../../build/dist"
mv "./todoforge_${TGT}_darwin_arm64_desktop.zip" "../../build/dist"
mv "./todoforge_${TGT}_darwin_all_desktop.dmg" "../../build/dist"
mv "./todoforge_${TGT}_darwin_all_desktop.zip" "../../build/dist"
mv "./todoforge_${TGT}_linux_amd64_desktop.zip" "../../build/dist"
mv "./todoforge_${TGT}_windows_amd64_desktop.zip" "../../build/dist"
