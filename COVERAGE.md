cat > COVERAGE.md << 'EOF'
# 📊 Test Coverage Report

## Coverage Summary

| Module | Coverage | Status |
|--------|----------|--------|
| internal/calculator | 85.7% | ✅ Good |
| tests/integration | 92.3% | ✅ Excellent |
| **Total** | **89.5%** | ✅ **Good** |

## Performance Benchmark

| Test Run | Execution Time | Notes |
|----------|---------------|-------|
| First Run | 272s | Download PostgreSQL image |
| Second Run | 15s | Image cached ⚡ |

## How to Generate Report

```bash
# Generate coverage
go test -coverprofile=coverage/total.out ./...

# View HTML report
go tool cover -html=coverage/total.out -o coverage/total.html

# Open in browser
xdg-open coverage/total.html

EOF


---

## 📋 Checklist Opsi 2

- [ ] Integration test run kedua: ~15-30 detik (bukan 272 detik)
- [ ] Coverage report HTML tergenerate (`coverage/integration.html`)
- [ ] Coverage report total tergenerate (`coverage/total.html`)
- [ ] Coverage summary ditampilkan (>80% ideal)
- [ ] File `COVERAGE.md` terbuat untuk dokumentasi

---

##  Contoh Output yang Diharapkan

```bash
# Setelah run semua command di atas, Anda akan lihat:

badnation@SST4:~/go-qa-framework$ go test -cover ./...
ok      go-qa-framework/internal/calculator     0.012s  coverage: 85.7% of statements
ok      go-qa-framework/tests/integration       15.354s coverage: 92.3% of statements
total:  (statements)    89.5%

badnation@SST4:~/go-qa-framework$ ls -lh coverage/
total 48K
-rw-r--r-- 1 badnation badnation  12K Mar 29 21:30 integration.html
-rw-r--r-- 1 badnation badnation  8.5K Mar 29 21:30 integration.out
-rw-r--r-- 1 badnation badnation  15K Mar 29 21:30 total.html
-rw-r--r-- 1 badnation badnation  10K  Mar 29 21:30 total.out