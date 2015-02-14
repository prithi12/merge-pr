package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	GitRemoteRegexp = regexp.MustCompile("(https|git)(@|://)github.com(:|/)([a-zA-Z0-9-_]+)/([a-zA-Z0-9-_]+).git")
)

func fetchRepoOwnerAndName() (string, string) {
	return extractOwnerAndNameFromRemote(gitOriginRemote(gitRemotes()))
}

func gitRemotes() []string {
	out, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		log.Fatal(err)
	}
	remotes := strings.Split(
		strings.TrimRight(string(out), "\n"),
		"\n",
	)
	return remotes
}

func extractOwnerAndNameFromRemote(url string) (string, string) {
	matches := GitRemoteRegexp.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", ""
	}
	return matches[len(matches)-2], matches[len(matches)-1]
}

func extractUrlFromRemote(remote string) string {
	return GitRemoteRegexp.FindString(remote)
}

func gitOriginRemote(remotes []string) string {
	var origin string
	for _, remote := range remotes {
		if strings.HasPrefix(remote, "origin") {
			origin = GitRemoteRegexp.FindString(remote)
			break
		}
	}
	return origin
}

func gitPull() {
	exec.Command("git", "pull", "--rebase").Run()
}

func gitPush() {
	exec.Command("git", "push").Run()
}

func commitChangesToHistoryFile(pr string) {
	exec.Command("git", "add", "History.markdown").Run()
	cmd := exec.Command(
		"git",
		"commit",
		"-m",
		"Update history to reflect merge of #"+pr+" [ci skip]",
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
