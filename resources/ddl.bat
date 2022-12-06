rem テーブル作成・データ投入を実行するスクリプト
rem 実行環境: Windows
rem 作成者: S.Tokumaru
rem 改訂履歴:
rem   2022.12.03 初版

@echo off

rem AWSの環境変数をセット
set AWS_ACCESS_KEY_ID=%1
set AWS_SECRET_ACCESS_KEY=%2
set AWS_DEFAULT_REGION=localhost

rem Create "Forum" Table
aws dynamodb create-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name Forum ^
    --attribute-definitions AttributeName=Name,AttributeType=S ^
    --key-schema AttributeName=Name,KeyType=HASH ^
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

rem Import data into "Forum"
aws dynamodb batch-write-item ^
    --endpoint-url http://localhost:8000 ^
    --request-items file://Forum.json

rem Create "Thread" Table
aws dynamodb create-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name Thread ^
    --attribute-definitions AttributeName=ForumName,AttributeType=S AttributeName=Subject,AttributeType=S ^
    --key-schema AttributeName=ForumName,KeyType=HASH AttributeName=Subject,KeyType=RANGE ^
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

rem Import data into "Thread"
aws dynamodb batch-write-item ^
    --endpoint-url http://localhost:8000 ^
    --request-items file://Thread.json

rem Create "Reply" Table
aws dynamodb create-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name Reply ^
    --attribute-definitions AttributeName=Id,AttributeType=S AttributeName=ReplyDateTime,AttributeType=S AttributeName=PostedBy,AttributeType=S ^
    --key-schema AttributeName=Id,KeyType=HASH AttributeName=ReplyDateTime,KeyType=RANGE ^
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 ^
    --local-secondary-indexes "IndexName=PostedBy-index,KeySchema=[{AttributeName=Id,KeyType=HASH},{AttributeName=PostedBy,KeyType=RANGE}],Projection={ProjectionType=KEYS_ONLY}"

rem Import data into "Reply"
aws dynamodb batch-write-item ^
    --endpoint-url http://localhost:8000 ^
    --request-items file://Reply.json

rem Create "ProductCatalog" Table
aws dynamodb create-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name ProductCatalog ^
    --attribute-definitions AttributeName=Id,AttributeType=N ^
    --key-schema AttributeName=Id,KeyType=HASH ^
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

rem Import data into "ProductCatalog"
aws dynamodb batch-write-item ^
    --endpoint-url http://localhost:8000 ^
    --request-items file://ProductCatalog.json

exit /b