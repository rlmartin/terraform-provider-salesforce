package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EventRelayConfigCreateRequest resource defined in the Terraform configuration
func EventRelayConfigCreateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: EventRelayConfigMetadata
			Elem: &schema.Resource{
				Schema: EventRelayConfigMetadataSchema(),
			},
			Required: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and EventRelayConfigCreateRequestSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceEventRelayConfigCreateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: EventRelayConfigMetadata
			Elem: &schema.Resource{
				Schema: EventRelayConfigMetadataSchema(),
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

// Update the underlying EventRelayConfigCreateRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventRelayConfigCreateRequestResourceData(d *schema.ResourceData, m *models.EventRelayConfigCreateRequest, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("full_name", m.FullName)
	d.Set("metadata", SetEventRelayConfigMetadataSubResourceData([]*models.EventRelayConfigMetadata{m.Metadata}))
}

// Iterate throught and update the EventRelayConfigCreateRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventRelayConfigCreateRequestSubResourceData(m []*models.EventRelayConfigCreateRequest) (d []*map[string]interface{}) {
	for _, eventRelayConfigCreateRequest := range m {
		if eventRelayConfigCreateRequest != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = eventRelayConfigCreateRequest.FullName
			properties["metadata"] = SetEventRelayConfigMetadataSubResourceData([]*models.EventRelayConfigMetadata{eventRelayConfigCreateRequest.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EventRelayConfigCreateRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventRelayConfigCreateRequestModel(d *schema.ResourceData) *models.EventRelayConfigCreateRequest {
	fullName := d.Get("full_name").(string)
	var metadata *models.EventRelayConfigMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = EventRelayConfigMetadataModelFromMap(MetadataMap)
	}

	return &models.EventRelayConfigCreateRequest{
		FullName: &fullName,
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func EventRelayConfigCreateRequestModelFromMap(m map[string]interface{}) *models.EventRelayConfigCreateRequest {
	fullName := m["full_name"].(string)
	var metadata *models.EventRelayConfigMetadata = nil //hit complex
	if m["metadata"] != nil {
		metadata = EventRelayConfigMetadataModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.EventRelayConfigCreateRequest{
		FullName: &fullName,
		Metadata: metadata,
	}
}

// Retrieve property field names for updating the EventRelayConfigCreateRequest resource
func GetEventRelayConfigCreateRequestPropertyFields() (t []string) {
	return []string{
		"full_name",
		"metadata",
	}
}
