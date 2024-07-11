package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mycli",
		Short: " My CLI",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("requires at least one arg")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcom to My CLI")
		},
	}

	var helloCmd = &cobra.Command{
		Use:   "hello [name]",
		Short: "Hello",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("requires at least one arg")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello %v", args[0])
		},
	}

	rootCmd.AddCommand(helloCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
