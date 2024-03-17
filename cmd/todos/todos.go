package todos

import (
	"github.com/spf13/cobra"
)

// TodosCmd represents the todos command
var TodosCmd = &cobra.Command{
	Use:   "todos",
	Short: "Todos is a palette that contains todos related commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
