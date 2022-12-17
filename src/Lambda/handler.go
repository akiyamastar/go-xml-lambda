//アプリケーションのディレクトリ直下はmainパッケージである必要があります。
package main
//①
import(
	“github.com/aws/aws-lambda-go/lambda”
	“TestLambda/greeting”
)
//②
func excuteFunction(){
	greeting.SayHello()
} 
//mainパッケージはmain関数を保持している必要があります。
//lambda.Start()はlambda関数での実装において記述してある必要があります。
//引数の関数名(excuteFunction)は任意です。
func main(){
	lambda.Start(excuteFunction)
} 
