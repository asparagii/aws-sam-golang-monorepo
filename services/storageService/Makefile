LAMBDA_SAVEMESSAGE_NAME=saveMessage
LAMBDA_SAVEMESSAGE=dist/functions/$(LAMBDA_SAVEMESSAGE_NAME)

build: lambdas

lambdas: $(LAMBDA_SAVEMESSAGE)

$(LAMBDA_SAVEMESSAGE):
	GOOS=linux GOARCH=amd64 go build -o $(LAMBDA_SAVEMESSAGE) functions/$(LAMBDA_SAVEMESSAGE_NAME)/main.go

clean:
	rm -rf dist
	
