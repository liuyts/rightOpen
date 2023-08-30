package main

import (
	"bufio"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
	"strings"
)

func CreateOpenBackground() {
	var openName string
	var openPath string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入右键打开的名称")
	if scanner.Scan() {
		openName = scanner.Text()
	}
	fmt.Println("请输入打开程序的路径，后缀为.exe")
	if scanner.Scan() {
		openPath = scanner.Text()
		openPath = strings.TrimSpace(openPath)
	}

	registryName := extractFileName(openPath)
	showPath := fmt.Sprintf(`Directory\Background\shell\%s`, registryName)
	commandPath := fmt.Sprintf(`Directory\Background\shell\%s\command`, registryName)
	openCommand := fmt.Sprintf(`"%s" "%%V"`, openPath)

	//创建或打开注册表项
	key, _, err := registry.CreateKey(registry.CLASSES_ROOT, showPath, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("无法创建或打开注册表项:", err)
		return
	}
	defer key.Close()

	// 设置注册表值
	err = key.SetStringValue(Empty, openName)
	if err != nil {
		fmt.Println("无法设置注册表值:", err)
		return
	}
	err = key.SetStringValue(Icon, openPath)
	if err != nil {
		fmt.Println("无法设置注册表值:", err)
		return
	}

	key, _, err = registry.CreateKey(registry.CLASSES_ROOT, commandPath, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("无法创建或打开注册表项:", err)
		return
	}
	err = key.SetStringValue(Empty, openCommand)
	if err != nil {
		fmt.Println("无法设置注册表值:", err)
		return
	}

	fmt.Println("HKEY_CLASSES_ROOT\\" + showPath)
	fmt.Println("注册表项添加成功！")
}

func DeleteOpenBackground() {
	fmt.Println("请输入打开程序的路径，后缀为.exe")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	openPath := scanner.Text()

	registryName := extractFileName(openPath)
	showPath := fmt.Sprintf(`HKEY_CLASSES_ROOT\Directory\Background\shell\%s`, registryName)
	fmt.Println(showPath)

	//删除注册表项
	cmd := exec.Command("cmd", "/c", "reg", "delete", showPath, "/f")
	err := cmd.Run()
	if err != nil {
		fmt.Println("注册表项删除失败:", err)
		return
	}

	fmt.Println("注册表项删除成功！")

}
