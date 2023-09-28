auth
https://api.real-exp-kasegi.com/auth?redirectUrl=${リダイレクト先}

google/callback
https://api.real-exp-kasegi.com/google/callback?redirectUrl={リダイレクト先（authを読んだときに指定したものと同じもの）}&code={googleからリダイレクトされたときにもらったcode}


createMyGameMatch
curl  -X POST -H "Content-Type: application/json" -d "{\"SessionToken\":\"$session\", \"GameMatch\":{\"GameId\" : \"hoge\", \"UserId\" :\"hoge\", \"CheckItems\":[{\"Title\":
 \"hoge\", \"IsChecked\" : true}]}}" localhost:8080/createMyGameMatch

 listMyGameMatch
 curl  -X POST -H "Content-Type: application/json" -d "{\"SessionToken\":\"$session\"}" localhost:8080/listMyGameMatch

createMyChecklistTemplate
curl  -X POST -H "Content-Type: application/json" -d "{\"SessionToken\":\"$session\", \"ChecklistTemplate\":{\"GameTitleId\" : \"hoge\", \"CheckItems\":[{\"Title\":
 \"hoge\", \"IsChecked\" : true}]}}" localhost:8080/createMyChecklistTemplate
 
listMyChecklistTemplate
curl  -X POST -H "Content-Type: application/json" -d "{\"SessionToken\":\"$session\"}" localhost:8080/listMyChecklistTemplate