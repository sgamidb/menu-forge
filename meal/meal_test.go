package meal

import "testing"

func TestMeal_Name(t *testing.T) {
	meal := New("Pasta with meatball", true)

	if meal.Name == "" {
		t.Errorf("Expected meal to have a name, but it was empty")
	}
}

func TestMeal_IsNotVeganWithMeat(t *testing.T) {

	meal := New("Pasta with meatball", true)

	if meal.IsVegan() == true {
		t.Errorf("A meal with meat is not vegan")
	}
}

func TestMeal_IsVeganWithoutMeal(t *testing.T) {
	meal := New("Quinoa with butternut", false)

	if meal.IsVegan() == false {
		t.Errorf("A meal without meat is vegan")
	}
}
