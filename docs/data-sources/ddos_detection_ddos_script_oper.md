---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_detection_ddos_script_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_detection_ddos_script_oper: Operational Status for the object ddos-script
  PLACEHOLDER
---

# thunder_ddos_detection_ddos_script_oper (Data Source)

`thunder_ddos_detection_ddos_script_oper`: Operational Status for the object ddos-script

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_detection_ddos_script_oper" "thunder_ddos_detection_ddos_script_oper" {
}
output "get_ddos_detection_ddos_script_oper" {
  value = ["${data.thunder_ddos_detection_ddos_script_oper.thunder_ddos_detection_ddos_script_oper}"]
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

- `file_list` (Block List) (see [below for nested schema](#nestedblock--oper--file_list))
- `total_records` (Number)

<a id="nestedblock--oper--file_list"></a>
### Nested Schema for `oper.file_list`

Optional:

- `file` (String)
- `file_size` (Number)
- `reference_count` (Number)


