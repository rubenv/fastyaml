# fastyaml

> Super fast YAML parser generator for Go

[![Build Status](https://github.com/rubenv/fastyaml/workflows/Test/badge.svg)](https://github.com/rubenv/fastyaml/actions) [![GoDoc](https://godoc.org/github.com/rubenv/fastyaml?status.png)](https://godoc.org/github.com/rubenv/fastyaml)

YAML is incredibly complex to parse and thus not very fast. But sometimes you
want it fast and you know that your input will only contain a small subset of
all the YAML features. That's where fastyaml comes in: it generates code for a
super fast parser for a very limited subset of the YAML specification.

Features:

* Did I mention fast?
* Generates proper Go code
* Almost no error handling, only use this on input that you know is well-formed
* Uses the same tag syntax as [go-yaml](https://github.com/go-yaml/yaml), so can be dropped in where possible with minimal changes

## Work in progress

This project is work in progress and not ready for use yet.

## License

This library is distributed under the [MIT](LICENSE) license.
