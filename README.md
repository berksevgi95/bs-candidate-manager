# BS Candidate Manager
This is a simple boilerplate of credential management API written in Go
## Installation & Run
```bash
# Download dependencies
$ go get -v -u github.com/gin-gonic/gin
$ go get -v -u github.com/swaggo/gin-swagger
$ go get -v -u github.com/gin-gonic/gin
$ go get -v -u github.com/alecthomas/template
$ go get -v -u github.com/swaggo/swag
$ go get -v -u go.mongodb.org/mongo-driver/bson
$ go get -v -u go.mongodb.org/mongo-driver/mongo

#Run the application
REPODB=Otsimo REPOHOST=localhost REPOPORT=27017 go run main.go

#Generate executable
go build main.go

#And run executable
REPODB=Otsimo REPOHOST=localhost REPOPORT=27017 ./main
```
