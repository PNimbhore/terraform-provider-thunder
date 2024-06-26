---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_policy_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_template_policy_stats: Statistics for the object policy
  PLACEHOLDER
---

# thunder_slb_template_policy_stats (Data Source)

`thunder_slb_template_policy_stats`: Statistics for the object policy

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_template_policy_stats" "thunder_slb_template_policy_stats" {

  name = "test"

}
output "get_slb_template_policy_stats" {
  value = ["${data.thunder_slb_template_policy_stats.thunder_slb_template_policy_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Policy template name

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `exp_client_hello_not_found` (Number) Expected Client HELLO requests not found
- `fwd_policy_dns_outstanding` (Number) Forward-policy current DNS outstanding requests
- `fwd_policy_dns_unresolved` (Number) Forward-policy unresolved DNS queries
- `fwd_policy_forward_to_internet` (Number) Number of forward-policy requests forwarded to internet
- `fwd_policy_forward_to_proxy` (Number) Number of forward-policy requests forwarded to proxy
- `fwd_policy_forward_to_service_group` (Number) Number of forward-policy requests forwarded to service group
- `fwd_policy_hits` (Number) Number of forward-policy requests for this policy template
- `fwd_policy_policy_drop` (Number) Number of forward-policy requests dropped
- `fwd_policy_snat_fail` (Number) Forward-policy source-nat translation failure
- `fwd_policy_source_match_not_found` (Number) Forward-policy requests without matching source rule


