---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_zone_ip_proto_proto_number_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_zone_ip_proto_proto_number_oper: Operational Status for the object proto-number
  PLACEHOLDER
---

# thunder_ddos_dst_zone_ip_proto_proto_number_oper (Data Source)

`thunder_ddos_dst_zone_ip_proto_proto_number_oper`: Operational Status for the object proto-number

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_ip_proto_proto_number_oper" "thunder_ddos_dst_zone_ip_proto_proto_number_oper" {
  zone_name    = "test"
  protocol_num = "80"

}
output "get_ddos_dst_zone_ip_proto_proto_number_oper" {
  value = ["${data.thunder_ddos_dst_zone_ip_proto_proto_number_oper.thunder_ddos_dst_zone_ip_proto_proto_number_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `protocol_num` (Number) Protocol Number
- `zone_name` (String) ZoneName

### Optional

- `ip_filtering_policy_oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip_filtering_policy_oper))
- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))
- `port_ind` (Block List, Max: 1) (see [below for nested schema](#nestedblock--port_ind))
- `progression_tracking` (Block List, Max: 1) (see [below for nested schema](#nestedblock--progression_tracking))
- `topk_destinations` (Block List, Max: 1) (see [below for nested schema](#nestedblock--topk_destinations))
- `topk_sources` (Block List, Max: 1) (see [below for nested schema](#nestedblock--topk_sources))

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



<a id="nestedblock--port_ind"></a>
### Nested Schema for `port_ind`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--port_ind--oper))

<a id="nestedblock--port_ind--oper"></a>
### Nested Schema for `port_ind.oper`

Optional:

- `active_time` (Number)
- `current_level` (String)
- `details` (Number)
- `detection_data_source` (String)
- `escalation_timestamp` (String)
- `indicators` (Block List) (see [below for nested schema](#nestedblock--port_ind--oper--indicators))
- `initial_learning` (String)
- `ipv6` (String)
- `sources` (Number)
- `sources_all_entries` (Number)
- `src_entry_list` (Block List) (see [below for nested schema](#nestedblock--port_ind--oper--src_entry_list))
- `subnet_ip_addr` (String)
- `subnet_ipv6_addr` (String)
- `total_score` (String)

<a id="nestedblock--port_ind--oper--indicators"></a>
### Nested Schema for `port_ind.oper.indicators`

Optional:

- `indicator_cfg` (Block List) (see [below for nested schema](#nestedblock--port_ind--oper--indicators--indicator_cfg))
- `indicator_index` (Number)
- `indicator_name` (String)
- `rate` (String)
- `score` (String)
- `src_maximum` (String)
- `zone_adaptive_threshold` (String)
- `zone_average` (String)
- `zone_maximum` (String)
- `zone_minimum` (String)
- `zone_non_zero_minimum` (String)

<a id="nestedblock--port_ind--oper--indicators--indicator_cfg"></a>
### Nested Schema for `port_ind.oper.indicators.indicator_cfg`

Optional:

- `level` (Number)
- `source_threshold` (String)
- `zone_threshold` (String)



<a id="nestedblock--port_ind--oper--src_entry_list"></a>
### Nested Schema for `port_ind.oper.src_entry_list`

Optional:

- `active_time` (Number)
- `current_level` (String)
- `detection_data_source` (String)
- `escalation_timestamp` (String)
- `indicators` (Block List) (see [below for nested schema](#nestedblock--port_ind--oper--src_entry_list--indicators))
- `initial_learning` (String)
- `src_address_str` (String)
- `src_level` (String)
- `total_score` (String)

<a id="nestedblock--port_ind--oper--src_entry_list--indicators"></a>
### Nested Schema for `port_ind.oper.src_entry_list.indicators`

Optional:

- `indicator_index` (Number)
- `indicator_name` (String)
- `rate` (String)
- `score` (String)
- `src_average` (String)
- `src_maximum` (String)
- `src_minimum` (String)
- `src_non_zero_minimum` (String)





<a id="nestedblock--progression_tracking"></a>
### Nested Schema for `progression_tracking`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--progression_tracking--oper))

<a id="nestedblock--progression_tracking--oper"></a>
### Nested Schema for `progression_tracking.oper`

Optional:

- `indicators` (Block List) (see [below for nested schema](#nestedblock--progression_tracking--oper--indicators))
- `learning_brief` (Number)
- `learning_details` (Number)
- `recommended_template` (Number)
- `template_debug_table` (Number)

<a id="nestedblock--progression_tracking--oper--indicators"></a>
### Nested Schema for `progression_tracking.oper.indicators`

Optional:

- `average` (String)
- `indicator_index` (Number)
- `indicator_name` (String)
- `maximum` (String)
- `minimum` (String)
- `num_sample` (Number)
- `standard_deviation` (String)




<a id="nestedblock--topk_destinations"></a>
### Nested Schema for `topk_destinations`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--topk_destinations--oper))

<a id="nestedblock--topk_destinations--oper"></a>
### Nested Schema for `topk_destinations.oper`

Optional:

- `details` (Number)
- `entry_list` (Block List) (see [below for nested schema](#nestedblock--topk_destinations--oper--entry_list))
- `finished` (Number)
- `indicators` (Block List) (see [below for nested schema](#nestedblock--topk_destinations--oper--indicators))
- `next_indicator` (Number)
- `top_k_key` (String)

<a id="nestedblock--topk_destinations--oper--entry_list"></a>
### Nested Schema for `topk_destinations.oper.entry_list`

Optional:

- `address_str` (String)
- `indicators` (Block List) (see [below for nested schema](#nestedblock--topk_destinations--oper--entry_list--indicators))

<a id="nestedblock--topk_destinations--oper--entry_list--indicators"></a>
### Nested Schema for `topk_destinations.oper.entry_list.indicators`

Optional:

- `indicator_index` (Number)
- `indicator_name` (String)
- `max_peak` (String)
- `psd_wdw_cnt` (Number)
- `rate` (String)



<a id="nestedblock--topk_destinations--oper--indicators"></a>
### Nested Schema for `topk_destinations.oper.indicators`

Optional:

- `destinations` (Block List) (see [below for nested schema](#nestedblock--topk_destinations--oper--indicators--destinations))
- `indicator_index` (Number)
- `indicator_name` (String)

<a id="nestedblock--topk_destinations--oper--indicators--destinations"></a>
### Nested Schema for `topk_destinations.oper.indicators.destinations`

Optional:

- `address` (String)
- `rate` (String)





<a id="nestedblock--topk_sources"></a>
### Nested Schema for `topk_sources`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--topk_sources--oper))

<a id="nestedblock--topk_sources--oper"></a>
### Nested Schema for `topk_sources.oper`

Optional:

- `details` (Number)
- `entry_list` (Block List) (see [below for nested schema](#nestedblock--topk_sources--oper--entry_list))
- `finished` (Number)
- `indicators` (Block List) (see [below for nested schema](#nestedblock--topk_sources--oper--indicators))
- `next_indicator` (Number)
- `top_k_key` (String)

<a id="nestedblock--topk_sources--oper--entry_list"></a>
### Nested Schema for `topk_sources.oper.entry_list`

Optional:

- `address_str` (String)
- `indicators` (Block List) (see [below for nested schema](#nestedblock--topk_sources--oper--entry_list--indicators))

<a id="nestedblock--topk_sources--oper--entry_list--indicators"></a>
### Nested Schema for `topk_sources.oper.entry_list.indicators`

Optional:

- `indicator_index` (Number)
- `indicator_name` (String)
- `max_peak` (String)
- `psd_wdw_cnt` (Number)
- `rate` (String)



<a id="nestedblock--topk_sources--oper--indicators"></a>
### Nested Schema for `topk_sources.oper.indicators`

Optional:

- `indicator_index` (Number)
- `indicator_name` (String)
- `sources` (Block List) (see [below for nested schema](#nestedblock--topk_sources--oper--indicators--sources))

<a id="nestedblock--topk_sources--oper--indicators--sources"></a>
### Nested Schema for `topk_sources.oper.indicators.sources`

Optional:

- `address` (String)
- `rate` (String)

