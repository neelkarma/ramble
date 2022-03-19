/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/erikgeiser/promptkit/textinput"
	"github.com/neelkarma/ramble/checks"
	"github.com/neelkarma/ramble/config"
	"github.com/neelkarma/ramble/constants"
	"github.com/neelkarma/ramble/enc"
	"github.com/neelkarma/ramble/passhash"
	"github.com/spf13/cobra"
)

// unlockCmd represents the unlock command
var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlocks your diary.",
	Run: func(cmd *cobra.Command, args []string) {
		if !checks.DiaryExists() {
			fmt.Println(constants.NoDiaryFoundMsg)
			return
		}
		if checks.DiaryUnlocked() {
			fmt.Println("Diary is already unlocked.")
			return
		}
		input := textinput.New("Enter passphrase: ")
		input.Hidden = true
		passphrase, _ := input.RunPrompt()
		cfg := config.ReadConfig()
		isValid := passhash.ValidatePassphrase(passphrase, cfg.Passhash)
		if !isValid {
			fmt.Println("Invalid passphrase.")
			return
		}
		salt, _ := hex.DecodeString(cfg.EncSalt)
		key := enc.DeriveKey(passphrase, salt)
		hexkey := hex.EncodeToString(key)
		file, _ := os.Create(constants.KeyPath)
		file.Write([]byte(hexkey))
		file.Close()
		fmt.Println("Diary unlocked successfully.")
	},
}

func init() {
	rootCmd.AddCommand(unlockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unlockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unlockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
