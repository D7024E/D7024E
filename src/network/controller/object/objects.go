package objectController

type Object struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

var Objects []Object
