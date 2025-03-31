#!/bin/bash

set -e

# # Mac
mkdir -p ./packages/npm-mcp-autotest-darwin-x64/bin
cp dist/mcp-autotest_darwin_amd64_v1/mcp-autotest ./packages/npm-mcp-autotest-darwin-x64/bin/mcp-autotest
chmod +x ./packages/npm-mcp-autotest-darwin-x64/bin/mcp-autotest
mkdir -p ./packages/npm-mcp-autotest-darwin-arm64/bin
cp dist/mcp-autotest_darwin_arm64_v8.0/mcp-autotest ./packages/npm-mcp-autotest-darwin-arm64/bin/mcp-autotest
chmod +x ./packages/npm-mcp-autotest-darwin-arm64/bin/mcp-autotest

# # Linux
mkdir -p ./packages/npm-mcp-autotest-linux-x64/bin
cp dist/mcp-autotest_linux_amd64_v1/mcp-autotest ./packages/npm-mcp-autotest-linux-x64/bin/mcp-autotest
chmod +x ./packages/npm-mcp-autotest-linux-x64/bin/mcp-autotest
mkdir -p ./packages/npm-mcp-autotest-linux-arm64/bin
cp dist/mcp-autotest_linux_arm64_v8.0/mcp-autotest ./packages/npm-mcp-autotest-linux-arm64/bin/mcp-autotest
chmod +x ./packages/npm-mcp-autotest-linux-arm64/bin/mcp-autotest

# # Windows
mkdir -p ./packages/npm-mcp-autotest-win32-x64/bin
cp dist/mcp-autotest_windows_amd64_v1/mcp-autotest.exe ./packages/npm-mcp-autotest-win32-x64/bin/mcp-autotest.exe
mkdir -p ./packages/npm-mcp-autotest-win32-arm64/bin
cp dist/mcp-autotest_windows_arm64_v8.0/mcp-autotest.exe ./packages/npm-mcp-autotest-win32-arm64/bin/mcp-autotest.exe

cd packages/npm-mcp-autotest-darwin-x64
npm publish --access public

cd ../npm-mcp-autotest-darwin-arm64
npm publish --access public

cd ../npm-mcp-autotest-linux-x64
npm publish --access public

cd ../npm-mcp-autotest-linux-arm64
npm publish --access public

cd ../npm-mcp-autotest-win32-x64
npm publish --access public

cd ../npm-mcp-autotest-win32-arm64
npm publish --access public

cd ../npm-mcp-autotest
npm publish --access public

cd -