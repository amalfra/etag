etag
========
[![GitHub release](https://img.shields.io/github/release/amalfra/etag.svg)](https://github.com/amalfra/etag/releases)
![Build Status](https://github.com/amalfra/etag/actions/workflows/test.yml/badge.svg?branch=main)
[![GoDoc](https://godoc.org/github.com/amalfra/etag/v4?status.svg)](https://godoc.org/github.com/amalfra/etag/v4)
[![Go Report Card](https://goreportcard.com/badge/github.com/amalfra/etag/v4)](https://goreportcard.com/report/github.com/amalfra/etag/v4)
[![Coverage Status](https://coveralls.io/repos/github/amalfra/etag/badge.svg?branch=main)](https://coveralls.io/github/amalfra/etag?branch=main)

A go package to create HTTP ETags (as defined in RFC 7232) for use in HTTP responses.

## Installation
You can download the package using
```sh
go get github.com/amalfra/etag/v4
```

## Usage
Next, import the package
``` go
import (
  "github.com/amalfra/etag/v4"
)
```
You can now call the ```generate``` function to create an ETag. The second parameter is a boolean which specifies whether the generated ETag must be weak or not.
```go
etag.Generate(body, false)
```
to generate a weak Etag
```go
etag.Generate(body, true)
```

## Development

Questions, problems or suggestions? Please post them on the [issue tracker](https://github.com/amalfra/etag/issues).

You can contribute changes by forking the project and submitting a pull request. You can ensure the tests are passing by running ```make test```. Feel free to contribute :heart_eyes:

## UNDER MIT LICENSE

The MIT License (MIT)

Copyright (c) 2017 Amal Francis

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
