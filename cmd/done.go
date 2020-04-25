package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

// doCmd represents the do command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Remove the task from your task list ",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := strconv.Atoi(task)
		if err != nil {
			log.Println(err)
		}
		collection.Find(bson.M{}, bson.M{}, &tasks)
		noOfTasks := len(tasks)
		if id == 0 || id > noOfTasks {
			fmt.Println("There is no such task")
		}
		for i, v := range tasks {
			if i == id-1 {
				fmt.Printf("Marking the task \"%s\" as done.\nRemoving it from your list...\n", v.Task)
				collection.DeleteOne(bson.M{"task": v.Task})
			}
		}
		fmt.Println()
		fmt.Println()
		collection.Find(bson.M{}, bson.M{}, &tasks)
		if tasks != nil {
			fmt.Println("You have the following tasks left:")
			for i, v := range tasks {
				fmt.Println(i+1, v.Task)
			}
		} else {
			fmt.Println("Well done!!\nYou have completed all of your tasks")

		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

}
