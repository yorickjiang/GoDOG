package datetime

import (
	"fmt"
)

/***** FUNCTION ********************************/

func isLeapYear(year int32) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

/***********************************************/

// Convert year/month/day to ordinal.
func ymd2ord(year int32, month, day uint8) int32 {
	if month < 1 || month > 12 {
		panic("month must be in 1..12")
	}

	var maxDim uint8

	if isLeapYear(year) && month == 2 {
		maxDim = 29
	} else {
		maxDim = _DAYS_IN_MONTH[month-1]
	}

	if day < 1 || day > maxDim {
		panic(fmt.Sprintf("day must be in 1..%d", maxDim))
	}

	var refYear int32 = year/400*400 + 1 // The year just after the nearest year divisible by 400 before 'year'

	if year < 0 {
		refYear -= 400
	}

	// number of days before January 1st of year
	var dby int32 = (year-refYear)*365 + (year-refYear)/4 - (year-refYear)/100 + (year-refYear)/400
	// number of days in year preceding first day of month
	var dbm uint16 = uint16(_DAYS_BEFORE_MONTH[month-1])

	if isLeapYear(year) && month > 2 {
		dbm++
	}

	var ord int32 = int32(dby) + int32(dbm) + int32(day)
	var refYear0 int32 = _ORD0_YEAR/400*400 + 1 // The year just after the nearest year divisible by 400 before 'year'

	if _ORD0_YEAR < 0 {
		refYear0 -= 400
	}

	dby = (_ORD0_YEAR-refYear0)*365 + (_ORD0_YEAR-refYear0)/4 - (_ORD0_YEAR-refYear0)/100 + (_ORD0_YEAR-refYear0)/400
	dbm = uint16(_DAYS_BEFORE_MONTH[_ORD0_MONTH-1])

	if isLeapYear(_ORD0_YEAR) && _ORD0_MONTH > 2 {
		dbm++
	}

	return ord + (refYear-refYear0)/400*(400*365+97) - int32(dby) - int32(dbm) - int32(_ORD0_DAY)
}

/***********************************************/

// Convert ordinal to year/month/day.
func ord2ymd(ord int32) (year int32, month, day uint8) {
	var di4y int32 = 4*365 + 1
	var di100y int32 = 25*di4y - 1
	var di400y int32 = 4*di100y + 1
	var refYear int32 = _ORD0_YEAR/400*400 + 1 // The year just after the nearest year divisible by 400 before 'year'

	if _ORD0_YEAR < 0 {
		refYear -= 400
	}

	ord += (_ORD0_YEAR-refYear)*365 + (_ORD0_YEAR-refYear)/4 - (_ORD0_YEAR-refYear)/100 + (_ORD0_YEAR-refYear)/400
	ord += int32(_DAYS_BEFORE_MONTH[_ORD0_MONTH-1]) + int32(_ORD0_DAY)

	if isLeapYear(_ORD0_YEAR) && _ORD0_MONTH > 2 {
		ord++
	}

	var n400, n100, n4, n1 int32

	if ord <= 0 {
		n400 = (1-ord)/(400*365+97) + 1
		refYear -= n400 * 400
		ord += n400 * (400*365 + 97)
	}

	ord--
	n400 = ord / di400y
	ord %= di400y
	n100 = ord / di100y
	ord %= di100y
	n4 = ord / di4y
	ord %= di4y
	n1 = ord / 365
	ord %= 365

	year = 400*n400 + 100*n100 + 4*n4 + n1 + refYear

	if (n1 == 4 || n100 == 4) && ord == 0 {
		year -= 1
		month = 12
		day = 31
	} else {
		month = uint8((ord + 50) >> 5)
		var dbm uint16 = uint16(_DAYS_BEFORE_MONTH[month-1])

		if isLeapYear(year) && month > 2 {
			dbm++
		}

		if int32(dbm) > ord {
			month -= 1
			dbm -= uint16(_DAYS_IN_MONTH[month-1])

			if isLeapYear(year) && month == 2 {
				dbm--
			}
		}

		day = uint8(ord - int32(dbm) + 1)
	}

	return
}

/***********************************************/

func getLeapSec(mjd int32) (value int8, total int16) {
	value = 0
	total = 0

	for _, item := range UTC_LEAP_SEC {
		if mjd+1 == item.Mjd {
			value = item.Value
		}

		if mjd >= item.Mjd {
			total = item.Total
			break
		}
	}

	return
}

/***********************************************/
