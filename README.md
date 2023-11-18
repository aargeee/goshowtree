# GoShowTree
> CLI tool to view your directory.

[![Go Version](https://img.shields.io/badge/Go-v1.21-blue.svg)](https://golang.org/doc/go1.21) <!-- TODO : get actual go version -->
[![Platform](https://img.shields.io/badge/Platform-Linux-green.svg?style=flat-square)](https://travis-ci.org/username/repository)

GoShowTree allows you to specify any directory and recursively view all the contents inside the directory in form of a tree.

## Installation

OS X & Linux:

```sh
git clone github.com/Rahul-NITD/goshowtree
cd goshowtree
sh build.sh
cd bin
```

Windows:

![Build Status](https://img.shields.io/badge/Build-Not%20Available-red.svg?style=flat-square)

## Usage example

<!-- TODO : usage example to be done in VM -->
```sh
# go to the bin folder
./goshowtree <path>
```
![Screenshot from 2023-11-18 14-53-20](https://github.com/Rahul-NITD/goshowtree/assets/96250420/1132993b-299b-4caf-b5d7-47fa7636e55b)

## Development setup

To install the development setup
```sh
git clone https://github.com/Rahul-NITD/goshowtree
cd goshowtree
sh build_dev.sh
```
To run tests
```sh
go test
```

## Release History
* 0.1.0
    * The first proper release
    * CHANGE: Built and tested `BuildTree` and `ShowTree` functions. 
* 0.0.1
    * Started Project

## Meta
[https://github.com/Rahul-NITD](https://github.com/Rahul-NITD)

<!--
## Contributing

1. Fork it (<https://github.com/Rahul-NITD/goshowtree/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
-->
<!-- Markdown link & img dfn's -->
[npm-image]: https://img.shields.io/npm/v/datadog-metrics.svg?style=flat-square
[npm-url]: https://npmjs.org/package/datadog-metrics
[npm-downloads]: https://img.shields.io/npm/dm/datadog-metrics.svg?style=flat-square
[travis-image]: https://img.shields.io/travis/dbader/node-datadog-metrics/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/dbader/node-datadog-metrics
[wiki]: https://github.com/yourname/yourproject/wiki
