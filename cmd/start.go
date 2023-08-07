package cmd

import (
	"github.com/Godvictory/douyin/cmd/flags"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动守护进程",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	initServer()
	initDaemon()
	if pid != -1 {
		_, err := os.FindProcess(pid)
		if err == nil {
			log.Info("抖音已经启动了, pid ", pid)
			return
		}
	}
	args := os.Args
	args[1] = "server"
	cmd := &exec.Cmd{
		Path: args[0],
		Args: args,
		Env:  os.Environ(),
	}
	stdout, err := os.OpenFile(filepath.Join(flags.ExPath, "data", "start.log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatal(os.Getpid(), ": 无法打开启动日志文件:", err)
	}
	cmd.Stderr = stdout
	cmd.Stdout = stdout
	err = cmd.Start()
	if err != nil {
		log.Fatal("未能启动子进程: ", err)
	}
	log.Infof("成功启动 pid: %d", cmd.Process.Pid)
	err = os.WriteFile(pidFile, []byte(strconv.Itoa(cmd.Process.Pid)), 0o666)
	if err != nil {
		log.Warn("pid 未能正常写入文件，您可能无法使用 `./douyin stop` 停止程序.")
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
}
