#!/bin/bash

set -e

# Mac
mkdir -p ./packages/npm-mcptest-darwin-x64/bin
cp dist/mcptest_darwin_amd64_v1/mcptest ./packages/npm-mcptest-darwin-x64/bin/mcptest
chmod +x ./packages/npm-mcptest-darwin-x64/bin/mcptest
mkdir -p ./packages/npm-mcptest-darwin-arm64/bin
cp dist/mcptest_darwin_arm64_v8.0/mcptest ./packages/npm-mcptest-darwin-arm64/bin/mcptest
chmod +x ./packages/npm-mcptest-darwin-arm64/bin/mcptest

# Linux
mkdir -p ./packages/npm-mcptest-linux-x64/bin
cp dist/mcptest_linux_amd64_v1/mcptest ./packages/npm-mcptest-linux-x64/bin/mcptest
chmod +x ./packages/npm-mcptest-linux-x64/bin/mcptest
mkdir -p ./packages/npm-mcptest-linux-arm64/bin
cp dist/mcptest_linux_arm64_v8.0/mcptest ./packages/npm-mcptest-linux-arm64/bin/mcptest
chmod +x ./packages/npm-mcptest-linux-arm64/bin/mcptest

# Windows
mkdir -p ./packages/npm-mcptest-win32-x64/bin
cp dist/mcptest_windows_amd64_v1/mcptest.exe ./packages/npm-mcptest-win32-x64/bin/mcptest.exe
mkdir -p ./packages/npm-mcptest-win32-arm64/bin
cp dist/mcptest_windows_arm64_v8.0/mcptest.exe ./packages/npm-mcptest-win32-arm64/bin/mcptest.exe

cd packages/npm-mcptest-darwin-x64
npm publish --access public

cd ../npm-mcptest-darwin-arm64
npm publish --access public

cd ../npm-mcptest-linux-x64
npm publish --access public

cd ../npm-mcptest-linux-arm64
npm publish --access public

cd ../npm-mcptest-win32-x64
npm publish --access public

cd ../npm-mcptest-win32-arm64
npm publish --access public

cd ../npm-mcptest
npm publish --access public

cd -