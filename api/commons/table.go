package commons

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
    region = "eu-west-2"
)

type Table struct {
    DynamoDbClient *dynamodb.Client
	TableName      *string
}

func GetTableInstance(context context.Context) Table {
    config, _ := awsConfig.LoadDefaultConfig(context, awsConfig.WithRegion(region))
    db := dynamodb.NewFromConfig(config)

    return Table {
        DynamoDbClient: db,
        TableName: aws.String(os.Getenv("CARDSTORAGE_TABLE_NAME")),
    }
}
