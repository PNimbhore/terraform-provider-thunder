---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_resource_usage Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_resource_usage: Configure CGNV6 Resource Usage
  PLACEHOLDER
---

# thunder_cgnv6_resource_usage (Resource)

`thunder_cgnv6_resource_usage`: Configure CGNV6 Resource Usage

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_resource_usage" "thunder_cgnv6_resource_usage" {

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `fixed_nat_inside_user_count` (Number) Total configurable CGNV6 Fixed NAT inside users
- `fixed_nat_ip_addr_count` (Number) Total configurable CGNV6 Fixed NAT addresses
- `lsn_nat_addr_count` (Number) Total configurable CGNV6 NAT Pool addresses
- `radius_table_size` (Number) Total configurable CGNV6 RADIUS Table entries
- `stateless_entries` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stateless_entries))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stateless_entries"></a>
### Nested Schema for `stateless_entries`

Optional:

- `l4_session_count` (Number) Helper size for CGN Stateless Technologies
- `uuid` (String) uuid of the object


