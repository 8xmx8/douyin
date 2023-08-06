package utils

import "os"

// Exists 确定指定的文件是否存在
func Exists(name string) bool {
	stat, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return !stat.IsDir()
}

// ExistsDir 确定指定的目录是否存在
func ExistsDir(name string) bool {
	stat, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return stat.IsDir()
}
