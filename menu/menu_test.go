package menu

import (
	"menuForge/meal"
	"testing"
)

func TestCreateWeeklyMenu_AsDailyMenuPerDay(t *testing.T) {
	meals := []meal.Meal{
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

	weekMenu := CreateWeekMenu(meals)

	mondayMenu := weekMenu.Monday()
	if mondayMenu.lunch.Name != "" {

	}
}
