package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	userfile, passfile, serverIP, serverPort string
	workers                                  int
)

var rootCmd = &cobra.Command{
	Use:   "goattacker",
	Short: "A brief description of your application",
	Long: `a tool made in golang to attack different protocols. 
	protocols worked:
	[x] ssh
	[x] ftp
	[x] smb
	[x] smtp
	[x] http y https
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {
// 	cobra.OnInitialize(initConfig)

// 	Here you will define your flags and configuration settings.
// 	Cobra supports persistent flags, which, if defined here,
// 	will be global for your application.

// 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goattacker.yaml)")

// 	Cobra also supports local flags, which will only run
// 	when this action is called directly.
// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// Search config in home directory with name ".goattacker" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".goattacker")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
