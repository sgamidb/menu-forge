package menu

import (
	"menuForge/meal"
	"testing"
)

func TestCreateWeeklyMenu_OneDailyMenuPerDay(t *testing.T) {
	meals := CreateSampleMeals()

	weekMenu := CreateWeekMenu(meals)
	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	for _, day := range days {
		dayMenu := weekMenu.GetDayMenu(day)
		if dayMenu == nil {
			t.Errorf("%s dayMenu is missing", day)
			return
		}

		if dayMenu.lunch.Name == "" {
			t.Errorf("lunch for %s is missing", day)
		}

		if dayMenu.dinner.Name == "" {
			t.Errorf("dinner for %s is missing", day)
		}
	}
}

func TestCreateWeeklyMenu_DontHaveTheSameMenuTwiceInTheWeek(t *testing.T) {
	meals := CreateSampleMeals()
	t.Errorf("something failed")
	menuMap := make(map[string]bool)
	weekMenu := CreateWeekMenu(meals)
	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	for _, day := range days {
		dayMenu := weekMenu.GetDayMenu(day)
		if _, ok := menuMap[dayMenu.lunch.Name]; ok {
			t.Errorf("%s still exists in week menu", dayMenu.lunch.Name)
		} else {
			menuMap[dayMenu.lunch.Name] = true
		}
		if _, ok := menuMap[dayMenu.dinner.Name]; ok {
			t.Errorf("%s still exists in week menu", dayMenu.dinner.Name)
		} else {
			menuMap[dayMenu.dinner.Name] = true
		}
	}
}

func TestCreateWeeklyMenu_DailyMenuDoesntHaveMeatForDinner(t *testing.T) {

}

func CreateSampleMeals() []meal.Meal {
	return []meal.Meal{
		meal.New("Pasta with meatball", true),
		meal.New("Quinoa with butternut", false),
		meal.New("Grilled chicken with vegetables", true),
		meal.New("Vegetable stir-fry", false),
		meal.New("Beef stew", true),
		meal.New("Lentil soup", false),
		meal.New("Chicken curry", true),
		meal.New("Vegetable curry", false),
		meal.New("Pork chops", true),
		meal.New("Mushroom risotto", false),
		meal.New("Fish tacos", true),
		meal.New("Vegetable tacos", false),
		meal.New("Lamb kebabs", true),
		meal.New("Falafel wrap", false),
		meal.New("Turkey sandwich", true),
		meal.New("Caprese salad", false),
		meal.New("Beef burger", true),
		meal.New("Veggie burger", false),
		meal.New("Chicken salad", true),
		meal.New("Greek salad", false),
	}
}
