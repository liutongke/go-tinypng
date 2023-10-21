package main

import (
	"tinypng/src"
)

func main() {
	go src.UploadNotifyChan()

	src.EchoSuccess(src.Welcome)
	src.DirExists(src.InputDir)
	src.DirExists(src.OutputDir)

	src.WalkDir()
	src.Wg.Wait()
	src.EchoSuccess("=============下载结束=============")
}
