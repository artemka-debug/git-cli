package gmacio

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/artemka-debug/git-cli/src/utils"
	"github.com/urfave/cli/v2"
	"github.com/lainio/err2"
)

func action(c *cli.Context) (err error) {
	defer err2.Handle(&err)

	agrumentLength := c.Args().Len()

	if agrumentLength != 1 {
		return errors.New("Please provide 1 arguments. Example git-merge-all-commits-into-one <commit message>")
	}

	commitMessage := c.Args().Get(0)

	fmt.Printf("Resetting git...\n")
	err = utils.RunWithError(exec.Command("bash", "-c", "git reset $(git merge-base master $(git branch --show-current))"))
	
	fmt.Printf("Adding files to git by pattern...\n")
 	err = utils.RunWithError(exec.Command("git", "add", "-A"))

	fmt.Printf("Commiting files to git with message <%s>...\n", commitMessage)
	err = utils.RunWithError(exec.Command("git", "commit", "-m", commitMessage))

	fmt.Println("Pushing files to git...")
	err = utils.RunWithError(exec.Command("git", "push", "--force"))

	return nil
}

var GMACIOCommand = &cli.Command{
	Name:      "git-merge-all-commits-into-one",
	Aliases:   []string{"gmacio"},
	Usage:     "Adds with -A flag, commits and force pushes to git",
	ArgsUsage: "<commit message>",
	Action: action,
}
