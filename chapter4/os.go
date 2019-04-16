package chapter4

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func OsUsage() {
	// 判断文件是否存在，获取文件信息
	fileMode, err := os.Stat("log.log")
	if os.IsNotExist(err) {
		return
	}
	fmt.Println(fileMode.Name(), fileMode.Mode(), fileMode.Size())
}

func OSUsageWith() {

	file, _ := os.OpenFile("os.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fmt.Println(file.Name())
	file.WriteString("Hello")
	file.WriteString("HelloWorld")
	ioutil.ReadFile("os.log")
}

func OSPathUsage() {
	dir, _ := os.Getwd()
	fmt.Println(dir, path.Base(dir))
	fmt.Println(dir, path.Dir(dir))
	parentDir := path.Dir(dir)
	fmt.Println(dir, path.Join(parentDir, "Chapter3"))
}

func OSPathWindows() {
	windowsPath := `C:\Windows\System\Chapter4`
	fmt.Println(path.Base(windowsPath))
	fmt.Println(path.Dir(windowsPath))

	windowsPathFormat := filepath.ToSlash(windowsPath)
	fmt.Println(filepath.Base(windowsPathFormat))
	fmt.Println(filepath.Dir(windowsPathFormat))
}

func OSExecUsage() {
	result := exec.Command("docker", "ps")
	result.Run()
	fmt.Println(result)
}
