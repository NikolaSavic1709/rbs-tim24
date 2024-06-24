package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"miniZanzibar/internal/model"
	"testing"
)

func TestAclCalculation(t *testing.T) {
	testNamespace := `{
  "name": "doc",
  "relation": [
    {
      "name": "owner"
    },
    {
      "name": "editor",
      "userset_rewrite": {
        "union": {
          "child": [
            {
              "_this": {}
            },
            {
              "computed_userset": {
                "relation": "owner"
              }
            }
          ]
        }
      }
    },
    {
      "name": "viewer",
      "userset_rewrite": {
        "union": {
          "child": [
            {
              "_this": {}
            },
            {
              "computed_userset": {
                "relation": "editor"
              }
            }
          ]
        }
      }
    }
  ]
}`
	data := []byte(testNamespace)

	// Create a new Document struct
	doc := model.Document{}

	// Unmarshal the JSON data into the Document struct
	err := json.Unmarshal(data, &doc)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}
	doc.GetMappedRelations()

	// Print the Document struct to verify it was populated correctly
	doc.Display()

	// Add an assertion to check if CheckRelation returns true
	assert.True(t, doc.CheckRelation("editor", "editor"), "CheckRelation should return true")
}
