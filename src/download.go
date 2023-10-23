package src

import (
	"fmt"
	"sync"
)

type Res struct {
	IsSucc bool
	Title  string
}

type files struct {
	Path string
	Name string
}

var (
	Wg                  sync.WaitGroup
	Chan1               = make(chan *Res, 10000)
	DownloadFailedQueue = make(map[string]*files) //下载失败等待重新下载的队列
	InputDir            = "./tinypng-input"       //需要压缩的文件夹位置
	OutputDir           = "./tinypng-output"      //压缩后的输入的文件夹
	filePaths           []*files

	readyDownloadNum = 0 //需要下载的文件数量
	progressNum      = 0

	Welcome = "\n               _   _                               \n              | | (_)                              \n  __ _  ___   | |_ _ _ __  _   _ _ __  _ __   __ _ \n / _` |/ _ \\  | __| | '_ \\| | | | '_ \\| '_ \\ / _` |\n| (_| | (_) | | |_| | | | | |_| | |_) | | | | (_| |\n \\__, |\\___/   \\__|_|_| |_|\\__, | .__/|_| |_|\\__, |\n  __/ |                     __/ | |           __/ |\n |___/                     |___/|_|          |___/ \n"
)

// SendUpload 开始下载
func SendUpload(filePath, fileName string) bool {

	_, err := UploadAndDownload(filePath, fileName) //开始下载

	if err != nil {
		DownloadFailedQueue[GetMD5Hash(filePath)] = &files{
			Path: filePath,
			Name: fileName,
		}

		EchoError(fmt.Sprintf("下载失败：%s进入队列稍后尝试重新下载", fileName))
		return false
	} else {
		progressNum++
		EchoSuccess(fmt.Sprintf("下载成功(%d/%d):下载后保存位置:%s", progressNum, readyDownloadNum, OutputDir+"/"+fileName))
		return true
	}

}
