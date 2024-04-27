package main

import (
	"context"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"io.github.taz03/api/commons"
)

func DeleteCard(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    table := commons.GetTableInstance(context)
    
    params := request.QueryStringParameters
    id, _ := strconv.ParseInt(params["id"], 10, 64)

    card := commons.Card {
        Name: params["name"],
        Id:   int64(id),
    }

    table.DynamoDbClient.DeleteItem(context, &dynamodb.DeleteItemInput {
        Key:       card.GetKey(),
        TableName: table.TableName,
    })

    return events.APIGatewayProxyResponse {
        StatusCode: 200,
    }, nil
}

func main() {
    lambda.Start(DeleteCard)
}

