---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_zone_web_gui_protection_port_zone_service_other Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_zone_web_gui_protection_port_zone_service_other: DDOS Port & Protocol configuration
  PLACEHOLDER
---

# thunder_ddos_dst_zone_web_gui_protection_port_zone_service_other (Resource)

`thunder_ddos_dst_zone_web_gui_protection_port_zone_service_other`: DDOS Port & Protocol configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_web_gui_protection_port_zone_service_other" "thunder_ddos_dst_zone_web_gui_protection_port_zone_service_other" {
  zone_name  = "testZone"
  pbe        = "91"
  port_other = "other"
  protocol   = "tcp"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `port_other` (String) 'other': other;
- `protocol` (String) 'tcp': TCP Port; 'udp': UDP Port;
- `zone_name` (String) ZoneName

### Optional

- `pbe` (String) Peak Bandwidth Expected
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

