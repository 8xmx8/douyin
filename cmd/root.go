package cmd

import (
	"fmt"
	"github.com/Godvictory/douyin/cmd/flags"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "douyin",
	Short:   "你所热爱的，就是你的生活。",
	Long:    `抖音让每一个人看见并连接更大的世界，鼓励表达、沟通和记录，激发创造，丰富人们的精神世界，让现实生活更美好。`,
	Version: "v0.7.26",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	var baseDir, dataDir string
	var err error
	rootCmd.PersistentFlags().StringVar(&dataDir, "data", "data", "修改配置文件路径")
	rootCmd.PersistentFlags().BoolVar(&flags.Debug, "debug", false, "Debug 模式（更多的日志输出）")
	rootCmd.PersistentFlags().BoolVar(&flags.Dev, "dev", false, "开发环境")
	rootCmd.PersistentFlags().BoolVar(&flags.LogStd, "log-std", false, "日志强制打印到控制台")
	rootCmd.PersistentFlags().BoolVar(&flags.Memory, "memory", false, "使用内存数据库")

	cobra.OnInitialize(func() {
		flags.Pro = !flags.Dev
		// 获取可执行文件路径
		if baseDir, err = os.Executable(); err != nil {
			logrus.Fatal("无法获取可执行文件路径", err)
		}
		flags.ExPath = filepath.Dir(baseDir)
		flags.DataDir = filepath.Join(flags.ExPath, dataDir)
	})

	// no-completion
	rootCmd.AddCommand(&cobra.Command{
		Use:    "completion",
		Hidden: true,
	})
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}
