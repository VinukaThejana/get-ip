package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "get-ip",
	Short: "Get the IP address assigned to the interface that connects to the internet",
	Run: func(_ *cobra.Command, _ []string) {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		localAddress := conn.LocalAddr().(*net.UDPAddr)
		fmt.Println(localAddress.IP.String())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
