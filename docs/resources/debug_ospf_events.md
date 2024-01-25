---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_debug_ospf_events Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_debug_ospf_events: OSPFv3 event
  PLACEHOLDER
---

# thunder_debug_ospf_events (Resource)

`thunder_debug_ospf_events`: OSPFv3 event

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ospf_events" "thunder_debug_ospf_events" {
  abr    = 0
  asbr   = 0
  os     = 0
  router = 0
  vlink  = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `abr` (Number) OSPF ABR events
- `asbr` (Number) OSPF ASBR events
- `os` (Number) OS events
- `router` (Number) Other router events
- `uuid` (String) uuid of the object
- `vlink` (Number) Virtual-Link event

### Read-Only

- `id` (String) The ID of this resource.

