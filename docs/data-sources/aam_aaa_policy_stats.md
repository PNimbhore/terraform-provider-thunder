---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_aaa_policy_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_aaa_policy_stats: Statistics for the object aaa-policy
  PLACEHOLDER
---

# thunder_aam_aaa_policy_stats (Data Source)

`thunder_aam_aaa_policy_stats`: Statistics for the object aaa-policy

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_aaa_policy_stats" "thunder_aam_aaa_policy_stats" {

  name = "test"
}
output "get_aam_aaa_policy_stats" {
  value = ["${data.thunder_aam_aaa_policy_stats.thunder_aam_aaa_policy_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify AAA policy name

### Optional

- `aaa_rule_list` (Block List) (see [below for nested schema](#nestedblock--aaa_rule_list))
- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--aaa_rule_list"></a>
### Nested Schema for `aaa_rule_list`

Required:

- `index` (Number) Specify AAA rule index

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--aaa_rule_list--stats))

<a id="nestedblock--aaa_rule_list--stats"></a>
### Nested Schema for `aaa_rule_list.stats`

Optional:

- `failure_bypass` (Number)
- `hit_auth` (Number)
- `hit_bypass` (Number)
- `hit_deny` (Number)
- `total_count` (Number)



<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `error` (Number) Error
- `failure_bypass` (Number) Auth Failure Bypass
- `req` (Number) Request
- `req_auth` (Number) Request Matching Authentication Template
- `req_bypass` (Number) Request Bypassed
- `req_reject` (Number) Request Rejected
- `req_skip` (Number) Request Skipped


