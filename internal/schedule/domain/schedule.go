package domain

//type Schedule struct{}

type YearSchedule struct {
	YearID  string `json:"_id" validate:"required,numeric,len=4"`
	groupID string `json:"group.id" validate:"required,numeric,len=1"`
}

type GroupSchedule struct {
	yearID  string `json:"year_id" validate:"required,numeric,len=4"`
	groupID string `json:"group_id" validate:"required,numeric,len=1"`
	Users   []User
}

type User struct {
	userID string `json:"user_id" validate:"required"`
}

type Schedule struct {
	Monday []Fields
}

type Fields struct {
}

//type T struct {
//	Year     string `json:"year"`
//	Semester string `json:"semester"`
//	Schedule struct {
//		Monday struct {
//			Field1 struct {
//				Group   int    `json:"group"`
//				Subject string `json:"subject"`
//				Teacher string `json:"teacher"`
//				Room    string `json:"room"`
//				EndHour string `json:"endHour"`
//			} `json:"8:00"`
//			Field2 struct {
//				Group   int    `json:"group"`
//				Subject string `json:"subject"`
//				Teacher string `json:"teacher"`
//				Room    string `json:"room"`
//				EndHour string `json:"endHour"`
//			} `json:"9:30"`
//		} `json:"Monday"`
//		Thursday struct {
//		} `json:"Thursday"`
//		Wednesday struct {
//		} `json:"Wednesday"`
//		Tuesday struct {
//		} `json:"Tuesday"`
//		Friday struct {
//		} `json:"Friday"`
//	} `json:"schedule"`
//}

//[
//  {
//    "_id": "yearID1",
//    "schedule": {jakies pola},
//    "groups": [
//      {
//        "id": "groupID1",
//        "schedule": {jakies pola},
//        "users": [
//          {
//            "id": "userID1",
//            "schedule": {jakies pola}
//          }
//        ]
//      }
//    ]
//  }
//]
