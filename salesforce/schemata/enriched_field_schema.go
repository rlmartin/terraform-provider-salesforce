package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EnrichedField resource defined in the Terraform configuration
func EnrichedFieldSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying EnrichedField resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEnrichedFieldResourceData(d *schema.ResourceData, m *models.EnrichedField, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("name", m.Name)
}

// Iterate throught and update the EnrichedField resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEnrichedFieldSubResourceData(m []*models.EnrichedField) (d []*map[string]interface{}) {
	for _, enrichedField := range m {
		if enrichedField != nil {
			properties := make(map[string]interface{})
			properties["name"] = enrichedField.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EnrichedField resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EnrichedFieldModel(d *schema.ResourceData) *models.EnrichedField {
	name := d.Get("name").(string)

	return &models.EnrichedField{
		Name: &name,
	}
}

// Function to perform the following actions:
func EnrichedFieldModelFromMap(m map[string]interface{}) *models.EnrichedField {
	name := m["name"].(string)

	return &models.EnrichedField{
		Name: &name,
	}
}

func EnrichedFieldModelFromArrayOfMap(m []interface{}) []*models.EnrichedField {
	mapped := make([]*models.EnrichedField, len(m))
	for i, v := range m {
		mapped[i] = EnrichedFieldModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the EnrichedField resource
func GetEnrichedFieldPropertyFields() (t []string) {
	return []string{
		"name",
	}
}
