---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_geo_location Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_geo_location: Configure a GSLB global geo-location object
  PLACEHOLDER
---

# thunder_gslb_geo_location (Resource)

`thunder_gslb_geo_location`: Configure a GSLB global geo-location object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_geo_location" "thunder_gslb_geo_location" {
  geo_locn_multiple_addresses {
    first_ipv6_address = "2001:db8:3333:4444:5555:6666:7777:8888"
    geol_ipv6_mask     = 122
  }
  geo_locn_obj_name = "test"
  user_tag          = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `geo_locn_obj_name` (String) Specify geo-location name, section range is (1-15)

### Optional

- `geo_locn_multiple_addresses` (Block List) (see [below for nested schema](#nestedblock--geo_locn_multiple_addresses))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--geo_locn_multiple_addresses"></a>
### Nested Schema for `geo_locn_multiple_addresses`

Optional:

- `first_ip_address` (String) Specify IP information (Specify IP address)
- `first_ipv6_address` (String) Specify IPv6 address
- `geol_ipv4_mask` (String) Specify IPv4 mask
- `geol_ipv6_mask` (Number) Specify IPv6 mask
- `ip_addr2` (String) Specify IP address range
- `ipv6_addr2` (String) Specify IPv6 address range

