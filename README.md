<!-- markdownlint-disable MD041 -->
<h1 align="center">:dvd: dependency visualization</h1>
<p align="center">
<!--  -->
<a href="https://github.com/jujili/depviz/releases"> <img src="https://img.shields.io/github/v/tag/jujili/depviz?include_prereleases&sort=semver" alt="Release" title="Release"></a>
<!--  -->
<a href="https://www.travis-ci.org/jujili/depviz"><img src="https://www.travis-ci.org/jujili/depviz.svg?branch=master"/></a>
<!--  -->
<a href="https://codecov.io/gh/jujili/depviz"><img src="https://codecov.io/gh/jujili/depviz/branch/master/graph/badge.svg"/></a>
<!--  -->
<a href="https://goreportcard.com/report/github.com/jujili/depviz"><img src="https://goreportcard.com/badge/github.com/jujili/depviz" alt="Go Report Card" title="Go Report Card"/></a>
<!--  -->
<a href="http://godoc.org/github.com/jujili/depviz"><img src="https://img.shields.io/badge/godoc-depviz-blue.svg" alt="Go Doc" title="Go Doc"/></a>
<!--  -->
<br/>
<!--  -->
<a href="https://github.com/jujili/depviz/blob/master/CHANGELOG.md"><img src="https://img.shields.io/badge/Change-Log-blueviolet.svg" alt="Change Log" title="Change Log"/></a>
<!--  -->
<a href="https://golang.google.cn"><img src="https://img.shields.io/github/go-mod/go-version/jujili/depviz" alt="Go Version" title="Go Version"/></a>
<!--  -->
<a href="https://github.com/jujili/depviz/blob/master/LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT License" title="MIT License"/></a>
<!--  -->
<br/>
<!--  -->
<a target="_blank" href="//shang.qq.com/wpa/qunwpa?idkey=7f61280435c41608fb8cb96cf8af7d31ef0007c44b223c9e3596ce84dec329bc"><img border="0" src="https://img.shields.io/badge/QQ%20群-23%2053%2000%2093-blue.svg" alt="jili交流QQ群:23530093" title="jili交流QQ群:23530093"></a>
<!--  -->
<a href="https://mp.weixin.qq.com/s?__biz=MzA4MDU4NDI5Mw==&mid=2455230332&idx=1&sn=8086c43e259b0012596ed63d6ecd7d10&chksm=88017c76bf76f5604f2f3280ffd96029b5ccaf99db48d18066d3e3bc9bc8a2e1a05de1a3225f&mpshare=1&scene=1&srcid=&sharer_sharetime=1578553397373&sharer_shareid=5ce52651949258759d82d1bf31b455b5#rd"><img src="https://img.shields.io/badge/微信公众号-jujili-success.svg" alt="微信公众号：jujili" title="微信公众号：jujili"/></a>
<!--  -->
<a href="https://zhuanlan.zhihu.com/jujili"><img src="https://img.shields.io/badge/知乎专栏-jili-blue.svg" alt="知乎专栏：jili" title="知乎专栏：jili"/></a>
<!--  -->
</p>

depviz 会把 Go 模块的依赖关系进可视化处理。

- [安装与更新](#%e5%ae%89%e8%a3%85%e4%b8%8e%e6%9b%b4%e6%96%b0)
- [选项](#%e9%80%89%e9%a1%b9)
- [说明](#%e8%af%b4%e6%98%8e)
	- [颜色](#%e9%a2%9c%e8%89%b2)

## 安装与更新

在命令行中输入以下内容，可以获取到最新版

```shell
go get -u github.com/jujili/depviz
```

## 选项

<https://graphviz.gitlab.io/_pages/doc/info/attrs.html>

## 说明

### 颜色

depviz 对不同种类的模块标记不同的颜色。

- 绿色：Go 标准模块。
- 蓝色：当前模块或忽略了前缀的模块。
- 黄色：其他模块。