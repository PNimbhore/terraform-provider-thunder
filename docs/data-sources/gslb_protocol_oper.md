---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_protocol_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_protocol_oper: Operational Status for the object protocol
  PLACEHOLDER
---

# thunder_gslb_protocol_oper (Data Source)

`thunder_gslb_protocol_oper`: Operational Status for the object protocol

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_protocol_oper" "thunder_gslb_protocol_oper" {
}
output "get_gslb_protocol_oper" {
  value = ["${data.thunder_gslb_protocol_oper.thunder_gslb_protocol_oper}"]
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

- `session_list` (Block List) (see [below for nested schema](#nestedblock--oper--session_list))

<a id="nestedblock--oper--session_list"></a>
### Nested Schema for `oper.session_list`

Optional:

- `connection_failed` (Number)
- `connection_succeeded` (Number)
- `keepalive_packet_received` (Number)
- `keepalive_packet_sent` (Number)
- `message_header_error` (Number)
- `notify_packet_received` (Number)
- `notify_packet_sent` (Number)
- `open_packet_received` (Number)
- `open_packet_sent` (Number)
- `open_session_failed` (Number)
- `open_session_succeeded` (Number)
- `protocol_info` (String)
- `retry` (Number)
- `secure_config` (String)
- `secure_negotiation_fail` (Number)
- `secure_negotiation_success` (Number)
- `secure_state` (String)
- `session_id` (Number)
- `sessions_dropped` (Number)
- `ssl_handshake_fail` (Number)
- `ssl_handshake_success` (Number)
- `state` (String)
- `update_packet_received` (Number)
- `update_packet_sent` (Number)

