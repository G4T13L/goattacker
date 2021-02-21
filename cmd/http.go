package cmd

import (
	"fmt"
	"os"

	"github.com/G4T13L/goattacker/attacks"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http mode url",
	Short: "Attacks by http/https protocol",
	Long: `
Attacks by http/https protocol`,
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
			attacks.AuthAttack(urlSite, postData, userfile, passfile, proxyURL, workers, false)
		} else if mode == "fileSearch" {
			if wordsFile == "" {
				fmt.Println("Specify words File to iterate")
				os.Exit(1)
			}
			if wordReplace == "" {
				fmt.Println("Specify a word to replace in url")
				os.Exit(1)
			}
			attacks.FileAttack(urlSite, postData, wordReplace, wordsFile, extFile, proxyURL, workers, redir)
		} else if mode == "formLogin" {
			if postData == "" {
				fmt.Println("Specify the post data for the form")
				os.Exit(1)
			}
			if wordReplace == "" {
				fmt.Println("Specify a word or phrase to search if the response is incorrect")
				os.Exit(1)
			}
			attacks.FormAttack(urlSite, postData, userfile, passfile, wordReplace, proxyURL, workers, true, showhtml)
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
	redir       bool
	showhtml    bool
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
	httpCmd.Flags().StringVarP(&postData, "post", "", "", "Post data\n[formLogin]")
	httpCmd.Flags().StringVarP(&wordReplace, "word", "w", "", "[fileSearch] word to replace in url\n[formLogin] phrase to search to correct validation")
	httpCmd.Flags().StringVarP(&wordsFile, "File", "f", "", "[fileSearch] File of words to iterate")
	httpCmd.Flags().StringVarP(&extFile, "ext", "e", "", "[fileSearch] File of extensions wanted to iterate")
	// httpCmd.Flags().StringVarP(&wordPhrase, "Phrase", "", "", "phrase to search to stop")
	httpCmd.Flags().StringVarP(&userfile, "userfile", "U", "", "User file")
	httpCmd.Flags().StringVarP(&passfile, "passfile", "L", "", "Pass file")
	httpCmd.Flags().IntVarP(&workers, "nWorkers", "n", 9, "Number of rutines at the same time")
	httpCmd.Flags().BoolVar(&redir, "redirect", false, "Redirect if code of response is 302 or 301")
	httpCmd.Flags().BoolVar(&showhtml, "show", false, "show html response")

	usage := httpCmd.UsageString()
	// fIndex := strings.Index(usage, "\nFlags:")
	httpCmd.SetUsageTemplate(usage + `
Modes:
    [*] auth
    [*] fileSearch
    [*] formLogin
`)
	// httpCmd.SetHelpTemplate(httpCmd.HelpTemplate())

}
