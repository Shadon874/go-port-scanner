---

# PortScanner

A tiny TCP port scanner in Go. Scan single or multiple hosts for open ports.

---

## Usage

Build:

```bash
go build -o portscanner main.go
```

Run:

```bash
./portscanner -t localhost -r 1-1024
./portscanner -t localhost,example.com -r 20-80
./portscanner -h   # show help
```

---

## Flags

| Flag             | Description                     |
| ---------------- | ------------------------------- |
| `-t`, `--target` | Target host(s), comma-separated |
| `-r`, `--range`  | Port range (start-end)          |
| `-h`, `--help`   | Show help                       |

---

## Notes

* Only TCP ports are scanned.
* Closed ports are skipped.
* Fun project to practice Go and networking.

---
