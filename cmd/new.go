/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/neelkarma/ramble/checks"
	"github.com/neelkarma/ramble/config"
	"github.com/neelkarma/ramble/constants"
	"github.com/neelkarma/ramble/editor"
	"github.com/neelkarma/ramble/enc"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new entry within your diary.",
	Run: func(cmd *cobra.Command, args []string) {
		if !checks.DiaryExists() {
			fmt.Println(constants.NoDiaryFoundMsg)
			return
		}

		if !checks.DiaryUnlocked() {
			fmt.Println(constants.DiaryNotUnlockedMsg)
			return
		}

		contents, _ := editor.Open("")
		if contents == "" {
			return
		}
		iv, _ := hex.DecodeString(config.ReadConfig().EncIv)
		encrypted, _ := enc.Encrypt(contents, iv)
		filename := time.Now().Format(constants.EntryFileTimeFormat)
		file, _ := os.Create(path.Join(constants.EntriesPath, filename))
		file.Write([]byte(encrypted))
		file.Close()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
