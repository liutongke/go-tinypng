package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	inputDir  = "./tinypng-input"  //输出的文件夹
	outputDir = "./tinypng-output" //输入的文件夹
	filePaths = []*files{}

	readyDownloadNum = 0 //需要下载的文件数量
	l                = []string{}
	progressNum      = 1
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
	DirExists(inputDir)
	DirExists(outputDir)
	echoSuccess("-----------------开始扫描文件夹------------------------")
	walkDir()
}

// 扫描文件夹中的文件
func walkDir() {
	err := filepath.Walk(inputDir,
		func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(info.Name(), path.Ext(filePath))
			//abs, _ := filepath.Abs(path)
			//fmt.Println(abs, info.Size())
			if path.Ext(filePath) == ".png" ||
				path.Ext(filePath) == ".jpg" ||
				path.Ext(filePath) == ".webp" ||
				path.Ext(filePath) == ".jpeg" {

				readyDownloadNum++
				fileAbsPath, _ := filepath.Abs(filePath)

				//fmt.Println(readyDownloadNum, filePath, path.Ext(filePath))
				filePaths = append(filePaths, &files{
					Path: fileAbsPath,
					Name: info.Name(),
				})
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	echoSuccess(fmt.Sprintf("总共扫描到%d个文件", readyDownloadNum))
	echoSuccess("-----------------扫描结束------------------------")
	//echoSuccess("-----------------开始压缩------------------------")
	echoSuccess("-----------------开始下载------------------------")
	for _, filePath := range filePaths {
		SendUpload(filePath.Path, filePath.Name)
	}

	echoSuccess("-----------------下载结束------------------------")
}

// 开始下载
func SendUpload(filePath, fileName string) {
	uploadErr, data := Uploads(filePath, fileName) //开始压缩
	if uploadErr != nil {
		echoError(fmt.Sprintf("压缩失败(%d/%d):%s,压缩失败文件名:%s", progressNum, readyDownloadNum, data.Url, fileName))
	} else {
		echoSuccess(fmt.Sprintf("压缩成功(%d/%d):压缩后保存位置:%s", progressNum, readyDownloadNum, filePath))
	}

	l = append(l, data.Url)
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
