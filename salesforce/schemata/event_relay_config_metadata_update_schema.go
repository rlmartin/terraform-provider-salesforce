package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EventRelayConfigMetadataUpdate resource defined in the Terraform configuration
func EventRelayConfigMetadataUpdateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"relay_option": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying EventRelayConfigMetadataUpdate resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventRelayConfigMetadataUpdateResourceData(d *schema.ResourceData, m *models.EventRelayConfigMetadataUpdate) {
	d.Set("relay_option", m.RelayOption)
}

// Iterate throught and update the EventRelayConfigMetadataUpdate resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventRelayConfigMetadataUpdateSubResourceData(m []*models.EventRelayConfigMetadataUpdate) (d []*map[string]interface{}) {
	for _, eventRelayConfigMetadataUpdate := range m {
		if eventRelayConfigMetadataUpdate != nil {
			properties := make(map[string]interface{})
			properties["relay_option"] = eventRelayConfigMetadataUpdate.RelayOption
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EventRelayConfigMetadataUpdate resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventRelayConfigMetadataUpdateModel(d *schema.ResourceData) *models.EventRelayConfigMetadataUpdate {
	relayOption := d.Get("relay_option").(string)

	return &models.EventRelayConfigMetadataUpdate{
		RelayOption: &relayOption,
	}
}

// Function to perform the following actions:
func EventRelayConfigMetadataUpdateModelFromMap(m map[string]interface{}) *models.EventRelayConfigMetadataUpdate {
	relayOption := m["relay_option"].(string)

	return &models.EventRelayConfigMetadataUpdate{
		RelayOption: &relayOption,
	}
}

// Retrieve property field names for updating the EventRelayConfigMetadataUpdate resource
func GetEventRelayConfigMetadataUpdatePropertyFields() (t []string) {
	return []string{
		"relay_option",
	}
}
