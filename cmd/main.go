/*
Copyright (c) 2020-2021 Tinet. All rights reserved.

Author: seanchann <zhouxq@ti-net.com.cn>

See docs/ for more information about this project.
*/

package main

import (
	goflag "flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Tinet-mucw/dtls-srtp-decrypt/cmd/app"
	"github.com/spf13/pflag"
	"k8s.io/component-base/logs"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewCommand()

	goflag.CommandLine.Parse([]string{})
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
