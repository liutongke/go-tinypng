package main

import (
	"tinypng/src"
)

func main() {

	src.EchoSuccess(src.Welcome)
	src.DirExists(src.InputDir)
	src.DirExists(src.OutputDir)

	src.ScanAndDownload()
	src.EchoSuccess("=============下载结束=============")
}
