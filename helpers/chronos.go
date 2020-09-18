package helpers

import (
	"time"
)

//GetTimestampTz to use in different places
func GetTimestampTz() time.Time {
	t := time.Now()
	return t.Local()

}
