package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	output := flag.String("o", "xcgui.dll", "output filename")                             // 输出文件名
	version := flag.String("v", "latest", "version number of xcgui.dll, example: 3.3.5.0") // xcgui.dll 的版本号
	bit := flag.Uint("b", 64, "bitness of xcgui.dll")                                      // xcgui.dll 的位数

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n\nPossible flags:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	// 获取最新版本号
	*version = strings.TrimSpace(*version)
	if *version == "" {
		*version = "latest"
	}

	// 删首尾空
	*output = strings.TrimSpace(*output)
	if *output == "" {
		*output = "xcgui.dll"
	}

	// 判断位数, 得到下载地址
	if *bit == 32 || *bit == 86 {
		*bit = 32
	} else {
		*bit = 64
	}

	addr := ""
	if *bit == 32 {
		addr = fmt.Sprintf("https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=%s", *version)
	} else {
		addr = fmt.Sprintf("https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=%s", *version)
	}

	// 开始下载dll
	fmt.Printf("start download %s, %d-bit, output: %s\n", *version, *bit, *output)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 1500; i++ { // 超过300秒就判定为下载失败
			select {
			case <-quit:
				return
			default:
				fmt.Print(".")
				time.Sleep(time.Millisecond * 200)
			}
		}
		fmt.Println("\ndownload failed")
		os.Exit(0)
	}()

	dll, err := getDll(addr)
	if err != nil {
		quit <- true
		fmt.Println(err)
		return
	}

	if len(dll) < 1.5*1024*1024 { // 小于1.5M肯定下载失败了
		quit <- true
		fmt.Println("download failed")
		return
	}

	err = os.WriteFile(*output, dll, 0666)
	if err != nil {
		quit <- true
		fmt.Println(err)
		return
	}

	quit <- true
	fmt.Println("\ndownload successful")
}

// 从指定网址下载dll
func getDll(addr string) ([]byte, error) {
	res, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return body, err
	}

	if bytes.Contains(body, []byte("File not found")) {
		return nil, errors.New("file not found")
	}
	return body, err
}
