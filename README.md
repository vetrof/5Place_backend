5Place
***

dev mode: 

for start postgis-db:  
```docker compose up -d```

.env:  
```
DB_HOST=localhost
DB_PORT=55000
DB_USER=postgres
DB_NAME=place5
DB_PASSWORD=postgrespw
PORT=8080
```

start project:  
```
go mod tidy
go run cmd/api/main.go
```
