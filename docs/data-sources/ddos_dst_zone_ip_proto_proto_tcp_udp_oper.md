---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_oper: Operational Status for the object proto-tcp-udp
  PLACEHOLDER
---

# thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_oper (Data Source)

`thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_oper`: Operational Status for the object proto-tcp-udp

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_oper" "thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_oper" {
  zone_name = "test"
  protocol  = "tcp"

}
output "get_ddos_dst_zone_ip_proto_proto_tcp_udp_oper" {
  value = ["${data.thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_oper.thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `protocol` (String) 'tcp': ip-proto tcp; 'udp': ip-proto udp;
- `zone_name` (String) ZoneName

### Optional

- `ip_filtering_policy_oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip_filtering_policy_oper))
- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--ip_filtering_policy_oper"></a>
### Nested Schema for `ip_filtering_policy_oper`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip_filtering_policy_oper--oper))

<a id="nestedblock--ip_filtering_policy_oper--oper"></a>
### Nested Schema for `ip_filtering_policy_oper.oper`

Optional:

- `rule_list` (Block List) (see [below for nested schema](#nestedblock--ip_filtering_policy_oper--oper--rule_list))

<a id="nestedblock--ip_filtering_policy_oper--oper--rule_list"></a>
### Nested Schema for `ip_filtering_policy_oper.oper.rule_list`

Optional:

- `hits` (Number)
- `seq` (Number)




<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `app_stat` (Number)
- `authenticated` (Number)
- `black_listed` (Number)
- `class_list` (String)
- `ddos_entry_list` (Block List) (see [below for nested schema](#nestedblock--oper--ddos_entry_list))
- `domain_name` (String)
- `entry_displayed_count` (Number)
- `exceeded` (Number)
- `hw_blacklisted` (Number)
- `indicator_detail` (Number)
- `indicators` (Number)
- `ipv6` (String)
- `level` (Number)
- `overflow_policy` (Number)
- `reporting_status` (Number)
- `service_displayed_count` (Number)
- `sources` (Number)
- `sources_all_entries` (Number)
- `subnet_ip_addr` (String)
- `subnet_ipv6_addr` (String)
- `suffix_request_rate` (Number)
- `white_listed` (Number)

<a id="nestedblock--oper--ddos_entry_list"></a>
### Nested Schema for `oper.ddos_entry_list`

Optional:

- `age` (Number)
- `bl_reasoning_rcode` (String)
- `bl_reasoning_timestamp` (String)
- `bw_state` (String)
- `connection_limit` (String)
- `connection_rate_limit` (String)
- `current_connection_rate` (String)
- `current_connections` (String)
- `current_frag_packet_rate` (String)
- `current_kbit_rate` (String)
- `current_packet_rate` (String)
- `debug_str` (String)
- `dst_address_str` (String)
- `dynamic_entry_count` (String)
- `dynamic_entry_limit` (String)
- `frag_packet_rate_limit` (String)
- `is_auth_passed` (String)
- `is_connection_rate_exceed` (Number)
- `is_connections_exceed` (Number)
- `is_frag_packet_rate_exceed` (Number)
- `is_kbit_rate_exceed` (Number)
- `is_packet_rate_exceed` (Number)
- `kbit_rate_limit` (String)
- `level` (Number)
- `lockup_time` (Number)
- `packet_rate_limit` (String)
- `sflow_source_id` (Number)


