package domain

type Schedule struct{}

type YearSchedule struct {
	yearID  string `validate:"required, numeric, len=4"`
	groupID string `validate:"required, numeric, len=1"`
}

type GroupSchedule struct {
	yearID  string `validate:"required, numeric, len=4"`
	groupID string `validate:"required, numeric, len=1"`
	userID  string `validate:"required, numeric, len=5"`
}
