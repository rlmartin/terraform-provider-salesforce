package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Record resource defined in the Terraform configuration
func RecordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and RecordSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceRecordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying Record resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetRecordResourceData(d *schema.ResourceData, m *models.Record, isDataResource bool) {
	if m.ID == "" && isDataResource {
		d.SetId("-")
	} else {
		d.SetId(m.ID)
	}
}

// Iterate throught and update the Record resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetRecordSubResourceData(m []*models.Record) (d []*map[string]interface{}) {
	for _, record := range m {
		if record != nil {
			properties := make(map[string]interface{})
			properties["id"] = record.ID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Record resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func RecordModel(d *schema.ResourceData) *models.Record {
	id := d.Get("id").(string)

	return &models.Record{
		ID: id,
	}
}

// Function to perform the following actions:
func RecordModelFromMap(m map[string]interface{}) *models.Record {
	id := m["id"].(string)

	return &models.Record{
		ID: id,
	}
}

// Retrieve property field names for updating the Record resource
func GetRecordPropertyFields() (t []string) {
	return []string{
		"id",
	}
}
