# getxcgui
<p>
	<a href="https://github.com/twgh/getxcgui/releases"><img src="https://img.shields.io/badge/release-0.0.1-blue" alt="release"></a>
	<a href="https://golang.org"> <img src="https://img.shields.io/badge/golang-1.17-blue" alt="golang"></a>
	<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-brightgreen" alt="License"></a>
</p>

## 介绍

本工具可以帮助下载xcgui.dll

## 获取

```bash
go install -ldflags="-s -w" github.com/twgh/getxcgui@latest
```

## Flags

```bash
-v	xcgui.dll 的版本号，为空时默认最新版本
-b	xcgui.dll 的位数，为空时默认64
-o	输出文件名，为空时默认“xcgui.dll”
```

## 使用

#### 默认下载最新版本64位的dll到当前目录

```bash
getxcgui
```

![cmd](https://s1.ax1x.com/2022/07/04/jJJNS1.png)

#### 下载最新版本32位的dll到当前目录

```bash
getxcgui -b 32
```

#### 下载3.3.5.0版本，64位的dll到当前目录

```bash
getxcgui -v 3.3.5.0
```

#### 下载3.3.5.0版本，64位的dll到当前目录，命名为xc.dll

```bash
getxcgui -v 3.3.5.0 -o xc.dll
```

