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
	Wg        sync.WaitGroup
	Chan1     = make(chan *Res, 10000)
	InputDir  = "./tinypng-input"  //输出的文件夹
	OutputDir = "./tinypng-output" //输入的文件夹
	filePaths []*files

	readyDownloadNum = 0 //需要下载的文件数量
	progressNum      = 1

	Welcome = "\n               _   _                               \n              | | (_)                              \n  __ _  ___   | |_ _ _ __  _   _ _ __  _ __   __ _ \n / _` |/ _ \\  | __| | '_ \\| | | | '_ \\| '_ \\ / _` |\n| (_| | (_) | | |_| | | | | |_| | |_) | | | | (_| |\n \\__, |\\___/   \\__|_|_| |_|\\__, | .__/|_| |_|\\__, |\n  __/ |                     __/ | |           __/ |\n |___/                     |___/|_|          |___/ \n"
)

// 开始下载
func SendUpload(filePath, fileName string) {
	Wg.Add(1)
	err, _ := Uploads(filePath, fileName) //开始下载
	if err != nil {
		Chan1 <- &Res{
			IsSucc: false,
			Title:  fmt.Sprintf("下载失败(%d/%d):下载失败文件名:%s", progressNum, readyDownloadNum, fileName),
		}
	} else {
		Chan1 <- &Res{
			IsSucc: true,
			Title:  fmt.Sprintf("下载成功(%d/%d):下载后保存位置:%s", progressNum, readyDownloadNum, OutputDir+"/"+fileName),
		}
	}

	progressNum++
}

func UploadNotifyChan() {
	for data := range Chan1 {
		if data.IsSucc {
			EchoSuccess(data.Title)
			Wg.Done()
		}
		if !data.IsSucc {
			EchoError(data.Title)
			Wg.Done()
		}
	}
}
