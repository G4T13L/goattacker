package cmd

import (
	"fmt"
	"os"
	"github.com/G4T13L/goattacker/attacks"
	"github.com/spf13/cobra"
)

// smbCmd represents the smb command
var smbCmd = &cobra.Command{
	Use:   "smb",
	Short: "Attacks by smb protocol, default port is 445",
	Long: `Attacks by ftp protocol, default port is 445

	serverIP: the IP of the machine to attack`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverIP = args[0]
		if (userfile == "" || passfile == ""){
			fmt.Println("Incorrect number of arguments in userfile or passfile")
			os.Exit(1)
		}
		if(serverPort == ""){
			serverPort = "445"
		}
		fmt.Println("smb called", userfile, passfile, serverIP, serverPort)
		attacks.Smb_attack_start(userfile, passfile, serverIP, serverPort, workers)
	},
}

func init() {
	rootCmd.AddCommand(smbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// smbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// smbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	smbCmd.Flags().StringVarP(&userfile, "userfile", "U", "", "User file (required)")
	// sshCmd.MarkFlagRequired("userfile")
	smbCmd.Flags().StringVarP(&passfile, "passfile", "L", "", "Pass file (required)")
	// sshCmd.Flags().StringVarP(&serverIP, "ip", "i", "", "Server IP (required)")
	smbCmd.Flags().StringVarP(&serverPort, "port", "p", "", "Server Port")
	smbCmd.Flags().IntVarP(&workers, "workers", "w", 9, "Number of rutines at the same time")
}