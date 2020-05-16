RESTful API для MyBank. В сервисе есть пользователи (id, balance). Можно:
Просмотреть баланс
````
Request:
GET /balance?user=101
 
Response:
200 OK
{“balance”:1000}
````
зачислять деньги на счет пользователям (создать пользователя, если еще не существует)
````
Request:
POST /deposit
{“user”:101, “amount”:100}

Response:
200 OK
```` 

снимать деньги со счетов
````
Request:
POST /withdraw
{“user”:101, “amount”:50} 

Response:
200 OK
````
 
переводить деньги от одного пользователя другому.
````
Request:
POST /transfer
{“from”:101, “to”:205, amount:25} 

Response:
200 OK
````
Данные необходимо хранить в postgresql. Реализовать валидацию
