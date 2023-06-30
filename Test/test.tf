terraform {
	required_version = ">= 0.12.0"
	required_providers {
		salesforce = {
			source  = "salesforce.com/com/salesforce"
			version = "0.2"
		}
	}
}

provider "salesforce" {
	client_id     = <ADD YOUR SALESFORCE CLIENT_ID HERE>
	client_secret = <ADD YOUR SALESFORCE CLIENT_SECRET HERE>
	subdomain     = <ADD YOUR SALESFORCE SUBDOMAIN HERE>
	token_url     = <ADD YOUR SALESFORCE TOKEN_URL HERE>
	username      = <ADD YOUR SALESFORCE USERNAME HERE>
	password      = <ADD YOUR PASSWORD HERE>
}

resource "salesforce_device" "my_channel" {
  foo_channel = {
    full_name = "my_channel"
    metadata  = {
      label        = "my_channel"
      channel_type = "my_channel_type"
		}
	}
}

data "salesforce_device" "my_devices" {
    filter = "displayName~\"Cisco Router\""
	// make sure that the device is created before we try to query for it
	depends_on = [
		salesforce_device.my_device
	]
}

output "my_channel_id" {
  description = "unique identifier for my_channel"
  value       = salesforce_device.my_channel.id
}
