package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the NamedCredentialUpdateRequest resource defined in the Terraform configuration
func NamedCredentialUpdateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: NamedCredentialMetadata
			Elem: &schema.Resource{
				Schema: NamedCredentialMetadataSchema(),
			},
			Optional: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and NamedCredentialUpdateRequestSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceNamedCredentialUpdateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: NamedCredentialMetadata
			Elem: &schema.Resource{
				Schema: NamedCredentialMetadataSchema(),
			},
			Optional: true,
			Computed: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying NamedCredentialUpdateRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNamedCredentialUpdateRequestResourceData(d *schema.ResourceData, m *models.NamedCredentialUpdateRequest, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("full_name", m.FullName)
	d.Set("metadata", SetNamedCredentialMetadataSubResourceData([]*models.NamedCredentialMetadata{m.Metadata}))
}

// Iterate throught and update the NamedCredentialUpdateRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNamedCredentialUpdateRequestSubResourceData(m []*models.NamedCredentialUpdateRequest) (d []*map[string]interface{}) {
	for _, namedCredentialUpdateRequest := range m {
		if namedCredentialUpdateRequest != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = namedCredentialUpdateRequest.FullName
			properties["metadata"] = SetNamedCredentialMetadataSubResourceData([]*models.NamedCredentialMetadata{namedCredentialUpdateRequest.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate NamedCredentialUpdateRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NamedCredentialUpdateRequestModel(d *schema.ResourceData) *models.NamedCredentialUpdateRequest {
	fullName := d.Get("full_name").(string)
	var metadata *models.NamedCredentialMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = NamedCredentialMetadataModelFromMap(MetadataMap)
	}

	return &models.NamedCredentialUpdateRequest{
		FullName: fullName,
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func NamedCredentialUpdateRequestModelFromMap(m map[string]interface{}) *models.NamedCredentialUpdateRequest {
	fullName := m["full_name"].(string)
	var metadata *models.NamedCredentialMetadata = nil //hit complex
	if m["metadata"] != nil {
		metadata = NamedCredentialMetadataModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.NamedCredentialUpdateRequest{
		FullName: fullName,
		Metadata: metadata,
	}
}

func NamedCredentialUpdateRequestModelFromArrayOfMap(m []interface{}) []*models.NamedCredentialUpdateRequest {
	mapped := make([]*models.NamedCredentialUpdateRequest, len(m))
	for i, v := range m {
		mapped[i] = NamedCredentialUpdateRequestModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the NamedCredentialUpdateRequest resource
func GetNamedCredentialUpdateRequestPropertyFields() (t []string) {
	return []string{
		"full_name",
		"metadata",
	}
}
