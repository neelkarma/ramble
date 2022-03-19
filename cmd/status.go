/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/neelkarma/ramble/checks"
	"github.com/neelkarma/ramble/constants"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets whether your diary is locked or not.",
	Run: func(cmd *cobra.Command, args []string) {
		if !checks.DiaryExists() {
			fmt.Println(constants.NoDiaryFoundMsg)
			return
		}

		if checks.DiaryUnlocked() {
			fmt.Println("Diary is unlocked.")
		} else {
			fmt.Println("Diary is locked.")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
