package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gclg",
		Usage: "Simple Git wrapper to work with Github repositories",
		Action: func(cCtx *cli.Context) error {
			for i := 0; i < cCtx.Args().Len(); i++ {
				arg := cCtx.Args().Get(i)
				vals := strings.Split(arg, "/")

				if len(vals) != 2 {
					fmt.Printf("Invalid argument: %s", arg)
					continue
				}

				owner := vals[0]
				repo := vals[1]

				fmt.Printf("Started cloning %s/%s...", owner, repo)

				if _, err := os.Stat("/path/to/whatever"); !os.IsNotExist(err) {
					fmt.Printf("Folder with name '%s' already exists. Skipping...", repo)
					continue
				}

				git := "git"
				clone := "clone"
				// git@github.com:<owner>/<repo>.git
				// https://github.com/<owner>/<repo>.git
				url := fmt.Sprintf("git@github.com:%s/%s.git", owner, repo)

				cmd := exec.Command(git, clone, url)
				stdout, err := cmd.Output()

				if err != nil {
					fmt.Println(err.Error())
					continue
				}

				fmt.Print(string(stdout))
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
