package main

import (
	"context"
	"guihub.com/guillospy92/utilsgolang/service-connect-documentdb/v1/mconnectv2"
	"guihub.com/guillospy92/utilsgolang/service-connect-documentdb/v1/utils"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

/**
basado en este tutorial
https://www.serverless.com/framework/docs/providers/aws/examples/hello-world/go
*/

// mongodb://gromo:Guillermo920517@docdb-2022-01-08-14-59-33.cluster-cgvrlgagsdcn.us-east-1.docdb.amazonaws.com/prime?replicaSet=rs0&readpreference=secondaryPreferred

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {

	//get dir
	err := utils.ReadingFiles()

	if err != nil {
		resp := Response{
			StatusCode:      http.StatusInternalServerError,
			IsBase64Encoded: false,
			Body:            err.Error(),
			Headers: map[string]string{
				"Content-Type":           "application/json",
				"X-MyCompany-Func-Reply": "v1-handler",
			},
		}
		return resp, nil
	}

	// init connect documentDb
	paramConnection := &mconnectv2.ParamConnectionDocumentDB{
		ConnectTimeOut:  mconnectv2.DefaultTimeOut,
		Cluster:         "docdb-2022-01-09-17-41-27.cluster-cgvrlgagsdcn.us-east-1.docdb.amazonaws.com",
		ReadPreference:  mconnectv2.DefaultReadPreference,
		UserName:        "gromo",
		Password:        "Guillermo920517",
		DBName:          "prime",
		Port:            27017,
		TLS:             true,
		NameCertificate: "rds-combined-ca-bundle.pem",
	}

	_, err = mconnectv2.NewCreateConnectionDocumentDB(paramConnection)

	if err != nil {
		resp := Response{
			StatusCode:      http.StatusInternalServerError,
			IsBase64Encoded: false,
			Body:            err.Error(),
			Headers: map[string]string{
				"Content-Type":           "application/json",
				"X-MyCompany-Func-Reply": "v1-handler",
			},
		}
		return resp, nil
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "hello hello connect",
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "v1-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
