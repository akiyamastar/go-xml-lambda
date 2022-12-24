package main

import(
	// "github.com/aws/aws-lambda-go/lambda"
	"app/database"
	"app/xml"
	// "app/logic"
)
func excuteFunction() {
	xmlUrls := xml.GetUrls()
	database.ReinsertNewUrls(xmlUrls)
	// dbUrls := database.GetUrls()
	// logic.CalcDiff(xmlUrls, dbUrls)
}
func main(){
	excuteFunction()
} 
