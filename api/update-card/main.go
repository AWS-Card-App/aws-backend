package main

import (
	"context"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"io.github.taz03/api/commons"
)

func UpdateCard(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    table := commons.GetTableInstance(context)

    params := request.QueryStringParameters
    id, _ := strconv.ParseInt(params["id"], 10, 64)
    card := commons.Card {
        Name: params["name"],
        Id: int64(id),
    }

    update := expression.Set(expression.Name("note"), expression.Value(request.Body))
    updateExpression, _ := expression.NewBuilder().WithUpdate(update).Build()

    table.DynamoDbClient.UpdateItem(context, &dynamodb.UpdateItemInput {
        TableName:                 table.TableName,
        Key:                       card.GetKey(),
        ExpressionAttributeNames:  updateExpression.Names(),
        ExpressionAttributeValues: updateExpression.Values(),
        UpdateExpression:          updateExpression.Update(),
    })

    return events.APIGatewayProxyResponse {
        StatusCode: 200,
    }, nil
}

func main() {
    lambda.Start(UpdateCard)
}

