package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannelMetadata resource defined in the Terraform configuration
func PlatformEventChannelMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"channel_type": {
			Type:     schema.TypeString,
			Required: true,
		},

		"label": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying PlatformEventChannelMetadata resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelMetadataResourceData(d *schema.ResourceData, m *models.PlatformEventChannelMetadata) {
	d.Set("channel_type", m.ChannelType)
	d.Set("label", m.Label)
}

// Iterate throught and update the PlatformEventChannelMetadata resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelMetadataSubResourceData(m []*models.PlatformEventChannelMetadata) (d []*map[string]interface{}) {
	for _, platformEventChannelMetadata := range m {
		if platformEventChannelMetadata != nil {
			properties := make(map[string]interface{})
			properties["channel_type"] = platformEventChannelMetadata.ChannelType
			properties["label"] = platformEventChannelMetadata.Label
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannelMetadata resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelMetadataModel(d *schema.ResourceData) *models.PlatformEventChannelMetadata {
	channelType := d.Get("channel_type").(string)
	label := d.Get("label").(string)

	return &models.PlatformEventChannelMetadata{
		ChannelType: &channelType,
		Label:       &label,
	}
}

// Function to perform the following actions:
func PlatformEventChannelMetadataModelFromMap(m map[string]interface{}) *models.PlatformEventChannelMetadata {
	channelType := m["channel_type"].(string)
	label := m["label"].(string)

	return &models.PlatformEventChannelMetadata{
		ChannelType: &channelType,
		Label:       &label,
	}
}

// Retrieve property field names for updating the PlatformEventChannelMetadata resource
func GetPlatformEventChannelMetadataPropertyFields() (t []string) {
	return []string{
		"channel_type",
		"label",
	}
}
