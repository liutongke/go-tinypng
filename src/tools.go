package src

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// DirExists 判断文件夹是否存在
func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, err
}

// ScanAndDownload 扫描文件夹中的文件
func ScanAndDownload() {
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
		os.Exit(1)
	}

	EchoSuccess(fmt.Sprintf("总共扫描到%d个文件", readyDownloadNum))

	EchoSuccess("=============开始下载=============")
	for _, filePath := range filePaths {
		SendUpload(filePath.Path, filePath.Name)
	}

	//错误重试
	retryDownloads()
}

func retryDownloads() {
	if len(DownloadFailedQueue) > 0 {
		EchoError(fmt.Sprintf("%d个文件下载失败,重试中...", len(DownloadFailedQueue)))
		for hashId, filePath := range DownloadFailedQueue {
			if SendUpload(filePath.Path, filePath.Name) {
				delete(DownloadFailedQueue, hashId)
			}
		}
	}

	if len(DownloadFailedQueue) > 0 {
		retryDownloads()
	}
}

func EchoSuccess(str string) {
	fmt.Printf("\033[1;32;40m%s\033[0m\n", str)
}

func EchoError(str string) {
	fmt.Printf("\033[1;31;40m%s\033[0m\n", str)
}

func GetMD5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
