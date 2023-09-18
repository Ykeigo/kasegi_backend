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