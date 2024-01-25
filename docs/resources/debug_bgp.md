---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_debug_bgp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_debug_bgp: Debug Border Gateway Protocol (BGP)
  PLACEHOLDER
---

# thunder_debug_bgp (Resource)

`thunder_debug_bgp`: Debug Border Gateway Protocol (BGP)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_bgp" "thunder_debug_bgp" {
  all        = 0
  bfd        = 0
  dampening  = 0
  events     = 0
  filters    = 0
  fsm        = 0
  in         = 0
  keepalives = 0
  nht        = 0
  nsm        = 0
  out        = 0
  updates    = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `all` (Number) all debugging
- `bfd` (Number) Bidirectional Forwarding Detection (BFD)
- `dampening` (Number) BGP Dampening
- `events` (Number) BGP events
- `filters` (Number) BGP filters
- `fsm` (Number) BGP Finite State Machine
- `in` (Number) Inbound updates
- `keepalives` (Number) BGP keepalives
- `nht` (Number) NHT message
- `nsm` (Number) NSM message
- `out` (Number) Outbound updates
- `updates` (Number) BGP updates
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

