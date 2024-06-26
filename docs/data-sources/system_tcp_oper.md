---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_tcp_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_tcp_oper: Operational Status for the object tcp
  PLACEHOLDER
---

# thunder_system_tcp_oper (Data Source)

`thunder_system_tcp_oper`: Operational Status for the object tcp

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_tcp_oper" "thunder_system_tcp_oper" {
}
output "get_system_tcp_oper" {
  value = ["${data.thunder_system_tcp_oper.thunder_system_tcp_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))
- `rate_limit_reset_unknown_conn` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rate_limit_reset_unknown_conn))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `cpu_count` (Number)
- `tcp_cpu_list` (Block List) (see [below for nested schema](#nestedblock--oper--tcp_cpu_list))

<a id="nestedblock--oper--tcp_cpu_list"></a>
### Nested Schema for `oper.tcp_cpu_list`

Optional:

- `activeopens` (Number)
- `attemptfails` (Number)
- `ax_rexmit_syn` (Number)
- `currclose` (Number)
- `currclsg` (Number)
- `currclsw` (Number)
- `currestab` (Number)
- `currfinw1` (Number)
- `currfinw2` (Number)
- `currlack` (Number)
- `currlstn` (Number)
- `currsynrcv` (Number)
- `currsyssnt` (Number)
- `currtimew` (Number)
- `estabresets` (Number)
- `exceedmss` (Number)
- `inerrs` (Number)
- `insegs` (Number)
- `mem_alloc` (Number)
- `noroute` (Number)
- `orphan_count` (Number)
- `outrsts` (Number)
- `outsegs` (Number)
- `passiveopens` (Number)
- `pawsactiverejected` (Number)
- `recv_mem` (Number)
- `retranssegs` (Number)
- `send_mem` (Number)
- `sock_alloc` (Number)
- `syn_rcv_ack` (Number)
- `syn_rcv_rst` (Number)
- `syn_rcv_rstack` (Number)
- `tcpabortontimeout` (Number)
- `tfo_actives` (Number)
- `tfo_conns` (Number)
- `tfo_denied` (Number)



<a id="nestedblock--rate_limit_reset_unknown_conn"></a>
### Nested Schema for `rate_limit_reset_unknown_conn`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rate_limit_reset_unknown_conn--oper))

<a id="nestedblock--rate_limit_reset_unknown_conn--oper"></a>
### Nested Schema for `rate_limit_reset_unknown_conn.oper`

Optional:

- `unknown_conn_current_rate` (Number)
- `unknown_conn_rate_limit` (Number)
- `unknown_conn_rate_limit_drop` (Number)


