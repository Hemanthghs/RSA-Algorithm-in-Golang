package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(Decoder([]int64{54, 115, 216}))
		encoded, _ := cmd.Flags().GetString("cipher")
		encoded_list := []int64{}
		s := strings.Split(encoded, " ")
		for _, val := range s {
			temp, _ := strconv.ParseInt(val, 10, 64)
			encoded_list = append(encoded_list, temp)
		}
		fmt.Println(Decoder(encoded_list))
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	decodeCmd.PersistentFlags().String("cipher", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
