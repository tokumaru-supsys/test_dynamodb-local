rem テーブルを削除するスクリプト
rem 実行環境: Windows
rem 作成者: S.Tokumaru
rem 改訂履歴:
rem   2022.12.03 初版

@echo off

rem AWSの環境変数をセット
set AWS_ACCESS_KEY_ID=%1
set AWS_SECRET_ACCESS_KEY=%2
set AWS_DEFAULT_REGION=localhost

rem Delete "Forum" Table
aws dynamodb delete-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name Forum 

rem Delete "Thread" Table
aws dynamodb delete-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name Thread 

rem Delete "Reply" Table
aws dynamodb delete-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name Reply 

rem Delete "ProductCatalog" Table
aws dynamodb delete-table ^
    --endpoint-url http://localhost:8000 ^
    --table-name ProductCatalog

exit /b