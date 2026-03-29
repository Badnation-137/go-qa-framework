# 🚀 Go QA Framework

[![CI/CD](https://github.com/Badnation-137/go-qa-framework/actions/workflows/ci.yml/badge.svg)](https://github.com/Badnation-137/go-qa-framework/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.26.1-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Made with ❤️](https://img.shields.io/badge/Made%20with-❤️-red.svg)]()

> **Professional QA Framework built with Go**  
> *Demonstrating Senior SDET capabilities with automated testing, containers & CI/CD*

---

## 📋 Overview

This project demonstrates **enterprise-level quality assurance** including unit testing, integration testing with PostgreSQL, and automated CI/CD pipelines.

### What This Demonstrates

| Capability | Technology |
|------------|------------|
| 🧪 Unit Testing | testify |
| 🐳 Integration Testing | Testcontainers + PostgreSQL |
| 🔄 CI/CD Pipeline | GitHub Actions |
| 🔒 Security Scanning | gosec |
| 📊 Multi-Go Version Testing | Go 1.21, 1.22, 1.23 |
| 📝 Code Linting | golangci-lint |

### ✨ Key Features

- ✅ Table-driven unit tests with assertions
- ✅ PostgreSQL integration via Docker containers
- ✅ Auto spin-up/teardown of test databases
- ✅ Multi-Go version compatibility testing
- ✅ Automated security vulnerability scanning
- ✅ Code linting with industry standards
- ✅ One-command test execution

---

## 🔄 CI/CD Pipeline

Automated testing with **GitHub Actions**:

| Job | Description |
|-----|-------------|
| **Lint & Security** | Code quality & vulnerability scan |
| **Unit Tests** | Multi-version Go testing (1.21, 1.22, 1.23) |
| **Integration Tests** | PostgreSQL container testing |
| **Build Binary** | Production compilation |

### Pipeline Flow
git push
│
▼
🚀 GitHub Actions Trigger
│
├──► 🔍 Lint & Security
│
├──► 🧪 Unit Tests (Go 1.21, 1.22, 1.23)
│
├──► 🐳 Integration Tests (PostgreSQL)
│
└──► ✅ All Checks Passed
│
▼
📦 Build Binary

---

## 🐳 Integration Testing

> Using **Testcontainers for Go** to spin up real PostgreSQL containers

### How It Works

1. 🐳 Start PostgreSQL container
2. 🗄️ Create schema & tables
3. 🧪 Run test cases
4. 🧹 Auto-terminate container

### Test Cases (All Passing ✅)

| Test Case | Description |
|-----------|-------------|
| `TestCreateUser` | INSERT + verify |
| `TestGetUserByEmail` | SELECT query |
| `TestUpdateUser` | UPDATE + check |
| `TestDeleteUser` | DELETE + verify |
| `TestDuplicateEmailConstraint` | UNIQUE constraint enforcement |
| `TestBulkInsertPerformance` | 100 inserts performance test |
| `TestFindUsersByAgeRange` | WHERE clause filtering |

### Performance Benchmark

| Run Type | Time | Notes |
|----------|------|-------|
| First Run | ~272s | Downloading Docker image |
| Cached Run | ~15s ⚡ | Image already local |
| CI/CD Runner | ~59s | Fresh environment |

### Benefits

- 🔄 Isolated test environments
- 🧹 Auto-cleanup after tests
- 🚀 Fast with caching (~15s after first run)
- 🎯 Real database, not mocks

---

## 🛠️ Tech Stack

| Category | Technology |
|----------|------------|
| Language | Go 1.26.1 |
| Testing | testify, testing package |
| Containers | Docker + Testcontainers |
| Database | PostgreSQL 15 Alpine |
| CI/CD | GitHub Actions |
| Linting | golangci-lint |
| Security | gosec |
| Environment | WSL2 (Ubuntu) on Windows |

---

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker (for integration tests)

---

╔════════════════════════════════════════════════════════════╗
║                                                            ║
║   🐍  Thanks for checking out Go QA Framework!  🐍        ║
║                                                            ║
║        Built with ❤️  using Go & Testcontainers           ║
║                                                            ║
║   [⭐ Star]  [🔔 Watch]  [🍴 Fork]  [💡 Issues]          ║
║                                                            ║
╚════════════════════════════════════════════════════════════╝

<div align="center">
https://img.shields.io/github/stars/Badnation-137/go-qa-framework?style=social
https://img.shields.io/github/watchers/Badnation-137/go-qa-framework?style=social
https://img.shields.io/github/forks/Badnation-137/go-qa-framework?style=social

</div> ```