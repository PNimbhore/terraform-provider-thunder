---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_player_id_global_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_player_id_global_oper: Operational Status for the object player-id-global
  PLACEHOLDER
---

# thunder_slb_player_id_global_oper (Data Source)

`thunder_slb_player_id_global_oper`: Operational Status for the object player-id-global

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_player_id_global_oper" "thunder_slb_player_id_global_oper" {
}
output "get_slb_player_id_global_oper" {
  value = ["${data.thunder_slb_player_id_global_oper.thunder_slb_player_id_global_oper}"]
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

- `state` (String)
- `table_count` (Number)
- `time_to_active` (Number)


