package cmd

import "github.com/spf13/cobra"

var pushCmd = &cobra.Command{
	Use:     "push",
	Aliases: []string{"push"},
	Short:   "Push a command to a Cobra Application",
	Long: `Push (cobra add) will create a new command, with a license and
the appropriate structure for a Cobra-based CLI application,
and register it to its parent (default rootCmd).
If you want your command to be public, pass in the command name
with an initial uppercase letter.
Example: cobra add server -> resulting in a new cmd/server.go`,
	Run: push,
}

func push(cmd *cobra.Command, args []string) {

}
