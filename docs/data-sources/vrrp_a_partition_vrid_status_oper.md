---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vrrp_a_partition_vrid_status_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vrrp_a_partition_vrid_status_oper: Operational Status for the object partition-vrid-status
  PLACEHOLDER
---

# thunder_vrrp_a_partition_vrid_status_oper (Data Source)

`thunder_vrrp_a_partition_vrid_status_oper`: Operational Status for the object partition-vrid-status

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vrrp_a_partition_vrid_status_oper" "thunder_vrrp_a_partition_vrid_status_oper" {

}
output "get_vrrp_a_partition_vrid_status_oper" {
  value = ["${data.thunder_vrrp_a_partition_vrid_status_oper.thunder_vrrp_a_partition_vrid_status_oper}"]
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

- `all_partition_list` (Block List) (see [below for nested schema](#nestedblock--oper--all_partition_list))

<a id="nestedblock--oper--all_partition_list"></a>
### Nested Schema for `oper.all_partition_list`

Optional:

- `active_device_id` (Number)
- `active_priority` (Number)
- `active_weight` (Number)
- `local_device_id` (Number)
- `partition_name` (String)
- `vrid` (Number)


