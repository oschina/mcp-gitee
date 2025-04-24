#!/usr/bin/env node

const childProcess = require('child_process');

const BINARY_MAP = {
  darwin_x64: {name: 'mcp-gitee-darwin-amd64', suffix: ''},
  darwin_arm64: {name: 'mcp-gitee-darwin-arm64', suffix: ''},
  linux_x64: {name: 'mcp-gitee-linux-amd64', suffix: ''},
  linux_arm64: {name: 'mcp-gitee-linux-arm64', suffix: ''},
  win32_x64: {name: 'mcp-gitee-windows-amd64', suffix: '.exe'},
  win32_arm64: {name: 'mcp-gitee-windows-arm64', suffix: '.exe'},
};

// Resolving will fail if the optionalDependency was not installed or the platform/arch is not supported
const resolveBinaryPath = () => {
  try {
    const binary = BINARY_MAP[`${process.platform}_${process.arch}`];
    return require.resolve(`@gitee/${binary.name}/bin/${binary.name}${binary.suffix}`);
  } catch (e) {
    console.error(`Could not resolve binary path for platform/arch: ${process.platform}/${process.arch}`);
    process.exit(1);
  }
};

childProcess.execFileSync(resolveBinaryPath(), process.argv.slice(2), {
  stdio: 'inherit',
});
