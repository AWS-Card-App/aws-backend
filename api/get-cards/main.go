package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"io.github.taz03/api/commons"
)

func GetCard(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

    var ss []string
    for _, v := range cards {
        ss = append(ss, v.String())
    }

    return events.APIGatewayProxyResponse {
        StatusCode: 200,
        Body: fmt.Sprintf("{\"cards\":[%v]}", strings.Join(ss, ",")),
    }, nil
}

func main() {
    lambda.Start(GetCard)
}
