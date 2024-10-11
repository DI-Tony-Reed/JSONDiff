[![Tests](https://github.com/DI-Tony-Reed/JSONDiff/actions/workflows/tests.yaml/badge.svg)](https://github.com/DI-Tony-Reed/JSONDiff/actions/workflows/tests.yaml)

# What is this
This tool accepts two JSON Snyk scans and returns the difference between them. It utilizes a slightly modified version of https://github.com/hezro/snyk-code-pr-diff for the comparison. 

# How do I use this?
Locally, you can clone this repository, build it via the Makefile, and run it by feeding it two JSON Snyk scan files:
```bash
make build
./bin/jsonDiff-darwin json1.json json2.json
```
Make sure to use the appropriate binary for your OS. The above example is assuming you are on a Mac.

# Viewing test coverage
```bash
make tests-coverage && open coverage.html
```
