# go-xml-lambda

## build 

```:sh
% docker exec -it go-xml-lambda_goapp_1 ash
# GOARCH=amd64 GOOS=linux go build -o handler handler.go
```
## zip for aws lambda

```:sh
% zip src/app/function.zip src/app/handler
```
