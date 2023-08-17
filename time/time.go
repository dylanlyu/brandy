package time

import "time"

func ChangeToUTC(lasting time.Time) time.Time {
	return lasting.UTC()
}

func ChangeToRFC3339(lasting time.Time) string {
	return lasting.Format(time.RFC3339)
}

func NowToUTC() time.Time {
	return time.Now().UTC()
}

func UTCToRFC3339() string {
	utc := NowToUTC()
	return ChangeToRFC3339(utc)
}

func GetDayLatest(endAt time.Time) time.Time {
	return time.Date(endAt.Year(), endAt.Month(), endAt.Day(), 23, 59, 59, 59, endAt.Location())
}

func GetYearEarliest(startAt time.Time) time.Time {
	return time.Date(startAt.Year(), 01, 01, 0, 0, 0, 0, startAt.Location())
}

func GetMonthEarliest(timeAt time.Time) time.Time {
	return time.Date(timeAt.Year(), timeAt.Month(), 01, 0, 0, 0, 0, timeAt.Location())
}

func GetMonthLatest(timeAt time.Time) time.Time {
	timeAt = timeAt.AddDate(0, 1, -timeAt.Day())
	return time.Date(timeAt.Year(), timeAt.Month(), timeAt.Day(), 23, 59, 59, 0, timeAt.Location())
}

func GetDayEarliest(timeAt time.Time) time.Time {
	return time.Date(timeAt.Year(), timeAt.Month(), timeAt.Day(), 0, 0, 0, 0, timeAt.Location())
}

func GetHourEarliest(timeAt time.Time) time.Time {
	return time.Date(timeAt.Year(), timeAt.Month(), timeAt.Day(), timeAt.Hour(), 0, 0, 0, timeAt.Location())
}

func GetMinEarliest(timeAt time.Time) time.Time {
	return time.Date(timeAt.Year(), timeAt.Month(), timeAt.Day(), timeAt.Hour(), timeAt.Minute(), 0, 0, timeAt.Location())
}

func GetAge(birthday time.Time) (age int) {
	if birthday.IsZero() {
		return 0
	}

	now := NowToUTC()
	age = now.Year() - birthday.Year()
	if int(now.Month()) < int(birthday.Month()) || int(now.Day()) < int(birthday.Day()) {
		age--
	}
	return age
}

func ChangeLocation(lasting time.Time, location string) (time.Time, error) {
	zone, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return lasting.In(zone), nil
}

func Calculation(lasting time.Time, hour, min int) time.Time {
	return lasting.Add(time.Hour*time.Duration(hour) + time.Minute*time.Duration(min))
}

func Started() time.Time {
	utc := NowToUTC()
	return time.Date(utc.Year(), utc.Month(), utc.Day(), 0, 0, 0, 0, time.UTC)
}

func Ended() time.Time {
	utc := NowToUTC()
	return time.Date(utc.Year(), utc.Month(), utc.Day(), 23, 59, 59, 59, time.UTC)
}
