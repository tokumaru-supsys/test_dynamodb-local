package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Forum のデータ構造
type Forum struct {
	Name             string `json:"Name" dynamodbav:"Name"`
	Category         string `json:"Category" dynamodbav:"Category"`
	Threads          uint   `json:"Threads" dynamodbav:"Threads"`
	Messages         uint   `json:"Messages" dynamodbav:"Messages"`
	Views            uint   `json:"Views" dynamodbav:"Views"`
	LastPostBy       string `json:"LastPostBy" dynamodbav:"LastPostBy"`
	LastPostDateTime string `json:"LastPostDateTime" dynamodbav:"LastPostDateTime"`
}

// Thread のデータ構造
type Thread struct {
	Id            string `json:"Id" dynamodbav:"Id"`
	ReplyDateTime string `json:"ReplyDateTime" dynamodbav:"ReplyDateTime"`
	Message       string `json:"Message" dynamodbav:"Message"`
	PostedBy      string `json:"PostedBy" dynamodbav:"PostedBy"`
}

// Reply のデータ構造
type Reply struct {
	Id            string `json:"Id" dynamodbav:"Id"`
	ReplyDateTime string `json:"ReplyDateTime" dynamodbav:"ReplyDateTime"`
	Message       string `json:"Message" dynamodbav:"Message"`
	PostedBy      string `json:"PostedBy" dynamodbav:"PostedBy"`
}

// ProductCatalog のデータ構造
type ProductCatalog struct {
	Id              uint     `json:"Id" dynamodbav:"Id"`
	Brand           string   `json:"Brand" dynamodbav:"Brand"`
	Description     string   `json:"Description" dynamodbav:"Description"`
	Price           uint     `json:"Price" dynamodbav:"Price"`
	Color           []string `json:"Color" dynamodbav:"Color"`
	ProductCategory string   `json:"ProductCategory" dynamodbav:"ProductCategory"`
	Title           string   `json:"Title" dynamodbav:"Title"`
	BicycleType     string   `json:"BicycleType" dynamodbav:"BicycleType"`
	InPublication   bool     `json:"InPublication" dynamodbav:"InPublication"`
	ISBN            string   `json:"ISBN" dynamodbav:"ISBN"`
	PageCount       uint     `json:"PageCount" dynamodbav:"PageCount"`
	Authors         []string `json:"Authors" dynamodbav:"Authors"`
	Dimensions      string   `json:"Dimensions" dynamodbav:"Dimensions"`
}

var (
	// コンテキスト
	ctx context.Context
	// dynamodbのクライアント
	client *dynamodb.Client
)

const (
	AWS_ACCESS_KEY_ID     = "jun9pag"
	AWS_SECRET_ACCESS_KEY = "00nbrm"
	AWS_DEFAULT_REGION    = "localhost"
)

// 初期化
func init() {

	ctx = context.TODO()

	// DynamoDB クライアントの生成
	cfg, err := config.LoadDefaultConfig(ctx,
		// CHANGE THIS TO ap-northeast-1 TO USE AWS proper
		config.WithRegion(AWS_DEFAULT_REGION),
		// Comment the below out if not using localhost
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000", SigningRegion: AWS_DEFAULT_REGION}, nil // The SigningRegion key was what7s was missing! D'oh.
			})),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.Credentials = credentials.NewStaticCredentialsProvider(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, "")
	})
}

func main() {
	scanByHashAndRangeKey()
	scanUsingLSI()
	fetchById()
	queryUsingGSI_index_ProductCategoryTitle()
	queryUsingGSI_index_BicycleTypePrice()
}

// PartitionKey & RangeKey で Scan してみる
func scanByHashAndRangeKey() {
	scanFilter := expression.Name("Id").Equal(expression.Value("Amazon DynamoDB#DynamoDB Thread 1")).And(expression.Name("ReplyDateTime").BeginsWith("2015-09-15"))
	expr, err := expression.NewBuilder().WithFilter(scanFilter).Build()
	if err != nil {
		log.Fatalf("creating expression builder failed. %s", err)
		return
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String("Reply"),
	}

	result, err := client.Scan(ctx, input)
	if err != nil {
		log.Fatalf("scanning Reply failed. %s", err)
		return
	}

	var reply []Reply
	err = attributevalue.UnmarshalListOfMaps(result.Items, &reply)
	if err != nil {
		log.Fatalf("unmarshal reply failed. %s", err)
		return
	}
	fmt.Printf("fetch result... %+v\n", reply)
}

// LSIを使って Scan してみる
func scanUsingLSI() {
	scanFilter := expression.Name("Id").Equal(expression.Value("Amazon DynamoDB#DynamoDB Thread 1")).And(expression.Name("PostedBy").Equal(expression.Value("User B")))
	expr, err := expression.NewBuilder().WithFilter(scanFilter).Build()
	if err != nil {
		log.Fatalf("creating expression builder failed. %s", err)
		return
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String("Reply"),
		IndexName:                 aws.String("PostedBy-index"),
	}

	result, err := client.Scan(ctx, input)
	if err != nil {
		log.Fatalf("scanning Reply failed. %s", err)
		return
	}

	var reply []Reply
	err = attributevalue.UnmarshalListOfMaps(result.Items, &reply)
	if err != nil {
		log.Fatalf("unmarshal reply failed. %s", err)
		return
	}
	fmt.Printf("fetch result... %+v\n", reply)
}

// 普通に query を投げてみる
func fetchById() {

	keycond := expression.Key("Id").Equal(expression.Value(103))
	expr, err := expression.NewBuilder().WithKeyCondition(keycond).Build()
	if err != nil {
		log.Fatalf("creating expression builder failed. %s", err)
		return
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String("ProductCatalog"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	result, err := client.Query(ctx, input)
	if err != nil {
		log.Fatalf("Query failed. %s", err)
		return
	}

	var productCatalog []ProductCatalog
	err = attributevalue.UnmarshalListOfMaps(result.Items, &productCatalog)
	if err != nil {
		log.Fatalf("unmarshal ProductCatalog failed. %s", err)
		return
	}
	fmt.Printf("fetch result... %+v\n", productCatalog)

}

// GSIを使って Query してみる
func queryUsingGSI_index_ProductCategoryTitle() {
	keycond := expression.Key("ProductCategory").Equal(expression.Value("Bicycle")).And(expression.Key("Title").BeginsWith("18-"))
	proj := expression.NamesList(expression.Name("Id"), expression.Name("Title"), expression.Name("Brand"), expression.Name("BicycleType"), expression.Name("Description"))
	expr, err := expression.NewBuilder().WithKeyCondition(keycond).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("creating expression builder failed. %s", err)
		return
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String("ProductCatalog"),
		IndexName:                 aws.String("ProductCategory-Title-index"),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	result, err := client.Query(ctx, input)
	if err != nil {
		log.Fatalf("Query failed. %s", err)
		return
	}

	var productCatalog []ProductCatalog
	err = attributevalue.UnmarshalListOfMaps(result.Items, &productCatalog)
	if err != nil {
		log.Fatalf("unmarshal ProductCatalog failed. %s", err)
		return
	}
	fmt.Printf("fetch result... %+v\n", productCatalog)
}

// GSIを使って Query してみる
func queryUsingGSI_index_BicycleTypePrice() {
	keycond := expression.Key("BicycleType").Equal(expression.Value("Road")).And(expression.Key("Price").LessThanEqual(expression.Value(200)))
	proj := expression.NamesList(expression.Name("Id"), expression.Name("Title"), expression.Name("Brand"), expression.Name("Price"), expression.Name("Description"))
	expr, err := expression.NewBuilder().WithKeyCondition(keycond).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("creating expression builder failed. %s", err)
		return
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String("ProductCatalog"),
		IndexName:                 aws.String("BicycleType-Price-index"),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	result, err := client.Query(ctx, input)
	if err != nil {
		log.Fatalf("Query failed. %s", err)
		return
	}

	var productCatalog []ProductCatalog
	err = attributevalue.UnmarshalListOfMaps(result.Items, &productCatalog)
	if err != nil {
		log.Fatalf("unmarshal ProductCatalog failed. %s", err)
		return
	}
	fmt.Printf("fetch result... %+v\n", productCatalog)
}
