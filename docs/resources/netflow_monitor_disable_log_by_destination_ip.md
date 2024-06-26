---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_netflow_monitor_disable_log_by_destination_ip Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_netflow_monitor_disable_log_by_destination_ip: Configure a filter IP enrty
  PLACEHOLDER
---

# thunder_netflow_monitor_disable_log_by_destination_ip (Resource)

`thunder_netflow_monitor_disable_log_by_destination_ip`: Configure a filter IP enrty

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor_disable_log_by_destination_ip" "thunder_netflow_monitor_disable_log_by_destination_ip" {

  name      = "a11"
  ipv4_addr = "10.10.10.10/24"
  tcp_list {
    tcp_port_start = 35324
    tcp_port_end   = 35324
  }
  udp_list {
    udp_port_start = 32422
    udp_port_end   = 32422
  }
  icmp     = 1
  others   = 1
  user_tag = "11"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ipv4_addr` (String) Configure an IP subnet
- `name` (String) Name

### Optional

- `icmp` (Number) Disable logging for icmp traffic
- `others` (Number) Disable logging for other L4 protocols
- `tcp_list` (Block List) (see [below for nested schema](#nestedblock--tcp_list))
- `udp_list` (Block List) (see [below for nested schema](#nestedblock--udp_list))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--tcp_list"></a>
### Nested Schema for `tcp_list`

Optional:

- `tcp_port_end` (Number) Port Range End
- `tcp_port_start` (Number) Destination Port (Single Destination Port or Port Range Start)


<a id="nestedblock--udp_list"></a>
### Nested Schema for `udp_list`

Optional:

- `udp_port_end` (Number) Port Range End
- `udp_port_start` (Number) Destination Port (Single Destination Port or Port Range Start)


