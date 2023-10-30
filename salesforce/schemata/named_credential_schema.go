package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the NamedCredential resource defined in the Terraform configuration
func NamedCredentialSchema() map[string]*schema.Schema {
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
			Type: schema.TypeList, //GoType: NamedCredentialMetadata
			Elem: &schema.Resource{
				Schema: NamedCredentialMetadataSchema(),
			},
			Required: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and NamedCredentialSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceNamedCredentialSchema() map[string]*schema.Schema {
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

// Update the underlying NamedCredential resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNamedCredentialResourceData(d *schema.ResourceData, m *models.NamedCredential, isDataResource bool) {
	d.Set("full_name", m.FullName)
	if m.ID == "" && isDataResource {
		d.SetId("-")
	} else {
		d.SetId(m.ID)
	}
	d.Set("metadata", SetNamedCredentialMetadataSubResourceData([]*models.NamedCredentialMetadata{m.Metadata}))
}

// Iterate throught and update the NamedCredential resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNamedCredentialSubResourceData(m []*models.NamedCredential) (d []*map[string]interface{}) {
	for _, namedCredential := range m {
		if namedCredential != nil {
			properties := make(map[string]interface{})
			properties["full_name"] = namedCredential.FullName
			properties["id"] = namedCredential.ID
			properties["metadata"] = SetNamedCredentialMetadataSubResourceData([]*models.NamedCredentialMetadata{namedCredential.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate NamedCredential resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NamedCredentialModel(d *schema.ResourceData) *models.NamedCredential {
	fullName := d.Get("full_name").(string)
	id := d.Get("id").(string)
	var metadata *models.NamedCredentialMetadata = nil //hit complex
	MetadataInterface, MetadataIsSet := d.GetOk("metadata")
	if MetadataIsSet {
		MetadataMap := MetadataInterface.([]interface{})[0].(map[string]interface{})
		metadata = NamedCredentialMetadataModelFromMap(MetadataMap)
	}

	return &models.NamedCredential{
		FullName: &fullName,
		ID:       id,
		Metadata: metadata,
	}
}

// Function to perform the following actions:
func NamedCredentialModelFromMap(m map[string]interface{}) *models.NamedCredential {
	fullName := m["full_name"].(string)
	id := m["id"].(string)
	var metadata *models.NamedCredentialMetadata = nil //hit complex
	if m["metadata"] != nil {
		metadata = NamedCredentialMetadataModelFromMap(m["metadata"].(map[string]interface{}))
	}

	return &models.NamedCredential{
		FullName: &fullName,
		ID:       id,
		Metadata: metadata,
	}
}

func NamedCredentialModelFromArrayOfMap(m []interface{}) []*models.NamedCredential {
	mapped := make([]*models.NamedCredential, len(m))
	for i, v := range m {
		mapped[i] = NamedCredentialModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the NamedCredential resource
func GetNamedCredentialPropertyFields() (t []string) {
	return []string{
		"full_name",
		"id",
		"metadata",
	}
}
