---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_rrd_cgn_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_rrd_cgn_oper: Operational Status for the object cgn
  PLACEHOLDER
---

# thunder_rrd_cgn_oper (Data Source)

`thunder_rrd_cgn_oper`: Operational Status for the object cgn

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rrd_cgn_oper" "thunder_rrd_cgn_oper" {
}
output "get_rrd_cgn_oper" {
  value = ["${data.thunder_rrd_cgn_oper.thunder_rrd_cgn_oper}"]
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

- `cgn_data` (Block List) (see [below for nested schema](#nestedblock--oper--cgn_data))
- `end_time` (Number)
- `start_time` (Number)

<a id="nestedblock--oper--cgn_data"></a>
### Nested Schema for `oper.cgn_data`

Optional:

- `dslite_user_quota_create` (Number)
- `dslite_user_quota_delete` (Number)
- `lsn_user_quota_create` (Number)
- `lsn_user_quota_delete` (Number)
- `nat64_user_quota_create` (Number)
- `nat64_user_quota_delete` (Number)
- `time` (Number)


