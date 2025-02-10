package creational

import "sync"

type Singleton struct {
	Data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Data: "I am a Singlton instance"}
	})

	return instance
}
