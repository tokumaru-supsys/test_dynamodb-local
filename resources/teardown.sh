# テーブルを削除するスクリプト
# 実行環境: Linux
# 作成者: S.Tokumaru
# 改訂履歴:
#   2022.12.03 初版
#!/bin/bash

# AWSの環境変数をセット
export AWS_ACCESS_KEY_ID=$1
export AWS_SECRET_ACCESS_KEY=$2
export AWS_DEFAULT_REGION=localhost

# Delete "Forum" Table
aws dynamodb delete-table /
    --endpoint-url http://localhost:8000 /
    --table-name Forum 

# Delete "Thread" Table
aws dynamodb delete-table /
    --endpoint-url http://localhost:8000 /
    --table-name Thread 

# Delete "Reply" Table
aws dynamodb delete-table /
    --endpoint-url http://localhost:8000 /
    --table-name Reply 

# Delete "ProductCatalog" Table
aws dynamodb delete-table /
    --endpoint-url http://localhost:8000 /
    --table-name ProductCatalog

