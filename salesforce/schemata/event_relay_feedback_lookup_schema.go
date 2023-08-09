package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EventRelayFeedbackLookup resource defined in the Terraform configuration
func EventRelayFeedbackLookupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"done": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"q": {
			Type:     schema.TypeString,
			Required: true,
		},

		"records": {
			Type: schema.TypeList, //GoType: []*Record
			Elem: &schema.Resource{
				Schema: RecordSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"total_size": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and EventRelayFeedbackLookupSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceEventRelayFeedbackLookupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"done": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},

		"q": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"records": {
			Type: schema.TypeList, //GoType: []*Record
			Elem: &schema.Resource{
				Schema: RecordSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			Computed:   true,
		},

		"total_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying EventRelayFeedbackLookup resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventRelayFeedbackLookupResourceData(d *schema.ResourceData, m *models.EventRelayFeedbackLookup, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("done", m.Done)
	d.Set("q", m.Q)
	d.Set("records", SetRecordSubResourceData(m.Records))
	d.Set("total_size", m.TotalSize)
}

// Iterate throught and update the EventRelayFeedbackLookup resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventRelayFeedbackLookupSubResourceData(m []*models.EventRelayFeedbackLookup) (d []*map[string]interface{}) {
	for _, eventRelayFeedbackLookup := range m {
		if eventRelayFeedbackLookup != nil {
			properties := make(map[string]interface{})
			properties["done"] = eventRelayFeedbackLookup.Done
			properties["q"] = eventRelayFeedbackLookup.Q
			properties["records"] = SetRecordSubResourceData(eventRelayFeedbackLookup.Records)
			properties["total_size"] = eventRelayFeedbackLookup.TotalSize
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EventRelayFeedbackLookup resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventRelayFeedbackLookupModel(d *schema.ResourceData) *models.EventRelayFeedbackLookup {
	done := d.Get("done").(bool)
	q := d.Get("q").(string)
	records := d.Get("records").([]*models.Record)
	totalSize := int32(d.Get("total_size").(int))

	return &models.EventRelayFeedbackLookup{
		Done:      done,
		Q:         &q,
		Records:   records,
		TotalSize: totalSize,
	}
}

// Function to perform the following actions:
func EventRelayFeedbackLookupModelFromMap(m map[string]interface{}) *models.EventRelayFeedbackLookup {
	done := m["done"].(bool)
	q := m["q"].(string)
	records := m["records"].([]*models.Record)
	totalSize := int32(m["total_size"].(int))

	return &models.EventRelayFeedbackLookup{
		Done:      done,
		Q:         &q,
		Records:   records,
		TotalSize: totalSize,
	}
}

// Retrieve property field names for updating the EventRelayFeedbackLookup resource
func GetEventRelayFeedbackLookupPropertyFields() (t []string) {
	return []string{
		"done",
		"q",
		"records",
		"total_size",
	}
}
