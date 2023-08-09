package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the NamedCredentialCreateResponse resource defined in the Terraform configuration
func NamedCredentialCreateResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and NamedCredentialCreateResponseSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceNamedCredentialCreateResponseSchema() map[string]*schema.Schema {
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

// Update the underlying NamedCredentialCreateResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNamedCredentialCreateResponseResourceData(d *schema.ResourceData, m *models.NamedCredentialCreateResponse, isDataResource bool) {
	if m.ID == "" && isDataResource {
		d.SetId("-")
	} else {
		d.SetId(m.ID)
	}
}

// Iterate throught and update the NamedCredentialCreateResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNamedCredentialCreateResponseSubResourceData(m []*models.NamedCredentialCreateResponse) (d []*map[string]interface{}) {
	for _, namedCredentialCreateResponse := range m {
		if namedCredentialCreateResponse != nil {
			properties := make(map[string]interface{})
			properties["id"] = namedCredentialCreateResponse.ID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate NamedCredentialCreateResponse resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NamedCredentialCreateResponseModel(d *schema.ResourceData) *models.NamedCredentialCreateResponse {
	id := d.Get("id").(string)

	return &models.NamedCredentialCreateResponse{
		ID: id,
	}
}

// Function to perform the following actions:
func NamedCredentialCreateResponseModelFromMap(m map[string]interface{}) *models.NamedCredentialCreateResponse {
	id := m["id"].(string)

	return &models.NamedCredentialCreateResponse{
		ID: id,
	}
}

// Retrieve property field names for updating the NamedCredentialCreateResponse resource
func GetNamedCredentialCreateResponsePropertyFields() (t []string) {
	return []string{
		"id",
	}
}
