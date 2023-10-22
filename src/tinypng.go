package src

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.57"
)

// TinyPngResponse 定义 TinyPNG 响应的结构
type TinyPngResponse struct {
	Input struct {
		Size int    `json:"size"`
		Type string `json:"type"`
	} `json:"input"`
	Output struct {
		Size   int     `json:"size"`
		Type   string  `json:"type"`
		Width  int     `json:"width"`
		Height int     `json:"height"`
		Ratio  float64 `json:"ratio"`
		Url    string  `json:"url"`
	} `json:"output"`
}

// UploadAndDownload 上传文件到 TinyPNG 并下载压缩后的图片
func UploadAndDownload(filePath, fileName string) (*Output, error) {
	url := "https://tinypng.com/backend/opt/shrink"
	method := "POST"

	payload, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer payload.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	setRequestHeaders(req)
	setRandomIPHeaders(req)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	TinyPngResponseData := TinyPngResponse{}
	err = json.Unmarshal(body, &TinyPngResponseData)
	if err != nil {
		return nil, err
	}

	err = downloadImage(TinyPngResponseData.Output.Url, fileName)
	if err != nil {
		return nil, err
	}

	return &Output{
		Url:  TinyPngResponseData.Output.Url,
		Name: fileName,
	}, nil
}

// downloadImage 下载图片
func downloadImage(url string, filename string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(OutputDir, filename), data, 0777)
	if err != nil {
		return err
	}

	return nil
}

// setRequestHeaders 设置请求头
func setRequestHeaders(req *http.Request) {
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Content-Type", "image/jpeg")
	req.Header.Add("Content-Length", "39761")
	req.Header.Add("Dnt", "1")
	req.Header.Add("Origin", "https://tinypng.com")
	req.Header.Add("Referer", "https://tinypng.com/")
	req.Header.Add("Sec-Ch-Ua", "\"Chromium\";v=\"118\", \"Microsoft Edge\";v=\"118\", \"Not=A?Brand\";v=\"99\"")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", UserAgent)
}

// setRandomIPHeaders 设置随机 IP 地址头部
func setRandomIPHeaders(req *http.Request) {
	ip := generateRandomIPv4().String()
	req.Header.Add("X-Forwarded-For", ip)
	req.Header.Add("X-Real-IP", ip)
	req.Header.Add("Remote-Addr", ip)
}

// Output 保存压缩图片的信息
type Output struct {
	Ratio float64 `json:"ratio"`
	Url   string  `json:"url"`
	Size  int     `json:"size"`
	Name  string  `json:"name"`
}

// generateRandomIPv4 生成随机的 IPv4 地址
func generateRandomIPv4() net.IP {
	ip := make(net.IP, 4)
	// TODO: 根据你的需求生成随机 IP 地址
	// 这里使用默认的随机生成方法，你可以根据需要进行替换
	return ip
}
