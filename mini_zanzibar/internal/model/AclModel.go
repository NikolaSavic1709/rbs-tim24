package model

import (
	"fmt"
	"strings"
)

type AclBody struct {
	Object   string `json:"object"`
	Relation string `json:"relation"`
	User     string `json:"user"`
}

func (a AclBody) ParseAcl() string {
	return fmt.Sprintf("%s#%s@%s", a.Object, a.Relation, a.User)
}
func (a *AclBody) ParseFromAcl(acl string) error {
	obj, rest := strings.Split(acl, "#")[0], strings.Split(acl, "#")[1]
	rel, usr := strings.Split(rest, "@")[0], strings.Split(rest, "@")[1]
	a.Object = obj
	a.Relation = rel
	a.User = usr
	return nil
}
