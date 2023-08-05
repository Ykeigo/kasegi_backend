sqlboilerにリアルDBのパスワードを書くな！！ローカルDBのは書いてもいいよ

DB起動」RUN docker run --rm -it -p 5432:5432 -e POSTGRES_PASSWORD=password --name postgres postgres