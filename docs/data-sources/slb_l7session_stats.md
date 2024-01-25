---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_l7session_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_l7session_stats: Statistics for the object l7session
  PLACEHOLDER
---

# thunder_slb_l7session_stats (Data Source)

`thunder_slb_l7session_stats`: Statistics for the object l7session

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_l7session_stats" "thunder_slb_l7session_stats" {
}
output "get_slb_l7session_stats" {
  value = ["${data.thunder_slb_l7session_stats.thunder_slb_l7session_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `client_fin` (Number) FIN from client
- `client_rst` (Number) RST from client
- `conn_not_exist` (Number) Conn does not exist
- `curr_proxy` (Number) Curr proxy conn
- `data_cb_failed` (Number) Data event callback fail
- `data_event` (Number) Data event from TCP
- `err_cb_failed` (Number) Err event callback failed
- `err_event` (Number) Err event from TCP
- `hps_fwdreq_fail` (Number) Fwd req fail
- `server_conn_failed` (Number) Server connection failed
- `server_fin` (Number) FIN from server
- `server_rst` (Number) RST from server
- `server_select_fail` (Number) Server selection fail
- `start_server_conn_succ` (Number) Start Server Conn Success
- `total_proxy` (Number) Total proxy conn
- `udp_data_event` (Number) Data event from UDP
- `wbuf_cb_failed` (Number) Wbuf event callback failed
- `wbuf_event` (Number) Wbuf event from TCP

