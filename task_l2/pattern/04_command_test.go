package pattern

import "testing"

func TestCommand(t *testing.T) {
	btn := &Button{}
	hk := &Hotkey{}
	cmd1 := &SaveCloudCommand{}
	cmd2 := &SaveLocalCommand{}

	btn.cmd = cmd1
	hk.cmd = cmd1
	btn.press()
	hk.press()

	btn.cmd = cmd2
	hk.cmd = cmd2
	btn.press()
	hk.press()
}
