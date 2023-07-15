package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the NamedCredentialMetadata resource defined in the Terraform configuration
func NamedCredentialMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_merge_fields_in_body": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"allow_merge_fields_in_header": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"endpoint": {
			Type:     schema.TypeString,
			Required: true,
		},

		"generate_authorization_header": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"label": {
			Type:     schema.TypeString,
			Required: true,
		},

		"principal_type": {
			Type:     schema.TypeString,
			Required: true,
		},

		"protocol": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying NamedCredentialMetadata resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNamedCredentialMetadataResourceData(d *schema.ResourceData, m *models.NamedCredentialMetadata) {
	d.Set("allow_merge_fields_in_body", m.AllowMergeFieldsInBody)
	d.Set("allow_merge_fields_in_header", m.AllowMergeFieldsInHeader)
	d.Set("endpoint", m.Endpoint)
	d.Set("generate_authorization_header", m.GenerateAuthorizationHeader)
	d.Set("label", m.Label)
	d.Set("principal_type", m.PrincipalType)
	d.Set("protocol", m.Protocol)
}

// Iterate throught and update the NamedCredentialMetadata resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNamedCredentialMetadataSubResourceData(m []*models.NamedCredentialMetadata) (d []*map[string]interface{}) {
	for _, namedCredentialMetadata := range m {
		if namedCredentialMetadata != nil {
			properties := make(map[string]interface{})
			properties["allow_merge_fields_in_body"] = namedCredentialMetadata.AllowMergeFieldsInBody
			properties["allow_merge_fields_in_header"] = namedCredentialMetadata.AllowMergeFieldsInHeader
			properties["endpoint"] = namedCredentialMetadata.Endpoint
			properties["generate_authorization_header"] = namedCredentialMetadata.GenerateAuthorizationHeader
			properties["label"] = namedCredentialMetadata.Label
			properties["principal_type"] = namedCredentialMetadata.PrincipalType
			properties["protocol"] = namedCredentialMetadata.Protocol
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate NamedCredentialMetadata resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NamedCredentialMetadataModel(d *schema.ResourceData) *models.NamedCredentialMetadata {
	allowMergeFieldsInBody := d.Get("allow_merge_fields_in_body").(bool)
	allowMergeFieldsInHeader := d.Get("allow_merge_fields_in_header").(bool)
	endpoint := d.Get("endpoint").(string)
	generateAuthorizationHeader := d.Get("generate_authorization_header").(bool)
	label := d.Get("label").(string)
	principalType := d.Get("principal_type").(string)
	protocol := d.Get("protocol").(string)

	return &models.NamedCredentialMetadata{
		AllowMergeFieldsInBody:      allowMergeFieldsInBody,
		AllowMergeFieldsInHeader:    allowMergeFieldsInHeader,
		Endpoint:                    &endpoint,
		GenerateAuthorizationHeader: generateAuthorizationHeader,
		Label:                       &label,
		PrincipalType:               &principalType,
		Protocol:                    &protocol,
	}
}

// Function to perform the following actions:
func NamedCredentialMetadataModelFromMap(m map[string]interface{}) *models.NamedCredentialMetadata {
	allowMergeFieldsInBody := m["allow_merge_fields_in_body"].(bool)
	allowMergeFieldsInHeader := m["allow_merge_fields_in_header"].(bool)
	endpoint := m["endpoint"].(string)
	generateAuthorizationHeader := m["generate_authorization_header"].(bool)
	label := m["label"].(string)
	principalType := m["principal_type"].(string)
	protocol := m["protocol"].(string)

	return &models.NamedCredentialMetadata{
		AllowMergeFieldsInBody:      allowMergeFieldsInBody,
		AllowMergeFieldsInHeader:    allowMergeFieldsInHeader,
		Endpoint:                    &endpoint,
		GenerateAuthorizationHeader: generateAuthorizationHeader,
		Label:                       &label,
		PrincipalType:               &principalType,
		Protocol:                    &protocol,
	}
}

// Retrieve property field names for updating the NamedCredentialMetadata resource
func GetNamedCredentialMetadataPropertyFields() (t []string) {
	return []string{
		"allow_merge_fields_in_body",
		"allow_merge_fields_in_header",
		"endpoint",
		"generate_authorization_header",
		"label",
		"principal_type",
		"protocol",
	}
}
