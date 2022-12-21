package main

import(
	"github.com/aws/aws-lambda-go/lambda"
	// "app/database"
	"app/xml"
)
func excuteFunction(){
	xml.GetUrls()
} 
func main(){
	lambda.Start(excuteFunction)
} 
