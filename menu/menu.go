package menu

import "menuForge/meal"

type DayMenu struct {
	lunch  meal.Meal
	dinner meal.Meal
}

type WeekMenu struct {
	monday DayMenu
}

func (weekMenu *WeekMenu) Monday() DayMenu {
	return weekMenu.monday
}

func CreateWeekMenu(meals []meal.Meal) WeekMenu {

	return WeekMenu{}
}
