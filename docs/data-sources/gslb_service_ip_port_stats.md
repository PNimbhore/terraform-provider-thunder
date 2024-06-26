---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_service_ip_port_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_service_ip_port_stats: Statistics for the object port
  PLACEHOLDER
---

# thunder_gslb_service_ip_port_stats (Data Source)

`thunder_gslb_service_ip_port_stats`: Statistics for the object port

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_service_ip_port_stats" "thunder_gslb_service_ip_port_stats" {

  node_name  = "vs2"
  port_num   = 30672
  port_proto = "tcp"
}
output "get_gslb_service_ip_port_stats" {
  value = ["${data.thunder_gslb_service_ip_port_stats.thunder_gslb_service_ip_port_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `node_name` (String) NodeName
- `port_num` (Number) Port Number
- `port_proto` (String) 'tcp': TCP Port; 'udp': UDP Port;

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `active` (Number) Active Servers
- `current` (Number) Current Connections


