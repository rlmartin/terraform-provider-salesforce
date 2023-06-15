terraform {
	required_version = ">= 0.12.00"
	required_providers {
		salesforce = {
			source  = "salesforce.com/com/salesforce"
			version = "0.2"
		}
	}
}

provider "salesforce" {
	api_id = <ADD YOUR LM API_ID HERE>
	api_key = <ADD YOUR LM API_KEY HERE>
	company = <ADD YOUR LM PORTAL NAME HERE>
}

resource "salesforce_device" "my_device"{
	preferred_collector_id = 1
	custom_properties = [
		{
			name = "addr"
      		value = "127.0.0.1"
		},
		{
			name = "host"
      		value = "localhost"
		}
	]
	description = "This is a Cisco Router"
	device_type  = 0
	disable_alerting = true
	display_name = "Cisco Router"
	enable_netflow = false
	link = "www.ciscorouter.com"
	name = "collector.host"
}

data "salesforce_device" "my_devices" {
    filter = "displayName~\"Cisco Router\""
	// make sure that the device is created before we try to query for it
	depends_on = [
		salesforce_device.my_device
	]
}

output "devices" {
  description = "devices list"
  value       = data.salesforce_device.my_devices
}