---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_tap Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_tap: DDoS TAP interface(s)
  PLACEHOLDER
---

# thunder_ddos_tap (Resource)

`thunder_ddos_tap`: DDoS TAP interface(s)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_tap" "thunder_ddos_tap" {
  ethernet_start_cfg {
    ethernet_start = 1
    ethernet_end   = 2
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ethernet_start_cfg` (Block List) (see [below for nested schema](#nestedblock--ethernet_start_cfg))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--ethernet_start_cfg"></a>
### Nested Schema for `ethernet_start_cfg`

Optional:

- `ethernet_end` (Number)
- `ethernet_start` (Number) Traffic receive from the ethernet port will be dropped

