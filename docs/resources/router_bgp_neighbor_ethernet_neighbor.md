---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_router_bgp_neighbor_ethernet_neighbor Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_router_bgp_neighbor_ethernet_neighbor: Specify an ethernet unnumbered neighbor
  PLACEHOLDER
---

# thunder_router_bgp_neighbor_ethernet_neighbor (Resource)

`thunder_router_bgp_neighbor_ethernet_neighbor`: Specify an ethernet unnumbered neighbor

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_neighbor_ethernet_neighbor" "thunder_router_bgp_neighbor_ethernet_neighbor" {

  as_number       = "300"
  ethernet        = 2
  unnumbered      = 1
  peer_group_name = "A10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `as_number` (String) AsNumber
- `ethernet` (Number) Ethernet interface number

### Optional

- `peer_group_name` (String)
- `unnumbered` (Number)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

