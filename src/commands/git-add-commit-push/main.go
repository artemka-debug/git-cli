package gacp

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/artemka-debug/git-cli/src/utils"
	"github.com/urfave/cli/v2"
)

func action(c *cli.Context) error {
	agrumentLength := c.Args().Len()

	if agrumentLength != 2 {
		return errors.New("Please provide 2 arguments. Example git-add-commit-push <add pattern> <commit message>")
	}

	addPattern := c.Args().Get(0)
	commitMessage := c.Args().Get(1)

	fmt.Printf("Adding files to git by pattern <%s>...\n", addPattern)

	err := utils.RunWithError(exec.Command("git", "add", addPattern))
	if err != nil {
		return err
	}

	fmt.Printf("Commiting files to git with message <%s>...\n", commitMessage)

	err = utils.RunWithError(exec.Command("git", "commit", "-m", commitMessage))
	if err != nil {
		return err
	}

	fmt.Println("Pushing files to git...")
 
	err = utils.RunWithError(exec.Command("git", "push"))
	if err != nil {
		return err
	}

	return nil
}

var GACPCommand = &cli.Command{
	Name:      "git-add-commit-push",
	Aliases:   []string{"gacp"},
	Usage:     "Add, commit and push to git",
	ArgsUsage: "<add pattern> <commit message>",
	Action: action, 
}
