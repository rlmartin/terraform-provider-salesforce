package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the QueryResult resource defined in the Terraform configuration
func QueryResultSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"done": {
			Type:     schema.TypeBool,
			Required: true,
		},

		"q": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"records": {
			Type: schema.TypeList, //GoType: []map[string]string
			Elem: &schema.Schema{
				Type:        schema.TypeMap,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			Required: true,
		},

		"total_size": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and QueryResultSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceQueryResultSchema() map[string]*schema.Schema {
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
			Type: schema.TypeList, //GoType: []map[string]string
			Elem: &schema.Schema{
				Type:        schema.TypeMap,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			Optional: true,
			Computed: true,
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

// Update the underlying QueryResult resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetQueryResultResourceData(d *schema.ResourceData, m *models.QueryResult, isDataResource bool) {
	if isDataResource {
		d.SetId("-")
	}
	d.Set("done", m.Done)
	d.Set("q", m.Q)
	d.Set("records", m.Records)
	d.Set("total_size", m.TotalSize)
}

// Iterate throught and update the QueryResult resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetQueryResultSubResourceData(m []*models.QueryResult) (d []*map[string]interface{}) {
	for _, queryResult := range m {
		if queryResult != nil {
			properties := make(map[string]interface{})
			properties["done"] = queryResult.Done
			properties["q"] = queryResult.Q
			properties["records"] = queryResult.Records
			properties["total_size"] = queryResult.TotalSize
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate QueryResult resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func QueryResultModel(d *schema.ResourceData) *models.QueryResult {
	done := d.Get("done").(bool)
	q := d.Get("q").(string)
	records := d.Get("records").([]map[string]string)
	totalSize := int32(d.Get("total_size").(int))

	return &models.QueryResult{
		Done:      &done,
		Q:         q,
		Records:   records,
		TotalSize: &totalSize,
	}
}

// Function to perform the following actions:
func QueryResultModelFromMap(m map[string]interface{}) *models.QueryResult {
	done := m["done"].(bool)
	q := m["q"].(string)
	records := m["records"].([]map[string]string)
	totalSize := int32(m["total_size"].(int))

	return &models.QueryResult{
		Done:      &done,
		Q:         q,
		Records:   records,
		TotalSize: &totalSize,
	}
}

// Retrieve property field names for updating the QueryResult resource
func GetQueryResultPropertyFields() (t []string) {
	return []string{
		"done",
		"q",
		"records",
		"total_size",
	}
}
