---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_overlay_tunnel_vtep_local_ip_address Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_overlay_tunnel_vtep_local_ip_address: IP Address of the local tunnel end point
  PLACEHOLDER
---

# thunder_overlay_tunnel_vtep_local_ip_address (Resource)

`thunder_overlay_tunnel_vtep_local_ip_address`: IP Address of the local tunnel end point

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_vtep_local_ip_address" "thunder_overlay_tunnel_vtep_local_ip_address" {

  id1        = 56
  ip_address = "10.10.101.10"
  vni_list {
    segment = 3
    lif     = "test"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id1` (String) Id1
- `ip_address` (String) Source Tunnel End Point IPv4 address

### Optional

- `uuid` (String) uuid of the object
- `vni_list` (Block List) (see [below for nested schema](#nestedblock--vni_list))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--vni_list"></a>
### Nested Schema for `vni_list`

Required:

- `segment` (Number) Id of the segment that is being extended

Optional:

- `gateway` (Number) This is a Gateway segment id
- `lif` (String) Logical interface (logical interface name)
- `partition` (String) Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)
- `uuid` (String) uuid of the object

