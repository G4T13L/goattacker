package cmd

import (
	"fmt"
	"os"
	"github.com/G4T13L/goattacker/attacks"
	"github.com/spf13/cobra"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh serverIP",
	Short: "Attacks by ssh protocol, default port is 22",
	Long: `Attacks by ssh protocol, default port is 22
	
	// userfile: file containing username list
	// passfile: file containing password list
	serverIP: the IP of the machine to attack
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverIP = args[0]
		if (userfile == "" || passfile == ""){
			fmt.Println("Incorrect number of arguments in userfile or passfile")
			os.Exit(1)
		}
		if(serverPort == ""){
			serverPort = "22"
		}
		fmt.Println("ssh called ", userfile, passfile, serverIP, serverPort)
		// workers = 9
		attacks.Ssh_bruteforce_start(userfile, passfile, serverIP, serverPort, workers)
	},
}

var (
    userfile, passfile, serverIP, serverPort string
    workers int
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

	sshCmd.Flags().StringVarP(&userfile, "userfile", "U", "", "User file (required)")
	// sshCmd.MarkFlagRequired("userfile")
	sshCmd.Flags().StringVarP(&passfile, "passfile", "L", "", "Pass file (required)")
	// sshCmd.Flags().StringVarP(&serverIP, "ip", "i", "", "Server IP (required)")
	sshCmd.Flags().StringVarP(&serverPort, "port", "p", "", "Server Port")
	sshCmd.Flags().IntVarP(&workers, "workers", "w", 9, "Number of rutines at the same time")
	
}
