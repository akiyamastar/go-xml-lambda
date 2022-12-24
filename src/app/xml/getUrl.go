package xml

import (
    "encoding/xml"
    "fmt"
    "io"
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

func GetUrls() []string {
    var jobUrls []string

    appEnv := os.Getenv("APP_ENV")
    filePath := "target.xml"
    if appEnv == "Lambda" {
        filePath = "/tmp/target.xml"
    }

    start := time.Now()
    fmt.Println("main start!")
    err := godotenv.Load()
    if err != nil {
        fmt.Println(err.Error())
    }
    xmlUrl := os.Getenv("XML_URL")
    if xmlUrl == "" {
        fmt.Println("環境変数 XML_URL が設定されていません。")
        return jobUrls
    }
    fmt.Println(xmlUrl)
    // data := httpGet(xmlUrl)
    xmlDownload(xmlUrl, filePath)
    xmlData := getXmls(filePath)

    result := XML{}
    // xmlErr := xml.Unmarshal([]byte(data), &result)
    xmlErr := xml.Unmarshal([]byte(xmlData), &result)
    
    if xmlErr != nil {
        fmt.Printf("error: %v", xmlErr)
        return jobUrls
    }

    fmt.Println("start!")
    for _, job := range result.Job {
        jobUrls = append(jobUrls, job.Url)
    }
    // fmt.Println(jobUrls)
    fmt.Println("end!")
    end := time.Now()
    fmt.Printf("xml url count: %d\n", len(result.Job))
    fmt.Printf("parse時間: %f秒\n", (end.Sub(start)).Seconds())

    return jobUrls
}

func httpGet(url string) string {
    fmt.Println("httpGet!")
    response, _ := http.Get(url)
    body, _ := ioutil.ReadAll(response.Body)
    defer response.Body.Close()
    return string(body)
}

func xmlDownload(url string, filePath string) string {
    fmt.Println("xmlDownload!")
    response, _ := http.Get(url)
    defer response.Body.Close()
    out, errOut := os.Create(filePath)
    if errOut != nil {
        return errOut.Error()
    }
    defer out.Close()
    io.Copy(out, response.Body)
    return ""
}

func getXmls(filePath string) string {
    xml, _ := ioutil.ReadFile(filePath)
    xmls := string(xml)
    return xmls
}
