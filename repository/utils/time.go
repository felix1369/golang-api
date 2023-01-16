package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type LocalTime struct {
	time.Time
}

func (localTime *LocalTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation("2006-01-02", s, time.Local)
	if err != nil {
		t2, err2 := time.ParseInLocation(time.RFC3339, s, time.Local)
		if err2 != nil {
			return errors.New(fmt.Sprint("Invalid Time Format (expected in YYYY-MM-DD format or RFC3339), ", err.Error(), ", ", err2.Error()))
		}
		t = t2
	}
	localTime.Time = t
	return nil
}

func (localTime *LocalTime) Value() (driver.Value, error) {
	if localTime == nil {
		return nil, nil
	}
	return localTime.Time, nil
}

func (localTime *LocalTime) Scan(src interface{}) (err error) {
	var ok bool
	localTime.Time, ok = src.(time.Time)
	if !ok {
		return errors.New("Incompatible type src to time.Time")
	}

	return nil
}

func (localTime *LocalTime) ScanString(src string) (err error) {
	localTime.Time, err = time.Parse(time.RFC3339, src)
	if err != nil {
		localTime.Time, err = time.Parse("2006-01-02", src)
		if err != nil {
			return err
		}
	}
	return nil
}

type LocalTimeHour struct {
	time.Time
}

func (localTime *LocalTimeHour) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation("03:04:05 PM", s, time.Local)
	if err != nil {
		t2, err2 := time.ParseInLocation(time.RFC3339, s, time.Local)
		if err2 != nil {
			return errors.New(fmt.Sprint("Invalid Time Format (expected in YYYY-MM-DD format or RFC3339), ", err.Error(), ", ", err2.Error()))
		}
		t = t2
	}
	localTime.Time = t
	return nil
}

func (localTime *LocalTimeHour) Value() (driver.Value, error) {
	if localTime == nil {
		return nil, nil
	}
	return localTime.Time, nil
}

func (localTime *LocalTimeHour) Scan(src interface{}) (err error) {
	var ok bool
	localTime.Time, ok = src.(time.Time)
	if !ok {
		return errors.New("Incompatible type src to time.Time")
	}

	return nil
}

func (localTime *LocalTimeHour) ScanString(src string) (err error) {
	localTime.Time, err = time.Parse(time.RFC3339, src)
	if err != nil {
		return err
	}
	return nil
}

func (localTime *LocalTime) ClockString() string {
	hours := localTime.Hour()
	var stringHours string
	minutes := localTime.Minute()
	var stringMinutes string
	if hours < 10 {
		stringHours = "0" + strconv.Itoa(hours)
	} else {
		stringHours = strconv.Itoa(hours)
	}
	if minutes < 10 {
		stringMinutes = "0" + strconv.Itoa(minutes)
	} else {
		stringMinutes = strconv.Itoa(minutes)
	}

	return stringHours + ":" + stringMinutes
}
