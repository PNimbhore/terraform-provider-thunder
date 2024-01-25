---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_scaleout_debug_traffic_map_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_scaleout_debug_traffic_map_oper: Operational Status for the object traffic-map
  PLACEHOLDER
---

# thunder_scaleout_debug_traffic_map_oper (Data Source)

`thunder_scaleout_debug_traffic_map_oper`: Operational Status for the object traffic-map

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_debug_traffic_map_oper" "thunder_scaleout_debug_traffic_map_oper" {
}
output "get_scaleout_debug_traffic_map_oper" {
  value = ["${data.thunder_scaleout_debug_traffic_map_oper.thunder_scaleout_debug_traffic_map_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `device_group_list` (Block List) (see [below for nested schema](#nestedblock--oper--device_group_list))

<a id="nestedblock--oper--device_group_list"></a>
### Nested Schema for `oper.device_group_list`

Optional:

- `buckets_list` (Block List) (see [below for nested schema](#nestedblock--oper--device_group_list--buckets_list))
- `buffer_len` (Number)
- `cmd` (String)
- `rc` (Number)

<a id="nestedblock--oper--device_group_list--buckets_list"></a>
### Nested Schema for `oper.device_group_list.buckets_list`

Optional:

- `active_device` (Number)
- `standby_device` (Number)
- `user_group` (Number)

