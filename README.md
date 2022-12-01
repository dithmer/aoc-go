# How to run?

## Prerequisites
- Environment variable `AOC_SESSION` (You can get it by logging in to the website and looking at the cookies.)
- Go >= 1.19

## Running
For running the tests completely and get output:
```bash
go test -v ./...
```

If you just want to see the solutions, I tend to use:
```bash
go test -v -json ./... | grep -i solution
```
