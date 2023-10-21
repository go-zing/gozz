<p align="center">
  <a href="https://github.com/go-zing/gozz" target="_blank">
    <img src="https://raw.githubusercontent.com/go-zing/gozz-doc/main/docs/.vuepress/public/logo.png" alt="logo">
  </a>
</p>

<div align=center>

[![Go](https://github.com/go-zing/gozz/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/go-zing/gozz/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zing/gozz)](https://goreportcard.com/report/github.com/go-zing/gozz)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-zing/gozz.svg)](https://pkg.go.dev/github.com/go-zing/gozz)

[![License: MIT](https://img.shields.io/github/license/go-zing/gozz)](https://github.com/go-zing/gozz/blob/master/LICENSE)
[![Last Commit](https://img.shields.io/github/last-commit/go-zing/gozz)](https://github.com/go-zing/gozz/commits)
[![codecov](https://codecov.io/gh/go-zing/gozz/branch/main/graph/badge.svg)](https://codecov.io/gh/go-zing/gozz)

</div>

## Introduction

[English](https://go-zing.github.io/gozz) | [简体中文](https://go-zing.github.io/gozz/zh)

### Fast and Simple

Intuitive annotation syntax, clean and fast command line tools,
runtime-dependencies-free generated codes.

#### Install

```shell
go install github.com/go-zing/gozz@latest
```

#### Usage

`Gozz` CLI is built with [cobra](https://github.com/spf13/cobra), command syntax as follows:

```shell
gozz [--GLOBAL-FLAGS] [COMMAND] [--COMMAND-FLAGS] [ARGS]
```

The main plugin execute command would be:

```shell
gozz run -p "[PLUGIN][:OPTIONS]" filename
```

#### Annotation

Annotations are comments that stick with declaration object, and match syntax as follows:

```go
// +zz:[PLUGIN][:ARGS][:OPTIONS]
type T interface{}
```

Check out [here](https://go-zing.github.io/gozz/guide/getting-started) for more details.

### Awesome Plugins

`Gozz` provides a series of awesome builtin plugins.
Rather than `Gozz` provides these built-in plugins,
Actually in order to integrate these plugins,
we develop the toolkits named `Gozz`.

- [Wire](https://go-zing.github.io/gozz/guide/plugins/wire) Autowire DI and AOP proxy。
- [Api](https://go-zing.github.io/gozz/guide/plugins/api) Generate API routing and invoker based on `interface`。
- [Impl](https://go-zing.github.io/gozz/guide/plugins/impl) Sync `interface` and `implement`。
- [Doc](https://go-zing.github.io/gozz/guide/plugins/doc) Generate runtime documentation based on comments。
- [Tag](https://go-zing.github.io/gozz/guide/plugins/tag) Manage structure field tags in templating。
- [Orm](https://go-zing.github.io/gozz/guide/plugins/orm) Generates data entity mapping code。
- [Option](https://go-zing.github.io/gozz/guide/plugins/option) Generates `Funcitonal Options` code。

### High Extensibility

We provide customizable generating templates,
[gozz-core](https://github.com/go-zing/gozz-core) for code analysis,
edit and generate.
External `.so` plugins and [official external](https://github.com/go-zing/gozz-plugins) supported also.

```shell
gozz install [--output/-o] [--filepath/-f] [repository] 
```

## Showcase

- [Gozz-Core](https://github.com/go-zing/gozz-core)
- [Gozz-Plugins](https://github.com/go-zing/gozz-plugins)
- [Gozz-Doc](https://github.com/go-zing/gozz-doc)
- [Gozz-Doc-Examples](https://github.com/go-zing/gozz-doc-examples)

## License

[Apache-2.0](https://github.com/go-zing/gozz/blob/main/LICENSE)