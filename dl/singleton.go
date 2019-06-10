package dl

import (
	"fmt"
	"sync"
)

type singleton struct {
	Name string
	Age int
	Sex string
	Score int
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("1")
		instance = &singleton{}
	})
	return instance
}