---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_app_protocol_port_tcp_passthrough Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_app_protocol_port_tcp_passthrough: Controls TCP ports filtering on all interfaces, Default mode is enabled
  PLACEHOLDER
---

# thunder_ip_app_protocol_port_tcp_passthrough (Resource)

`thunder_ip_app_protocol_port_tcp_passthrough`: Controls TCP ports filtering on all interfaces, Default mode is enabled

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_app_protocol_port_tcp_passthrough" "thunder_ip_app_protocol_port_tcp_passthrough" {
  disable = 0
  enable  = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `disable` (Number) Disable passthrough mode
- `enable` (Number) Enables passthrough mode
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


