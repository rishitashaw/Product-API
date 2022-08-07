package adapter

import (
	"github/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type Interface interface {

}  

func (db *Database) Health() bool{

}

func (db *Database) NewAdapter()  {
}

func (db *Database) FindAll(){

}

func (db *Database) FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) {
	
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input:=&dynamodb.GetItemInput{
		TableName: aws.String("tableName"),
		Key: conditionParsed,
	}

	return db.connection.GetItem(input)
}

func (db *Database) CreateOrUpdate(){

}

func (db *Database) Delete(){}