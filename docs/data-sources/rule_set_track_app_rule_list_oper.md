---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_rule_set_track_app_rule_list_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_rule_set_track_app_rule_list_oper: Operational Status for the object track-app-rule-list
  PLACEHOLDER
---

# thunder_rule_set_track_app_rule_list_oper (Data Source)

`thunder_rule_set_track_app_rule_list_oper`: Operational Status for the object track-app-rule-list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_track_app_rule_list_oper" "thunder_rule_set_track_app_rule_list_oper" {
  oper {
    rule_list {
    }
  }

}
output "get_rule_set_track_app_rule_list_oper" {
  value = ["${data.thunder_rule_set_track_app_rule_list_oper.thunder_rule_set_track_app_rule_list_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `rule_list` (Block List) (see [below for nested schema](#nestedblock--oper--rule_list))

<a id="nestedblock--oper--rule_list"></a>
### Nested Schema for `oper.rule_list`

Optional:

- `name` (String)


