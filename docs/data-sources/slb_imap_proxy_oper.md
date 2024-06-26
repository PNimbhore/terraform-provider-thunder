---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_imap_proxy_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_imap_proxy_oper: Operational Status for the object imap-proxy
  PLACEHOLDER
---

# thunder_slb_imap_proxy_oper (Data Source)

`thunder_slb_imap_proxy_oper`: Operational Status for the object imap-proxy

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_imap_proxy_oper" "thunder_slb_imap_proxy_oper" {
}
output "get_slb_imap_proxy_oper" {
  value = ["${data.thunder_slb_imap_proxy_oper.thunder_slb_imap_proxy_oper}"]
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
- `imap_proxy_cpu_list` (Block List) (see [below for nested schema](#nestedblock--oper--imap_proxy_cpu_list))

<a id="nestedblock--oper--imap_proxy_cpu_list"></a>
### Nested Schema for `oper.imap_proxy_cpu_list`

Optional:

- `alloc_error` (Number)
- `bad_seq` (Number)
- `boundary_error` (Number)
- `capability_packet` (Number)
- `client_est_state_err` (Number)
- `client_rq_state_err` (Number)
- `control_chn_ssl` (Number)
- `current_proxy_conns` (Number)
- `imap_line_too_long` (Number)
- `inv_start_line` (Number)
- `login_packet` (Number)
- `negative_error` (Number)
- `no_route_failure` (Number)
- `other_cmd` (Number)
- `realloc_error` (Number)
- `request_line_freed` (Number)
- `serv_ctng_state_err` (Number)
- `serv_resp_state_err` (Number)
- `serv_sel_ins_tpl_fail` (Number)
- `serv_sel_persist_fail` (Number)
- `serv_sel_smpv4_fail` (Number)
- `serv_sel_smpv6_fail` (Number)
- `server_selection_failure` (Number)
- `source_nat_failure` (Number)
- `start_tls_cmd` (Number)
- `total_imap_request` (Number)
- `total_proxy_conns` (Number)


