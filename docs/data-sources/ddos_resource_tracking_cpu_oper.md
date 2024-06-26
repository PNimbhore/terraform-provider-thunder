---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_resource_tracking_cpu_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_resource_tracking_cpu_oper: Operational Status for the object cpu
  PLACEHOLDER
---

# thunder_ddos_resource_tracking_cpu_oper (Data Source)

`thunder_ddos_resource_tracking_cpu_oper`: Operational Status for the object cpu

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_resource_tracking_cpu_oper" "thunder_ddos_resource_tracking_cpu_oper" {
}
output "get_ddos_resource_tracking_cpu_oper" {
  value = ["${data.thunder_ddos_resource_tracking_cpu_oper.thunder_ddos_resource_tracking_cpu_oper}"]
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

- `error_str` (String)
- `if_show_error_str` (Number)
- `max_count` (Number)
- `timestamps` (Block List) (see [below for nested schema](#nestedblock--oper--timestamps))

<a id="nestedblock--oper--timestamps"></a>
### Nested Schema for `oper.timestamps`

Optional:

- `average_cpu_percent` (String)
- `entries` (Block List) (see [below for nested schema](#nestedblock--oper--timestamps--entries))
- `timestamp` (String)

<a id="nestedblock--oper--timestamps--entries"></a>
### Nested Schema for `oper.timestamps.entries`

Optional:

- `absolute_cpu_percent` (String)
- `address` (String)
- `entry` (String)
- `relative_cpu_percent` (String)


