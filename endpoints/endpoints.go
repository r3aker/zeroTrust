package endpoints

import (
	"crypto/md5"

	"github.com/pritunl/mongo-go-driver/bson/primitive"
	"github.com/pritunl/pritunl-zero/database"
)

type Doc interface {
	GetCollection(*database.Database) *database.Collection
	SetEndpoint(primitive.ObjectID)
}

func GenerateId(endpoint, clientId primitive.ObjectID) []byte {
	hash := md5.New()
	hash.Write([]byte(endpoint.Hex()))
	hash.Write([]byte(clientId.Hex()))
	return hash.Sum(nil)
}
