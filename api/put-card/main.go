package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"io.github.taz03/api/commons"
)

func PutCard(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    table := commons.GetTableInstance(context)

    id := time.Now().UnixMilli()
    card := commons.Card {
        Name: request.QueryStringParameters["name"],
        Id:   id,
        Note: request.Body,
    }

    item, _ := attributevalue.MarshalMap(card)
    table.DynamoDbClient.PutItem(context, &dynamodb.PutItemInput {
        TableName: table.TableName,
        Item: item,
    })

    var body struct {
        Id int64 `json:"id"`
    }
    body.Id = id
    response, _ := json.Marshal(body)

    return events.APIGatewayProxyResponse {
        Body:       string(response),
        StatusCode: 200,
    }, nil
}

func main() {
    lambda.Start(PutCard)
}
