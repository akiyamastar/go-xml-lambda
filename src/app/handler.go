package main

import(
	"github.com/aws/aws-lambda-go/lambda"
	"app/database"
	"app/xml"
)
func excuteFunction(){
	xmlUrls := xml.GetUrls()
	database.ReinsertNewUrls(xmlUrls)
} 
func main(){
	lambda.Start(excuteFunction)
} 
