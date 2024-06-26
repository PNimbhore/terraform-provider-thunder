---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_stateful_firewall_global Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_stateful_firewall_global: Stateful Firewall Configuration (default:disabled)
  PLACEHOLDER
---

# thunder_cgnv6_stateful_firewall_global (Resource)

`thunder_cgnv6_stateful_firewall_global`: Stateful Firewall Configuration (default:disabled)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_global" "thunder_cgnv6_stateful_firewall_global" {
  respond_to_user_mac = 0
  sampling_enable {
    counters1 = "all"
  }
  stateful_firewall_value = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `respond_to_user_mac` (Number) Use the user's source MAC for the next hop rather than the routing table (default: off)
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `stateful_firewall_value` (String) 'enable': Enable stateful firewall;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'tcp_packet_process': TCP Packet Process; 'udp_packet_process': UDP Packet Process; 'other_packet_process': Other Packet Process; 'packet_inbound_deny': Inbound Packet Denied; 'packet_process_failure': Packet Error Drop; 'outbound_session_created': Outbound Session Created; 'outbound_session_freed': Outbound Session Freed; 'inbound_session_created': Inbound Session Created; 'inbound_session_freed': Inbound Session Freed; 'tcp_session_created': TCP Session Created; 'tcp_session_freed': TCP Session Freed; 'udp_session_created': UDP Session Created; 'udp_session_freed': UDP Session Freed; 'other_session_created': Other Session Created; 'other_session_freed': Other Session Freed; 'session_creation_failure': Session Creation Failure; 'no_fwd_route': No Forward Route; 'no_rev_route': No Reverse Route; 'packet_standby_drop': Standby Drop; 'tcp_fullcone_created': TCP Full-cone Created; 'tcp_fullcone_freed': TCP Full-cone Freed; 'udp_fullcone_created': UDP Full-cone Created; 'udp_fullcone_freed': UDP Full-cone Freed; 'fullcone_creation_failure': Full-Cone Creation Failure; 'eif_process': Endpnt-Independent Filter Matched; 'one_arm_drop': One-Arm Drop; 'no_class_list_match': No Class-List Match Drop; 'outbound_session_created_shadow': Outbound Session Created Shadow; 'outbound_session_freed_shadow': Outbound Session Freed Shadow; 'inbound_session_created_shadow': Inbound Session Created Shadow; 'inbound_session_freed_shadow': Inbound Session Freed Shadow; 'tcp_session_created_shadow': TCP Session Created Shadow; 'tcp_session_freed_shadow': TCP Session Freed Shadow; 'udp_session_created_shadow': UDP Session Created Shadow; 'udp_session_freed_shadow': UDP Session Freed Shadow; 'other_session_created_shadow': Other Session Created Shadow; 'other_session_freed_shadow': Other Session Freed Shadow; 'session_creation_failure_shadow': Session Creation Failure Shadow; 'bad_session_freed': Bad Session Proto on Free; 'ctl_mem_alloc': Memory Alloc; 'ctl_mem_free': Memory Free; 'tcp_fullcone_created_shadow': TCP Full-cone Created Shadow; 'tcp_fullcone_freed_shadow': TCP Full-cone Freed Shadow; 'udp_fullcone_created_shadow': UDP Full-cone Created Shadow; 'udp_fullcone_freed_shadow': UDP Full-cone Freed Shadow; 'fullcone_in_del_q': Full-cone Found in Delete Queue; 'fullcone_overflow_eim': EIM Overflow; 'fullcone_overflow_eif': EIF Overflow; 'fullcone_free_found': Full-cone Free Found From Conn; 'fullcone_free_retry_lookup': Full-cone Retry Look-up; 'fullcone_free_not_found': Full-cone Free Not Found; 'eif_limit_exceeded': EIF Limit Exceeded; 'eif_disable_drop': EIF Disable Drop; 'eif_process_failure': EIF Process Failure; 'eif_filtered': EIF Filtered; 'ha_standby_session_created': HA Standby Session Created; 'ha_standby_session_eim': HA Standby Session EIM; 'ha_standby_session_eif': HA Standby Session EIF;


