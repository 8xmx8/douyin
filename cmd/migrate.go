package cmd

import (
	"github.com/Godvictory/douyin/internal/bootstrap"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "迁移数据库",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.InitConf()
		bootstrap.InitDb()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
