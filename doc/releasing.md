<!--- Content managed by Project Forge, see [projectforge.md] for details. -->
# Releasing

TODO Forge uses `goreleaser` to create build artifacts.

You can release to GitHub using `./bin/build/release.sh`, or test the release by running `./bin/build/release-test.sh`.

Your releases are available at https://github.com/kyleu/todoforge/releases

### Checksums

All release binaries are checksummed, available in `checksums.txt` in the root of the release

### Docker Images

Multiple Docker images will be created. The main image is `ghcr.io/kyleu/todoforge/x.x.x`, and a debug image is provided at `ghcr.io/kyleu/todoforge/x.x.x-debug` that includes `delve` for debugging

### Homebrew

Packages for macOS and Linux will be pushed to Homebrew at `kyleu/homebrew-kyleu`

### NFPMS

The build will produce `apk`, `deb`, and `rpm` packages for each supported Linux architecture

### BOM

The build will create a bill of materials for each binary

### Notarization

Release binaries for macOS and iOS are notarized using Apple Notarization services

### Signing

Release binaries and the checksum file are signed using `gpg`

### Snapcraft

The build will produce `snap` packages for each supported Linux architecture

### Source Code

The source code will be bundled in the release, available as `todoforge_x.x.x_source.zip`

### Universal Binaries

A universal macOS app will be created, containing the complete app for x86-64 and ARM

### Desktop Build

A standalone desktop application, bundling the server and a web view, will be built for Linux, macOS, and Windows

### Mobile Build

A standalone mobile app, bundling the server and a web view, will be built for Android and iOS

