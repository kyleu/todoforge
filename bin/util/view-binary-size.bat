@ECHO OFF
rem Content managed by Project Forge, see [projectforge.md] for details.

rem Visualizes space usage for the release binary

cd %~dp0\..\..

@ECHO ON
make build-release
go tool nm -size build\release\todoforge
