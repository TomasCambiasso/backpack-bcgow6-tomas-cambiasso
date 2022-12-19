package product

import (
	"context"
	"db-implementation/domain"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var LastInsertId int = 0

type dynamoRepository struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

const (
	TABLE_NAME = "products"
)


func (r *dynamoRepository) Store(ctx context.Context, name string, ptype string, count int, price float64, warehouse_id int) (int, error) {
	LastInsertId += 1
	newProd := domain.Product{
		Id:           LastInsertId,
		Name:         name,
		Ptype:        ptype,
		Count:        count,
		Price:        price,
		Id_warehouse: warehouse_id,
	}
	av, err := dynamodbattribute.MarshalMap(&newProd)

	if err != nil {
		return 0, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.table),
	}

	_, err = r.dynamo.PutItemWithContext(ctx, input)

	if err != nil {
		return 0, err
	}

	return LastInsertId, nil
}

func (r *dynamoRepository) GetByID(ctx context.Context, id int) (domain.Product, error) {
	result, err := r.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
	})
	if err != nil {
		return domain.Product{}, err
	}

	if result.Item == nil {
		return domain.Product{}, nil
	}
	prodPtr, err := domain.ItemToProd(result.Item)
	return *prodPtr, err
}

func (r *dynamoRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	result, err := r.dynamo.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.table),
		Select:    aws.String("ALL_ATTRIBUTES"),
	},
	)
	if err != nil {
		return nil, err
	}

	if result.Items == nil {
		return nil, nil
	}
	var prodList []domain.Product
	for _, prod := range result.Items {
		prodPtr, err := domain.ItemToProd(prod)
		if err != nil {
			return nil, err
		}
		prodList = append(prodList, *prodPtr)
	}
	return prodList, err
}

func (r *repository) DeleteS(ctx context.Context, id int) error {

	stmt, err := r.db.Prepare(DELETE)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id) // retorna un sql.Result y un error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateS(ctx context.Context, name string, ptype string, count int, price float64, warehouse_id int, id int) (int, error) {
	stm, err := r.db.Prepare(UPDATE_PRODUCT) //preparamos la consulta
	if err != nil {
		return 0, err
	}
	defer stm.Close()
	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(name, ptype, count, price, warehouse_id, id)
	if err != nil {
		return 0, err
	}

	//obtenemos el ultimo id
	idAffected, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(idAffected), nil
}
