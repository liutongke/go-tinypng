package src

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
)

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

func Uploads(filePath, fileName string) (error, *Output) {

	url := "https://tinypng.com/backend/opt/shrink"
	method := "POST"

	payload, err := os.Open(filePath)
	if err != nil {
		return err, nil
	}
	defer payload.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err, nil
	}

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
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.57")
	ip := generateRandomIPv4().String()
	fmt.Printf("ip 地址:%s \n", ip)
	req.Header.Add("X-Forwarded-For", ip)
	req.Header.Add("X-Real-IP", ip)
	req.Header.Add("Remote-Addr", ip)
	//req.Header.Add("accept", " */*")
	//req.Header.Add("accept-encoding", " gzip, deflate, br")
	//req.Header.Add("accept-language", " zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	////req.Header.Add("content-length", " 1288562")
	//req.Header.Add("content-type", " image/png")
	////req.Header.Add("dnt", " 1")
	//req.Header.Add("origin", " https://tinypng.com")
	//req.Header.Add("referer", " https://tinypng.com/")
	////req.Header.Add("referer", " https://tinypng.com/cn/")
	//req.Header.Add("sec-fetch-dest", " empty")
	//req.Header.Add("sec-fetch-mode", " cors")
	//req.Header.Add("sec-fetch-site", " same-origin")
	//req.Header.Add("user-agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.57")
	//req.Header.Add("X-Forwarded-For", generateRandomIPv4().String())

	res, err := client.Do(req)

	if err != nil {
		return err, nil
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err, nil
	}

	TinyPngResponseData := TinyPngResponse{}
	err = json.Unmarshal(body, &TinyPngResponseData)

	if err != nil {
		return err, nil
	}

	err = download(TinyPngResponseData.Output.Url, fileName)
	if err != nil {
		return err, nil
	}

	return nil, &Output{
		Url:  TinyPngResponseData.Output.Url,
		Name: fileName,
	}
}

func download(url string, filename string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = os.WriteFile(OutputDir+"/"+filename, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

type Output struct {
	Ratio float64 `json:"ratio"`
	Url   string  `json:"url"`
	Size  int     `json:"size"`
	Name  string  `json:"name"`
}

func generateRandomIPv4() net.IP {
	ip := make(net.IP, 4)
	rand.Read(ip)
	return ip
}
