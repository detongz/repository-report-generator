package graphql

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/tbbrave/repository-report-generator/graphql/model"
)

func TestPullGithubData(t *testing.T) {
	buf := make([]byte, 0, 100)
	out := bytes.NewBuffer(buf)
	cmd := exec.Command("/bin/sh", "-c", "git config -l | grep insteadof=https://github.com | awk -F'@' '{ print $1}' | awk -F':' '{ print $NF}'")
	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var token = strings.TrimSpace(out.String())
	fmt.Println("+++++", token)

	res1, res2 := pullGithubData("apache", "incubator-seatunnel", "2022-03-29T13:33:46.000Z", "2022-04-05T13:33:46.000Z", token, model.PullRequestStateOpen)

	fmt.Println(res1, res2)
}
