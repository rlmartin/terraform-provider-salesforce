//// Code generated by go-swagger; DO NOT EDIT.

package resources

import (
	"context"
	"log"
	"vestahealthcare/client/platform_event_channel"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
PlatformEventChannel platform event channel API
*/

func PlatformEventChannel() *schema.Resource {
	return &schema.Resource{
		ReadContext: getPlatformEventChannel,
		Schema:      schemata.PlatformEventChannelSchema(),
	}
}

func DataResourcePlatformEventChannel() *schema.Resource {
	return &schema.Resource{
		Schema: schemata.DataSourcePlatformEventChannelSchema(),
	}
}

func getPlatformEventChannel(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := platform_event_channel.NewGetPlatformEventChannelParams()

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		params.ID = idVal.(string)
	} else {
		diags = append(diags, diag.Errorf("unexpected: Missing parameter - Id")...)
		diags = append(diags, diag.Errorf("ending operation")...)
		return diags
	}

	client := m.(*client.SalesforceRESTAPI)

	resp, err := client.PlatformEventChannel.GetPlatformEventChannel(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetPlatformEventChannelResourceData(d, respModel)

	return diags
}
