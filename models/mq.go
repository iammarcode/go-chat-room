package models

type Message struct {
	Type  string
	Table string
	Data  map[string]interface{}
}

var Mapper = map[string]interface{}{
	"user": &User{},
}
