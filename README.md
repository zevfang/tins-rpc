<p align="center">
  <img src="resource/logo.png" width="200" alt="TinsRPC Logo">
</p>

<div align=center>

[![Go](https://github.com/zevfang/tins-rpc/workflows/Go/badge.svg?branch=master)](https://github.com/zevfang/tins-rpc/actions)
[![Release](https://img.shields.io/github/v/release/zevfang/tins-rpc.svg?style=flat-square)](https://github.com/zevfang/tins-rpc/releases)

</div>

# TinsRPC

English | [简体中文](README-CN.md)

TinsRPC is an RPC client tool that satisfies developers' RPC local debugging. Currently, the supported functions are relatively simple, and more practical functions are still planned.

<p align="center">
  <img src="resource/preview.gif" style="max-width: 100%; display: inline-block;" data-target="animated-image.originalImage">
</p>

### Features

* Cross platform RPC client
* Support .proto file import
* Send requests to a RPCx service(Supports unary)
* Send requests to a gRPC service(Supports unary)
* Support Chinese and English themes
* Support dark theme

### Installation

This packages requires Go 1.18 or later. It can be installed by running the command below:

```
go get github.com/zevfang/tins-rpc
```


### Usage

First you need [Fyne](https://github.com/fyne-io/fyne) installed, then clone this repository and compile it:
```
fyne package
```

### Build

fyne build
```
fyne package -os windows -icon ./resource/logo.png -name TinsRPC
fyne package -os darwin -icon ./resource/logo.png -name TinsRPC
fyne package -os linux -icon ./resource/logo.png -name TinsRPC
```

local build

```
go build -ldflags="-H windowsgui"
```

You can go to the [release](https://github.com/zevfang/tins-rpc/releases) page for the latest binary.


