package commons

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Card struct {
    Id   int64  `json:"id" dynamodbav:"id"`
    Name string `json:"name" dynamodbav:"name"`
    Note string `json:"note" dynamodbav:"note"`
}

func (card *Card) GetKey() map[string]types.AttributeValue {
    id, _ := attributevalue.Marshal(card.Id)
    name, _ := attributevalue.Marshal(card.Name)

    return map[string]types.AttributeValue {
        "id": id,
        "name": name,
    }
}
