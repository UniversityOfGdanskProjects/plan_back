package repository

import (
	"context"
	"github.com/UniversityOfGdanskProjects/plan_back/internal/schedule/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchScheduleForYear(coll *mongo.Collection, yearID string) (domain.YearSchedule, error) {
	var result domain.YearSchedule
	query := bson.D{{"_id", yearID}}

	cursor, err := coll.Find(context.Background(), query)
	if err != nil {
		return domain.YearSchedule{}, err
	}

	err = cursor.Decode(&result)
	if err != nil {
		return domain.YearSchedule{}, err
	}

	return result, nil
}

func RemoveScheduleForYear(coll *mongo.Collection, yearID string) error {
	query := bson.D{{"_id", yearID}}

	_, err := coll.DeleteOne(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func UpsertScheduleForYear(coll *mongo.Collection, yearID string, schedule domain.YearSchedule) error {
	query := bson.D{{"_id", yearID}}

	_, err := coll.UpdateOne(context.Background(), query, schedule, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

func FetchScheduleForGroup(coll *mongo.Collection, yearID string, groupID string) (domain.GroupSchedule, error) {
	var result domain.GroupSchedule
	query := bson.D{{"_id", yearID}, {"groups.id", groupID}}

	cursor, err := coll.Find(context.Background(), query)
	if err != nil {
		return domain.GroupSchedule{}, err
	}

	err = cursor.Decode(&result)
	if err != nil {
		return domain.GroupSchedule{}, err
	}

	return result, nil
}

// below under work
//func RemoveScheduleForGroup(coll *mongo.Collection, yearID string) error {
//	query := bson.D{{"_id", yearID}}
//
//	_, err := coll.DeleteOne(context.Background(), query)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func UpsertScheduleForGroup(coll *mongo.Collection, yearID string, schedule domain.GroupSchedule) error {
//	query := bson.D{{"_id", yearID}}
//
//	_, err := coll.UpdateOne(context.Background(), query, schedule, options.Update().SetUpsert(true))
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func FetchScheduleForUser(coll *mongo.Collection,
//	yearID string, groupID string, userID string) (domain.Schedule, error) {
//
//	var result domain.Schedule
//	query := bson.D{{fmt.Sprintf("%s.groups.%s.users.%s", yearID, groupID, userID), nil}} // the query might be bad
//
//	cursor, err := coll.Find(context.Background(), query)
//	if err != nil {
//		return domain.Schedule{}, err
//	}
//
//	err = cursor.Decode(&result)
//	if err != nil {
//		return domain.Schedule{}, err
//	}
//
//	return result, nil
//}
//
//func UpsertScheduleForUser(coll *mongo.Collection,
//	yearID string, groupID string, userID string, schedule domain.Schedule) error {
//
//	update := schedule // Placeholder
//	query := bson.D{}  // the query might be bad
//
//	_, err := coll.UpdateOne(context.Background(), query, update, options.Update().SetUpsert(true))
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
