package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Tinypng struct {
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

	url := "https://tinypng.com/web/shrink"
	method := "POST"

	payload, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	defer payload.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	req.Header.Add("accept", " */*")
	req.Header.Add("accept-encoding", " gzip, deflate, br")
	req.Header.Add("accept-language", " zh-CN,zh;q=0.9")
	req.Header.Add("content-length", " 1288562")
	req.Header.Add("content-type", " image/png")
	req.Header.Add("dnt", " 1")
	req.Header.Add("origin", " https://tinypng.com")
	req.Header.Add("referer", " https://tinypng.com/")
	req.Header.Add("sec-fetch-dest", " empty")
	req.Header.Add("sec-fetch-mode", " cors")
	req.Header.Add("sec-fetch-site", " same-origin")
	req.Header.Add("user-agent", " Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	req.Header.Add("X-Forwarded-For", genIpaddr())

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	tinypngData := Tinypng{}
	jsonerr := json.Unmarshal(body, &tinypngData)
	if jsonerr != nil {
		fmt.Println("json数据有误", jsonerr, string(body))
		return jsonerr, nil
	}

	download(tinypngData.Output.Url, fileName)

	return nil, &Output{
		Url:  tinypngData.Output.Url,
		Name: fileName,
	}
}

func download(url string, filename string) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("http.Get -> %v", err)
		return
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("ioutil.ReadAll -> %s", err.Error())
		return
	}
	defer res.Body.Close()

	if err = os.WriteFile(outputDir+"/"+filename, data, 0777); err != nil {
		log.Println("Error Saving:", filename, err)
	} else {
		//log.Println("Saved:", filename)
	}
}

// {"input":{"size":1288562,"type":"image/png"},"output":{"size":485959,"type":"image/png","width":4167,"height":4167,"ratio":0.3771,"url":"https://tinypng.com/web/output/uucc3u30rnfn9xzdb5zf00ab58z19h2k"}}
type Output struct {
	Ratio float64 `json:"ratio"`
	Url   string  `json:"url"`
	Size  int     `json:"size"`
	Name  string  `json:"name"`
}

func genIpaddr() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}
