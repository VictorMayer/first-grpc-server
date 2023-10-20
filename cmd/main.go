package main

import (
	"os"

	"github.com/VictorMayer/first-grpc-server/application/grpc"
	"github.com/VictorMayer/first-grpc-server/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
