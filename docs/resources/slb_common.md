---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_common Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_common: SLB related commands
  PLACEHOLDER
---

# thunder_slb_common (Resource)

`thunder_slb_common`: SLB related commands

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_common" "thunder_slb_common" {
  aflex_table_entry_aging_interval = 1
  aflex_table_entry_sync {
    aflex_table_entry_sync_enable        = 0
    aflex_table_entry_sync_max_key_len   = 1000
    aflex_table_entry_sync_max_value_len = 1000
    aflex_table_entry_sync_min_lifetime  = 0
  }
  after_disable              = 0
  allow_in_gateway_mode      = 0
  attack_resp_code           = 410
  auto_nat_no_ip_refresh     = "enable"
  auto_translate_port        = 0
  buff_thresh                = 0
  buff_thresh_hw_buff        = 814147845
  buff_thresh_relieve_thresh = 204157524
  buff_thresh_sys_buff_high  = 596889450
  buff_thresh_sys_buff_low   = 1233878464
  cache_expire_time          = 1
  cancel_stream_loop_limit   = 5
  cert_pinning {
    ttl = 144
    candidate_list_feedback_opt_in {
      enable   = 0
      schedule = 0
      week_day = "Monday"
      daily    = 0
    }
  }
  clientside_ip       = "10.10.10.10"
  clientside_ipv6     = "2001:db8:3333:4444:5555:6666:7777:8888"
  compress_block_size = 33117
  conn_rate_limit {
    src_ip_list {
      disable_ipv6_support = disable_ipv6_support
      protocol             = "tcp"
      limit                = 165808
      limit_period         = "100"
      exceed_action        = 0
      lock_out             = 1722
    }
  }
  custom_message        = "589"
  custom_signal_clist   = "14"
  ddos_pkt_count_thresh = 100
  ddos_pkt_size_thresh  = 64
  ddos_protection {
    ipd_enable_toggle = "disable"
    logging {
      ipd_logging_toggle = "enable"
    }
    packets_per_second {
      ipd_tcp = 200
      ipd_udp = 200
    }
  }
  disable_adaptive_resource_check    = 0
  disable_persist_scoring            = 0
  disable_port_masking               = 0
  disable_server_auto_reselect       = 0
  dns_cache_age                      = 300
  dns_cache_age_min_threshold        = 0
  dns_cache_aging_weight             = 1
  dns_cache_enable                   = 0
  dns_cache_entry_size               = 256
  dns_cache_sync                     = 0
  dns_cache_sync_entry_size          = 256
  dns_cache_sync_ttl_threshold       = 0
  dns_cache_ttl_adjustment_enable    = 0
  dns_negative_cache_enable          = 0
  dns_persistent_cache_enable        = 0
  dns_persistent_cache_hit_threshold = 0
  dns_persistent_cache_ttl_threshold = 0
  dns_response_rate_limiting {
    max_table_entries = 564259
  }
  dns_vip_stateless               = 0
  drop_icmp_to_vip_when_vip_down  = 0
  dsr_health_check_enable         = 0
  ecmp_hash                       = "system-default"
  enable_ddos                     = 0
  enable_fast_path_rerouting      = 0
  enable_l7_req_acct              = 0
  entity                          = "server"
  exclude_destination             = "local"
  extended_stats                  = 0
  fast_path_disable               = 0
  gateway_health_check            = 0
  graceful_shutdown               = 6336
  graceful_shutdown_enable        = 0
  health_check_to_all_vip         = 0
  honor_server_response_ttl       = 0
  http_fast_enable                = 0
  hw_compression                  = 0
  hw_syn_rr                       = 430867
  interval                        = 5
  ipv4_offset                     = 0
  l2l3_trunk_lb_disable           = 0
  log_for_reset_unknown_conn      = 0
  low_latency                     = 0
  max_buff_queued_per_conn        = 1000
  max_http_header_count           = 90
  max_local_rate                  = 32
  max_persistent_cache            = max_persistent_cache
  max_remote_rate                 = 15000
  monitor_mode_enable             = 0
  msl_time                        = 2
  mss_table                       = 536
  multi_cpu                       = 0
  n5_new                          = 0
  n5_old                          = 0
  ngwaf_proxy_ipv4                = "10.10.10.10"
  ngwaf_proxy_port                = 56626
  no_auto_up_on_aflex             = 0
  odd_even_nat_enable             = 0
  one_server_conn_hm_rate         = 15
  override_port                   = 0
  pbslb_entry_age                 = 6
  pbslb_overflow_glid             = "350"
  per_thr_percent                 = 5
  ping_sweep_detection            = "disable"
  pkt_rate_for_reset_unknown_conn = 68960
  player_id_check_enable          = 0
  port                            = 56279
  port_scan_detection             = "disable"
  pre_process_enable              = 0
  qat                             = 0
  quic {
    cid_len          = 4
    signature        = "105"
    signature_len    = 3
    signature_offset = 4
    cpu_offset       = 12
    quic_lb_offset   = 8
    enable_hash      = 0
    enable_signature = 0
  }
  range                               = 3
  range_end                           = 17880
  range_start                         = 25364
  rate_limit_logging                  = 0
  recursive_ns_cache                  = "honor-packet-ttl"
  reset_stale_session                 = 0
  resolve_port_conflict               = 0
  response_type                       = "single-answer"
  scale_out                           = 0
  scale_out_traffic_map               = 0
  serverside_ip                       = "10.10.10.10"
  serverside_ipv6                     = "2001:db8:3333:4444:5555:6666:7777:8888"
  service_group_on_no_dest_nat_vports = "enforce-different"
  show_slb_server_legacy_cmd          = 0
  show_slb_service_group_legacy_cmd   = 0
  show_slb_virtual_server_legacy_cmd  = 0
  snat_gwy_for_l3                     = 0
  snat_on_vip                         = 0
  snat_preserve {
    range {
      port1 = 1025
      port2 = 1025
    }
  }
  sort_res                = 0
  ssl_module_usage_enable = 0
  ssl_n5_delay_tx_enable  = 0
  ssl_ratelimit_cfg {
    disable_rate = 0
  }
  ssli_cert_not_ready_inspect_limit   = 2000
  ssli_cert_not_ready_inspect_timeout = 10
  ssli_silent_termination_enable      = 0
  ssli_sni_hash_enable                = 0
  stateless_sg_multi_binding          = 0
  stats_data_disable                  = 0
  substitute_source_mac               = 0
  timeout                             = 15
  traffic_map_type                    = "vport"
  ttl_threshold                       = 9712904
  use_default_sess_count              = 0
  use_https_proxy                     = 0
  use_mss_tab                         = 0
  vport_global                        = 75
  vport_l3v                           = 125
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `aflex_table_entry_aging_interval` (Number) aFleX table entry aging interval in second
- `aflex_table_entry_sync` (Block List, Max: 1) (see [below for nested schema](#nestedblock--aflex_table_entry_sync))
- `after_disable` (Number) Graceful shutdown after disable server/port and/or virtual server/port
- `allow_in_gateway_mode` (Number) Use source NAT gateway for L3 traffic for gateway mode
- `attack_resp_code` (Number) Custom response code
- `auto_nat_no_ip_refresh` (String) 'enable': enable; 'disable': disable;
- `auto_translate_port` (Number) Auto Translate Port range
- `buff_thresh` (Number) Set buffer threshold
- `buff_thresh_hw_buff` (Number) Set hardware buffer threshold
- `buff_thresh_relieve_thresh` (Number) Relieve threshold
- `buff_thresh_sys_buff_high` (Number) Set high water mark of system buffer
- `buff_thresh_sys_buff_low` (Number) Set low water mark of system buffer
- `cache_expire_time` (Number) Cache expiration time, default is 1 minute
- `cancel_stream_loop_limit` (Number) Set global cancel stream loop limit (cancel stream loop limit, default is 5)
- `cert_pinning` (Block List, Max: 1) (see [below for nested schema](#nestedblock--cert_pinning))
- `clientside_ip` (String) Clientside IP address
- `clientside_ipv6` (String) Clientside IPv6 address
- `compress_block_size` (Number) Set compression block size (Compression block size in bytes)
- `conn_rate_limit` (Block List, Max: 1) (see [below for nested schema](#nestedblock--conn_rate_limit))
- `custom_message` (String) Block message
- `custom_page` (String) Specify the custom webpage name
- `custom_signal_clist` (String) Provide custom signal names
- `ddos_pkt_count_thresh` (Number) Set packet count threshold for DDOS, default is 100
- `ddos_pkt_size_thresh` (Number) Set data packet size threshold for DDOS, default is 64 bytes
- `ddos_protection` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ddos_protection))
- `disable_adaptive_resource_check` (Number) Disable adaptive resource check based on buffer usage
- `disable_persist_scoring` (Number) Disable Persist Scoring
- `disable_port_masking` (Number) Disable masking of ports for CPU hashing
- `disable_server_auto_reselect` (Number) Disable auto reselection of server
- `dns_cache_age` (Number) Set DNS cache entry age, default is 300 seconds (1-1000000 seconds, default is 300 seconds)
- `dns_cache_age_min_threshold` (Number) Set DNS cache entry age minimum threshold, default is 0 seconds (1-1000000 seconds, default is 0 seconds)
- `dns_cache_aging_weight` (Number) Set DNS cache entry weight, default is 1
- `dns_cache_enable` (Number) Enable DNS cache
- `dns_cache_entry_size` (Number) Set DNS cache entry size, default is 256 bytes (1-4096 bytes, default is 256 bytes)
- `dns_cache_sync` (Number) Enable DNS cache HA sync
- `dns_cache_sync_entry_size` (Number) Only sync DNS cache with smaller size (1-4096 bytes, default is 256 bytes)
- `dns_cache_sync_ttl_threshold` (Number) Only sync DNS cache with longer TTL (0-10000000 seconds, default is 0 second)
- `dns_cache_ttl_adjustment_enable` (Number) Enable DNS cache response ttl adjustment
- `dns_negative_cache_enable` (Number) Enable DNS negative cache
- `dns_persistent_cache_enable` (Number) Enable persistent DNS cache
- `dns_persistent_cache_hit_threshold` (Number) Only save DNS cache with larger hit count (0-10000000, default is 0)
- `dns_persistent_cache_ttl_threshold` (Number) Only save DNS cache with longer TTL (0-10000000 seconds, default is 0 second)
- `dns_response_rate_limiting` (Block List, Max: 1) (see [below for nested schema](#nestedblock--dns_response_rate_limiting))
- `dns_vip_stateless` (Number) Enable DNS VIP stateless mode
- `drop_icmp_to_vip_when_vip_down` (Number) Drop ICMP to VIP when VIP down
- `dsr_health_check_enable` (Number) Enable dsr-health-check (direct server return health check)
- `ecmp_hash` (String) 'system-default': Use system default ecmp hashing algorithm; 'connection-based': Use connection information for hashing;
- `enable_ddos` (Number) Enable DDoS protection
- `enable_fast_path_rerouting` (Number) Enable Fast-Path Rerouting
- `enable_l7_req_acct` (Number) Enable L7 request accounting
- `entity` (String) 'server': Graceful shutdown server/port only; 'virtual-server': Graceful shutdown virtual server/port only;
- `exclude_destination` (String) 'local': Maximum local rate; 'remote': Maximum remote rate;  (Maximum rates)
- `extended_stats` (Number) Enable global slb extended statistics
- `fast_path_disable` (Number) Disable fast path in SLB processing
- `gateway_health_check` (Number) Enable gateway health check
- `graceful_shutdown` (Number) 1-65535, in unit of seconds
- `graceful_shutdown_enable` (Number) Enable graceful shutdown
- `health_check_to_all_vip` (Number)
- `honor_server_response_ttl` (Number) Honor the server reponse TTL
- `http_fast_enable` (Number) Enable Http Fast in SLB processing
- `hw_compression` (Number) Use hardware compression
- `hw_syn_rr` (Number) Configure hardware SYN round robin (range 1-500000)
- `interval` (Number) Specify the healthcheck interval, default is 5 seconds (Interval Value, in seconds (default 5))
- `ipv4_offset` (Number) IPv4 Octet Offset for Hash
- `ipv6_subnet` (Number) IPv6 Octet Valid Subnet Length for Hash
- `l2l3_trunk_lb_disable` (Number) Disable L2/L3 trunk LB
- `log_for_reset_unknown_conn` (Number) Log when rate exceed
- `low_latency` (Number) Enable low latency mode
- `max_buff_queued_per_conn` (Number) Set per connection buffer threshold (Buffer value range 128-4096)
- `max_http_header_count` (Number) Set maximum number of HTTP headers allowed
- `max_local_rate` (Number) Set maximum local rate
- `max_persistent_cache` (Number) Define maximum persistent cache (Maximum persistent cache entry)
- `max_remote_rate` (Number) Set maximum remote rate
- `monitor_mode_enable` (Number) Enable NG-WAF monitor mode
- `msl_time` (Number) Configure maximum session life, default is 2 seconds (1-39 seconds, default is 2 seconds)
- `mss_table` (Number) Set MSS table (128-750, default is 536)
- `multi_cpu` (Number) Specific NGWAF CPU
- `n5_new` (Number) HW assisted N5 SSL module with TLS 1.3 and TLS 1.2 support using OpenSSL 1.1.1
- `n5_old` (Number) HW assisted N5 SSL module with TLS 1.2 support using OpenSSL 0.9.7
- `ngwaf_proxy_ipv4` (String) IPv4 address
- `ngwaf_proxy_ipv6` (String) IPv6 address
- `ngwaf_proxy_port` (Number) Port
- `no_auto_up_on_aflex` (Number) Don't automatically mark vport up when aFleX is bound
- `odd_even_nat_enable` (Number) Enable odd even nat pool allocation in dual blade systems
- `one_server_conn_hm_rate` (Number) One Server Conn Health Check Rate
- `override_port` (Number) Enable override port in DSR health check mode
- `pbslb_entry_age` (Number) Set global pbslb entry age (minute)
- `pbslb_overflow_glid` (String) Apply global limit id to overflow pbslb entry
- `per_thr_percent` (Number) Percentage of default session count to use for per thread session table size
- `ping_sweep_detection` (String) 'enable': Enable ping sweep detection; 'disable': Disable ping sweep detection(default);
- `pkt_rate_for_reset_unknown_conn` (Number)
- `player_id_check_enable` (Number) Enable the Player id check
- `port` (Number) Serverside port number for SNI transmission
- `port_scan_detection` (String) 'enable': Enable port scan detection; 'disable': Disable port scan detection(default);
- `pre_process_enable` (Number) Enable NG-WAF pre-processing
- `qat` (Number) HW assisted QAT SSL module
- `quic` (Block List, Max: 1) (see [below for nested schema](#nestedblock--quic))
- `range` (Number) auto translate port range
- `range_end` (Number) port range end
- `range_start` (Number) port range start
- `rate_limit_logging` (Number) Configure rate limit logging
- `recursive_ns_cache` (String) 'honor-packet-ttl': Honor the lowest TTL among NS records in the server response; 'honor-age-config': Honor the ttl/age settings based on acos dns cache configuration;
- `reset_stale_session` (Number) Send reset if session in delete queue receives a SYN packet
- `resolve_port_conflict` (Number) Enable client port service port conflicts
- `response_type` (String) 'single-answer': Only cache DNS response with single answer; 'round-robin': Round robin;
- `scale_out` (Number) Enable SLB scale out
- `scale_out_traffic_map` (Number) Set SLB scaleout traffic-map
- `serverside_ip` (String) Serverside IP address
- `serverside_ipv6` (String) Serverside IPv6 address
- `service_group_on_no_dest_nat_vports` (String) 'allow-same': Allow the binding service-group on no-dest-nat virtual ports; 'enforce-different': Enforce that the same service-group can not be bound on different no-dest-nat virtual ports;
- `show_slb_server_legacy_cmd` (Number) Enable show slb server legacy command
- `show_slb_service_group_legacy_cmd` (Number) Enable show slb service-group legacy command
- `show_slb_virtual_server_legacy_cmd` (Number) Enable show slb virtual-server legacy command
- `snat_gwy_for_l3` (Number) Use source NAT gateway for L3 traffic for transparent mode
- `snat_on_vip` (Number) Enable source NAT traffic against VIP
- `snat_preserve` (Block List, Max: 1) (see [below for nested schema](#nestedblock--snat_preserve))
- `software` (Number) Software
- `software_tls13` (Number) Software TLS1.3
- `software_tls13_offload` (Number) Software TLS1.3 with CPU Offload Support
- `sort_res` (Number) Enable SLB sorting of resource names
- `ssl_module_usage_enable` (Number) Enable SSL module usage calculations for QAT
- `ssl_n5_delay_tx_enable` (Number) Enable delay transmission for N5-new
- `ssl_ratelimit_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ssl_ratelimit_cfg))
- `ssli_cert_not_ready_inspect_limit` (Number) SSLI asynchronized connection max number, default is 2000 (set to 0 for unlimited size)
- `ssli_cert_not_ready_inspect_timeout` (Number) SSLI asynchronized connection timeout, default is 10 seconds (seconds, set to 0 for never timeout)
- `ssli_silent_termination_enable` (Number) Terminate the SSLi sessions silently without sending RST/FIN packet
- `ssli_sni_hash_enable` (Number) Enable SSLi SNI hash table
- `stateless_sg_multi_binding` (Number) Enable stateless service groups to be assigned to multiple L2/L3 DSR VIPs
- `stats_data_disable` (Number) Disable global slb data statistics
- `substitute_source_mac` (Number) Substitute Source MAC Address to that of the outgoing interface
- `timeout` (Number) Specify the healthcheck timeout value, default is 15 seconds (Timeout Value, in seconds (default 15))
- `traffic_map_type` (String) 'vport': traffic-map per vport; 'global': global traffic-map;
- `ttl_threshold` (Number) Only cache DNS response with longer TTL
- `use_default_sess_count` (Number) Use default session count
- `use_https_proxy` (Number) NG-WAF connects to Cloud through proxy server
- `use_mgmt_port` (Number) Use management port to connect
- `use_mss_tab` (Number) Use MSS based on internal table for SLB processing
- `uuid` (String) uuid of the object
- `vport_global` (Number) Configure periodic showtech vport paging global limit
- `vport_l3v` (Number) Configure periodic showtech vport paging l3v limit

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--aflex_table_entry_sync"></a>
### Nested Schema for `aflex_table_entry_sync`

Optional:

- `aflex_table_entry_sync_enable` (Number) Enable aflex table sync
- `aflex_table_entry_sync_max_key_len` (Number) aflex table entry max key length to sync
- `aflex_table_entry_sync_max_value_len` (Number) aflex table entry max value length to sync
- `aflex_table_entry_sync_min_lifetime` (Number) aflex table entry minimum lifetime to sync
- `uuid` (String) uuid of the object


<a id="nestedblock--cert_pinning"></a>
### Nested Schema for `cert_pinning`

Optional:

- `candidate_list_feedback_opt_in` (Block List, Max: 1) (see [below for nested schema](#nestedblock--cert_pinning--candidate_list_feedback_opt_in))
- `ttl` (Number) The ttl of local cert pinning candidate list, multiple of 10 minutes, default is 144 (1440 minutes)
- `uuid` (String) uuid of the object

<a id="nestedblock--cert_pinning--candidate_list_feedback_opt_in"></a>
### Nested Schema for `cert_pinning.candidate_list_feedback_opt_in`

Optional:

- `daily` (Number) Every day
- `day_time` (String) Time of day to update (hh:mm) in 24 hour local time
- `enable` (Number) Enable the feedback function
- `schedule` (Number) schedule the uploading time, default is daily 00:00
- `use_mgmt_port` (Number) Use management port to connect
- `uuid` (String) uuid of the object
- `week_day` (String) 'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;
- `week_time` (String) Time of day to update (hh:mm) in 24 hour local time
- `weekly` (Number) Every week



<a id="nestedblock--conn_rate_limit"></a>
### Nested Schema for `conn_rate_limit`

Optional:

- `src_ip_list` (Block List) (see [below for nested schema](#nestedblock--conn_rate_limit--src_ip_list))

<a id="nestedblock--conn_rate_limit--src_ip_list"></a>
### Nested Schema for `conn_rate_limit.src_ip_list`

Required:

- `disable_ipv6_support` (Number)
- `protocol` (String) 'tcp': Set TCP connection rate limit; 'udp': Set UDP packet rate limit;

Optional:

- `exceed_action` (Number) Set action if threshold exceeded
- `limit` (Number) Set max connections per period
- `limit_period` (String) '100': 100 ms; '1000': 1000 ms;
- `lock_out` (Number) Set lockout period in seconds if threshold exceeded
- `log` (Number) Send log if threshold exceeded
- `shared` (Number) Set threshold shared amongst all virtual ports
- `uuid` (String) uuid of the object



<a id="nestedblock--ddos_protection"></a>
### Nested Schema for `ddos_protection`

Optional:

- `ipd_enable_toggle` (String) 'enable': Enable SLB DDoS protection; 'disable': Disable SLB DDoS protection (default);
- `logging` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ddos_protection--logging))
- `packets_per_second` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ddos_protection--packets_per_second))

<a id="nestedblock--ddos_protection--logging"></a>
### Nested Schema for `ddos_protection.logging`

Optional:

- `ipd_logging_toggle` (String) 'enable': Enable SLB DDoS protection logging (default); 'disable': Disable SLB DDoS protection logging;


<a id="nestedblock--ddos_protection--packets_per_second"></a>
### Nested Schema for `ddos_protection.packets_per_second`

Optional:

- `ipd_tcp` (Number) Configure packets-per-second threshold per TCP port (default: 200)
- `ipd_udp` (Number) Configure packets-per-second threshold per UDP port (default: 200)



<a id="nestedblock--dns_response_rate_limiting"></a>
### Nested Schema for `dns_response_rate_limiting`

Optional:

- `max_table_entries` (Number) Maximum number of entries allowed
- `uuid` (String) uuid of the object


<a id="nestedblock--quic"></a>
### Nested Schema for `quic`

Optional:

- `cid_len` (Number) Length of CID
- `cpu_offset` (Number) Offset for Encoded CPU
- `enable_hash` (Number) Enable CID Hashing
- `enable_signature` (Number) Enable CID Signature Validation
- `quic_lb_offset` (Number) Offset for QUIC-LB
- `signature` (String) Set CID Signature
- `signature_len` (Number) Offset for CID Signature
- `signature_offset` (Number) Offset for CID Signature
- `uuid` (String) uuid of the object


<a id="nestedblock--snat_preserve"></a>
### Nested Schema for `snat_preserve`

Optional:

- `range` (Block List) (see [below for nested schema](#nestedblock--snat_preserve--range))

<a id="nestedblock--snat_preserve--range"></a>
### Nested Schema for `snat_preserve.range`

Optional:

- `port1` (Number) start port
- `port2` (Number) end port which is greater than start



<a id="nestedblock--ssl_ratelimit_cfg"></a>
### Nested Schema for `ssl_ratelimit_cfg`

Optional:

- `disable_rate` (Number) Disable HW SSL Rate limit for N5-new
- `tls12_rate` (Number) Enabling Rateliming for TLS1.2 HW requests per chip in 1K - default 120
- `tls13_rate` (Number) Enabling Rateliming for TLS1.3 HW requests per chip in 1K - default 72

