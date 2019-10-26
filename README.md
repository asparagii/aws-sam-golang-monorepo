# AWS SAM Golang Monorepo

An example golang monorepo with multiple services using AWS SAM

```bash
.
├── bin
│   └── deploy.sh                   <-- SAM deploy script (package+deploy)
├── go.mod                          <-- Single go module
├── go.sum
├── Makefile                        <-- Build whole project
├── README.md
├── services                        <-- Every subapplication
│   ├── storageService              <-- Simple S3-based storage service
│   │   ├── functions               <-- Lambda functions code
│   │   │   └── saveMessage         <-- Simple lambda
│   │   │       └── main.go         <-- saveMessage code
│   │   ├── Makefile                <-- Subapplication makefile
│   │   └── template.yaml           <-- Subapplication sam template
│   └── webService                  <-- Another service (with api gateway)
│       ├── functions               
│       │   └── index               <-- GET /
│       │       └── main.go
│       ├── Makefile                <-- Subapplication makefile
│       └── template.yaml           <-- Subapplication sam template
└── template.yaml                   <-- Main application sam template
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Golang](https://golang.org)

## Setup process
##### IMPORTANT
Make sure to change `bin/deploy.sh`'s variable `OUTPUT_BUCKET` to your own artifacts dump s3 bucket

### Building
`make build`

### Deploy
`make deploy`
