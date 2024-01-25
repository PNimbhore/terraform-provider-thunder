---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_rule_set_track_app_rule_list_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_rule_set_track_app_rule_list_stats: Statistics for the object track-app-rule-list
  PLACEHOLDER
---

# thunder_rule_set_track_app_rule_list_stats (Data Source)

`thunder_rule_set_track_app_rule_list_stats`: Statistics for the object track-app-rule-list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_track_app_rule_list_stats" "thunder_rule_set_track_app_rule_list_stats" {

  name = "test"
}
output "get_rule_set_track_app_rule_list_stats" {
  value = ["${data.thunder_rule_set_track_app_rule_list_stats.thunder_rule_set_track_app_rule_list_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `dummy` (Number) Entry for a10countergen

