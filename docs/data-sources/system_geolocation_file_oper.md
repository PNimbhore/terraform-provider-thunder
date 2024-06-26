---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_geolocation_file_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_geolocation_file_oper: Operational Status for the object geolocation-file
  PLACEHOLDER
---

# thunder_system_geolocation_file_oper (Data Source)

`thunder_system_geolocation_file_oper`: Operational Status for the object geolocation-file

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_geolocation_file_oper" "thunder_system_geolocation_file_oper" {
}
output "get_system_geolocation_file_oper" {
  value = ["${data.thunder_system_geolocation_file_oper.thunder_system_geolocation_file_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `error_info` (Block List, Max: 1) (see [below for nested schema](#nestedblock--error_info))
- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--error_info"></a>
### Nested Schema for `error_info`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--error_info--oper))

<a id="nestedblock--error_info--oper"></a>
### Nested Schema for `error_info.oper`

Optional:

- `error_list` (Block List) (see [below for nested schema](#nestedblock--error_info--oper--error_list))
- `filename` (String)

<a id="nestedblock--error_info--oper--error_list"></a>
### Nested Schema for `error_info.oper.error_list`

Optional:

- `error` (String)
- `line` (Number)
- `offset` (Number)




<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `geofiles` (Block List) (see [below for nested schema](#nestedblock--oper--geofiles))

<a id="nestedblock--oper--geofiles"></a>
### Nested Schema for `oper.geofiles`

Optional:

- `comment` (Number)
- `error_warning` (Number)
- `filename` (String)
- `lines` (Number)
- `percentage_loaded` (Number)
- `success` (Number)
- `template` (String)
- `type` (String)


