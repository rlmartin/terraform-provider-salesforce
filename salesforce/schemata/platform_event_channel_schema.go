package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannel resource defined in the Terraform configuration
func PlatformEventChannelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: PlatformEventChannelMetadata
			Elem: &schema.Resource{
				Schema: PlatformEventChannelMetadataSchema(),
			},
			Required: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and PlatformEventChannelSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourcePlatformEventChannelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
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

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying PlatformEventChannel resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelResourceData(d *schema.ResourceData, m *models.PlatformEventChannel) {
	d.Set("full_name", m.FullName)
	d.Set("id", m.ID)
	d.Set("metadata", SetPlatformEventChannelMetadataSubResourceData([]*models.PlatformEventChannelMetadata{m.Metadata}))
}

// Iterate throught and update the PlatformEventChannel resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelSubResourceData(m []*models.PlatformEventChannel) (d []*map[string]interface{}) {
	for _, platformEventChannel := range m {
		if platformEventChannel != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = platformEventChannel.FullName
			properties["id"] = platformEventChannel.ID
			properties["metadata"] = SetPlatformEventChannelMetadataSubResourceData([]*models.PlatformEventChannelMetadata{platformEventChannel.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannel resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelModel(d *schema.ResourceData) *models.PlatformEventChannel {
	fullName := d.Get("full_name").(string)
	id := d.Get("id").(string)
	var metadata *models.PlatformEventChannelMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		var m schema.ResourceData
		m.Set("meta", MetadataMap)
		metadata = PlatformEventChannelMetadataModel(&m)
	}

	return &models.PlatformEventChannel{
		FullName: &fullName,
		ID:       id,
		Metadata: metadata,
	}
}

// Retrieve property field names for updating the PlatformEventChannel resource
func GetPlatformEventChannelPropertyFields() (t []string) {
	return []string{
		"full_name",
		"id",
		"metadata",
	}
}
