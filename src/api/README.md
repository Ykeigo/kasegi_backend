createMyGameMatch
curl  -X POST -H "Content-Type: application/json" -d '{"SessionToken":"this is session", "GameMatch":{"GameId" : "hoge", "UserId" :"hoge", "CheckItems":[{"Title":
 "hoge", "IsChecked" : true}]}}' localhost:8080/createMyGameMatch

 listMyGameMatch
 curl  -X POST -H "Content-Type: application/json" -d '{"SessionToken":"this is session"}' localhost:8080/listMyGameMatch