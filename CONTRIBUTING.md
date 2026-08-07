# Contributing to Kavrok

First off, thank you for considering contributing to Kavrok!

Kavrok is an open-source engineering intelligence platform focused on helping developers and platform engineers diagnose, understand, and automate Kubernetes operations.

Every contribution, whether it's code, documentation, testing, or ideas, is appreciated.

---

# Ways to Contribute

There are many ways to contribute:

- Report bugs
- Suggest new features
- Improve documentation
- Improve error messages
- Add diagnostic rules
- Improve performance
- Add unit tests
- Improve developer experience
- Review pull requests

---

# Before You Start

Please:

- Search existing issues before opening a new one.
- Discuss significant changes before starting implementation.
- Keep pull requests focused on a single feature or fix.

---

# Development Setup

## Prerequisites

- Go 1.26 or later
- Git
- Make
- Docker (optional)
- Kubernetes cluster (Kind, Minikube, or any CNCF-compliant cluster)

---

## Clone the Repository

```bash
git clone https://github.com/mohammedimrankasab/kavrok.git

cd kavrok
```

---

## Install Dependencies

```bash
go mod download
```

---

## Verify Installation

```bash
make test

make build

make run
```

---

# Project Structure

```text
cmd/
internal/
pkg/
docs/
examples/
scripts/
```

## Directory Guidelines

### cmd/

CLI entrypoints.

### internal/

Private application packages.

### pkg/

Reusable public packages.

Only place code here if it is intended to be imported by external applications.

---

# Coding Standards

Kavrok follows standard Go conventions.

Please:

- Keep functions focused.
- Prefer composition over inheritance.
- Prefer interfaces where appropriate.
- Write idiomatic Go.
- Keep packages cohesive.
- Avoid unnecessary abstractions.

---

# Formatting

Run:

```bash
make fmt
```

---

# Linting

Run:

```bash
make lint
```

---

# Static Analysis

Run:

```bash
make vet
```

---

# Testing

Run:

```bash
make test
```

Before opening a pull request ensure:

- All tests pass.
- New features include tests.
- Existing tests remain green.

---

# Benchmarks

Performance-sensitive changes should include benchmark results.

Run:

```bash
make benchmark
```

---

# Commit Messages

Kavrok follows Conventional Commits.

Examples:

```text
feat: add OOMKilled analyzer

fix: prevent nil pointer in renderer

docs: update installation guide

test: improve analyzer coverage

refactor: simplify rule evaluation

perf: optimize event processing

chore: update dependencies
```

---

# Branch Naming

Examples:

```text
feature/doctor-command

feature/oom-rule

fix/pod-rendering

docs/readme

refactor/rule-engine
```

---

# Pull Requests

A good pull request should:

- Solve one problem.
- Include tests where appropriate.
- Update documentation.
- Keep commits clean.
- Pass CI.

---

# Pull Request Checklist

Before submitting:

- [ ] Code builds successfully
- [ ] Tests pass
- [ ] Lint passes
- [ ] Documentation updated
- [ ] CHANGELOG updated (if applicable)
- [ ] No unnecessary files included

---

# Reporting Bugs

When reporting bugs, please include:

- Kavrok version
- Go version
- Kubernetes version
- Operating system
- Steps to reproduce
- Expected behavior
- Actual behavior
- Relevant logs

---

# Suggesting Features

Feature requests should describe:

- The problem
- Why it matters
- Proposed solution
- Alternatives considered
- Additional context

---

# Design Principles

Kavrok is built around several core principles:

- Engineering-first
- Explain before exposing raw data
- Deterministic analysis over guesswork
- Small, composable packages
- Modular architecture
- Extensible rule engine
- Production-ready code quality

Please keep these principles in mind when contributing.

---

# Code Review

Pull requests are reviewed for:

- Correctness
- Maintainability
- Readability
- Performance
- Testing
- Documentation

Feedback is intended to improve the project and should always remain respectful and constructive.

---

# Community

Be respectful.

Be welcoming.

Help others learn.

Constructive discussions are encouraged.

---

# Questions

If you have questions, feel free to open a GitHub Discussion or start an issue.

---

Thank you for helping make Kavrok better!