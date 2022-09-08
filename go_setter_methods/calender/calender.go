package calender

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

// Setter Method to validate the field values
// Notice we pass a pointer receiver to update the original value

func (d *Date) SetYear(y int) error {
	// If the given year is invalid, return the error
	if y < 1 {
		return errors.New("invalid year")

	}
	// Otherwise set the year
	d.year = y
	// And return an error of nil, meaning a valid year was given
	return nil
}

// getter method
// Use a pointer receiver type
// Same name as the struct field but Capitalized to enable to be Exported
func (d *Date) Year() int {
	// return the field value
	return d.year
}

func (d *Date) SetMonth(m int) error {
	if m < 1 || m > 12 {
		return errors.New("invalid month")
	}
	d.month = m

	return nil
}

// getter method
// Use a pointer receiver type
// Same name as the struct field but Capitalized to enable to be Exported
func (d *Date) Month() int {
	// return the field value
	return d.month
}

func (d *Date) SetDay(dy int) error {
	if dy < 0 || dy > 31 {
		return errors.New("invalid day")
	}
	d.day = dy

	return nil
}

// getter method
// Use a pointer receiver type
// Same name as the struct field but Capitalized to enable to be Exported
func (d *Date) Day() int {
	// return the field value
	return d.day
}
