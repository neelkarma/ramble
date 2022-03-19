/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/neelkarma/ramble/checks"
	"github.com/neelkarma/ramble/constants"
	"github.com/spf13/cobra"
)

// lockCmd represents the lock command
var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Locks your diary.",
	Run: func(cmd *cobra.Command, args []string) {
		if !checks.DiaryExists() {
			fmt.Println(constants.NoDiaryFoundMsg)
			return
		}

		if !checks.DiaryUnlocked() {
			fmt.Println("Diary is already locked.")
			return
		}

		os.Remove(constants.KeyPath)
		fmt.Println("Diary successfully locked.")
	},
}

func init() {
	rootCmd.AddCommand(lockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
