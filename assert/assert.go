package assert

import "log"

func True(condition bool, message string) {
	if !condition {
		log.Fatal(message)
	}
}

func False(condition bool, message string) {
	if condition {
		log.Fatal(message)
	}
}

func Nil(value interface{}, message string) {
	if value != nil {
		log.Fatal(message)
	}
}

func NotNil(value interface{}, message string) {
	if value == nil {
		log.Fatal(message)
	}
}

func Equal(a interface{}, b interface{}, message string) {
	if a != b {
		log.Fatal(message)
	}
}

func NotEqual(a interface{}, b interface{}, message string) {
	if a == b {
		log.Fatal(message)
	}
}