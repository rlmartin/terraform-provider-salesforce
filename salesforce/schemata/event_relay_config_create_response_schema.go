package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EventRelayConfigCreateResponse resource defined in the Terraform configuration
func EventRelayConfigCreateResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and EventRelayConfigCreateResponseSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceEventRelayConfigCreateResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying EventRelayConfigCreateResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventRelayConfigCreateResponseResourceData(d *schema.ResourceData, m *models.EventRelayConfigCreateResponse, isDataResource bool) {
	if m.ID == "" && isDataResource {
		d.SetId("-")
	} else {
		d.SetId(m.ID)
	}
}

// Iterate throught and update the EventRelayConfigCreateResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventRelayConfigCreateResponseSubResourceData(m []*models.EventRelayConfigCreateResponse) (d []*map[string]interface{}) {
	for _, eventRelayConfigCreateResponse := range m {
		if eventRelayConfigCreateResponse != nil {
			properties := make(map[string]interface{})
			properties["id"] = eventRelayConfigCreateResponse.ID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EventRelayConfigCreateResponse resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventRelayConfigCreateResponseModel(d *schema.ResourceData) *models.EventRelayConfigCreateResponse {
	id := d.Get("id").(string)

	return &models.EventRelayConfigCreateResponse{
		ID: id,
	}
}

// Function to perform the following actions:
func EventRelayConfigCreateResponseModelFromMap(m map[string]interface{}) *models.EventRelayConfigCreateResponse {
	id := m["id"].(string)

	return &models.EventRelayConfigCreateResponse{
		ID: id,
	}
}

func EventRelayConfigCreateResponseModelFromArrayOfMap(m []interface{}) []*models.EventRelayConfigCreateResponse {
	mapped := make([]*models.EventRelayConfigCreateResponse, len(m))
	for i, v := range m {
		mapped[i] = EventRelayConfigCreateResponseModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the EventRelayConfigCreateResponse resource
func GetEventRelayConfigCreateResponsePropertyFields() (t []string) {
	return []string{
		"id",
	}
}
