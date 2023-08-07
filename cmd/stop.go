package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止守护进程",
	Run: func(cmd *cobra.Command, args []string) {
		stop()
	},
}

func stop() {
	initDaemon()
	if pid == -1 {
		fmt.Println("似乎还没有启动。尝试使用 `./douyin start` 启动服务器.")
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("无法按pid找到进程：%d，原因: %v", pid, process)
		return
	}
	err = process.Kill()
	if err != nil {
		fmt.Println("无法终止进程 %d: %v", pid, err)
	} else {
		fmt.Println("杀死进程: ", pid)
	}
	err = os.Remove(pidFile)
	if err != nil {
		fmt.Println("pid 文件未能正常删除")
	}
	pid = -1
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
