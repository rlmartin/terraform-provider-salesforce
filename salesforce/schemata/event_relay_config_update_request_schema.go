package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EventRelayConfigUpdateRequest resource defined in the Terraform configuration
func EventRelayConfigUpdateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metadata": {
			Type: schema.TypeList, //GoType: EventRelayConfigMetadataUpdate
			Elem: &schema.Resource{
				Schema: EventRelayConfigMetadataUpdateSchema(),
			},
			Optional: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and EventRelayConfigUpdateRequestSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceEventRelayConfigUpdateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metadata": {
			Type: schema.TypeList, //GoType: EventRelayConfigMetadataUpdate
			Elem: &schema.Resource{
				Schema: EventRelayConfigMetadataUpdateSchema(),
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

// Update the underlying EventRelayConfigUpdateRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventRelayConfigUpdateRequestResourceData(d *schema.ResourceData, m *models.EventRelayConfigUpdateRequest, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("metadata", SetEventRelayConfigMetadataUpdateSubResourceData([]*models.EventRelayConfigMetadataUpdate{m.Metadata}))
}

// Iterate throught and update the EventRelayConfigUpdateRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventRelayConfigUpdateRequestSubResourceData(m []*models.EventRelayConfigUpdateRequest) (d []*map[string]interface{}) {
	for _, eventRelayConfigUpdateRequest := range m {
		if eventRelayConfigUpdateRequest != nil {
			properties := make(map[string]interface{})
			properties["metadata"] = SetEventRelayConfigMetadataUpdateSubResourceData([]*models.EventRelayConfigMetadataUpdate{eventRelayConfigUpdateRequest.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EventRelayConfigUpdateRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventRelayConfigUpdateRequestModel(d *schema.ResourceData) *models.EventRelayConfigUpdateRequest {
	var metadata *models.EventRelayConfigMetadataUpdate = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = EventRelayConfigMetadataUpdateModelFromMap(MetadataMap)
	}

	return &models.EventRelayConfigUpdateRequest{
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func EventRelayConfigUpdateRequestModelFromMap(m map[string]interface{}) *models.EventRelayConfigUpdateRequest {
	var metadata *models.EventRelayConfigMetadataUpdate = nil //hit complex
	if m["metadata"] != nil {
		metadata = EventRelayConfigMetadataUpdateModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.EventRelayConfigUpdateRequest{
		Metadata: metadata,
	}
}

// Retrieve property field names for updating the EventRelayConfigUpdateRequest resource
func GetEventRelayConfigUpdateRequestPropertyFields() (t []string) {
	return []string{
		"metadata",
	}
}
