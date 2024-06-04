package pattern

import "testing"

func TestStrategy(t *testing.T) {
	p := &PaymentProcessor{}
	amount := 10.0

	methods := make(map[string]PaymentStrategy)
	methods["via_card"] = &CreditCardStrategy{"1321 3123 5435 6921"}
	methods["via_paypal"] = &PayPalStrategy{"keklul@mail.com"}
	methods["via_btc"] = &BitcoinStrategy{"1Lbcfr7sAHTD9CgdQo3HTMTkV8LK4ZnX71"}

	p.SetStrategy(methods["via_card"])
	p.Pay(amount)

	p.SetStrategy(methods["via_paypal"])
	p.Pay(amount)

	p.SetStrategy(methods["via_btc"])
	p.Pay(amount)
}
