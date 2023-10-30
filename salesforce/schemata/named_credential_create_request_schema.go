package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the NamedCredentialCreateRequest resource defined in the Terraform configuration
func NamedCredentialCreateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: NamedCredentialMetadata
			Elem: &schema.Resource{
				Schema: NamedCredentialMetadataSchema(),
			},
			Required: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and NamedCredentialCreateRequestSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceNamedCredentialCreateRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"metadata": {
			Type: schema.TypeList, //GoType: NamedCredentialMetadata
			Elem: &schema.Resource{
				Schema: NamedCredentialMetadataSchema(),
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

// Update the underlying NamedCredentialCreateRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNamedCredentialCreateRequestResourceData(d *schema.ResourceData, m *models.NamedCredentialCreateRequest, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("full_name", m.FullName)
	d.Set("metadata", SetNamedCredentialMetadataSubResourceData([]*models.NamedCredentialMetadata{m.Metadata}))
}

// Iterate throught and update the NamedCredentialCreateRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNamedCredentialCreateRequestSubResourceData(m []*models.NamedCredentialCreateRequest) (d []*map[string]interface{}) {
	for _, namedCredentialCreateRequest := range m {
		if namedCredentialCreateRequest != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = namedCredentialCreateRequest.FullName
			properties["metadata"] = SetNamedCredentialMetadataSubResourceData([]*models.NamedCredentialMetadata{namedCredentialCreateRequest.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate NamedCredentialCreateRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NamedCredentialCreateRequestModel(d *schema.ResourceData) *models.NamedCredentialCreateRequest {
	fullName := d.Get("full_name").(string)
	var metadata *models.NamedCredentialMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = NamedCredentialMetadataModelFromMap(MetadataMap)
	}

	return &models.NamedCredentialCreateRequest{
		FullName: &fullName,
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func NamedCredentialCreateRequestModelFromMap(m map[string]interface{}) *models.NamedCredentialCreateRequest {
	fullName := m["full_name"].(string)
	var metadata *models.NamedCredentialMetadata = nil //hit complex
	if m["metadata"] != nil {
		metadata = NamedCredentialMetadataModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.NamedCredentialCreateRequest{
		FullName: &fullName,
		Metadata: metadata,
	}
}

func NamedCredentialCreateRequestModelFromArrayOfMap(m []interface{}) []*models.NamedCredentialCreateRequest {
	mapped := make([]*models.NamedCredentialCreateRequest, len(m))
	for i, v := range m {
		mapped[i] = NamedCredentialCreateRequestModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the NamedCredentialCreateRequest resource
func GetNamedCredentialCreateRequestPropertyFields() (t []string) {
	return []string{
		"full_name",
		"metadata",
	}
}
