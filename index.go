package main

import (
	"github.com/urfave/cli/v2"

	"fmt"
	"os"
	"os/exec"
	"strings"
)

func index(c *cli.Context) error {
	recordFlag := c.Bool("record")
	if recordFlag {
		panic("Not Support")
	}

	first := c.Args().First()
	tail := c.Args().Tail()

	cmd := exec.Command(first, tail...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fullCmd := strings.Join(c.Args().Slice(), " ")
	client.Email(fmt.Sprintf("✅ %s", fullCmd), "")
	return nil
}
