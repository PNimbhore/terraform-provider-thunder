---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_reroute_suppress_protocols Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_reroute_suppress_protocols: suppress protocols
  PLACEHOLDER
---

# thunder_ip_reroute_suppress_protocols (Resource)

`thunder_ip_reroute_suppress_protocols`: suppress protocols

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_reroute_suppress_protocols" "thunder_ip_reroute_suppress_protocols" {
  connected = 0
  ebgp      = 0
  ibgp      = 0
  isis      = 0
  ospf      = 0
  rip       = 0
  static    = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `connected` (Number)
- `ebgp` (Number) EBGP
- `ibgp` (Number) IBGP
- `isis` (Number) ISIS
- `ospf` (Number) OSPF
- `rip` (Number) RIP
- `static` (Number)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

