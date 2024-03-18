package todos

import (
	"fmt"

	"github.com/priyank-amagi/toolbox/config"
	"github.com/priyank-amagi/toolbox/lib/http_helper"
	"github.com/priyank-amagi/toolbox/pkg/todos"
	"github.com/spf13/cobra"
)

var (
	todosCount int
	todoType   string
	todoFilter string
	todoSvc    todos.ITodoSvc
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List is a command that lists todos based on provided flags",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		todoSvc = todos.NewTodoSvc(config.GetJsonApiHost(), http_helper.NewHttpSvc())
		listTodos()
	},
}

func init() {
	listCmd.Flags().IntVarP(&todosCount, "count", "c", 20, "The number of todos to be listed")
	listCmd.Flags().StringVarP(&todoType, "type", "t", "", "The todo-type to be listed (complete/incomplete)")
	listCmd.Flags().StringVarP(&todoFilter, "filter", "f", "", "The filtered todos to be listed (even/odd)")

	// Here you will define your flags and configuration settings.
	TodosCmd.AddCommand(listCmd)
}

func listTodos() {
	if todos, err := todoSvc.GetTodos(); err == nil {
		count := 0
		for i, todo := range todos {
			if count >= todosCount {
				break
			}
			// check if todo matches the filter criteria
			if (todoType == "" || (todoType == "complete" && todo.Completed) || (todoType == "incomplete" && !todo.Completed)) &&
				(todoFilter == "" || (todoFilter == "even" && (i+1)%2 == 0) || (todoFilter == "odd" && (i+1)%2 != 0)) {
				count++
				fmt.Printf("%2d. Title: %-65s | Completed: %t\n", count, todo.Title, todo.Completed)
			}
		}
	} else {
		fmt.Println(err.Error())
	}
}
