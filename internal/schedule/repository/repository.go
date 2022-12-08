package repository

import (
	"context"
	"fmt"
	"github.com/UniversityOfGdanskProjects/plan_back/internal/schedule/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchScheduleForYear(coll *mongo.Collection, yearID string) (domain.YearSchedule, error) {

}

func UpsertScheduleForYear(coll *mongo.Collection, yearID string, schedule domain.YearSchedule) error {
	query := bson.D{}

	_, err := coll.UpdateOne(context.Background(), query, schedule, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

func FetchScheduleForUser(coll *mongo.Collection,
	yearID string, groupID string, userID string) (domain.Schedule, error) {

	var result domain.Schedule
	query := bson.D{{fmt.Sprintf("%s.groups.%s.users.%s", yearID, groupID, userID), nil}} // the query might be bad

	cursor, err := coll.Find(context.Background(), query)
	if err != nil {
		return domain.Schedule{}, err
	}

	err = cursor.Decode(&result)
	if err != nil {
		return domain.Schedule{}, err
	}

	return result, nil
}

func UpsertScheduleForUser(coll *mongo.Collection,
	yearID string, groupID string, userID string, schedule domain.Schedule) error {

	update := schedule // Placeholder
	query := bson.D{}  // the query might be bad

	_, err := coll.UpdateOne(context.Background(), query, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

//{
//	yearID: {
//		plan {},
//		grupy: [
//		id_grupy: {
//			plan: {
//
//			}
//			users: {
//				id: {
//
//				}
//			}
//		}]
//	}
//}
