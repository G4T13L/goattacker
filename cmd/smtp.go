package cmd

import (
	"fmt"
	"os"
	"github.com/G4T13L/goattacker/attacks"
	"github.com/spf13/cobra"
)

var smtpCmd = &cobra.Command{
	Use:   "smtp",
	Short: "Attacks by stmp protocol, default port is 25",
	Long: `Attacks by stmp protocol, default port is 25

	serverIP: the IP of the machine to attack`,
	Args: cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		serverIP = args[0]
		if (userfile == "" || passfile == ""){
			// panic ("Incorrect number of arguments in userfile or passfile")
			fmt.Println("Incorrect number of arguments in userfile or passfile")
			os.Exit(1)

		}
		if(serverPort == ""){
			serverPort = "25"
		}
		fmt.Println("Smtp called: ", userfile, passfile, serverIP, serverPort)
		attacks.Smtp_attack_start(userfile, passfile, serverIP, serverPort, workers)
	},
}

func init() {
	rootCmd.AddCommand(smtpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// smtpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// smtpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	smtpCmd.Flags().StringVarP(&userfile, "userfile", "U", "", "User file (required)")
	// sshCmd.MarkFlagRequired("userfile")
	smtpCmd.Flags().StringVarP(&passfile, "passfile", "L", "", "Pass file (required)")
	// sshCmd.Flags().StringVarP(&serverIP, "ip", "i", "", "Server IP (required)")
	smtpCmd.Flags().StringVarP(&serverPort, "port", "p", "", "Server Port")
	smtpCmd.Flags().IntVarP(&workers, "workers", "w", 9, "Number of rutines at the same time")
	// smtpCmd.Flags().IntVarP(&timeOut, "timeOut", "t", 5, "Time limit")
}