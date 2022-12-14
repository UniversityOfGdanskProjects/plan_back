package domain

type Schedule struct{}

type YearSchedule struct {
	YearID  string `json:"_id" validate:"required,numeric,len=4"`
	groupID string `json:"group.id" validate:"required,numeric,len=1"`
}

type GroupSchedule struct {
	yearID  string `json:"year_id" validate:"required,numeric,len=4"`
	groupID string `json:"group_id" validate:"required,numeric,len=1"`
	userID  string `json:"user_id" validate:"required"`
}
