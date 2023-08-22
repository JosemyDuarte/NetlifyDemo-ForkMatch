## ForkMatch

ForkMatch is a demo app with the required code to deploy a Golang endpoint on [Netlify](https://www.netlify.com/) using [Netlify Functions](https://www.netlify.com/products/functions/).

## Deploy

To deploy the app, you need to have a [Netlify](https://www.netlify.com/) account and a [GitHub](https://github.com/) account.
Make sure to follow the [required steps](https://www.netlify.com/blog/2016/09/29/a-step-by-step-guide-deploying-on-netlify/) to fork and deploy on your own account.

Define the following environment variables on your Netlify account:

```bash
ENVIRONMENT=aws
```

This is because under the hood, Netlify uses [AWS Lambda](https://aws.amazon.com/lambda/) to run their serverless functions, so we need to implement our 
code to be compatible with AWS Lambda.

## Run locally

To run the app locally, you need to have [Go](https://golang.org/) and [docker](https://www.docker.com/) installed.

```bash
make docker/run
```

## Endpoints

If you are running locally, the base URL is `http://localhost:8080` and if you are using Netlify, the base URL is `https://<your-netlify-app>.netlify.com/.netlify/functions/ForkMatch`.

Take a look at `app.go` to see the available endpoints. Keep in mind that there are different "App" depending on the environment. Meaning that the `App` for AWS Lambda is different than the `App` for running locally.
