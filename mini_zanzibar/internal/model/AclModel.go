package model

import "fmt"

type AclBody struct {
	Object   string `json:"object"`
	Relation string `json:"relation"`
	User     string `json:"user"`
}

func (a AclBody) ParseAcl() string {
	return fmt.Sprintf("%s#%s@%s", a.Object, a.Relation, a.User)
}
