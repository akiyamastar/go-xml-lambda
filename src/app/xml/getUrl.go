package xml

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
    "os"
    "github.com/joho/godotenv"
)

type XML struct {
    Job []struct {
        Date  string `xml:"date"`
        Title string `xml:"title"`
        Url  string `xml:"url"`
    } `xml:"job"`
}

func GetUrls() {
    start := time.Now()
    fmt.Println("main start!")
    err := godotenv.Load()
    if err != nil {
        fmt.Println(err.Error())
    }
    xmlUrl := os.Getenv("XML_URL")
    if xmlUrl == "" {
		fmt.Println("環境変数 XML_URL が設定されていません")
		return
	}
	data := httpGet(xmlUrl)
    result := XML{}
    xmlErr := xml.Unmarshal([]byte(data), &result)
    if xmlErr != nil {
        fmt.Printf("error: %v", xmlErr)
        return
    }
    var jobUrls []string

    fmt.Println("start!")
    for _, job := range result.Job {
        jobUrls = append(jobUrls, job.Url)
    }

    fmt.Println(jobUrls)
    fmt.Println("end!")
    end := time.Now()
    fmt.Printf("xml数: %d\n", len(result.Job))
    fmt.Printf("parse時間: %f秒\n", (end.Sub(start)).Seconds())
    
}

func httpGet(url string) string {
    response, _ := http.Get(url)
    body, _ := ioutil.ReadAll(response.Body)
    defer response.Body.Close()
    return string(body)
}
