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