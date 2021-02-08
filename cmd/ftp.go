package cmd

import (
	"fmt"
	"os"
	"github.com/G4T13L/goattacker/attacks"
	"github.com/spf13/cobra"
)

// var (
//     userfile, passfile, serverIP, serverPort string
//     workers int
// )
var timeOut int

// ftpCmd represents the ftp command
var ftpCmd = &cobra.Command{
	Use:   "ftp",
	Short: "Attacks by ftp protocol, default port is 21",
	Long: `Attacks by ftp protocol, default port is 21

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
			serverPort = "21"
		}
		if(timeOut == 0){
			timeOut = 5
		}
		fmt.Println("Ftp called: ", userfile, passfile, serverIP, serverPort, timeOut)
		attacks.Ftp_attack_start(userfile, passfile, serverIP, serverPort, workers, timeOut)
	},
}

func init() {
	rootCmd.AddCommand(ftpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ftpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ftpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ftpCmd.Flags().StringVarP(&userfile, "userfile", "U", "", "User file (required)")
	// sshCmd.MarkFlagRequired("userfile")
	ftpCmd.Flags().StringVarP(&passfile, "passfile", "L", "", "Pass file (required)")
	// sshCmd.Flags().StringVarP(&serverIP, "ip", "i", "", "Server IP (required)")
	ftpCmd.Flags().StringVarP(&serverPort, "port", "p", "", "Server Port")
	ftpCmd.Flags().IntVarP(&workers, "workers", "w", 9, "Number of rutines at the same time")
	ftpCmd.Flags().IntVarP(&timeOut, "timeOut", "t", 5, "time limit")
}
