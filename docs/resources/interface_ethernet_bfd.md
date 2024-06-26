---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_ethernet_bfd Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_interface_ethernet_bfd: Configure BFD (Bidirectional Forwarding Detection)
  PLACEHOLDER
---

# thunder_interface_ethernet_bfd (Resource)

`thunder_interface_ethernet_bfd`: Configure BFD (Bidirectional Forwarding Detection)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_bfd" "thunder_interface_ethernet_bfd" {
  ifnum = 1
  authentication {
    key_id   = 255
    method   = "md5"
    password = "a10net"
  }
  demand = 0
  echo   = 0
  interval_cfg {
    interval   = 208
    min_rx     = 280
    multiplier = 47
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ifnum` (String) Ifnum

### Optional

- `authentication` (Block List, Max: 1) (see [below for nested schema](#nestedblock--authentication))
- `demand` (Number) Demand mode
- `echo` (Number) Enable BFD Echo
- `interval_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--interval_cfg))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--authentication"></a>
### Nested Schema for `authentication`

Optional:

- `key_id` (Number) Key ID
- `method` (String) 'md5': Keyed MD5; 'meticulous-md5': Meticulous Keyed MD5; 'meticulous-sha1': Meticulous Keyed SHA1; 'sha1': Keyed SHA1; 'simple': Simple Password;
- `password` (String) Key String


<a id="nestedblock--interval_cfg"></a>
### Nested Schema for `interval_cfg`

Optional:

- `interval` (Number) Transmit interval between BFD packets (Milliseconds)
- `min_rx` (Number) Minimum receive interval capability (Milliseconds)
- `multiplier` (Number) Multiplier value used to compute holddown (value used to multiply the interval)


