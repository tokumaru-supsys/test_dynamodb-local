# テーブル作成・データ投入を実行するスクリプト
# 実行環境: Linux
# 作成者: S.Tokumaru
# 改訂履歴:
#   2022.12.03 初版
#!/bin/bash

# AWSの環境変数をセット
export AWS_ACCESS_KEY_ID=$1
export AWS_SECRET_ACCESS_KEY=$2
export AWS_DEFAULT_REGION=localhost

# Create "Forum" Table
aws dynamodb create-table \
    --endpoint-url http://localhost:8000 \
    --table-name Forum \
    --attribute-definitions AttributeName=Name,AttributeType=S \
    --key-schema AttributeName=Name,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

# Import data into "Forum"
aws dynamodb batch-write-item \
    --endpoint-url http://localhost:8000 \
    --request-items file://Forum.json

# Create "Thread" Table
aws dynamodb create-table \
    --endpoint-url http://localhost:8000 \
    --table-name Thread \
    --attribute-definitions AttributeName=ForumName,AttributeType=S AttributeName=Subject,AttributeType=S \
    --key-schema AttributeName=ForumName,KeyType=HASH AttributeName=Subject,KeyType=RANGE \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

# Import data into "Thread"
aws dynamodb batch-write-item \
    --endpoint-url http://localhost:8000 \
    --request-items file://Thread.json

# Create "Reply" Table
aws dynamodb create-table \
    --endpoint-url http://localhost:8000 \
    --table-name Reply \
    --attribute-definitions AttributeName=Id,AttributeType=S AttributeName=ReplyDateTime,AttributeType=S AttributeName=PostedBy,AttributeType=S \
    --key-schema AttributeName=Id,KeyType=HASH AttributeName=ReplyDateTime,KeyType=RANGE \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --local-secondary-indexes "IndexName=PostedBy-index,KeySchema=[{AttributeName=Id,KeyType=HASH},{AttributeName=PostedBy,KeyType=RANGE}],Projection={ProjectionType=KEYS_ONLY}"

# Import data into "Reply"
aws dynamodb batch-write-item \
    --endpoint-url http://localhost:8000 \
    --request-items file://Reply.json

# Create "ProductCatalog" Table
aws dynamodb create-table \
    --endpoint-url http://localhost:8000 \
    --table-name ProductCatalog \
    --attribute-definitions AttributeName=Id,AttributeType=N \
    --key-schema AttributeName=Id,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

# Import data into "ProductCatalog"
aws dynamodb batch-write-item \
    --endpoint-url http://localhost:8000 \
    --request-items file://ProductCatalog.json
