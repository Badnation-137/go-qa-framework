# 🚀 Go QA Framework


[![CI/CD](https://github.com/Badnation-137/go-qa-framework/actions/workflows/ci.yml/badge.svg)](https://github.com/Badnation-137/go-qa-framework/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.26.1-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Made with ❤️](https://img.shields.io/badge/Made%20with-❤️-red.svg)]()

> 🚀 **Professional QA Framework built with Go**  
> *Demonstrating Senior SDET capabilities with automated testing, containers & CI/CD*

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

---

## 📋 Overview

This project demonstrates **enterprise-level quality assurance** including unit testing, integration testing with PostgreSQL, and automated CI/CD pipelines.

### What This Demonstrates:
- 🧪 Unit Testing with testify
- 🐳 Integration Testing with Testcontainers + PostgreSQL  
- 🔄 CI/CD Pipeline with GitHub Actions
- 🔒 Security Scanning with gosec
- 📊 Multi-Go Version Testing (1.21, 1.22, 1.23)

┌─────────────────────────────────────────┐
│ 🎯 What This Project Demonstrates: │
├─────────────────────────────────────────┤
│ 🧪 Unit Testing │ testify │
│ 🐳 Integration Test │ Testcontainers │
│ 🔄 CI/CD Pipeline │ GitHub Actions │
│ 🔒 Security Scan │ gosec │
│ 📊 Code Quality │ golangci-lint │
│ 🔢 Multi-Version │ Go 1.21-1.23 │
└─────────────────────────────────────────┘

### ✨ Key Features
✅ Table-driven unit tests with assertions
✅ PostgreSQL integration via Docker containers
✅ Auto spin-up/teardown of test databases
✅ Multi-Go version compatibility testing
✅ Automated security vulnerability scanning
✅ Code linting with industry standards
✅ One-command test execution

---

### **Part 3: CI/CD Section**
```markdown
---
                ┌─────────────────┐
                │   git push      │
                └────────┬────────┘
                         ▼
        ┌─────────────────────────────┐
        │  🚀 GitHub Actions Trigger  │
        └─────────────────────────────┘
                         │
    ┌────────────────────┼────────────────────┐
    ▼                    ▼                    ▼
    ┌───────────────┐ ┌───────────────┐ ┌───────────────┐
│ 🔍 Lint & │ │ 🧪 Unit Tests │ │ 🐳 Integration│
│ Security │ │ Go 1.21/22/23 │ │ PostgreSQL │
└───────┬───────┘ └───────┬───────┘ └───────┬───────┘
│ │ │
▼ ▼ ▼
┌─────────────────────────────────────────────┐
│ ✅ All Checks Passed │
│ ▼ │
│ ┌─────────────────┐ │
│ │ 📦 Build Binary │ │
│ └─────────────────┘ │
└─────────────────────────────────────────────┘

## 🔄 CI/CD Pipeline

Automated testing with **GitHub Actions**:

| Job | Description |
|-----|-------------|
| **Lint & Security** | Code quality & vulnerability scan |
| **Unit Tests** | Multi-version Go testing |
| **Integration Tests** | PostgreSQL container testing |
| **Build Binary** | Production compilation |


### Pipeline Status

[![CI/CD Pipeline](https://github.com/Badnation-137/go-qa-framework/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/Badnation-137/go-qa-framework/actions)

---

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

## 🐳 Integration Testing

> Using **Testcontainers for Go** to spin up real PostgreSQL containers 🐘

### How It Works
┌─────────────────────────────────────┐
│ 1. 🐳 Start PostgreSQL container │
│ 2. 🗄️ Create schema & tables │
│ 3. 🧪 Run test cases │
│ 4. 🧹 Auto-terminate container │
└─────────────────────────────────────┘

### ✅ Test Cases (All Passing)
┌────────────────────────────────┬────────────────┐
│ Test Case │ Description │
├────────────────────────────────┼────────────────┤
│ TestCreateUser │ INSERT + verify│
│ TestGetUserByEmail │ SELECT query │
│ TestUpdateUser │ UPDATE + check │
│ TestDeleteUser │ DELETE + verify│
│ TestDuplicateEmailConstraint │ UNIQUE enforce │
│ TestBulkInsertPerformance │ 100 inserts │
│ TestFindUsersByAgeRange │ WHERE clause │
└────────────────────────────────┴────────────────┘

### Test Cases (All Passing ✅):
- TestCreateUser
- TestGetUserByEmail
- TestUpdateUser
- TestDeleteUser
- TestDuplicateEmailConstraint
- TestBulkInsertPerformance
- TestFindUsersByAgeRange

### 📊 Performance Benchmark
┌─────────────────┬─────────────┬────────────────────┐
│ Run Type │ Time │ Notes │
├─────────────────┼─────────────┼────────────────────┤
│ First Run │ ~272s │ Download image │
│ Cached Run │ ~15s ⚡ │ Image already local│
│ CI/CD Runner │ ~59s │ Fresh environment │
└─────────────────┴─────────────┴────────────────────┘

### Benefits:
- 🔄 Isolated test environments
- 🧹 Auto-cleanup after tests  
- 🚀 Fast with caching (~15s after first run)
- 🎯 Real database, not mocks

---

## 🛠️ Tech Stack
┌────────────────┬────────────────────────┐
│ Category │ Technology │
├────────────────┼────────────────────────┤
│ Language │ Go 1.26.1 🐹 │
│ Testing │ testify, testing pkg │
│ Containers │ Docker + Testcontainers│
│ Database │ PostgreSQL 15 Alpine │
│ CI/CD │ GitHub Actions │
│ Linting │ golangci-lint │
│ Security │ gosec │
│ Environment │ WSL2 (Ubuntu) on Win │
└────────────────┴────────────────────────┘

- **Language:** Go 1.26.1
- **Testing:** testify, Go testing package
- **Database:** PostgreSQL 15 (Testcontainers)
- **CI/CD:** GitHub Actions
- **Tools:** golangci-lint, gosec

---

## 🚀 Quick Start

git clone https://github.com/Badnation-137/go-qa-framework.git
cd go-qa-framework
go mod download
go test -v ./... -timeout=10m

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

🔗 View CI/CD Pipeline
📁 Explore Code
Made with 🐍 by Badnation-137
</div>
```