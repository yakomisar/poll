# Тестовое задание - Poll service

- Written in Go
- Stores data in PostgreSQL with Gorm

# Quickstart (Mac OS)
```shell
git clone https://github.com/yakomisar/poll poll
cd poll
docker-compose up -d
```

# Usage

To create a poll with multiple choices
```shell
curl --request POST \
	--url "localhost:8000/api/createpoll" \
	--header "Content-Type: application/json" \
	--data "{\"id\":1,\"name\":\"Как зовут вашего друга?\",\"choice\":[{\"name\":\"Артем\"},{\"name\":\"Иннокентий\"},{\"name\":\"У меня нет друзей\"}]}"
```
Response
```shell
"Succesfully created."
```

To vote for a specific choice
```shell
curl --request POST \
	--url "localhost:8000/api/poll" \
	--header "Content-Type: application/json" \
	--data "{\"id\":1,\"name\":\"Иннокентий\"}"
```
Response
```shell
"Successfully voted to Иннокентий"
```

To get the result for a specific vote
```shell
curl --request POST \
	--url "localhost:8000/api/getresult" \
	--header "Content-Type: application/json" \
	--data "{\"id\":1}"
```
Response
```shell
{
  "ID": 1,
  "name": "Как зовут вашего друга",
  "choice": [
    {
      "id": 1,
      "name": "Артем",
      "votes": 0
    },
    {
      "id": 1,
      "name": "Иннокентий",
      "votes": 1
    },
    {
      "id": 1,
      "name": "У меня нет друзей",
      "votes": 0
    }
  ]
}
```