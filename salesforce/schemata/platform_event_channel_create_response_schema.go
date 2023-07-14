package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannelCreateResponse resource defined in the Terraform configuration
func PlatformEventChannelCreateResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and PlatformEventChannelCreateResponseSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourcePlatformEventChannelCreateResponseSchema() map[string]*schema.Schema {
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

// Update the underlying PlatformEventChannelCreateResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelCreateResponseResourceData(d *schema.ResourceData, m *models.PlatformEventChannelCreateResponse) {
	d.SetId(m.ID)
}

// Iterate throught and update the PlatformEventChannelCreateResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelCreateResponseSubResourceData(m []*models.PlatformEventChannelCreateResponse) (d []*map[string]interface{}) {
	for _, platformEventChannelCreateResponse := range m {
		if platformEventChannelCreateResponse != nil {
			properties := make(map[string]interface{})
			properties["id"] = platformEventChannelCreateResponse.ID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannelCreateResponse resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelCreateResponseModel(d *schema.ResourceData) *models.PlatformEventChannelCreateResponse {
	id := d.Get("id").(string)

	return &models.PlatformEventChannelCreateResponse{
		ID: id,
	}
}

// Function to perform the following actions:
func PlatformEventChannelCreateResponseModelFromMap(m map[string]interface{}) *models.PlatformEventChannelCreateResponse {
	id := m["id"].(string)

	return &models.PlatformEventChannelCreateResponse{
		ID: id,
	}
}

// Retrieve property field names for updating the PlatformEventChannelCreateResponse resource
func GetPlatformEventChannelCreateResponsePropertyFields() (t []string) {
	return []string{
		"id",
	}
}
