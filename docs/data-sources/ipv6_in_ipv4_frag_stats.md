---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ipv6_in_ipv4_frag_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ipv6_in_ipv4_frag_stats: Statistics for the object frag
  PLACEHOLDER
---

# thunder_ipv6_in_ipv4_frag_stats (Data Source)

`thunder_ipv6_in_ipv4_frag_stats`: Statistics for the object frag

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ipv6_in_ipv4_frag_stats" "thunder_ipv6_in_ipv4_frag_stats" {
}
output "get_ipv6_in_ipv4_frag_stats" {
  value = ["${data.thunder_ipv6_in_ipv4_frag_stats.thunder_ipv6_in_ipv4_frag_stats}"]
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

- `bad_ip_len` (Number) Bad IP Length
- `cpu_threshold_drop` (Number) High CPU Drop
- `duplicate_first_frag` (Number) Duplicate First Fragment
- `duplicate_last_frag` (Number) Duplicate Last Fragment
- `error_drop` (Number) Fragment Processing Drop
- `exceeded_len` (Number) Payload Length Out of Bounds
- `fast_aging_set` (Number) Fragmentation Fast Aging Set
- `fast_aging_unset` (Number) Fragmentation Fast Aging Unset
- `first_gtp_packet_too_small` (Number) First GTP Fragment Too Small Drop
- `first_l4_too_small` (Number) First L4 Fragment Too Small Drop
- `first_tcp_too_small` (Number) First TCP Fragment Too Small Drop
- `fragment_queue_failure` (Number) Fragment Queue Failure
- `fragment_queue_success` (Number) Fragment Queue Success
- `high_cpu_threshold` (Number) High CPU Threshold Reached
- `icmp_dropped` (Number) ICMP Dropped
- `icmp_rcv` (Number) ICMP Received
- `icmpv6_dropped` (Number) ICMPv6 Dropped
- `icmpv6_rcv` (Number) ICMPv6 Received
- `ipd_entry_drop` (Number) DDoS Protection Drop
- `ipip_dropped` (Number) IP-in-IP Dropped
- `ipip_rcv` (Number) IP-in-IP Received
- `ipv6ip_dropped` (Number) IPv6-in-IP Dropped
- `ipv6ip_rcv` (Number) IPv6-in-IP Received
- `low_cpu_threshold` (Number) Low CPU Threshold Reached
- `max_len_exceeded` (Number) Fragment Max Data Length Exceeded
- `max_packets_exceeded` (Number) Too Many Packets Per Reassembly Drop
- `no_session_memory` (Number) Out of Session Memory
- `other_dropped` (Number) Other Dropped
- `other_rcv` (Number) Other Received
- `overlap_error` (Number) Overlapping Fragment Dropped
- `policy_drop` (Number) MTU Exceeded Policy Drop
- `reassembly_failure` (Number) Fragment Reassembly Failure
- `reassembly_success` (Number) Fragment Reassembly Success
- `sctp_dropped` (Number) SCTP Dropped
- `sctp_rcv` (Number) SCTP Received
- `session_expired` (Number) Session Expired
- `session_inserted` (Number) Session Inserted
- `session_packets_exceeded` (Number) Session Max Packets Exceeded
- `tcp_dropped` (Number) TCP Dropped
- `tcp_rcv` (Number) TCP Received
- `too_small` (Number) Fragment Too Small Drop
- `total_fragments_exceeded` (Number) Total Queued Fragments Exceeded
- `total_sessions_exceeded` (Number) Total Sessions Exceeded Drop
- `udp_dropped` (Number) UDP Dropped
- `udp_rcv` (Number) UDP Received
- `unaligned_len` (Number) Payload Length Unaligned


