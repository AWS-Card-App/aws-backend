package main

import (
	"context"
	"fmt"
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
        Id: id,
        Name: request.PathParameters["name"],
        Note: request.Body,
    }

    item, _ := attributevalue.MarshalMap(card)
    table.DynamoDbClient.PutItem(context, &dynamodb.PutItemInput {
        TableName: table.TableName,
        Item: item,
    })

    return events.APIGatewayProxyResponse {
        StatusCode: 200,
        Body: fmt.Sprintf("Added card at id %v", id),
    }, nil
}

func main() {
    lambda.Start(PutCard)
}
