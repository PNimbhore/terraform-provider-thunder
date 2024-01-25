---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_lsn_rule_list_domain_list_name_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_lsn_rule_list_domain_list_name_stats: Statistics for the object domain-list-name
  PLACEHOLDER
---

# thunder_cgnv6_lsn_rule_list_domain_list_name_stats (Data Source)

`thunder_cgnv6_lsn_rule_list_domain_list_name_stats`: Statistics for the object domain-list-name

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_rule_list_domain_list_name_stats" "thunder_cgnv6_lsn_rule_list_domain_list_name_stats" {

  name             = "test"
  name_domain_list = "testing"

}
output "get_cgnv6_lsn_rule_list_domain_list_name_stats" {
  value = ["${data.thunder_cgnv6_lsn_rule_list_domain_list_name_stats.thunder_cgnv6_lsn_rule_list_domain_list_name_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name
- `name_domain_list` (String) Configure a Specific Rule-Set (Domain List Name)

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

