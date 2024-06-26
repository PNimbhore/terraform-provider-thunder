---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_traffic_control_rule_set_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_traffic_control_rule_set_oper: Operational Status for the object rule-set
  PLACEHOLDER
---

# thunder_traffic_control_rule_set_oper (Data Source)

`thunder_traffic_control_rule_set_oper`: Operational Status for the object rule-set

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_traffic_control_rule_set_oper" "thunder_traffic_control_rule_set_oper" {

  name = "testing_rule"
}
output "get_traffic_control_rule_set_oper" {
  value = ["${data.thunder_traffic_control_rule_set_oper.thunder_traffic_control_rule_set_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Rule set name

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))
- `rule_list` (Block List) (see [below for nested schema](#nestedblock--rule_list))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `policy_rule_count` (Number)
- `policy_status` (String)
- `rule_stats` (Block List) (see [below for nested schema](#nestedblock--oper--rule_stats))

<a id="nestedblock--oper--rule_stats"></a>
### Nested Schema for `oper.rule_stats`

Optional:

- `rule_hitcount` (Number)
- `rule_name` (String)
- `rule_status` (String)



<a id="nestedblock--rule_list"></a>
### Nested Schema for `rule_list`

Required:

- `name` (String) Rule name

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rule_list--oper))

<a id="nestedblock--rule_list--oper"></a>
### Nested Schema for `rule_list.oper`

Optional:

- `hitcount` (Number)
- `status` (String)


