package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"io.github.taz03/api/commons"
)

func GetCards(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    table := commons.GetTableInstance(context)

    key := expression.Key("name").Equal(expression.Value(request.QueryStringParameters["name"]))
    keyExpression, _ := expression.NewBuilder().WithKeyCondition(key).Build()

    queryPaginator := dynamodb.NewQueryPaginator(table.DynamoDbClient, &dynamodb.QueryInput {
        TableName:                 table.TableName,
        ExpressionAttributeNames:  keyExpression.Names(),
        ExpressionAttributeValues: keyExpression.Values(),
        KeyConditionExpression:    keyExpression.KeyCondition(),
    })

    var cards []commons.Card
    for queryPaginator.HasMorePages() {
        response, _ := queryPaginator.NextPage(context)

        var cardPage []commons.Card
        attributevalue.UnmarshalListOfMaps(response.Items, &cardPage)
        cards = append(cards, cardPage...)
    }

    var body struct {
        Cards []commons.Card `json:"cards"`
    }
    body.Cards = append(body.Cards, cards...)
    response, _ := json.Marshal(body)

    return events.APIGatewayProxyResponse {
        Body:       string(response),
        StatusCode: 200,
    }, nil
}

func main() {
    lambda.Start(GetCards)
}
