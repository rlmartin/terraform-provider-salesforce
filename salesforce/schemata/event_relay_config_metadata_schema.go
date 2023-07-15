package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EventRelayConfigMetadata resource defined in the Terraform configuration
func EventRelayConfigMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"destination_resource_name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"event_channel": {
			Type:     schema.TypeString,
			Required: true,
		},

		"relay_option": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying EventRelayConfigMetadata resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventRelayConfigMetadataResourceData(d *schema.ResourceData, m *models.EventRelayConfigMetadata) {
	d.Set("destination_resource_name", m.DestinationResourceName)
	d.Set("event_channel", m.EventChannel)
	d.Set("relay_option", m.RelayOption)
}

// Iterate throught and update the EventRelayConfigMetadata resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventRelayConfigMetadataSubResourceData(m []*models.EventRelayConfigMetadata) (d []*map[string]interface{}) {
	for _, eventRelayConfigMetadata := range m {
		if eventRelayConfigMetadata != nil {
			properties := make(map[string]interface{})
			properties["destination_resource_name"] = eventRelayConfigMetadata.DestinationResourceName
			properties["event_channel"] = eventRelayConfigMetadata.EventChannel
			properties["relay_option"] = eventRelayConfigMetadata.RelayOption
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EventRelayConfigMetadata resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventRelayConfigMetadataModel(d *schema.ResourceData) *models.EventRelayConfigMetadata {
	destinationResourceName := d.Get("destination_resource_name").(string)
	eventChannel := d.Get("event_channel").(string)
	relayOption := d.Get("relay_option").(string)

	return &models.EventRelayConfigMetadata{
		DestinationResourceName: &destinationResourceName,
		EventChannel:            &eventChannel,
		RelayOption:             &relayOption,
	}
}

// Function to perform the following actions:
func EventRelayConfigMetadataModelFromMap(m map[string]interface{}) *models.EventRelayConfigMetadata {
	destinationResourceName := m["destination_resource_name"].(string)
	eventChannel := m["event_channel"].(string)
	relayOption := m["relay_option"].(string)

	return &models.EventRelayConfigMetadata{
		DestinationResourceName: &destinationResourceName,
		EventChannel:            &eventChannel,
		RelayOption:             &relayOption,
	}
}

// Retrieve property field names for updating the EventRelayConfigMetadata resource
func GetEventRelayConfigMetadataPropertyFields() (t []string) {
	return []string{
		"destination_resource_name",
		"event_channel",
		"relay_option",
	}
}
