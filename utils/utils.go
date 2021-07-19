package utils

import "time"

var loc *time.Location

func GetTimeZone() *time.Location {
	return loc
}

func SetTimeZone(name string) error {
	var err error
	loc, err = time.LoadLocation(name)
	if err != nil {
		return err
	}
	return nil
}
