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

	// Create a new Namespace struct
	doc := model.Namespace{}

	// Unmarshal the JSON data into the Namespace struct
	err := json.Unmarshal(data, &doc)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Print the Namespace struct to verify it was populated correctly
	doc.Display()

	// Add an assertion to check if CheckRelation returns true
	assert.True(t, doc.CheckRelation("editor", "editor"), "CheckRelation should return true")
	assert.True(t, doc.CheckRelation("owner", "viewer"), "CheckRelation should return true")
	assert.True(t, doc.CheckRelation("editor", "viewer"), "CheckRelation should return true")
	assert.True(t, doc.CheckRelation("editor", "editor"), "CheckRelation should return true")
}
