package repository

import (
	"context"
	"github.com/UniversityOfGdanskProjects/plan_back/internal/schedule/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchPlanForUser(coll *mongo.Collection,
	yearID string, groupID string, userID string) (domain.Schedule, error) {

	var result domain.Schedule
	filter := bson.D{}

	cursor, err := coll.Find(context.TODO(), filter) // ! but change this context
	if err != nil {
		return domain.Schedule{}, err
	}

	err = cursor.Decode(&result)
	if err != nil {
		return domain.Schedule{}, err
	}

	return result, nil
}
