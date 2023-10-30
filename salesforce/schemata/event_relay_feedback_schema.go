package schemata

import (
	"vestahealthcare/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EventRelayFeedback resource defined in the Terraform configuration
func EventRelayFeedbackSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"event_relay_number": {
			Type:     schema.TypeString,
			Required: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"is_deleted": {
			Type:     schema.TypeBool,
			Required: true,
		},

		"remote_resource": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"status": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and EventRelayFeedbackSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceEventRelayFeedbackSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"event_relay_number": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"is_deleted": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},

		"remote_resource": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"status": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying EventRelayFeedback resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventRelayFeedbackResourceData(d *schema.ResourceData, m *models.EventRelayFeedback, isDataResource bool) {
	d.Set("event_relay_number", m.EventRelayNumber)
	if m.ID == "" && isDataResource {
		d.SetId("-")
	} else {
		d.SetId(m.ID)
	}
	d.Set("is_deleted", m.IsDeleted)
	d.Set("remote_resource", m.RemoteResource)
	d.Set("status", m.Status)
}

// Iterate throught and update the EventRelayFeedback resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventRelayFeedbackSubResourceData(m []*models.EventRelayFeedback) (d []*map[string]interface{}) {
	for _, eventRelayFeedback := range m {
		if eventRelayFeedback != nil {
			properties := make(map[string]interface{})
			properties["event_relay_number"] = eventRelayFeedback.EventRelayNumber
			properties["id"] = eventRelayFeedback.ID
			properties["is_deleted"] = eventRelayFeedback.IsDeleted
			properties["remote_resource"] = eventRelayFeedback.RemoteResource
			properties["status"] = eventRelayFeedback.Status
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EventRelayFeedback resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventRelayFeedbackModel(d *schema.ResourceData) *models.EventRelayFeedback {
	eventRelayNumber := d.Get("event_relay_number").(string)
	id := d.Get("id").(string)
	isDeleted := d.Get("is_deleted").(bool)
	remoteResource := d.Get("remote_resource").(string)
	status := d.Get("status").(string)

	return &models.EventRelayFeedback{
		EventRelayNumber: &eventRelayNumber,
		ID:               id,
		IsDeleted:        &isDeleted,
		RemoteResource:   remoteResource,
		Status:           &status,
	}
}

// Function to perform the following actions:
func EventRelayFeedbackModelFromMap(m map[string]interface{}) *models.EventRelayFeedback {
	eventRelayNumber := m["event_relay_number"].(string)
	id := m["id"].(string)
	isDeleted := m["is_deleted"].(bool)
	remoteResource := m["remote_resource"].(string)
	status := m["status"].(string)

	return &models.EventRelayFeedback{
		EventRelayNumber: &eventRelayNumber,
		ID:               id,
		IsDeleted:        &isDeleted,
		RemoteResource:   remoteResource,
		Status:           &status,
	}
}

func EventRelayFeedbackModelFromArrayOfMap(m []interface{}) []*models.EventRelayFeedback {
	mapped := make([]*models.EventRelayFeedback, len(m))
	for i, v := range m {
		mapped[i] = EventRelayFeedbackModelFromMap(v.(map[string]interface{}))
	}
	return mapped
}

// Retrieve property field names for updating the EventRelayFeedback resource
func GetEventRelayFeedbackPropertyFields() (t []string) {
	return []string{
		"event_relay_number",
		"id",
		"is_deleted",
		"remote_resource",
		"status",
	}
}
