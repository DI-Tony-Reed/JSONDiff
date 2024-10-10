[![Tests](https://github.com/DI-Tony-Reed/JSONDiff/actions/workflows/tests.yaml/badge.svg)](https://github.com/DI-Tony-Reed/JSONDiff/actions/workflows/tests.yaml)

# What is this
This tool simply accepts two JSON files and returns the difference between them. It utilizes https://github.com/go-test/deep for the comparison. 

# How do I use this?
Locally, you can clone this repository, build it via the Makefile, and run it by feeding it two JSON files:
```bash
make build
./bin/jsonDiff-darwin json1.json json2.json
```
Make sure to use the appropriate binary for your OS. The above example is assuming you are on a Mac.

## Additional CLI flags
You can also use the `--byteskip` to skip the comparison if the second file has fewer bytes than the first.
```bash
./bin/jsonDiff-darwin json1.json json2.json --byteskip
```


# Viewing test coverage
```bash
make tests-coverage && open coverage.html
```
