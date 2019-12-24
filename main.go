package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Markdown struct {
	Path string
	cmd  *exec.Cmd
}

func NewMarkdown(path string) *Markdown {
	return &Markdown{
		Path: path,
	}
}

func (m *Markdown) setCmd(cmd string, args string) {
	m.cmd = exec.Command("bash", "-c", fmt.Sprintf("%s %s", cmd, args))
}

func (m Markdown) ListFiles() []string {
	m.setCmd("ls", m.Path)
	results, e := m.cmd.CombinedOutput()
	if e != nil {
		log.Println(e)
		return nil
	}
	str := string(results)
	replacer := strings.NewReplacer("\n", " ", "\t", " ")
	list := strings.Split(replacer.Replace(str), " ")
	var response []string
	for _, i := range list {
		if strings.Contains(i, ".md") {
			response = append(response, i)
		}
	}
	return response
}

func (m Markdown) ToDoc(source string, dst string) {
	os.RemoveAll(dst)
	os.MkdirAll(dst, os.ModePerm)
	rename := func(v string) string {
		replacer := strings.NewReplacer(".md", ".docx")
		return replacer.Replace(v)
	}
	for _, i := range m.ListFiles() {
		log.Println("pandoc", fmt.Sprintf(" -s %s/%s -o %s/%s", source, i, dst, rename(i)))
		m.setCmd("pandoc", fmt.Sprintf(" -s %s/%s -o %s/%s", source, i, dst, rename(i)))
		m.cmd.CombinedOutput()
	}
}

func (m Markdown) ToHtml(source string, dst string) {
	os.RemoveAll(dst)
	os.MkdirAll(dst, os.ModePerm)
	rename := func(v string) string {
		replacer := strings.NewReplacer(".md", ".html")
		return replacer.Replace(v)
	}
	for _, i := range m.ListFiles() {
		log.Println("pandoc", fmt.Sprintf(" --ascii %s/%s -o %s/%s", source, i, dst, rename(i)))
		m.setCmd("pandoc", fmt.Sprintf(" --ascii %s/%s -o %s/%s", source, i, dst, rename(i)))
		m.cmd.CombinedOutput()
	}
}

func (m Markdown) ToPdf(dst string) {
	// pandoc not work
}

func (m Markdown) Copy(source string, dst string) {
	os.RemoveAll(dst)
	os.MkdirAll(dst, os.ModePerm)
	for _, i := range m.ListFiles() {
		m.setCmd("cp", fmt.Sprintf("%s/%s %s", source, i, dst))
		m.cmd.CombinedOutput()
	}
}
func main() {
	md := NewMarkdown("./Book")
	results := md.ListFiles()
	fmt.Println(results)
	// copy
	md.Copy("./Book", "./BookTest/markdown")
	// todocx
	md.ToDoc("./Book", "./BookTest/docx")
	// tohtml
	md.ToHtml("./Book", "./BookTest/html")
}
