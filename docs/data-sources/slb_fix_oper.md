---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_fix_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_fix_oper: Operational Status for the object fix
  PLACEHOLDER
---

# thunder_slb_fix_oper (Data Source)

`thunder_slb_fix_oper`: Operational Status for the object fix

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_fix_oper" "thunder_slb_fix_oper" {
}
output "get_slb_fix_oper" {
  value = ["${data.thunder_slb_fix_oper.thunder_slb_fix_oper}"]
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

- `cpu_count` (Number)
- `fix_cpu_list` (Block List) (see [below for nested schema](#nestedblock--oper--fix_cpu_list))

<a id="nestedblock--oper--fix_cpu_list"></a>
### Nested Schema for `oper.fix_cpu_list`

Optional:

- `client_err` (Number)
- `client_tls_conn` (Number)
- `curr_proxy` (Number)
- `default_switching` (Number)
- `insert_clientip` (Number)
- `noroute` (Number)
- `sender_switching` (Number)
- `server_err` (Number)
- `server_tls_conn` (Number)
- `snat_fail` (Number)
- `svrsel_fail` (Number)
- `target_switching` (Number)
- `total_proxy` (Number)


