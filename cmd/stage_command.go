/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stageCmd represents the stage command
var stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Add files to the staging area",
	Long: `Add files to the staging area
	eg: got stage <file1> <file2> <file3>
It also works with wildcards, folders or globs
	eg: got stage *.txt
	    got stage folder/*
	    got stage .
Recursive behavior can be DISABLED with the -r flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stage called")
	},
}

func init() {
	rootCmd.AddCommand(stageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
