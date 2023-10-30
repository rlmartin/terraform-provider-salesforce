package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannelMemberCreateResponse resource defined in the Terraform configuration
func PlatformEventChannelMemberCreateResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and PlatformEventChannelMemberCreateResponseSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourcePlatformEventChannelMemberCreateResponseSchema() map[string]*schema.Schema {
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

// Update the underlying PlatformEventChannelMemberCreateResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelMemberCreateResponseResourceData(d *schema.ResourceData, m *models.PlatformEventChannelMemberCreateResponse, isDataResource bool) {
	if m.ID == "" && isDataResource {
		d.SetId("-")
	} else {
		d.SetId(m.ID)
	}
}

// Iterate throught and update the PlatformEventChannelMemberCreateResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelMemberCreateResponseSubResourceData(m []*models.PlatformEventChannelMemberCreateResponse) (d []*map[string]interface{}) {
	for _, platformEventChannelMemberCreateResponse := range m {
		if platformEventChannelMemberCreateResponse != nil {
			properties := make(map[string]interface{})
			properties["id"] = platformEventChannelMemberCreateResponse.ID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannelMemberCreateResponse resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelMemberCreateResponseModel(d *schema.ResourceData) *models.PlatformEventChannelMemberCreateResponse {
	id := d.Get("id").(string)

	return &models.PlatformEventChannelMemberCreateResponse{
		ID: id,
	}
}

// Function to perform the following actions:
func PlatformEventChannelMemberCreateResponseModelFromMap(m map[string]interface{}) *models.PlatformEventChannelMemberCreateResponse {
	id := m["id"].(string)

	return &models.PlatformEventChannelMemberCreateResponse{
		ID: id,
	}
}

func PlatformEventChannelMemberCreateResponseModelFromArrayOfMap(m []interface{}) []*models.PlatformEventChannelMemberCreateResponse {
	mapped := make([]*models.PlatformEventChannelMemberCreateResponse, len(m))
	for i, v := range m {
		mapped[i] = PlatformEventChannelMemberCreateResponseModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the PlatformEventChannelMemberCreateResponse resource
func GetPlatformEventChannelMemberCreateResponsePropertyFields() (t []string) {
	return []string{
		"id",
	}
}
