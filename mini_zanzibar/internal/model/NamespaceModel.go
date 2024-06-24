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

type Document struct {
	Name            string     `json:"doc"`
	Relations       []Relation `json:"relation"`
	MappedRelations map[string]Relation
}

func (d Document) Display() {
	fmt.Printf("Name: %s\n", d.Name)
	for _, relation := range d.Relations {
		fmt.Printf("Relation Name: %s\n", relation.Name)
		if relation.UsersetRewrite != nil {
			for _, child := range relation.UsersetRewrite.Union.Child {
				fmt.Printf("Child: %+v\n", child)
			}
		}
	}
}

func (d *Document) GetMappedRelations() {
	d.MappedRelations = make(map[string]Relation)
	for _, relation := range d.Relations {
		d.MappedRelations[relation.Name] = relation
	}
}

func (d Document) CheckRelation(knownRelation string, wantedRelation string) bool {
	relation, exists := d.MappedRelations[wantedRelation]
	if !exists {
		return false
	}

	if relation.UsersetRewrite != nil && relation.UsersetRewrite.Union != nil {
		for _, child := range relation.UsersetRewrite.Union.Child {
			if child.This != nil && knownRelation == wantedRelation {
				return true
			}

			if child.ComputedUserset != nil && child.ComputedUserset.Relation == knownRelation {
				return true
			} else if child.ComputedUserset != nil {
				found := d.CheckRelation(knownRelation, child.ComputedUserset.Relation)
				if found {
					return true
				}
			}
		}
	}
	return false

}
