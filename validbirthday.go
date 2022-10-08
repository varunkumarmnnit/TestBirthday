package main

import (
	"errors"
	"fmt"
	"jaxf-github.fanatics.corp/TestBirthdayNew/Model"
	"jaxf-github.fanatics.corp/TestBirthdayNew/Utils"
)

type ListTodayPeopleBirthday [][]string

func findTodayPeopleBirthday(pDetails *Model.PersonDobDetails) (ListTodayPeopleBirthday, error) {
	var ans ListTodayPeopleBirthday
	var isTodayLeapYear bool = false
	if len(pDetails.Values) == 0 {
		fmt.Println("No inpout Provided")

	}
	yearToday, monthToday, dayToday, err := Utils.GetYearMonthDay("")
	if err != nil {
		return nil, err
	}
	if yearToday%4 == 0 {
		isTodayLeapYear = true
	}

	for _, eachPersonDetail := range pDetails.Values {
		if len(eachPersonDetail) != 3 {
			return nil, errors.New(Utils.IncorrectInputDataFormat)

		}
		_, month, day, err := Utils.GetYearMonthDay("")
		if err != nil {
			return nil, err
		}

		if monthToday == 2 {
			if isTodayLeapYear {
				ValidateBirthday(monthToday, dayToday, month, day, eachPersonDetail, ans)
			} else {
				if day == 29 && dayToday == 28 {
					ans = append(ans, []string{eachPersonDetail[0], eachPersonDetail[1], eachPersonDetail[2]})
				} else {
					ValidateBirthday(monthToday, dayToday, month, day, eachPersonDetail, ans)
				}
			}
		} else {
			ValidateBirthday(monthToday, dayToday, month, day, eachPersonDetail, ans)
		}

	}

	return ans, nil
}

func ValidateBirthday(monthToday, dayToday, month, day int, eachPersonDetail []string, ans ListTodayPeopleBirthday) {
	if monthToday == monthToday && dayToday == day {
		ans = append(ans, []string{eachPersonDetail[0], eachPersonDetail[1], eachPersonDetail[2]})

	}

}
