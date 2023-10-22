# go-tinypng

**实现思路：**

* 1.递归遍历本地`tinypng-input`文件夹里的文件
* 2.获取遍历文件名的后缀和文件体积，格式必须是`.WebP` `.PNG ` `.JPEG`且文件体积低于5MB
* 3.每次上传文件随机生成一个IP地址（tinypng 对用户上传数量有限制，使用了 `X-Forwarded-For` 头绕过该限制）
* 4.处理返回数据拿到远程压缩图片地址
* 5.通过远程压缩地址下载图片至本地`tinypng-output`文件夹

#### **使用说明**

### **直接运行**

* 1.将需要压缩的图片放入项目根目录`tinypng-input`文件夹中

* 2.打开`CMD`，输入以下命令执行

```
go run main.go
```

Windows系统免安装客户端版：[Releases · liutongke/go-tinypng (github.com)](https://github.com/liutongke/go-tinypng/releases)

**声明：仅供学习讨论。**

## **免责声明**

该仓库仅用于学习，如有商业用途，请购买官方的 pro
版：[https://tinify.com/checkout/web-pro](https://tinify.com/checkout/web-pro)

**Implementation steps：**

* 1.Recursively traverse files in the local `tinypng-input` folder.
* 2.Get the file extension and file size of each file in the traversal, which must be in the formats
  of `.WebP` `.PNG ` `.JPEG` and have a file size less than 5MB.
* 3.Generate a random IP address each time a file is uploaded (as Tinypng limits the number of
  uploads per user, `X-Forwarded-For` headers are used to bypass this limit).
* 4.Process the returned data to obtain the remote compressed image address.
* 5.Download the image from the remote compressed address to the local `tinypng-output` folder.

#### **Instructions for use**

### **Run directly**

1. Place the images that need to be compressed into the `tinypng-input` folder in the project root
   directory.

2. Open `CMD` and enter the following command to execute.

```
go run main.go
```

Windows system portable client version (no need to
install)：[Releases · liutongke/go-tinypng (github.com)](https://github.com/liutongke/go-tinypng/releases)

**Disclaimer: For learning and discussion purposes only.**

## **Disclaimer**

This repository is only for learning purposes. If you want to use it for commercial purposes, please
purchase the official pro
version.：[https://tinify.com/checkout/web-pro](https://tinify.com/checkout/web-pro)