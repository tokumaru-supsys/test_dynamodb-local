rem ProductCatalogにGSIを追加するスクリプト
rem   GSI1 --> ProductCategory, Title
rem   GSI2 --> BicycleType, Price
rem 実行環境: Windows
rem 作成者: S.Tokumaru
rem 改訂履歴:
rem   2022.12.03 初版

@echo off

rem AWSの環境変数をセット
set AWS_ACCESS_KEY_ID=%1
set AWS_SECRET_ACCESS_KEY=%2
set AWS_DEFAULT_REGION=localhost

rem Add GSI1 
aws dynamodb update-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name ProductCatalog ^
    --attribute-definitions ^
      AttributeName=ProductCategory,AttributeType=S ^
      AttributeName=Title,AttributeType=S ^
      AttributeName=BicycleType,AttributeType=S ^
      AttributeName=Price,AttributeType=N ^
    --global-secondary-index-updates file://gsi1.json

rem Add GSI2
aws dynamodb update-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name ProductCatalog ^
    --attribute-definitions ^
      AttributeName=ProductCategory,AttributeType=S ^
      AttributeName=Title,AttributeType=S ^
      AttributeName=BicycleType,AttributeType=S ^
      AttributeName=Price,AttributeType=N ^
    --global-secondary-index-updates file://gsi2.json
