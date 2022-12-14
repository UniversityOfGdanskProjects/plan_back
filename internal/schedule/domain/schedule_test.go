package domain

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestYearSchedule(t *testing.T) {
	yearSch := YearSchedule{
		YearID:  "2021",
		groupID: "3",
	}
	validate := validator.New()
	err := validate.Struct(yearSch)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGroupSchedule(t *testing.T) {
	groupSch := GroupSchedule{
		yearID:  "2022",
		groupID: "4",
		userID:  "hbsduab",
	}
	validate := validator.New()
	err := validate.Struct(groupSch)
	if err != nil {
		t.Errorf(err.Error())
	}
}
