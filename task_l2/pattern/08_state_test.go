package pattern

import "testing"

func TestState(t *testing.T) {
	machine := &CoffeeMachine{}
	machine.SetState(&IdleState{machine})

	machine.Dispense()
	machine.PressButton()
	machine.InsertCoin()
	machine.InsertCoin()
	machine.Dispense()
	machine.PressButton()
	machine.InsertCoin()
	machine.PressButton()
	machine.Dispense()
}
