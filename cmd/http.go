package cmd

import (
	"fmt"
	"os"

	"github.com/G4T13L/goattacker/attacks"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http mode",
	Short: "Attacks by http/https protocol",
	Long: `
Attacks by http/https protocol
	3 modes to use
	[-] auth
	[-] fileSearch
	[-] formLogin

serverIP: the IP of the machine to attack`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mode := args[0]
		urlSite = args[1]
		if urlSite == "" {
			fmt.Println("You need to enter an url")
			os.Exit(1)
		}

		fmt.Println("http/https called: ", urlSite, mode)
		if mode == "auth" {
			attacks.AuthAttack(urlSite, postData, userfile, passfile, proxyURL, workers)
		} else if mode == "fileSearch" {

		} else if mode == "formLogin" {

		} else {
			fmt.Println("Bad mode operation")
			os.Exit(1)
		}

	},
}

var (
	urlSite     string
	proxyURL    string
	postData    string
	wordReplace string
	wordsFile   string
	extFile     string
	wordPhrase  string
)

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// [-] auth
	// [-] fileSearch
	// [-] formLogin
	httpCmd.Flags().StringVarP(&proxyURL, "proxy", "p", "", "Proxy url: \"proxyUrl:port\"")
	httpCmd.Flags().StringVarP(&postData, "post", "", "", "Post data")
	httpCmd.Flags().StringVarP(&wordReplace, "word", "w", "", "[fileSearch] word to replace in url\n[formLogin] phrase to search to correct validation")
	httpCmd.Flags().StringVarP(&wordsFile, "File", "f", "", "[fileSearch] File of words to iterate")
	httpCmd.Flags().StringVarP(&extFile, "ext", "e", "", "[fileSearch] File of extensions wanted to iterate")
	// httpCmd.Flags().StringVarP(&wordPhrase, "Phrase", "", "", "phrase to search to stop")
	httpCmd.Flags().StringVarP(&userfile, "userfile", "U", "", "User file")
	httpCmd.Flags().StringVarP(&passfile, "passfile", "L", "", "Pass file")
	httpCmd.Flags().IntVarP(&workers, "nWorkers", "n", 9, "Number of rutines at the same time")
}
