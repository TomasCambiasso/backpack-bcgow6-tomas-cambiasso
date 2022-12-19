package domain

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Ptype        string  `json:"type"`
	Count        int     `json:"count"`
	Price        float64 `json:"price"`
	Id_warehouse int     `json:"id_warehouse"`
}

// Por que usar una funcion para un unmarshal?
func ItemToProd(input map[string]*dynamodb.AttributeValue) (*Product, error) {
	var item Product
	err := dynamodbattribute.UnmarshalMap(input, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
