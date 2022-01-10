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

// getNs represents the getNameServer command
var getNs = &cobra.Command{
	Use:   "get-ns",
	Short: "Busca o NS do host DNS",
	Long:  `O comando get-ns busca os Name Servers de um host DNS.`,
	Run: func(cmd *cobra.Command, args []string) {
		dns, err := cmd.Flags().GetString("dns")
		if err != nil {
			log.Fatal(err)
		}

		servidores, err := net.LookupNS(dns)
		if err != nil {
			log.Fatal(err)
		}

		for _, servidor := range servidores {
			fmt.Println(servidor.Host)
		}
	},
}

func init() {
	getNs.Flags().StringP("dns", "d", "", "Host DNS para buscar o NS")
	rootCmd.AddCommand(getNs)
}
