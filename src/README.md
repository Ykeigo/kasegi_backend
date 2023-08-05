sqlboilerにリアルDBのパスワードを書くな！！ローカルDBのは書いてもいいよ


ここのコピペ：https://dev.classmethod.jp/articles/db-migrate-with-golang-migrate/

DB起動
docker run --rm -it -p 5432:5432 -e POSTGRES_PASSWORD=password --name postgres postgres

接続
psql -h localhost -p 5432 -U postgres -d postgres

# sql開いてるときの話
DB一覧
\l

DB切り替え
\c

テーブル一覧
\dt（らしい）

DB作成（多分他操作もだいたい同じ、小文字で書かないと動かない）
create database {DB名};
