package booking

import (
	"time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	// const layout = "Jan 2, 2006 at 3:04pm (MST)"
	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "1/2/2006 15:04:05"
	v, _ := time.Parse(layout, date)
	return v
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const layout = "January _2, 2006 15:04:05"
	v, _ := time.Parse(layout, date)
	// println(v.String());
	return v.Unix() < time.Now().Unix()
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January _2, 2006 15:4:5"
	v, _ := time.Parse(layout, date)
	// println(v.String());
	return v.Hour() >= 12 && v.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	// Description("7/25/2019 13:45:00")
	const layout = "1/2/2006 15:04:05"
	v, _ := time.Parse(layout, date)
	println(v.String());
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", v.Weekday(), v.Month(), v.Day(), v.Year(), v.Hour(), v.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	v, _ := time.Parse("2006-1-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return v
}

// const (
//     ANSIC       = "Mon Jan _2 15:04:05 2006"
//     UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
//     RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
//     RFC822      = "02 Jan 06 15:04 MST"
//     RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
//     RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
//     RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
//     RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
//     RFC3339     = "2006-01-02T15:04:05Z07:00"
//     RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
//     Kitchen     = "3:04PM"
//     // Handy time stamps.
//     Stamp      = "Jan _2 15:04:05"
//     StampMilli = "Jan _2 15:04:05.000"
//     StampMicro = "Jan _2 15:04:05.000000"
//     StampNano  = "Jan _2 15:04:05.000000000"
// )
