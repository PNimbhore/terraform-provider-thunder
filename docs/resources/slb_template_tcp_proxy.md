---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_tcp_proxy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_template_tcp_proxy: TCP Proxy
  PLACEHOLDER
---

# thunder_slb_template_tcp_proxy (Resource)

`thunder_slb_template_tcp_proxy`: TCP Proxy

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_tcp_proxy" "test" {
  name                       = "tcp_proxy1"
  server_down_action         = "FIN"
  dynamic_buffer_allocation  = 1
  nagle                      = 1
  initial_window_size        = 65535
  reset_rev                  = 1
  disable                    = 1
  maxburst                   = 90
  half_close_idle_timeout    = 100
  ack_aggressiveness         = "low"
  backend_wscale             = 14
  del_session_on_server_down = 1
  disable_abc                = 1
  disable_sack               = 1
  disable_tcp_timestamps     = 1
  disable_window_scale       = 1
  early_retransmit           = 1
  fin_timeout                = 55
  force_delete_timeout       = 3
  half_open_idle_timeout     = 30
  idle_timeout               = 2097120
  init_cwnd                  = 15
  insert_client_ip           = 1
  invalid_rate_limit         = 60000000
  keepalive_interval         = 12000
  keepalive_probes           = 5
  limited_slowstart          = 214748364
  min_rto                    = 900
  mss                        = 8999
  psh_flag_optimization      = 1
  qos                        = 63
  reassembly_timeout         = 299
  reassembly_limit           = 5
  receive_buffer             = 2147483647
  reno                       = 1
  reset_fwd                  = 1
  retransmit_retries         = 15
  syn_retries                = 15
  timewait                   = 59
  transmit_buffer            = 2147483647
  user_tag                   = "modify tcp proxy template"
  proxy_header {
    proxy_header_action = "insert"
    version             = "v1"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) TCP Proxy Template Name

### Optional

- `ack_aggressiveness` (String) 'low': Delayed ACK; 'medium': Delayed ACK, with ACK on each packet with PUSH flag; 'high': ACK on each packet;
- `alive_if_active` (Number) keep connection alive if active traffic
- `backend_wscale` (Number) The TCP window scale used for the server side, default is off (number)
- `del_session_on_server_down` (Number) Delete session if the server/port goes down (either disabled/hm down)
- `disable` (Number) send reset to client when server is disabled
- `disable_abc` (Number) Appropriate Byte Counting RFC 3465 Disabled, default is enabled (Appropriate Byte Counting (ABC) is enabled by default)
- `disable_sack` (Number) disable Selective Ack Option
- `disable_tcp_timestamps` (Number) disable TCP Timestamps Option
- `disable_window_scale` (Number) disable TCP Window-Scale Option
- `down` (Number) send reset to client when server is down
- `dynamic_buffer_allocation` (Number) Optimally adjust the transmit and receive buffer sizes of TCP proxy while keeping their sum constant
- `early_retransmit` (Number) Configure the Early-Retransmit Algorithm (RFC 5827) (Early-Retransmit is disabled by default)
- `fin_timeout` (Number) FIN timeout (sec), default is disabled (number)
- `force_delete_timeout` (Number) The maximum time that a session can stay in the system before being deleted, default is off (number (second))
- `force_delete_timeout_100ms` (Number) The maximum time that a session can stay in the system before being deleted, default is off (number in 100ms)
- `half_close_idle_timeout` (Number) TCP Half Close Idle Timeout (sec), default is off (cmd is deprecated, use fin-timeout instead) (number)
- `half_open_idle_timeout` (Number) TCP Half Open Idle Timeout (sec), default is off (number)
- `idle_timeout` (Number) Idle Timeout (Interval of 60 seconds), default is 600 (idle timeout in second, default 600)
- `init_cwnd` (Number) The initial congestion control window size (packets), default is 10 (init-cwnd in packets, default 10)
- `initial_window_size` (Number) Set the initial window size, default is off (number)
- `insert_client_ip` (Number) Insert client ip into TCP option
- `invalid_rate_limit` (Number) Invalid Packet Response Rate Limit (ms), default is 500 (number default 500 challenges)
- `keepalive_interval` (Number) Interval between keepalive probes (sec), default is off (number (seconds))
- `keepalive_probes` (Number) Number of keepalive probes sent, default is off
- `limited_slowstart` (Number) RFC 3742 Limited Slow-Start for TCP with Large Congestion Windows (number)
- `maxburst` (Number) The max packet count sent per transmission event (number)
- `min_rto` (Number) The minmum retransmission timeout, default is 200ms (number)
- `mss` (Number) Responding MSS to use if client MSS is large, default is off (number)
- `nagle` (Number) Enable Nagle Algorithm
- `naked_ack_on_handshake` (Number) Send naked ack before data during 3-way handshake
- `proxy_header` (Block List, Max: 1) (see [below for nested schema](#nestedblock--proxy_header))
- `psh_flag_optimization` (Number) Enable Optimized PSH Flag Use
- `qos` (Number) QOS level (number)
- `reassembly_limit` (Number) The reassembly queuing limit, default is 25 segments (number)
- `reassembly_timeout` (Number) The reassembly timeout, default is 30sec (number)
- `receive_buffer` (Number) TCP Receive Buffer (default 200k) (number default 200000 bytes)
- `reno` (Number) Enable Reno Congestion Control Algorithm
- `reset_fwd` (Number) send reset to server if error happens
- `reset_rev` (Number) send reset to client if error happens
- `retransmit_retries` (Number) Number of Retries for Retransmit, default is 5
- `server_down_action` (String) 'FIN': FIN Connection; 'RST': Reset Connection;
- `syn_retries` (Number) SYN Retry Numbers, default is 5
- `timewait` (Number) Timewait Threshold (sec), default 5 (number)
- `transmit_buffer` (Number) TCP Transmit Buffer (default 200k) (number default 200000 bytes)
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--proxy_header"></a>
### Nested Schema for `proxy_header`

Optional:

- `proxy_header_action` (String) 'insert': Insert proxy header;
- `version` (String) 'v1': version 1; 'v2': version 2;


