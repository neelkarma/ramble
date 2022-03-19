/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/neelkarma/ramble/checks"
	"github.com/neelkarma/ramble/config"
	"github.com/neelkarma/ramble/constants"
	"github.com/neelkarma/ramble/editor"
	"github.com/neelkarma/ramble/enc"
	"github.com/neelkarma/ramble/entries"
	"github.com/spf13/cobra"
)

// openCmd represents the view command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens an entry from your diary.",
	Run: func(cmd *cobra.Command, args []string) {
		if !checks.DiaryExists() {
			fmt.Println(constants.NoDiaryFoundMsg)
			return
		}

		if !checks.DiaryUnlocked() {
			fmt.Println(constants.DiaryNotUnlockedMsg)
			return
		}

		filemap := entries.GetAll()

		if len(filemap) == 0 {
			fmt.Println("You have no entries! Use \"ramble new\" to create a new entry.")
			return
		}

		displayValues := make([]string, 0, len(filemap))
		for k := range filemap {
			displayValues = append(displayValues, k)
		}

		input := selection.New("Pick an entry:", selection.Choices(displayValues))
		choice, _ := input.RunPrompt()

		filename := filemap[choice.String]

		if delete, _ := cmd.Flags().GetBool("delete"); delete {
			os.Remove(path.Join(constants.EntriesPath, filename))
			fmt.Println("Entry deleted successfully.")
			return
		}

		file, _ := os.Open(path.Join(constants.EntriesPath, filename))
		defer file.Close()
		bytes, _ := ioutil.ReadAll(file)
		iv, _ := hex.DecodeString(config.ReadConfig().EncIv)
		contents, _ := enc.Decrypt(string(bytes), iv)
		newContents, _ := editor.Open(contents)
		if newContents == contents {
			return
		}
		encrypted, _ := enc.Encrypt(newContents, iv)
		file.Write([]byte(encrypted))
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	openCmd.Flags().BoolP("delete", "d", false, "Deletes the selected entry.")
}
