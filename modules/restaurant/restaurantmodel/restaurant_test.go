package restaurantmodel

import "testing"

type DataTable struct {
	Input  RestaurantCreate
	Expect error
}

func TestValidate(t *testing.T) {

	table := []DataTable{
		{Input: RestaurantCreate{Name: ""}, Expect: ErrNameCannotBeEmpty},
		{Input: RestaurantCreate{Name: "dsf"}, Expect: nil},
	}

	for _, d := range table {
		err := d.Input.Validate()

		if err != d.Expect {
			t.Errorf("Test Valdate() failed, expected %v, but got %v", d.Expect, err)
		}
	}

	//dataTest := RestaurantCreate{
	//	Name: "",
	//}
	//err := dataTest.Validate()
	//if err != ErrNameCannotBeEmpty {
	//	t.Errorf("Test Validata() failed, expected ErrNameCannotBeEmpty, but got %v", err)
	//}
	//dataTest.Name = "das"
	//err = dataTest.Validate()
	//if err != nil {
	//	t.Errorf("Test Validata() failed, expected ErrNameCannotBeEmpty, but got %v", err)
	//}

	t.Log("Test validate() passed")
}
