---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_tcp_syn_per_sec_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_tcp_syn_per_sec_oper: Operational Status for the object tcp-syn-per-sec
  PLACEHOLDER
---

# thunder_system_tcp_syn_per_sec_oper (Data Source)

`thunder_system_tcp_syn_per_sec_oper`: Operational Status for the object tcp-syn-per-sec

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_tcp_syn_per_sec_oper" "thunder_system_tcp_syn_per_sec_oper" {
}
output "get_system_tcp_syn_per_sec_oper" {
  value = ["${data.thunder_system_tcp_syn_per_sec_oper.thunder_system_tcp_syn_per_sec_oper}"]
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

- `tcp_syn_per_sec` (Number)


