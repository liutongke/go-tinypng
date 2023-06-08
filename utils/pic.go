package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"unsafe"

	"github.com/lxn/win"
)

func ListenImages() {
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

	// 获取设备上下文
	hdc := win.GetDC(0)
	defer win.ReleaseDC(0, hdc)

	// 创建位图缓冲区
	bitmapData := make([]byte, bitmapInfo.BmWidthBytes*bitmapInfo.BmHeight)

	// 获取位图像素数据
	bi := win.BITMAPINFO{}
	bi.BmiHeader.BiSize = uint32(unsafe.Sizeof(bi.BmiHeader))
	bi.BmiHeader.BiWidth = bitmapInfo.BmWidth
	bi.BmiHeader.BiHeight = -bitmapInfo.BmHeight
	bi.BmiHeader.BiPlanes = bitmapInfo.BmPlanes
	bi.BmiHeader.BiBitCount = bitmapInfo.BmBitsPixel
	bi.BmiHeader.BiCompression = win.BI_RGB
	ret := win.GetDIBits(hdc, hBitmap, 0, uint32(bitmapInfo.BmHeight), &bitmapData[0], &bi, win.DIB_RGB_COLORS)
	if ret == 0 {
		fmt.Println("Failed to get bitmap bits")
		return
	}

	// 在这里可以使用位图像素数据进行进一步处理
	// 例如，可以保存为文件或进行图像处理等操作
	// ...

	// 打印位图信息
	fmt.Println("Bitmap Width:", bitmapInfo.BmWidth)
	fmt.Println("Bitmap Height:", bitmapInfo.BmHeight)
	fmt.Println("Bitmap BitsPerPixel:", bitmapInfo.BmBitsPixel)
	fmt.Println("Bitmap Data Size:", len(bitmapData))

	saveImg(bitmapData, int(bitmapInfo.BmWidth), int(bitmapInfo.BmHeight))

}

func saveImg(bitmapData []byte, bitmapWidth, bitmapHeight int) {

	// 假设你已经获取到了位图像素数据，存储在 bitmapData 变量中

	// 假设你已经获取到了位图的宽度和高度，存储在 bitmapWidth 和 bitmapHeight 变量中

	// 创建一个 RGBA 图像对象
	img := image.NewRGBA(image.Rect(0, 0, bitmapWidth, bitmapHeight))

	// 将位图像素数据填充到图像对象中
	for y := 0; y < bitmapHeight; y++ {
		for x := 0; x < bitmapWidth; x++ {
			// 计算位图像素数据中的索引
			index := (y*bitmapWidth + x) * 4

			// 获取红、绿、蓝、透明度通道的值
			red := bitmapData[index]
			green := bitmapData[index+1]
			blue := bitmapData[index+2]
			alpha := bitmapData[index+3]

			// 设置像素点的颜色值
			img.SetRGBA(x, y, color.RGBA{red, green, blue, alpha})
		}
	}

	// 创建输出文件
	file, err := os.Create("bitmap.png")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	// 将图像数据保存为 PNG 格式
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Failed to encode image:", err)
		return
	}

	fmt.Println("Image saved to bitmap.png")
}
