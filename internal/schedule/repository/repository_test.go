package repository

import (
	"github.com/UniversityOfGdanskProjects/plan_back/internal/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"testing"
)

func setup() *mongo.Client {
	return db.MongoInit()
}

func shutdown() {

}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestFetchYearWithExistingID(t *testing.T) {

}

func TestFetchYearWithNonExistingID(t *testing.T) {

}
