package utils

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/lxn/win"
	"time"
	"unsafe"
)

func Listen() {
	// 打开剪贴板
	if !win.OpenClipboard(0) {
		fmt.Println("Failed to open clipboard")
		return
	}
	defer win.CloseClipboard()

	// 获取位图句柄
	hBitmap := win.HBITMAP(win.GetClipboardData(win.CF_BITMAP))
	if hBitmap == 0 {
		fmt.Println("No bitmap data found in clipboard")
		return
	}

	// 创建位图对象
	bitmap := hBitmap
	bitmapInfo := win.BITMAP{}
	if win.GetObject(win.HGDIOBJ(bitmap), unsafe.Sizeof(bitmapInfo), unsafe.Pointer(&bitmapInfo)) == 0 {
		fmt.Println("Failed to get bitmap information")
		return
	}

	// 打印位图信息
	fmt.Println("Bitmap Width:", bitmapInfo.BmWidth)
	fmt.Println("Bitmap Height:", bitmapInfo.BmHeight)
	fmt.Println("bitmapInfo.BmType:", bitmapInfo.BmType)
	fmt.Println("bitmapInfo.BmBits:", bitmapInfo.BmBits)
	fmt.Println("bitmapInfo.BmBitsPixel:", bitmapInfo.BmBitsPixel)
	// 在这里可以使用其他函数对位图进行处理，例如保存到文件或显示在窗口中

}

// 监听图片
func ListenPic() {
	// 打开剪贴板
	if !win.OpenClipboard(0) {
		fmt.Println("Failed to open clipboard")
		return
	}
	defer win.CloseClipboard()

	// 获取位图句柄
	hBitmap := win.HBITMAP(win.GetClipboardData(win.CF_BITMAP))
	if hBitmap == 0 {
		fmt.Println("No bitmap data found in clipboard")
		return
	}

	// 创建位图对象
	bitmap := win.HBITMAP(hBitmap)
	bitmapInfo := win.BITMAP{}
	if win.GetObject(win.HGDIOBJ(bitmap), unsafe.Sizeof(bitmapInfo), unsafe.Pointer(&bitmapInfo)) == 0 {
		fmt.Println("Failed to get bitmap information")
		return
	}

	// 打印位图信息
	fmt.Println("Bitmap Width:", bitmapInfo.BmWidth)
	fmt.Println("Bitmap Height:", bitmapInfo.BmHeight)
	fmt.Println(bitmapInfo)
	// 在这里可以使用其他函数对位图进行处理，例如保存到文件或显示在窗口中
}

// 监听文本内容
func ListenText() {
	// 创建一个用于接收剪贴板内容变化的通道
	clipboardCh := make(chan string)

	// 启动一个 goroutine 持续监听剪贴板变化
	go func() {
		// 初始剪贴板内容
		prevClipboard, _ := clipboard.ReadAll()

		for {
			// 读取当前剪贴板内容
			currentClipboard, _ := clipboard.ReadAll()

			// 比较当前剪贴板内容与上一次的内容是否不同
			if currentClipboard != prevClipboard {
				// 将变化的剪贴板内容发送到通道
				clipboardCh <- currentClipboard
				prevClipboard = currentClipboard
			}

			// 等待一段时间后继续检查剪贴板内容
			time.Sleep(1 * time.Second)
		}
	}()

	// 在主 goroutine 中读取剪贴板内容变化
	for {
		clipboardText := <-clipboardCh
		fmt.Println("剪贴板内容变化:", clipboardText)
	}
}
