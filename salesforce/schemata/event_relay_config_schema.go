package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EventRelayConfig resource defined in the Terraform configuration
func EventRelayConfigSchema() map[string]*schema.Schema {
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
			Type: schema.TypeList, //GoType: EventRelayConfigMetadata
			Elem: &schema.Resource{
				Schema: EventRelayConfigMetadataSchema(),
			},
			Required: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and EventRelayConfigSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceEventRelayConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"id": {
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

// Update the underlying EventRelayConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventRelayConfigResourceData(d *schema.ResourceData, m *models.EventRelayConfig, isDataResource bool) {
	d.Set("full_name", m.FullName)
	if m.ID == "" && isDataResource {
		d.SetId("-")
	} else {
		d.SetId(m.ID)
	}
	d.Set("metadata", SetEventRelayConfigMetadataSubResourceData([]*models.EventRelayConfigMetadata{m.Metadata}))
}

// Iterate throught and update the EventRelayConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventRelayConfigSubResourceData(m []*models.EventRelayConfig) (d []*map[string]interface{}) {
	for _, eventRelayConfig := range m {
		if eventRelayConfig != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = eventRelayConfig.FullName
			properties["id"] = eventRelayConfig.ID
			properties["metadata"] = SetEventRelayConfigMetadataSubResourceData([]*models.EventRelayConfigMetadata{eventRelayConfig.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EventRelayConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventRelayConfigModel(d *schema.ResourceData) *models.EventRelayConfig {
	fullName := d.Get("full_name").(string)
	id := d.Get("id").(string)
	var metadata *models.EventRelayConfigMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = EventRelayConfigMetadataModelFromMap(MetadataMap)
	}

	return &models.EventRelayConfig{
		FullName: &fullName,
		ID:       id,
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func EventRelayConfigModelFromMap(m map[string]interface{}) *models.EventRelayConfig {
	fullName := m["full_name"].(string)
	id := m["id"].(string)
	var metadata *models.EventRelayConfigMetadata = nil //hit complex
	if m["metadata"] != nil {
		metadata = EventRelayConfigMetadataModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.EventRelayConfig{
		FullName: &fullName,
		ID:       id,
		Metadata: metadata,
	}
}

func EventRelayConfigModelFromArrayOfMap(m []interface{}) []*models.EventRelayConfig {
	mapped := make([]*models.EventRelayConfig, len(m))
	for i, v := range m {
		mapped[i] = EventRelayConfigModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the EventRelayConfig resource
func GetEventRelayConfigPropertyFields() (t []string) {
	return []string{
		"full_name",
		"id",
		"metadata",
	}
}
