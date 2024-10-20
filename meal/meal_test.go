package meal

import "testing"

func TestMeal(t *testing.T) {

	meal := New("Pasta")

	if meal.Name == "" {
		t.Errorf("Expected meal to have a name, but it was empty")
	}
}
