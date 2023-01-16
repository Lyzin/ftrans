/*
	Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Lyzin/ftrans/serve"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动文件服务",
	Long: `启动文件服务，默认会将当前目录作为根目录启动服务，端口是8000`,
	Run: func(cmd *cobra.Command, args []string) {
		serve.StartServe()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}