# go-tinypng


**实现思路：**

*   1.递归遍历本地`tinypng-input`文件夹里的文件
*   2.获取遍历文件名的后缀和文件体积，格式必须是`.WebP` `.PNG ` `.JPEG`且文件体积低于5MB
*   3.每次上传文件随机生成一个IP地址（tinypng 对用户上传数量有限制，使用了 `X-Forwarded-For` 头绕过该限制）
*   4.处理返回数据拿到远程压缩图片地址
*   5.通过远程压缩地址下载图片至本地`tinypng-output`文件夹

#### **使用说明**

### **直接运行**

*  1.将需要压缩的图片放入项目根目录`tinypng-input`文件夹中

*  2.打开`CMD`，输入以下命令执行

```
go run main.go Tinypng.go
```
Windows系统免安装客户端版：[Releases · liutongke/go-tinypng (github.com)](https://github.com/liutongke/go-tinypng/releases)


**声明：仅供学习讨论。**

## **免责声明**

该仓库仅用于学习，如有商业用途，请购买官方的 pro 版：[https://tinify.com/checkout/web-pro](https://tinify.com/checkout/web-pro)

This Repo is only for study.
