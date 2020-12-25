package databasex

import "time"

const layOutDateTime string = "2006-01-02 15:04:05"

// Student is type
type tStudentTest struct {
	ID        string    `fieldtbl:"id"`
	Name      string    `fieldtbl:"name"`
	Age       int       `fieldtbl:"age"`
	Grade     int       `fieldtbl:"grade"`
	CreatedAt time.Time `fieldtbl:"created_at"`
	//CreatedAt string `fieldtbl:"created_at"`
}

func isDateDifferent(time1, time2 time.Time) bool {
	return time1.Year() != time2.Year() && time1.Month() != time2.Month() && time1.Day() != time2.Day()
}

func isTimeDifferent(time1, time2 time.Time) bool {
	return time1.Hour() != time2.Hour() && time1.Minute() != time2.Minute() && time1.Second() != time2.Second()
}

func isDateTimeDifferent(time1, time2 time.Time) bool {
	return isDateDifferent(time1, time2) || isTimeDifferent(time1, time2)
}
