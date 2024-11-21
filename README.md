# serverless

This repository contains the code for a aws lambda function which is triggered by sns and sends email to the user.

## Instructions to use the repo

### clone the repo

```bash
git clone https://github.com/csye-6225-gaurav/serverless.git
```

### Build the binary with name bootstrap

```bash
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go
```

### Zip the binary

```bash
zip myFunction.zip bootstrap 
```

use this zip to deploy your lambda function.

### Required ENV's

- Sendgrid API key
- URL