package menu

import (
	"math/rand/v2"
	"menuForge/meal"
)

type DayMenu struct {
	lunch  meal.Meal
	dinner meal.Meal
}

type WeekMenu struct {
	monday    DayMenu
	tuesday   DayMenu
	wednesday DayMenu
	thursday  DayMenu
	friday    DayMenu
	saturday  DayMenu
	sunday    DayMenu
}

func (weekMenu *WeekMenu) GetDayMenu(day string) *DayMenu {
	switch day {
	case "monday":
		return &weekMenu.monday
	case "tuesday":
		return &weekMenu.tuesday
	case "wednesday":
		return &weekMenu.wednesday
	case "thursday":
		return &weekMenu.thursday
	case "friday":
		return &weekMenu.friday
	case "saturday":
		return &weekMenu.saturday
	case "sunday":
		return &weekMenu.sunday
	default:
		return nil
	}
}

func CreateWeekMenu(meals []meal.Meal) WeekMenu {
	var weekMenu WeekMenu

	weekMenu.monday = DayMenu{
		lunch:  meals[rand.IntN(len(meals))],
		dinner: meals[rand.IntN(len(meals))],
	}
	weekMenu.tuesday = DayMenu{
		lunch:  meals[rand.IntN(len(meals))],
		dinner: meals[rand.IntN(len(meals))],
	}
	weekMenu.wednesday = DayMenu{
		lunch:  meals[rand.IntN(len(meals))],
		dinner: meals[rand.IntN(len(meals))],
	}
	weekMenu.thursday = DayMenu{
		lunch:  meals[rand.IntN(len(meals))],
		dinner: meals[rand.IntN(len(meals))],
	}
	weekMenu.friday = DayMenu{
		lunch:  meals[rand.IntN(len(meals))],
		dinner: meals[rand.IntN(len(meals))],
	}
	weekMenu.saturday = DayMenu{
		lunch:  meals[rand.IntN(len(meals))],
		dinner: meals[rand.IntN(len(meals))],
	}
	weekMenu.sunday = DayMenu{
		lunch:  meals[rand.IntN(len(meals))],
		dinner: meals[rand.IntN(len(meals))],
	}
	return weekMenu
}
