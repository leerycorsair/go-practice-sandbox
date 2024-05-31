package pattern

import "testing"

func TestFactoryMethod(t *testing.T) {
	types := []string{"eng", "ru"}
	for _, t := range types {
		agency := NewTravelAgency(t)
		agency.Welcome()
	}
}
