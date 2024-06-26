---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_tuple_filter_filter_rule Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_tuple_filter_filter_rule: Specify filter matching rule
  PLACEHOLDER
---

# thunder_tuple_filter_filter_rule (Resource)

`thunder_tuple_filter_filter_rule`: Specify filter matching rule

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tuple_filter_filter_rule" "thunder_tuple_filter_filter_rule" {

  name     = "test"
  id1      = 2
  src_addr = "10.10.10.10/32"
  dst_addr = "10.10.10.12/32"
  src_port = 20
  dst_port = 22
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id1` (Number) filter rule id
- `name` (String) Name

### Optional

- `dst_addr` (String) Destination IPv4 address with prefix
- `dst_port` (Number) Destination port
- `dst_v6_addr` (String) Destination IPv6 address with prefix
- `dst_v6_port` (Number) Destination port
- `src_addr` (String) Source IPv4 address with prefix
- `src_port` (Number) Source port
- `src_v6_addr` (String) Source IPv6 address with prefix
- `src_v6_port` (Number) Source port
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


