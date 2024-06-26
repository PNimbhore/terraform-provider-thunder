---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_entry_port_range_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_entry_port_range_oper: Operational Status for the object port-range
  PLACEHOLDER
---

# thunder_ddos_dst_entry_port_range_oper (Data Source)

`thunder_ddos_dst_entry_port_range_oper`: Operational Status for the object port-range

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_entry_port_range_oper" "thunder_ddos_dst_entry_port_range_oper" {

  port_range_end   = 90
  port_range_start = 80
  protocol         = "tcp"
  dst_entry_name   = "test"
}
output "get_ddos_dst_entry_port_range_oper" {
  value = ["${data.thunder_ddos_dst_entry_port_range_oper.thunder_ddos_dst_entry_port_range_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dst_entry_name` (String) DstEntryName
- `port_range_end` (Number) Port-Range End Port Number
- `port_range_start` (Number) Port-Range Start Port Number
- `protocol` (String) 'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port;

### Optional

- `ip_filtering_policy_oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip_filtering_policy_oper))
- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))
- `pattern_recognition` (Block List, Max: 1) (see [below for nested schema](#nestedblock--pattern_recognition))
- `pattern_recognition_pu_details` (Block List, Max: 1) (see [below for nested schema](#nestedblock--pattern_recognition_pu_details))
- `port_ind` (Block List, Max: 1) (see [below for nested schema](#nestedblock--port_ind))
- `progression_tracking` (Block List, Max: 1) (see [below for nested schema](#nestedblock--progression_tracking))
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

- `all_ip_protos` (Number)
- `all_ports` (Number)
- `all_src_ports` (Number)
- `app_stat` (Number)
- `ddos_entry_list` (Block List) (see [below for nested schema](#nestedblock--oper--ddos_entry_list))
- `domain_name` (String)
- `entry_displayed_count` (Number)
- `hw_blacklisted` (String)
- `l4_ext_rate` (Number)
- `port_protocol` (String)
- `reource_limit_alloc` (String)
- `reporting_status` (Number)
- `resource_limit_config` (String)
- `resource_limit_remain` (String)
- `resource_usage` (Number)
- `service_displayed_count` (Number)
- `sflow_source_id` (Number)
- `suffix_request_rate` (Number)

<a id="nestedblock--oper--ddos_entry_list"></a>
### Nested Schema for `oper.ddos_entry_list`

Optional:

- `age_str` (String)
- `app_stat1_limit` (String)
- `app_stat2_limit` (String)
- `app_stat3_limit` (String)
- `app_stat4_limit` (String)
- `app_stat5_limit` (String)
- `app_stat6_limit` (String)
- `app_stat7_limit` (String)
- `app_stat8_limit` (String)
- `connection_limit` (String)
- `connection_rate_limit` (String)
- `current_app_stat1` (String)
- `current_app_stat2` (String)
- `current_app_stat3` (String)
- `current_app_stat4` (String)
- `current_app_stat5` (String)
- `current_app_stat6` (String)
- `current_app_stat7` (String)
- `current_app_stat8` (String)
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
- `kbit_rate_limit` (String)
- `level_str` (String)
- `lockup_time_str` (String)
- `packet_rate_limit` (String)
- `port_str` (String)
- `sflow_source_id` (String)
- `src_address_str` (String)
- `state_str` (String)



<a id="nestedblock--pattern_recognition"></a>
### Nested Schema for `pattern_recognition`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--pattern_recognition--oper))

<a id="nestedblock--pattern_recognition--oper"></a>
### Nested Schema for `pattern_recognition.oper`

Optional:

- `filter_count` (Number)
- `filter_list` (Block List) (see [below for nested schema](#nestedblock--pattern_recognition--oper--filter_list))
- `filter_threshold` (Number)
- `peace_pkt_count` (Number)
- `state` (String)
- `timestamp` (String)
- `war_pkt_count` (Number)
- `war_pkt_percentage` (Number)

<a id="nestedblock--pattern_recognition--oper--filter_list"></a>
### Nested Schema for `pattern_recognition.oper.filter_list`

Optional:

- `filter_desc` (String)
- `filter_enabled` (Number)
- `filter_expr` (String)
- `hardware_filter` (Number)
- `processing_unit` (String)
- `sample_ratio` (Number)




<a id="nestedblock--pattern_recognition_pu_details"></a>
### Nested Schema for `pattern_recognition_pu_details`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--pattern_recognition_pu_details--oper))

<a id="nestedblock--pattern_recognition_pu_details--oper"></a>
### Nested Schema for `pattern_recognition_pu_details.oper`

Optional:

- `all_filters` (Block List) (see [below for nested schema](#nestedblock--pattern_recognition_pu_details--oper--all_filters))

<a id="nestedblock--pattern_recognition_pu_details--oper--all_filters"></a>
### Nested Schema for `pattern_recognition_pu_details.oper.all_filters`

Optional:

- `filter_count` (Number)
- `filter_list` (Block List) (see [below for nested schema](#nestedblock--pattern_recognition_pu_details--oper--all_filters--filter_list))
- `filter_threshold` (Number)
- `peace_pkt_count` (Number)
- `processing_unit` (String)
- `state` (String)
- `timestamp` (String)
- `war_pkt_count` (Number)
- `war_pkt_percentage` (Number)

<a id="nestedblock--pattern_recognition_pu_details--oper--all_filters--filter_list"></a>
### Nested Schema for `pattern_recognition_pu_details.oper.all_filters.filter_list`

Optional:

- `filter_desc` (String)
- `filter_enabled` (Number)
- `filter_expr` (String)
- `hardware_filter` (Number)
- `sample_ratio` (Number)





<a id="nestedblock--port_ind"></a>
### Nested Schema for `port_ind`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--port_ind--oper))

<a id="nestedblock--port_ind--oper"></a>
### Nested Schema for `port_ind.oper`

Optional:

- `detection_data_source` (String)
- `indicators` (Block List) (see [below for nested schema](#nestedblock--port_ind--oper--indicators))

<a id="nestedblock--port_ind--oper--indicators"></a>
### Nested Schema for `port_ind.oper.indicators`

Optional:

- `entry_average` (String)
- `entry_maximum` (String)
- `entry_minimum` (String)
- `entry_non_zero_minimum` (String)
- `indicator_index` (Number)
- `indicator_name` (String)
- `rate` (String)
- `src_maximum` (String)




<a id="nestedblock--progression_tracking"></a>
### Nested Schema for `progression_tracking`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--progression_tracking--oper))

<a id="nestedblock--progression_tracking--oper"></a>
### Nested Schema for `progression_tracking.oper`

Optional:

- `indicators` (Block List) (see [below for nested schema](#nestedblock--progression_tracking--oper--indicators))

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


