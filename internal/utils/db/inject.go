package db

import (
	"chat-app-fiber/internal/utils/db/mongodb"
	"github.com/google/wire"
)

var Set = wire.NewSet(mongodb.OpenMongoDBConnection)
