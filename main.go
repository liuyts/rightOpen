package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

var (
	Empty = ""
	Icon  = "Icon"
)

func main() {
	var option string
	fmt.Println("1.右键文件")
	fmt.Println("2.右键文件夹")
	fmt.Println("3.右键文件夹背景")
	fmt.Println("请输入编号")
	fmt.Scanln(&option)
	switch option {
	case "1":
		OpenFile()
	case "2":
		OpenDir()
	case "3":
		OpenBackground()
	}
}

func extractFileName(path string) string {
	base := filepath.Base(path)
	filename := strings.TrimSuffix(base, filepath.Ext(base))
	return filename
}

func OpenFile() {
	var option string
	fmt.Println("1.创建")
	fmt.Println("2.删除")
	fmt.Println("请输入编号")
	fmt.Scanln(&option)
	switch option {
	case "1":
		CreateOpenFile()
	case "2":
		DeleteOpenFile()
	}
}

func OpenDir() {
	var option string
	fmt.Println("1.创建")
	fmt.Println("2.删除")
	fmt.Println("请输入编号")
	fmt.Scanln(&option)
	switch option {
	case "1":
		CreateOpenDir()
	case "2":
		DeleteOpenDir()
	}
}

func OpenBackground() {
	var option string
	fmt.Println("1.创建")
	fmt.Println("2.删除")
	fmt.Println("请输入编号")
	fmt.Scanln(&option)
	switch option {
	case "1":
		CreateOpenBackground()
	case "2":
		DeleteOpenBackground()
	}
}
