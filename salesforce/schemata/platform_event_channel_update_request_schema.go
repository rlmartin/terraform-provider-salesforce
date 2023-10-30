package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannelUpdateRequest resource defined in the Terraform configuration
func PlatformEventChannelUpdateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: PlatformEventChannelMetadata
			Elem: &schema.Resource{
				Schema: PlatformEventChannelMetadataSchema(),
			},
			Optional: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and PlatformEventChannelUpdateRequestSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourcePlatformEventChannelUpdateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: PlatformEventChannelMetadata
			Elem: &schema.Resource{
				Schema: PlatformEventChannelMetadataSchema(),
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

// Update the underlying PlatformEventChannelUpdateRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelUpdateRequestResourceData(d *schema.ResourceData, m *models.PlatformEventChannelUpdateRequest, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("full_name", m.FullName)
	d.Set("metadata", SetPlatformEventChannelMetadataSubResourceData([]*models.PlatformEventChannelMetadata{m.Metadata}))
}

// Iterate throught and update the PlatformEventChannelUpdateRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelUpdateRequestSubResourceData(m []*models.PlatformEventChannelUpdateRequest) (d []*map[string]interface{}) {
	for _, platformEventChannelUpdateRequest := range m {
		if platformEventChannelUpdateRequest != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = platformEventChannelUpdateRequest.FullName
			properties["metadata"] = SetPlatformEventChannelMetadataSubResourceData([]*models.PlatformEventChannelMetadata{platformEventChannelUpdateRequest.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannelUpdateRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelUpdateRequestModel(d *schema.ResourceData) *models.PlatformEventChannelUpdateRequest {
	fullName := d.Get("full_name").(string)
	var metadata *models.PlatformEventChannelMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = PlatformEventChannelMetadataModelFromMap(MetadataMap)
	}

	return &models.PlatformEventChannelUpdateRequest{
		FullName: fullName,
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func PlatformEventChannelUpdateRequestModelFromMap(m map[string]interface{}) *models.PlatformEventChannelUpdateRequest {
	fullName := m["full_name"].(string)
	var metadata *models.PlatformEventChannelMetadata = nil //hit complex
	if m["metadata"] != nil {
		metadata = PlatformEventChannelMetadataModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.PlatformEventChannelUpdateRequest{
		FullName: fullName,
		Metadata: metadata,
	}
}

func PlatformEventChannelUpdateRequestModelFromArrayOfMap(m []interface{}) []*models.PlatformEventChannelUpdateRequest {
	mapped := make([]*models.PlatformEventChannelUpdateRequest, len(m))
	for i, v := range m {
		mapped[i] = PlatformEventChannelUpdateRequestModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the PlatformEventChannelUpdateRequest resource
func GetPlatformEventChannelUpdateRequestPropertyFields() (t []string) {
	return []string{
		"full_name",
		"metadata",
	}
}
