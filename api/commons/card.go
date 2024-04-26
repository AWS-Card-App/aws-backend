package commons

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Card struct {
    Name string `json:"name" dynamodbav:"name"`
    Id   int64  `json:"id" dynamodbav:"id"`
    Note string `json:"note" dynamodbav:"note"`
}

func (card *Card) GetKey() map[string]types.AttributeValue {
    name, _ := attributevalue.Marshal(card.Name)
    id, _ := attributevalue.Marshal(card.Id)

    return map[string]types.AttributeValue {
        "name": name,
        "id": id,
    }
}
