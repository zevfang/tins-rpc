<p align="center">
  <img src="resource/logo.png" width="300" alt="TinsRPC Logo">
</p>

# TinsRPC

English | [简体中文](README-CN.md)

TinsRPC is an rpc client tool. To use it, you must import the proto file. It is not yet mature, and there may be abnormal phenomena, but it will definitely get better and better. If you are interested in it, welcome to join us.

### Features

* Support simple rpcx service calls
* ……

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
fyne package -os windows -icon ./theme/icon.png -name TinsRPC
fyne package -os darwin -icon ./theme/icon.png -name TinsRPC
fyne package -os linux -icon ./theme/icon.png -name TinsRPC
```

local build

```
go build -ldflags="-H windowsgui"
```

You can go to the [release](https://github.com/zevfang/tins-rpc/releases) page for the latest binary.


