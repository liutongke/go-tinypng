package src

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

// DirExists 判断文件夹是否存在
func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			return true, nil
		}
	}
	return false, err
}

// 扫描文件夹中的文件
func WalkDir() {
	err := filepath.Walk(InputDir,
		func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if (path.Ext(filePath) == ".png" ||
				path.Ext(filePath) == ".jpg" ||
				path.Ext(filePath) == ".webp" ||
				path.Ext(filePath) == ".jpeg") && info.Size() <= 5242880 {

				readyDownloadNum++
				fileAbsPath, _ := filepath.Abs(filePath)

				filePaths = append(filePaths, &files{
					Path: fileAbsPath,
					Name: info.Name(),
				})
			}
			return nil
		})
	if err != nil {
		EchoError("扫描文件夹失败")
		os.Exit(0)
	}

	EchoSuccess(fmt.Sprintf("总共扫描到%d个文件", readyDownloadNum))

	EchoSuccess("=============开始下载=============")
	for _, filePath := range filePaths {
		go SendUpload(filePath.Path, filePath.Name)
		time.Sleep(1 * time.Second)
	}

}

func EchoSuccess(str string) {
	fmt.Printf("\033[1;32;40m%s\033[0m\n", str)
}

func EchoError(str string) {
	fmt.Printf("\033[1;31;40m%s\033[0m\n", str)
}
