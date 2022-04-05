package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/tbbrave/repository-report-generator/graphql"
	"gopkg.in/yaml.v1"
)

func main() {
	config := parseConfig()

	gitInfo := RequestForCommitInfo(config)

	weeks, authors, progress := RequestForPullRequests(config)

	templ := `
# 1. 统计数据

共%d位作者提交了%s个 Commit 。感谢以下作者的贡献：

%s

最近%d周，共修改新增代码%s行 ，删除代码%s行 。

%s
`

	data := fmt.Sprintf(templ,
		len(authors), gitInfo.commitNumber, strings.Join(authors, ","),
		weeks, gitInfo.addLines, gitInfo.modifiedDel, progress)

	os.WriteFile("output.md", []byte(data), 777)

	fmt.Println(">>>output.md")
	fmt.Println(data)
}

type Config struct {
	//	Template string
	Github struct {
		Token      string `yaml:"token"`
		Repository string `yaml:"repository"`
		StartTime  string `yaml:"start_time"`
		EndTime    string `yaml:"end_time"`
		repoName   string
		repoOwner  string
	}
}

func parseConfig() *Config {
	f, err := os.Open("./config.yml")
	if err != nil {
		panic("config.yml not found")
	}
	defer f.Close()

	configData, _ := io.ReadAll(f)
	config := &Config{}
	yaml.Unmarshal(configData, &config)

	buf := make([]byte, 0, 100)
	out := bytes.NewBuffer(buf)
	cmd := exec.Command("/bin/sh", "-c", "git config -l | grep insteadof=https://github.com | awk -F'@' '{ print $1}' | awk -F':' '{ print $NF}'")
	cmd.Stdout = out
	cmd.Run()

	config.Github.Token = strings.TrimSpace(out.String())

	repos := strings.Split(config.Github.Repository, "/")
	config.Github.repoOwner, config.Github.repoName = repos[0], repos[1]

	return config
}

type GitCommandData struct {
	modifiedAdd  string
	modifiedDel  string
	addLines     string
	commitNumber string
}

type PullRequestData struct {
	authorNumber int
	authorsList  string //merged only

	WIP     string
	BugFix  string
	Feat    string
	Other   string
	Improve string
}

func RequestForPullRequests(config *Config) (int, []string, string) {
	prCollection := graphql.PullGithubData(config.Github.repoOwner, config.Github.repoName,
		config.Github.StartTime, config.Github.EndTime, config.Github.Token)

	templ := `
# 2. 主要进展
## 2.1 新增功能
%s

## 2.2 Bug 修复
%s

## 2.3 功能改进
%s

## 2.4 WIP
%s

## 2.5 其他
%s
`

	return prCollection.Weeks, prCollection.Authors, fmt.Sprintf(templ,
		strings.Join(prCollection.Feature, string('\n')),
		strings.Join(prCollection.BugFix, string('\n')),
		strings.Join(prCollection.Improve, string('\n')),
		strings.Join(prCollection.WIP, string('\n')),
		strings.Join(prCollection.Other, string('\n')),
	)
}

func RequestForCommitInfo(config *Config) (data GitCommandData) {
	err := exec.Command("/bin/sh", "-c", fmt.Sprintf("git clone https://github.com/%s %s", config.Github.Repository, config.Github.repoName)).Run()
	fmt.Println("git clone err", err)

	buf := make([]byte, 0, 256)
	out := bytes.NewBuffer(buf)

	collectCommitLineNumbers := fmt.Sprintf(`cd %s && git log --pretty=tformat: --numstat --since="%s" --until="%s" | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { print  add, subs, loc }';`, config.Github.repoName, config.Github.StartTime, config.Github.EndTime)
	cmd := exec.Command("/bin/sh", "-c", collectCommitLineNumbers)

	cmd.Stdout = out
	err = cmd.Run()
	if err != nil {
		fmt.Println("git log err:", err)
	}

	collectCommitLineNumbersData := out.String()

	splited := strings.Split(collectCommitLineNumbersData, " ")
	data.modifiedDel, data.modifiedDel, data.addLines = splited[0], splited[1], splited[2]

	buf1 := make([]byte, 0, 256)
	out1 := bytes.NewBuffer(buf1)

	commitCommand := fmt.Sprintf(`cd %s && git log --since="%s" --until="%s" | grep "^commit" | wc -l`, config.Github.repoName, config.Github.StartTime, config.Github.EndTime)
	//	fmt.Println(commitCommand)
	cmd = exec.Command("/bin/sh", "-c", commitCommand)
	cmd.Stdout = out1

	err = cmd.Run()

	if err != nil {
		fmt.Println("git log err: ", err)
	}

	data.commitNumber = strings.TrimSpace(out1.String())

	return
}
