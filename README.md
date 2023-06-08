# go-tinypng

**实现思路：**

* 1.监听windwos剪贴板.
* 2.获取剪贴板数据，判断是是否新数据且是图片。
* 3.将图片上传tinypng压缩。
* 4.通过远程压缩地址下载图片至本地`tinypng-output`文件夹

#### **使用说明**

电脑需要安装[snipaste](https://zh.snipaste.com/)截图软件

### **直接运行**

* 1.打开`CMD`进入项目根目录，输入以下命令执行

```
go run main.go
```

**声明：仅供学习讨论。**

## **免责声明**

该仓库仅用于学习，如有商业用途，请购买官方的 pro
版：[https://tinify.com/checkout/web-pro](https://tinify.com/checkout/web-pro)