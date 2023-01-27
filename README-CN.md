<p align="center">
  <img src="resource/logo.png" width="200" alt="TinsRPC Logo">
</p>

<div align=center>

[![Go](https://github.com/zevfang/tins-rpc/workflows/Go/badge.svg?branch=master)](https://github.com/zevfang/tins-rpc/actions)
[![Release](https://img.shields.io/github/v/release/zevfang/tins-rpc.svg?style=flat-square)](https://github.com/zevfang/tins-rpc)

</div>

# TinsRPC

[English](README.md) | 简体中文

TinsRPC是一个rpc的客户端工具，使用它你必须导入proto文件，目前它并不成熟，可能会有异常现象，但它一定会越来越好，如果你对它感兴趣欢迎加入我们。

### Features

* 跨平台RPC客户端
* 支持向RPCx服务发送请求（一元）
* 支持向gRPC服务发送请求（一元）


### Installation

此软件包需要 Go 1.18 或更高版本。可以通过运行以下命令来安装它：

```
go get github.com/zevfang/tins-rpc
```


### Usage

首先你需要安装 [Fyne](https://github.com/fyne-io/fyne) ，然后克隆这个仓库并编译它：
```
fyne package
```

### Build

fyne打包编译
```
fyne package -os windows -icon ./resource/logo.png -name TinsRPC
fyne package -os darwin -icon ./resource/logo.png -name TinsRPC
fyne package -os linux -icon ./resource/logo.png -name TinsRPC
```

本地编译

```
go build -ldflags="-H windowsgui"
```

您可以前往 [release](https://github.com/zevfang/tins-rpc/releases) 页面获取最新的二进制文件。
