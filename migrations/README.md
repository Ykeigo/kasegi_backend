以下はローカルでdocker立ち上げたときの話
リモートのDBでやるなら接続情報が変わる
テーブル定義間違えてたときとかに気付くために一回手元でやったほうがいいと思う

postgleSQLのパスをexportしておく
postgressqlのURL（？）はこうらしい
一番最初のpostgres://っていうのはどういうこと？httpとかじゃなくてpostgres形式で通信します的なこと？
kasegiはDB名
export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/kasegi?sslmode=disable'

バージョンアップ
migrate -database $POSTGRESQL_URL -path example1 up ${上げるバージョン数（指定しない場合全部）}

バージョンダウン
migrate -database $POSTGRESQL_URL -path example1 down ${下げるバージョン数（指定しない場合全部）}

バージョンアップでミスった時
migrate -database $POSTGRESQL_URL -path example1 force ${VERSION}
${VERSION}の部分はエラーで出ているVERSIONを数字で指定。https://qiita.com/juchilian/items/0bfed79cc1229deb4c62


ローカルDB起動 -> srcのreadme

テーブル作成
createdb -h postgres -U postgres 






DB操作

ここのコピペ：https://dev.classmethod.jp/articles/db-migrate-with-golang-migrate/

DB起動
docker run --rm -it -p 5432:5432 -e POSTGRES_PASSWORD=password --name postgres postgres

接続
psql -h localhost -p 5432 -U postgres -d postgres

# sql開いてるときの話
DB一覧
\l

DB切り替え
\c ${DB名}

テーブル一覧
\dt（らしい）

DB作成（多分他操作もだいたい同じ、小文字で書かないと動かない）
create database {DB名};

ORM生成
sqlboiler psql