package pattern

import (
	"fmt"
	"testing"
)

func TestChainOfResp(t *testing.T) {
	h1 := &HandlerA{}
	h2 := &HandlerB{}
	h3 := &HandlerC{}

	h1.SetNext(h2)
	h2.SetNext(h3)

	data := []string{"a", "b", "c", "d"}
	for _, v := range data {
		fmt.Printf("Handling data: %s\n", v)
		h1.Execute(v)
	}
}
