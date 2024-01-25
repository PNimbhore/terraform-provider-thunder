---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_stateful_firewall_global_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_stateful_firewall_global_stats: Statistics for the object global
  PLACEHOLDER
---

# thunder_cgnv6_stateful_firewall_global_stats (Data Source)

`thunder_cgnv6_stateful_firewall_global_stats`: Statistics for the object global

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_stateful_firewall_global_stats" "thunder_cgnv6_stateful_firewall_global_stats" {
}
output "get_cgnv6_stateful_firewall_global_stats" {
  value = ["${data.thunder_cgnv6_stateful_firewall_global_stats.thunder_cgnv6_stateful_firewall_global_stats}"]
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

- `eif_process` (Number) Endpnt-Independent Filter Matched
- `fullcone_creation_failure` (Number) Full-Cone Creation Failure
- `inbound_session_created` (Number) Inbound Session Created
- `inbound_session_freed` (Number) Inbound Session Freed
- `no_class_list_match` (Number) No Class-List Match Drop
- `no_fwd_route` (Number) No Forward Route
- `no_rev_route` (Number) No Reverse Route
- `one_arm_drop` (Number) One-Arm Drop
- `other_packet_process` (Number) Other Packet Process
- `other_session_created` (Number) Other Session Created
- `other_session_freed` (Number) Other Session Freed
- `outbound_session_created` (Number) Outbound Session Created
- `outbound_session_freed` (Number) Outbound Session Freed
- `packet_inbound_deny` (Number) Inbound Packet Denied
- `packet_process_failure` (Number) Packet Error Drop
- `packet_standby_drop` (Number) Standby Drop
- `session_creation_failure` (Number) Session Creation Failure
- `tcp_fullcone_created` (Number) TCP Full-cone Created
- `tcp_fullcone_freed` (Number) TCP Full-cone Freed
- `tcp_packet_process` (Number) TCP Packet Process
- `tcp_session_created` (Number) TCP Session Created
- `tcp_session_freed` (Number) TCP Session Freed
- `udp_fullcone_created` (Number) UDP Full-cone Created
- `udp_fullcone_freed` (Number) UDP Full-cone Freed
- `udp_packet_process` (Number) UDP Packet Process
- `udp_session_created` (Number) UDP Session Created
- `udp_session_freed` (Number) UDP Session Freed

