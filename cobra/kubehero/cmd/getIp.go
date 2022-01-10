/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

// getIpCmd represents the getIp command
var getIpCmd = &cobra.Command{
	Use:   "get-ip",
	Short: "Busca IP host DNS",
	Long:  `O comando get-ip busca os endereços IP de um host DNS.`,
	Run: func(cmd *cobra.Command, args []string) {
		dns, err := cmd.Flags().GetString("dns")
		if err != nil {
			log.Fatal(err)
		}

		ips, err := net.LookupIP(dns)
		if err != nil {
			log.Fatal(err)
		}

		for _, ip := range ips {
			fmt.Println(ip)
		}
	},
}

func init() {
	getIpCmd.Flags().StringP("dns", "d", "", "Host DNS para buscar IP")
	rootCmd.AddCommand(getIpCmd)
}
