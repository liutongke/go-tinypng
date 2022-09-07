package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

var (
	inputDir  = "./tinypng-input"  //输出的文件夹
	outputDir = "./tinypng-output" //输入的文件夹
	filePaths []*files

	readyDownloadNum = 0 //需要下载的文件数量
	progressNum      = 1

	welcome = "\n               _   _                               \n              | | (_)                              \n  __ _  ___   | |_ _ _ __  _   _ _ __  _ __   __ _ \n / _` |/ _ \\  | __| | '_ \\| | | | '_ \\| '_ \\ / _` |\n| (_| | (_) | | |_| | | | | |_| | |_) | | | | (_| |\n \\__, |\\___/   \\__|_|_| |_|\\__, | .__/|_| |_|\\__, |\n  __/ |                     __/ | |           __/ |\n |___/                     |___/|_|          |___/ \n"
)

type files struct {
	Path string
	Name string
}

func echoSuccess(str string) {
	fmt.Printf("\033[1;32;40m%s\033[0m\n", str)
}

func echoError(str string) {
	fmt.Printf("\033[1;31;40m%s\033[0m\n", str)
}

func main() {
	echoSuccess(welcome)
	DirExists(inputDir)
	DirExists(outputDir)
	walkDir()
}

// 扫描文件夹中的文件
func walkDir() {
	err := filepath.Walk(inputDir,
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
		echoError("扫描文件夹失败")
		os.Exit(0)
	}

	echoSuccess(fmt.Sprintf("总共扫描到%d个文件", readyDownloadNum))

	echoSuccess("=============开始下载=============")
	for _, filePath := range filePaths {
		SendUpload(filePath.Path, filePath.Name)
	}

	echoSuccess("=============下载结束=============")
}

// 开始下载
func SendUpload(filePath, fileName string) {
	err, _ := Uploads(filePath, fileName) //开始下载
	if err != nil {
		echoError(fmt.Sprintf("下载失败(%d/%d):下载失败文件名:%s", progressNum, readyDownloadNum, fileName))
	} else {
		echoSuccess(fmt.Sprintf("下载成功(%d/%d):下载后保存位置:%s", progressNum, readyDownloadNum, outputDir+"/"+fileName))
	}

	progressNum++
}

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
