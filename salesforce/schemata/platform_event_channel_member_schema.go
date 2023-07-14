package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the PlatformEventChannelMember resource defined in the Terraform configuration
func PlatformEventChannelMemberSchema() map[string]*schema.Schema {
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
			Type: schema.TypeList, //GoType: PlatformEventChannelMemberMetadata
			Elem: &schema.Resource{
				Schema: PlatformEventChannelMemberMetadataSchema(),
			},
			Required: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and PlatformEventChannelMemberSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourcePlatformEventChannelMemberSchema() map[string]*schema.Schema {
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

// Update the underlying PlatformEventChannelMember resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPlatformEventChannelMemberResourceData(d *schema.ResourceData, m *models.PlatformEventChannelMember) {
	d.Set("full_name", m.FullName)
	d.Set("id", m.ID)
	d.Set("metadata", SetPlatformEventChannelMemberMetadataSubResourceData([]*models.PlatformEventChannelMemberMetadata{m.Metadata}))
}

// Iterate throught and update the PlatformEventChannelMember resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPlatformEventChannelMemberSubResourceData(m []*models.PlatformEventChannelMember) (d []*map[string]interface{}) {
	for _, platformEventChannelMember := range m {
		if platformEventChannelMember != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = platformEventChannelMember.FullName
			properties["id"] = platformEventChannelMember.ID
			properties["metadata"] = SetPlatformEventChannelMemberMetadataSubResourceData([]*models.PlatformEventChannelMemberMetadata{platformEventChannelMember.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PlatformEventChannelMember resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PlatformEventChannelMemberModel(d *schema.ResourceData) *models.PlatformEventChannelMember {
	fullName := d.Get("full_name").(string)
	id := d.Get("id").(string)
	var metadata *models.PlatformEventChannelMemberMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = PlatformEventChannelMemberMetadataModelFromMap(MetadataMap)
	}

	return &models.PlatformEventChannelMember{
		FullName: &fullName,
		ID:       id,
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func PlatformEventChannelMemberModelFromMap(m map[string]interface{}) *models.PlatformEventChannelMember {
	fullName := m["full_name"].(string)
	id := m["id"].(string)
	var metadata *models.PlatformEventChannelMemberMetadata = nil //hit complex
	if m["metadata"] != nil {
		metadata = PlatformEventChannelMemberMetadataModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.PlatformEventChannelMember{
		FullName: &fullName,
		ID:       id,
		Metadata: metadata,
	}
}

// Retrieve property field names for updating the PlatformEventChannelMember resource
func GetPlatformEventChannelMemberPropertyFields() (t []string) {
	return []string{
		"full_name",
		"id",
		"metadata",
	}
}
