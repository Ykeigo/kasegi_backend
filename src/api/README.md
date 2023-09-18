createGameMatch
curl  -X POST -H "Content-Type: application/json" -d '{"SessionToken":"this is session", "GameMatch":{"GameId" : "hoge", "UserId" :"hoge", "CheckItems":[{"Title":
 "hoge", "IsChecked" : true}]}}' localhost:8080/createGameMatch