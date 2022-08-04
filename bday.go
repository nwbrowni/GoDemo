package main

import (
	"strconv"
	"time"
)

// bday struct and methods
type bday struct {
	day   int
	month int
	year  int
}

func (b *bday) SetDay(day int) {
	// at most there are only 31 days in a month
	if day > 31 {
		return
	} else if b.month == 4 || b.month == 6 || b.month == 9 || b.month == 11 {
		// if the month is April, June, September, or November, it only has 30 days
		if day > 30 {
			return
		}
	} else if b.month == 2 {
		// if the month is February, it only has 28 days, except leap years
		if day == 29 {
			currentyear, _, _ := time.Now().Date() // only really interested in the current year
			for i := 1582; i < currentyear; i = i + 4 {
				if i == b.year {
					b.day = day
				}
			}

			return
		} else if day > 28 {
			return
		}
	} else {
		b.day = day
	}
}

func (b *bday) SetMonth(month int) {
	if month > 12 {
		return
	} else {
		b.month = month
	}
}

func (b *bday) SetYear(year int) {
	if year > time.Now().Year() {
		return
	} else {
		b.year = year
	}
}

func (b *bday) SetBirthday(day int, month int, year int) {
	b.SetYear(year)
	b.SetMonth(month)
	b.SetDay(day)
}

func (b *bday) GetBirthdayStrings() []string {
	return []string{strconv.Itoa(b.day), strconv.Itoa(b.month), strconv.Itoa(b.year)}
}
