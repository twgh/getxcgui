## 介绍

本工具可以帮助下载xcgui.dll

## 获取

```bash
go install github.com/twgh/getxcgui@latest
```

## Flags

```bash
-v		xcgui.dll 的版本号，默认为最新版本
-b		xcgui.dll 的位数，默认为64
-o		输出文件名，默认为“xcgui.dll”
```

## 使用

##### 默认下载最新版本64位的dll到当前目录

```bash
getxcgui
```

![cmd](https://s1.ax1x.com/2022/07/03/j8WP41.jpg)

##### 下载最新版本32位的dll到当前目录

```bash
getxcgui -b 32
```

##### 下载3.3.5.0版本，64位的dll到当前目录

```bash
getxcgui -v 3.3.5.0
```

##### 下载3.3.5.0版本，64位的dll到当前目录，命名为xc.dll

```bash
getxcgui -v 3.3.5.0 -o xc.dll
```

