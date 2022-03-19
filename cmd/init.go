/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/erikgeiser/promptkit/textinput"
	"github.com/neelkarma/ramble/checks"
	"github.com/neelkarma/ramble/config"
	"github.com/neelkarma/ramble/constants"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes your diary in the current directory.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		if checks.DiaryExists() {
			fmt.Println("There is already a diary initialized in this directory.")
			return
		}

		fmt.Println("Choose a passphrase - make sure you don't forget this, as you cannot change or reset it in any way.")
		input := textinput.New("Passphrase:")
		input.Hidden = true
		passphrase, _ := input.RunPrompt()
		config.Init(passphrase)
		os.Mkdir(constants.EntriesPath, os.ModePerm)

		fmt.Println("Successfully initialized in current directory.")
		fmt.Println("Use \"ramble unlock\" and then \"ramble new\" to write your first entry.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
