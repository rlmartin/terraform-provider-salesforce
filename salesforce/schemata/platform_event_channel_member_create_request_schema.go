package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannelMemberCreateRequest resource defined in the Terraform configuration
func PlatformEventChannelMemberCreateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: PlatformEventChannelMemberMetadata
			Elem: &schema.Resource{
				Schema: PlatformEventChannelMemberMetadataSchema(),
			},
			Required: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and PlatformEventChannelMemberCreateRequestSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourcePlatformEventChannelMemberCreateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: PlatformEventChannelMemberMetadata
			Elem: &schema.Resource{
				Schema: PlatformEventChannelMemberMetadataSchema(),
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

// Update the underlying PlatformEventChannelMemberCreateRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelMemberCreateRequestResourceData(d *schema.ResourceData, m *models.PlatformEventChannelMemberCreateRequest, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("full_name", m.FullName)
	d.Set("metadata", SetPlatformEventChannelMemberMetadataSubResourceData([]*models.PlatformEventChannelMemberMetadata{m.Metadata}))
}

// Iterate throught and update the PlatformEventChannelMemberCreateRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelMemberCreateRequestSubResourceData(m []*models.PlatformEventChannelMemberCreateRequest) (d []*map[string]interface{}) {
	for _, platformEventChannelMemberCreateRequest := range m {
		if platformEventChannelMemberCreateRequest != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = platformEventChannelMemberCreateRequest.FullName
			properties["metadata"] = SetPlatformEventChannelMemberMetadataSubResourceData([]*models.PlatformEventChannelMemberMetadata{platformEventChannelMemberCreateRequest.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannelMemberCreateRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelMemberCreateRequestModel(d *schema.ResourceData) *models.PlatformEventChannelMemberCreateRequest {
	fullName := d.Get("full_name").(string)
	var metadata *models.PlatformEventChannelMemberMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = PlatformEventChannelMemberMetadataModelFromMap(MetadataMap)
	}

	return &models.PlatformEventChannelMemberCreateRequest{
		FullName: &fullName,
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func PlatformEventChannelMemberCreateRequestModelFromMap(m map[string]interface{}) *models.PlatformEventChannelMemberCreateRequest {
	fullName := m["full_name"].(string)
	var metadata *models.PlatformEventChannelMemberMetadata = nil //hit complex
	if m["metadata"] != nil {
		metadata = PlatformEventChannelMemberMetadataModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.PlatformEventChannelMemberCreateRequest{
		FullName: &fullName,
		Metadata: metadata,
	}
}

// Retrieve property field names for updating the PlatformEventChannelMemberCreateRequest resource
func GetPlatformEventChannelMemberCreateRequestPropertyFields() (t []string) {
	return []string{
		"full_name",
		"metadata",
	}
}
