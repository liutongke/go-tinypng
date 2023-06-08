package utils

import (
	"fmt"
	"os"
)

var Conf = map[string]string{"input": "tinypng-input", "output": "tinypng-output"}

func GetConfDir() map[string]string {
	return Conf
}

func Init() {
	for _, fileName := range GetConfDir() {
		mkDir(fileName)
	}
}
func mkDir(dirPath string) {
	// 检查文件夹是否存在
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败:", err)
			return
		}
		fmt.Println("创建文件夹成功:", dirPath)
	} else if err != nil {
		// 其他错误，无法确定文件夹是否存在
		fmt.Println("其他错误，无法确定文件夹是否存在:", err)
		return
	} else {
		// 文件夹已经存在
		fmt.Println("文件夹已经存在:", dirPath)
	}
}
