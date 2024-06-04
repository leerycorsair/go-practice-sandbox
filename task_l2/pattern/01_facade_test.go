package pattern

import "testing"

func TestOrderHandler(t *testing.T) {
	h := OrderHandler{
		os: &OrderService{},
		cs: &CustomerService{},
		ss: &SellerService{},
	}
	h.CreateOrder()
}
