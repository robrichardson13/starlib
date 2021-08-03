<h1 align="center">Starlib - Starlark's Missing standard Library.</h1>

[![Qri](https://img.shields.io/badge/made%20by-qri-magenta.svg?style=flat-square)](https://qri.io) [![GoDoc](https://godoc.org/github.com/robrichardson13/starlib?status.svg)](http://godoc.org/github.com/robrichardson13/starlib) [![License](https://img.shields.io/github/license/qri-io/starlib.svg?style=flat-square)](./LICENSE) [![Codecov](https://img.shields.io/codecov/c/github/qri-io/starlib.svg?style=flat-square)](https://codecov.io/gh/qri-io/starlib) [![CI](https://img.shields.io/circleci/project/github/qri-io/starlib.svg?style=flat-square)](https://circleci.com/gh/qri-io/starlib)

<div align="center">
  <h3>
    <a href="https://qri.io">
      Website
    </a>
    <span> | </span>
    <a href="#packages">
      Packages
    </a>
    <span> | </span>
    <a href="https://github.com/robrichardson13/starlib/CONTRIBUTOR.md">
      Contribute
    </a>
    <span> | </span>
    <a href="https://github.com/robrichardson13/starlib/issues">
      Issues
    </a>
     <span> | </span>
    <a href="https://qri.io/docs/starlark/starlib">
      Docs
    </a>
  </h3>
</div>

<div align="center">
  <!-- Build Status -->
</div>

## Welcome

This is a community-driven project to bring a standard library to the starlark programming dialect. We here at Qri need a standard library, and we thought it might benefit others to structure this library in a reusable way. We are a little biased towards our needs, and will be shaping the library primarily toward's Qri's use case.

| Question                                     | Answer                                                                                                                                            |
| -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| "What's starlark?"                           | It's a python-like scripting language open-sourced by Google. [Here are the docs](https://docs.bazel.build/versions/master/skylark/language.html) |
| "What's the use-case for this?"              | [We're building it for Qri ('query')](https://qri.io)                                                                                             |
| "I want to play with starlib outside of Qri" | [Checkout the starlark playground](https://github.com/qri-io/skypg)                                                                               |
| "I have a question"                          | [Create an issue](https://github.com/robrichardson13/starlib/issues)                                                                                       |
| "I found a bug"                              | [Create an issue](https://github.com/robrichardson13/starlib/issues)                                                                                       |
| "I would like to propose a new package"      | You should think about [creating an RFC](https://github.com/qri-io/rfcs)                                                                          |

## Packages

The following is a list of the packages currently in the standard library

| Package                                                                            | Go Docs                                                                                                                                                           | Description                                     |
| ---------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------- |
| [`bsoup`](https://github.com/robrichardson13/starlib/tree/master/bsoup)                     | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/bsoup?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/bsoup)                     | a beautiful-soup-like API for working with HTML |
| [`encoding/base64`](https://github.com/robrichardson13/starlib/tree/master/encoding/base64) | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/encoding/base64?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/encoding/base64) | base64 de/serialization                         |
| [`encoding/csv`](https://github.com/robrichardson13/starlib/tree/master/encoding/csv)       | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/encoding/csv?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/encoding/csv)       | csv de/serialization                            |
| [`encoding/json`](https://github.com/robrichardson13/starlib/tree/master/encoding/json)     | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/encoding/json?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/encoding/json)     | json de/serialization                           |
| [`geo`](https://github.com/robrichardson13/starlib/tree/master/geo)                         | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/geo?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/geo)                         | 2d geographic operations                        |
| [`hash`](https://github.com/robrichardson13/starlib/tree/master/hash)                       | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/hash?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/hash)                       | hash operations (md5/sha1/sha256)               |
| [`html`](https://github.com/robrichardson13/starlib/tree/master/html)                       | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/html?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/html)                       | html text processing                            |
| [`http`](https://github.com/robrichardson13/starlib/tree/master/http)                       | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/http?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/http)                       | http client operations                          |
| [`math`](https://github.com/robrichardson13/starlib/tree/master/math)                       | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/math?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/math)                       | mathematical functions & values                 |
| [`re`](https://github.com/robrichardson13/starlib/tree/master/re)                           | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/re?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/re)                           | regular expressions                             |
| [`time`](https://github.com/robrichardson13/starlib/tree/master/time)                       | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/time?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/time)                       | time operations                                 |
| [`xlsx`](https://github.com/robrichardson13/starlib/tree/master/xlsx)                       | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/xlsx?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/xlsx)                       | xlsx file format reading                        |
| [`zipfile`](https://github.com/robrichardson13/starlib/tree/master/zipfile)                 | <img width=190/>[![Go Docs](https://godoc.org/github.com/robrichardson13/starlib/zipfile?status.svg)](https://godoc.org/github.com/robrichardson13/starlib/zipfile)                 | support for zip archives                        |

###### This documentation has been adapted from the [Cycle.js](https://github.com/cyclejs/cyclejs) documentation.
