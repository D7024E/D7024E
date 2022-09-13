package objectController

type Object struct {
	Name string `json:"name"` // json name as string.
	Hash string `json:"hash"` // json hash as string.
}

var Objects []Object // slice of object.
