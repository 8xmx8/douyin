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
	/*
		O_RDONLY 打开只读文件
		O_WRONLY 打开只写文件
		O_RDWR 打开既可以读取又可以写入文件
		O_APPEND 写入文件时将数据追加到文件尾部
		O_CREATE 如果文件不存在，则创建一个新的文件
		0o666：表示文件权限的八进制数。0o666 表示文件所有者、所属组和其他用户都具有读写权限。
	*/
	/*
		在八进制表示法中，0o 前缀表示八进制数。数字 766 对应了文件权限 rw-rw-rw-。
		7 表示所有者（owner）具有读取、写入和执行权限。
		6 表示所属组（group）具有读取和写入权限。
		6 表示其他用户（others）具有读取和写入权限。
	*/
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
