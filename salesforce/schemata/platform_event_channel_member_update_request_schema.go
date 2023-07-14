package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannelMemberUpdateRequest resource defined in the Terraform configuration
func PlatformEventChannelMemberUpdateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: PlatformEventChannelMemberMetadata
			Elem: &schema.Resource{
				Schema: PlatformEventChannelMemberMetadataSchema(),
			},
			Optional: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and PlatformEventChannelMemberUpdateRequestSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourcePlatformEventChannelMemberUpdateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: PlatformEventChannelMemberMetadata
			Elem: &schema.Resource{
				Schema: PlatformEventChannelMemberMetadataSchema(),
			},
			Optional: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying PlatformEventChannelMemberUpdateRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelMemberUpdateRequestResourceData(d *schema.ResourceData, m *models.PlatformEventChannelMemberUpdateRequest) {
	d.Set("full_name", m.FullName)
	d.Set("metadata", SetPlatformEventChannelMemberMetadataSubResourceData([]*models.PlatformEventChannelMemberMetadata{m.Metadata}))
}

// Iterate throught and update the PlatformEventChannelMemberUpdateRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelMemberUpdateRequestSubResourceData(m []*models.PlatformEventChannelMemberUpdateRequest) (d []*map[string]interface{}) {
	for _, platformEventChannelMemberUpdateRequest := range m {
		if platformEventChannelMemberUpdateRequest != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = platformEventChannelMemberUpdateRequest.FullName
			properties["metadata"] = SetPlatformEventChannelMemberMetadataSubResourceData([]*models.PlatformEventChannelMemberMetadata{platformEventChannelMemberUpdateRequest.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannelMemberUpdateRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelMemberUpdateRequestModel(d *schema.ResourceData) *models.PlatformEventChannelMemberUpdateRequest {
	fullName := d.Get("full_name").(string)
	var metadata *models.PlatformEventChannelMemberMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		var m schema.ResourceData
		m.Set("meta", MetadataMap)
		metadata = PlatformEventChannelMemberMetadataModel(&m)
	}

	return &models.PlatformEventChannelMemberUpdateRequest{
		FullName: fullName,
		Metadata: metadata,
	}
}

// Retrieve property field names for updating the PlatformEventChannelMemberUpdateRequest resource
func GetPlatformEventChannelMemberUpdateRequestPropertyFields() (t []string) {
	return []string{
		"full_name",
		"metadata",
	}
}
