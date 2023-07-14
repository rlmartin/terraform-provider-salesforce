package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannelMemberMetadata resource defined in the Terraform configuration
func PlatformEventChannelMemberMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"event_channel": {
			Type:     schema.TypeString,
			Required: true,
		},

		"selected_entity": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying PlatformEventChannelMemberMetadata resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelMemberMetadataResourceData(d *schema.ResourceData, m *models.PlatformEventChannelMemberMetadata) {
	d.Set("event_channel", m.EventChannel)
	d.Set("selected_entity", m.SelectedEntity)
}

// Iterate throught and update the PlatformEventChannelMemberMetadata resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelMemberMetadataSubResourceData(m []*models.PlatformEventChannelMemberMetadata) (d []*map[string]interface{}) {
	for _, platformEventChannelMemberMetadata := range m {
		if platformEventChannelMemberMetadata != nil {
			properties := make(map[string]interface{})
			properties["event_channel"] = platformEventChannelMemberMetadata.EventChannel
			properties["selected_entity"] = platformEventChannelMemberMetadata.SelectedEntity
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannelMemberMetadata resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelMemberMetadataModel(d *schema.ResourceData) *models.PlatformEventChannelMemberMetadata {
	eventChannel := d.Get("event_channel").(string)
	selectedEntity := d.Get("selected_entity").(string)

	return &models.PlatformEventChannelMemberMetadata{
		EventChannel:   &eventChannel,
		SelectedEntity: &selectedEntity,
	}
}

// Function to perform the following actions:
func PlatformEventChannelMemberMetadataModelFromMap(m map[string]interface{}) *models.PlatformEventChannelMemberMetadata {
	eventChannel := m["event_channel"].(string)
	selectedEntity := m["selected_entity"].(string)

	return &models.PlatformEventChannelMemberMetadata{
		EventChannel:   &eventChannel,
		SelectedEntity: &selectedEntity,
	}
}

// Retrieve property field names for updating the PlatformEventChannelMemberMetadata resource
func GetPlatformEventChannelMemberMetadataPropertyFields() (t []string) {
	return []string{
		"event_channel",
		"selected_entity",
	}
}
