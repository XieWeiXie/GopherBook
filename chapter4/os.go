package chapter4

import (
	"bytes"
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

	windowsPathFormat := filepath.FromSlash(windowsPath)
	fmt.Println(filepath.Base(windowsPathFormat))
	fmt.Println(filepath.Dir(windowsPathFormat))
}

func OSDirUsage() {
	path, _ := os.Getwd()
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		//fmt.Println("file:", info.Name(), "in directory:", path)

		return nil
	})

}

func OSExecUsage() {
	dockerPath, err := exec.LookPath("docker")
	if err != nil {
		return
	}
	fmt.Println(dockerPath)

	pwdPath, err := exec.LookPath("pwd")
	if err != nil {
		return
	}
	fmt.Println(pwdPath)

	// 1
	cmd := exec.Command("docker", "ps")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	opBytes, err := ioutil.ReadAll(stdout)
	fmt.Println(cmd.Dir, cmd.Path, string(opBytes))

	// 2
	pwd, _ := os.Getwd()
	cmd2 := exec.Command("ls", pwd)
	var buf bytes.Buffer
	cmd2.Stdout = &buf
	cmd2.Run()
	fmt.Println(buf.String())

	// 3
	cmd3 := exec.Command("cat", "log.log")
	out, _ := cmd3.Output()
	//out2, _ := cmd3.CombinedOutput()
	fmt.Println(string(out))

	// 4
	cmd4 := exec.Command("sh", "os.sh")
	out4, _ := cmd4.CombinedOutput()
	fmt.Println(string(out4))

	// 5
	cmd5 := exec.Command("ls", pwd)
	stdout5, _ := cmd5.StdoutPipe()
	if err := cmd5.Start(); err != nil {
		fmt.Println(err)
	}

	bytes, err := ioutil.ReadAll(stdout5)

	err = ioutil.WriteFile("file.log", bytes, 0644)
	if err != nil {
		panic(err)
	}

}
