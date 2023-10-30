package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannelCreateRequest resource defined in the Terraform configuration
func PlatformEventChannelCreateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Required: true,
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
// Only difference between this and PlatformEventChannelCreateRequestSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourcePlatformEventChannelCreateRequestSchema() map[string]*schema.Schema {
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

// Update the underlying PlatformEventChannelCreateRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelCreateRequestResourceData(d *schema.ResourceData, m *models.PlatformEventChannelCreateRequest, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("full_name", m.FullName)
	d.Set("metadata", SetPlatformEventChannelMetadataSubResourceData([]*models.PlatformEventChannelMetadata{m.Metadata}))
}

// Iterate throught and update the PlatformEventChannelCreateRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelCreateRequestSubResourceData(m []*models.PlatformEventChannelCreateRequest) (d []*map[string]interface{}) {
	for _, platformEventChannelCreateRequest := range m {
		if platformEventChannelCreateRequest != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = platformEventChannelCreateRequest.FullName
			properties["metadata"] = SetPlatformEventChannelMetadataSubResourceData([]*models.PlatformEventChannelMetadata{platformEventChannelCreateRequest.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannelCreateRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelCreateRequestModel(d *schema.ResourceData) *models.PlatformEventChannelCreateRequest {
	fullName := d.Get("full_name").(string)
	var metadata *models.PlatformEventChannelMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = PlatformEventChannelMetadataModelFromMap(MetadataMap)
	}

	return &models.PlatformEventChannelCreateRequest{
		FullName: &fullName,
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func PlatformEventChannelCreateRequestModelFromMap(m map[string]interface{}) *models.PlatformEventChannelCreateRequest {
	fullName := m["full_name"].(string)
	var metadata *models.PlatformEventChannelMetadata = nil //hit complex
	if m["metadata"] != nil {
		metadata = PlatformEventChannelMetadataModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.PlatformEventChannelCreateRequest{
		FullName: &fullName,
		Metadata: metadata,
	}
}

func PlatformEventChannelCreateRequestModelFromArrayOfMap(m []interface{}) []*models.PlatformEventChannelCreateRequest {
	mapped := make([]*models.PlatformEventChannelCreateRequest, len(m))
	for i, v := range m {
		mapped[i] = PlatformEventChannelCreateRequestModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the PlatformEventChannelCreateRequest resource
func GetPlatformEventChannelCreateRequestPropertyFields() (t []string) {
	return []string{
		"full_name",
		"metadata",
	}
}
