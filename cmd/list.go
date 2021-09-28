/*
Copyright © 2021 Terry Osayawe <terrymarcy2000@yahoo.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all task",
	RunE: func(cmd *cobra.Command, args []string) error {
		flag, err := cmd.Flags().GetBool("todo_only")
		if err != nil {
			return fmt.Errorf("could not parse flags: %v", err)
		}
		return list(context.Background(), flag)
	},
}

func init() {
	listCmd.Flags().BoolP("todo_only", "t", false, "show only task that are not completed yet")
	rootCmd.AddCommand(listCmd)
}

func list(ctx context.Context, filterDone bool) error {

	l, err := client.List(ctx, new(emptypb.Empty))
	if err != nil {
		return fmt.Errorf("could not fetch task: %v", err)
	}

	for _, task := range l.Tasks {
		if task.Done {
			if filterDone {
				continue
			}
			fmt.Printf("✅")
		} else {
			fmt.Printf("❌")
		}

		fmt.Printf(" %s\n", task.Text)
	}
	return nil
}
