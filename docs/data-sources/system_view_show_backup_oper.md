---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_view_show_backup_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_view_show_backup_oper: Operational Status for the object show-backup
  PLACEHOLDER
---

# thunder_system_view_show_backup_oper (Data Source)

`thunder_system_view_show_backup_oper`: Operational Status for the object show-backup

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_view_show_backup_oper" "thunder_system_view_show_backup_oper" {
}
output "get_system_view_show_backup_oper" {
  value = ["${data.thunder_system_view_show_backup_oper.thunder_system_view_show_backup_oper}"]
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

- `backup_show_1` (String)
- `backup_show_2` (String)
- `backup_show_3` (String)


