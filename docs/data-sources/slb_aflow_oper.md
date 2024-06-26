---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_aflow_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_aflow_oper: Operational Status for the object aflow
  PLACEHOLDER
---

# thunder_slb_aflow_oper (Data Source)

`thunder_slb_aflow_oper`: Operational Status for the object aflow

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_aflow_oper" "thunder_slb_aflow_oper" {
}
output "get_slb_aflow_oper" {
  value = ["${data.thunder_slb_aflow_oper.thunder_slb_aflow_oper}"]
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

- `aflow_cpu_list` (Block List) (see [below for nested schema](#nestedblock--oper--aflow_cpu_list))
- `cpu_count` (Number)

<a id="nestedblock--oper--aflow_cpu_list"></a>
### Nested Schema for `oper.aflow_cpu_list`

Optional:

- `error_resume_conn` (Number)
- `event_resume_conn` (Number)
- `inc_aflow_limit` (Number)
- `open_new_server_conn` (Number)
- `pause_conn` (Number)
- `pause_conn_fail` (Number)
- `resume_conn` (Number)
- `retry_resume_conn` (Number)
- `reuse_server_idle_conn` (Number)
- `timer_resume_conn` (Number)
- `try_to_resume_conn` (Number)


