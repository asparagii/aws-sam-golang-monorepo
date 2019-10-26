LAMBDA_INDEX_NAME=index
LAMBDA_INDEX=dist/functions/$(LAMBDA_INDEX_NAME) 

build: lambdas

lambdas: $(LAMBDA_INDEX)

$(LAMBDA_INDEX):
	GOOS=linux GOARCH=amd64 go build -o $(LAMBDA_INDEX) functions/$(LAMBDA_INDEX_NAME)/main.go

clean:
	rm -rf dist
