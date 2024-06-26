---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ipv6_nat_pool_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ipv6_nat_pool_group: IPv6 pool group name
  PLACEHOLDER
---

# thunder_ipv6_nat_pool_group (Resource)

`thunder_ipv6_nat_pool_group`: IPv6 pool group name

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_nat_pool_group" "thunder_ipv6_nat_pool_group" {

  member_list {
    pool_name = "11"
  }
  pool_group_name = "35"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "8"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `pool_group_name` (String) Specify pool group name

### Optional

- `member_list` (Block List) (see [below for nested schema](#nestedblock--member_list))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `vrid` (Number) Specify VRRP-A vrid (Specify ha VRRP-A vrid)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--member_list"></a>
### Nested Schema for `member_list`

Required:

- `pool_name` (String) Specify NAT pool name

Optional:

- `uuid` (String) uuid of the object


<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'Failed': Failed;


