package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Применимость:
// 1) Когда нужно динамически менять поведение объектов.
// 2) Когда нужно организовать очередь действий/передавать операцию как объект.
// 3) Когда нужно хранить историю операций (например, для отмены).

// Плюсы:
// 1) Отсутсвие зависимости между объектами, вызывающими операции,
// и объектами, которые выполняют операции.
// 2) Возможность реализации отмены/повтора операций.
// 3) Возможность отложенного запуска операций.
// 4) Возможность объединения простых команд в сложные.

// Минусы:
// 1) Разрастание числа интерфейсов/типов.

type Command interface {
	execute()
}

type Button struct {
	cmd Command
}

func (b *Button) press() {
	b.cmd.execute()
}

type Hotkey struct {
	cmd Command
}

func (hk *Hotkey) press() {
	hk.cmd.execute()
}

type SaveCloudCommand struct{}

func (cmd *SaveCloudCommand) execute() {
	fmt.Println("Saving data on cloud storage...")
}

type SaveLocalCommand struct{}

func (cmd *SaveLocalCommand) execute() {
	fmt.Println("Saving data on local storage...")
}
