package node

type Object struct {
	Name string `json:"name"` // json name as string.
	Hash string `json:"hash"` // json hash as string.
}

type Objects []Object
