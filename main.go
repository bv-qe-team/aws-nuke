package main

import (
	"os"

	"github.com/bv-qe-team/aws-nuke/cmd"
)

type NukeParameters struct {
	ConfigPath string

	Profile         string
	AccessKeyID     string
	SecretAccessKey string

	NoDryRun   bool
	Force      bool
	ForceSleep int

	MaxWaitRetries int
}

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(-1)
	}
}
