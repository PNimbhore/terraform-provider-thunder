---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_nat_range_list Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_nat_range_list: IP Source NAT Static range list
  PLACEHOLDER
---

# thunder_ip_nat_range_list (Resource)

`thunder_ip_nat_range_list`: IP Source NAT Static range list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_range_list" "thunder_ip_nat_range_list" {
  name                   = 11
  global_start_ipv4_addr = "10.0.2.6"
  global_netmaskv4       = "255.255.255.255"
  local_start_ipv4_addr  = "10.0.2.7"
  local_netmaskv4        = "255.255.255.255"
  v4_count               = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name for this Static List

### Optional

- `global_netmaskv4` (String) Mask for this Address range
- `global_start_ipv4_addr` (String) Global Start IPv4 Address of this list
- `global_start_ipv6_addr` (String) Global Start IPv6 Address of this list
- `local_netmaskv4` (String) Mask for this Address range
- `local_start_ipv4_addr` (String) Local Start IPv4 Address of this list
- `local_start_ipv6_addr` (String) Local Start IPv6 Address of this list
- `uuid` (String) uuid of the object
- `v4_acl_id` (Number) Access list ID
- `v4_acl_name` (String) Access list name
- `v4_count` (Number) Number of addresses to be translated in this range
- `v4_vrid` (Number) VRRP-A vrid (Specify ha VRRP-A vrid)
- `v6_acl_name` (String) Access list name
- `v6_count` (Number) Number of addresses to be translated in this range
- `v6_vrid` (Number) VRRP-A vrid (Specify ha VRRP-A vrid)

### Read-Only

- `id` (String) The ID of this resource.


