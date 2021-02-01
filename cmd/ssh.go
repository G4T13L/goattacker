package cmd

import (
	"fmt"
	"github.com/G4T13L/goattacker/attacks"
	"github.com/spf13/cobra"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh userfile passfile serverIP",
	Short: "Ataques en el protocolo ssh, por defecto en el puerto 22",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Longer description of ssh help.`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		userfile = args[0]
		passfile = args[1]
		serverIP = args[2]
		if serverPort == ""{
			serverPort = "22"
		}
		fmt.Println("###########\nssh called\n###########", userfile, passfile, serverIP, serverPort)
		ssh_bruteforce_start(userfile, passfile, serverIP, serverPort)
	},
}

var (
    userfile, passfile, serverIP, serverPort string
)

func init() {
	rootCmd.AddCommand(sshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// sshCmd.Flags().StringVarP(&userfile, "userfile", "u", "", "User file (required)")
	// // sshCmd.MarkFlagRequired("userfile")
	// sshCmd.Flags().StringVarP(&passfile, "passfile", "l", "", "Pass file (required)")
	// sshCmd.Flags().StringVarP(&serverIP, "ip", "i", "", "Server IP (required)")
	sshCmd.Flags().StringVarP(&serverPort, "port", "p", "", "Server Port (optional)")
	
}
