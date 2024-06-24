package model

import "fmt"

type Child struct {
	This            *struct{} `json:"_this,omitempty"`
	ComputedUserset *struct {
		Relation string `json:"relation"`
	} `json:"computed_userset,omitempty"`
}

type UsersetRewrite struct {
	Union *struct {
		Child []Child `json:"child"`
	} `json:"union,omitempty"`
}

type Relation struct {
	Name           string          `json:"name"`
	UsersetRewrite *UsersetRewrite `json:"userset_rewrite,omitempty"`
}

type Namespace struct {
	Name      string     `json:"name"`
	Relations []Relation `json:"relation"`
}

func (n Namespace) Display() {
	fmt.Printf("Name: %s\n", n.Name)
	for _, relation := range n.Relations {
		fmt.Printf("Relation Name: %s\n", relation.Name)
		if relation.UsersetRewrite != nil {
			for _, child := range relation.UsersetRewrite.Union.Child {
				fmt.Printf("Child: %+v\n", child)
			}
		}
	}
}

func (n Namespace) CheckRelation(currentRelation string, wantedRelation string) bool {
	for _, relation := range n.Relations {
		if relation.Name == wantedRelation {
			if relation.UsersetRewrite != nil && relation.UsersetRewrite.Union != nil {
				for _, child := range relation.UsersetRewrite.Union.Child {
					if child.This != nil && currentRelation == wantedRelation {
						return true
					}
					if child.ComputedUserset != nil && child.ComputedUserset.Relation == currentRelation {
						return true
					} else if child.ComputedUserset != nil {
						found := n.CheckRelation(currentRelation, child.ComputedUserset.Relation)
						if found {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func (n Namespace) CheckValid() bool {
	isValid := true
	if n.Name == "" || len(n.Relations) <= 0 {
		isValid = false
	}
	return isValid
}
