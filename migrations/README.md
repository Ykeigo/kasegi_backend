バージョンアップ
migrate -database ${POSTGRESQL_URL} -path migrations/example1 up ${上げるバージョン数（指定しない場合全部）}

バージョンダウン
migrate -database ${POSTGRESQL_URL} -path migrations/example1 down ${下げるバージョン数（指定しない場合全部）}

バージョンアップでミスった時
migrate -database ${POSTGRESQL_URL} -path migrations/example1 force ${VERSION}
${VERSION}の部分はエラーで出ているVERSIONを数字で指定。https://qiita.com/juchilian/items/0bfed79cc1229deb4c62


ローカルDB起動 -> test-appのreadme

テーブル作成
createdb -h postgres -U postgres 