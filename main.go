//go:generate goversioninfo
package main

import (
	"bytes"
	"crypto/md5"
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
	// 输出文件名
	output := flag.String("o", "xcgui.dll", "output filename")
	// xcgui.dll 的版本号
	version := flag.String("v", "", "version number of xcgui.dll, example: 3.3.5.0")
	// xcgui.dll 的位数
	bit := flag.Uint("b", 64, "bitness of xcgui.dll")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n\nPossible flags:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	// 获取最新版本号
	*version = strings.TrimSpace(*version)
	if *version == "" {
		latest, err := getLatestVersion()
		if err != nil {
			fmt.Println(err)
			return
		}
		*version = latest
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
		addr = fmt.Sprintf("https://cnb.cool/twgh521/xcguidll/-/releases/download/%s/xcgui-32.dll", *version)
	} else {
		addr = fmt.Sprintf("https://cnb.cool/twgh521/xcguidll/-/releases/download/%s/xcgui.dll", *version)
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

	data, err := getDll(addr)
	if err != nil {
		quit <- true
		fmt.Println("\n" + err.Error())
		return
	}

	if len(data) < 1.5*1024*1024 { // 小于1.5M肯定下载失败了
		quit <- true
		fmt.Println("\ndownload failed")
		return
	}

	err = os.WriteFile(*output, data, 0777)
	if err != nil {
		quit <- true
		fmt.Println(err)
		return
	}

	// 计算data的md5
	md5Str := fmt.Sprintf("%x", md5.Sum(data))
	fmt.Printf("\nMD5: %s\n", strings.ToUpper(md5Str))

	quit <- true
	fmt.Println("download successful")
}

// 从指定网址下载dll
func getDll(addr string) ([]byte, error) {
	res, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, errors.New("file not found")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	if bytes.Contains(body, []byte("NoSuchKey")) {
		return nil, errors.New("file not found")
	}
	return body, err
}

// 获取最新版本号
func getLatestVersion() (string, error) {
	res, err := http.Get("https://cnb.cool/twgh521/xcguidll/-/git/raw/main/version.txt?download=true")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	version := strings.TrimSpace(string(body))
	if version == "" {
		return "", errors.New("failed to get the latest version number")
	}
	return version, nil
}
