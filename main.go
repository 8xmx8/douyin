package main

import (
	"github.com/Godvictory/douyin/cmd"
)

func main() {
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	log.Fatalf("failed to create trace output file: %v", err)
	//}
	//defer func() {
	//	f.Close()
	//}()
	//// 2. trace绑定文件句柄
	//if err := trace.Start(f); err != nil {
	//	log.Fatalf("failed to start trace: %v", err)
	//}
	//defer trace.Stop()

	cmd.Execute()
}
