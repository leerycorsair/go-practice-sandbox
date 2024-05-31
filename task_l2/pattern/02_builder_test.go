package pattern

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	director := DatabaseDirector{}
	fBuilder := firstBuilder{db: &Database{}}
	director.Init(&fBuilder)
	db := fBuilder.GetResult()
	fmt.Println(db)

	sBuilder := secondBuilder{db: &Database{}}
	director.Init(&sBuilder)
	db = sBuilder.GetResult()
	fmt.Println(db)
}
