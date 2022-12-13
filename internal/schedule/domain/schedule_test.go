package domain

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestYearSchedule(t *testing.T) {
	yearSch := YearSchedule{
		yearID:  "2001",
		groupID: "4",
	}
	validate := validator.New()
	err := validate.Struct(yearSch)
	if err != nil {
		t.Errorf(err.Error())
	}
}

//type YearSchedule struct {
//	yearID  string `validate:"required, numeric, len=4"`
//	groupID string `validate:"required, numeric, len=1"`
//}
//
//type GroupSchedule struct {
//	yearID  string `validate:"required, numeric, len=4"`
//	groupID string `validate:"required, numeric, len=1"`
//	userID  string `validate:"required, numeric, len=5"`
//}
