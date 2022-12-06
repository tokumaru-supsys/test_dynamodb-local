# ProductCatalogにGSIを追加するスクリプト
#   GSI1 --> ProductCategory, Title
#   GSI2 --> BicycleType, Price
# 実行環境: Windows
# 作成者: S.Tokumaru
# 改訂履歴:
#   2022.12.03 初版
#!/bin/bash

# AWSの環境変数をセット
export AWS_ACCESS_KEY_ID=$1
export AWS_SECRET_ACCESS_KEY=$2
export AWS_DEFAULT_REGION=localhost

# Add GSI1 
aws dynamodb update-table \
    --endpoint-url http://localhost:8000 \
    --table-name ProductCatalog \
    --attribute-definitions \
      AttributeName=ProductCategory,AttributeType=S \
      AttributeName=Title,AttributeType=S \
      AttributeName=BicycleType,AttributeType=S \
      AttributeName=Price,AttributeType=N \
    --global-secondary-index-updates file://gsi1.json
    

# Add GSI2
aws dynamodb update-table \
    --endpoint-url http://localhost:8000 \
    --table-name ProductCatalog \
    --attribute-definitions \
      AttributeName=ProductCategory,AttributeType=S \
      AttributeName=Title,AttributeType=S \
      AttributeName=BicycleType,AttributeType=S \
      AttributeName=Price,AttributeType=N \
    --global-secondary-index-updates file://gsi2.json