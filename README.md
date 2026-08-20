# Data Profiler

![CI](https://github.com/Qyroxen/Data-Profiler/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Data-Profiler/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Data-Profiler?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Data-Profiler)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Data-Profiler)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Data-Profiler?style=social)](https://github.com/Qyroxen/Data-Profiler/stargazers)

## What is it?

Data Profiler is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Data-Profiler.git
cd Data-Profiler
go build -o dataprofiler .

# Run
./dataprofiler --help
```

## CLI Usage

```bash
# Basic usage
./dataprofiler

# With flags
./dataprofiler --verbose --output json

# Get help
./dataprofiler --help
```

## Examples

```bash
# Example 1
./dataprofiler example1

# Example 2
./dataprofiler example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o dataprofiler .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Data-Profiler/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Data-Profiler?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Data-Profiler/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Data-Profiler?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Data-Profiler/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Data-Profiler" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Data-Profiler/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Data-Profiler" alt="Pull Requests">
  </a>
</p>
