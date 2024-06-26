---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_lsn_endpoint_independent_filtering_tcp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_lsn_endpoint_independent_filtering_tcp: Set behavior for TCP
  PLACEHOLDER
---

# thunder_cgnv6_lsn_endpoint_independent_filtering_tcp (Resource)

`thunder_cgnv6_lsn_endpoint_independent_filtering_tcp`: Set behavior for TCP

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_endpoint_independent_filtering_tcp" "thunder_cgnv6_lsn_endpoint_independent_filtering_tcp" {

  port_list {
    port     = 100
    port_end = 1000
  }
  session_limit = 65535
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `port_list` (Block List) (see [below for nested schema](#nestedblock--port_list))
- `session_limit` (Number) Limit number of EIF sessions that can be created per port
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--port_list"></a>
### Nested Schema for `port_list`

Optional:

- `port` (Number) Single Destination Port or Port Range Start
- `port_end` (Number) Port Range End


