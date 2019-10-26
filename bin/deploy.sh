#!/bin/bash

STAGE=${STAGE:-dev}
OUTPUT_BUCKET=sam-artifacts-asparagii
STACK_NAME=aws-sam-go-monorepo-$STAGE

sam package --template-file template.yaml --output-template-file packaged.template.yaml --s3-bucket $OUTPUT_BUCKET 
sam deploy --template-file packaged.template.yaml --stack-name $STACK_NAME --capabilities CAPABILITY_AUTO_EXPAND CAPABILITY_IAM --parameter-overrides Stage=$STAGE
