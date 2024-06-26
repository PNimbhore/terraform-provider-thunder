---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_template_udp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_template_udp: UDP template configuration
  PLACEHOLDER
---

# thunder_ddos_template_udp (Resource)

`thunder_ddos_template_udp`: UDP template configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_udp" "thunder_ddos_template_udp" {
  name = "test"
  age  = 8
  drop_known_resp_src_port_cfg {
    drop_known_resp_src_port = 1
    exclude_src_resp_port    = 1
  }
  drop_ntp_monlist = 1
  filter_list {
    udp_filter_seq       = 1
    udp_filter_regex     = "1209"
    udp_filter_unmatched = 1
    udp_filter_action    = "blacklist-src"
    user_tag             = "16"
  }
  max_payload_size        = 1413
  min_payload_size        = 1054
  per_conn_pkt_rate_limit = 1246
  per_conn_rate_interval  = "1sec"
  previous_salt_timeout   = 1
  spoof_detect_cfg {
    spoof_detect                        = 1
    min_retry_gap_interval              = "1sec"
    spoof_detect_retry_timeout_val_only = 5
    min_retry_gap                       = 3
  }
  tunnel_encap {
    ip_encap = 1
    always {
      ipv4_addr         = "10.10.10.10"
      preserve_src_ipv4 = 1
    }
  }
  user_tag = "31"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) DDOS UDP Template Name

### Optional

- `age` (Number) Configure session age(in minutes) for UDP sessions
- `drop_known_resp_src_port_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--drop_known_resp_src_port_cfg))
- `drop_ntp_monlist` (Number) Drop NTP monlist request/response
- `filter_list` (Block List) (see [below for nested schema](#nestedblock--filter_list))
- `max_payload_size` (Number) Maximum UDP payload size for each single packet
- `min_payload_size` (Number) Minimum UDP payload size for each single packet
- `per_conn_pkt_rate_limit` (Number) Packet rate limit per connection per rate-interval
- `per_conn_rate_interval` (String) '100ms': 100ms; '1sec': 1sec;
- `previous_salt_timeout` (Number) Token-Authentication previous salt-prefix timeout in minutes, default is 1 min
- `public_ipv4_addr` (String) IP address
- `public_ipv6_addr` (String) IPV6 address
- `spoof_detect_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spoof_detect_cfg))
- `token_authentication` (Number) Enable Token Authentication
- `token_authentication_formula` (String) 'md5_Salt-SrcIp-SrcPort-DstIp-DstPort': md5 of Salt-SrcIp-SrcPort-DstIp-DstPort; 'md5_Salt-DstIp-DstPort': md5 of Salt-DstIp-DstPort; 'md5_Salt-SrcIp-DstIp': md5 of Salt-SrcIp-DstIp; 'md5_Salt-SrcPort-DstPort': md5 of Salt-SrcPort-DstPort; 'md5_Salt-UintDstIp-DstPort': Using the uint value of IP for md5 of Salt-DstIp-DstPort; 'sha1_Salt-SrcIp-SrcPort-DstIp-DstPort': sha1 of Salt-SrcIp-SrcPort-DstIp-DstPort; 'sha1_Salt-DstIp-DstPort': sha1 of Salt-DstIp-DstPort; 'sha1_Salt-SrcIp-DstIp': sha1 of Salt-SrcIp-DstIp; 'sha1_Salt-SrcPort-DstPort': sha1 of Salt-SrcPort-DstPort; 'sha1_Salt-UintDstIp-DstPort': Using the uint value of IP for sha1 of Salt-DstIp-DstPort;
- `token_authentication_hw_assist_disable` (Number) token-authentication disable hardware assistance
- `token_authentication_public_address` (Number) The server public IP address
- `token_authentication_salt_prefix` (Number) token-authentication salt-prefix
- `token_authentication_salt_prefix_curr` (Number)
- `token_authentication_salt_prefix_prev` (Number)
- `tunnel_encap` (Block List, Max: 1) (see [below for nested schema](#nestedblock--tunnel_encap))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--drop_known_resp_src_port_cfg"></a>
### Nested Schema for `drop_known_resp_src_port_cfg`

Optional:

- `drop_known_resp_src_port` (Number) Drop well-known if src-port is less than 1024
- `exclude_src_resp_port` (Number) excluding src port equal destination port


<a id="nestedblock--filter_list"></a>
### Nested Schema for `filter_list`

Required:

- `udp_filter_seq` (Number) Sequence number

Optional:

- `byte_offset_filter` (String) Filter Expression using Berkeley Packet Filter syntax
- `udp_filter_action` (String) 'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;
- `udp_filter_regex` (String) Regex Expression
- `udp_filter_unmatched` (Number) action taken when it does not match
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object


<a id="nestedblock--spoof_detect_cfg"></a>
### Nested Schema for `spoof_detect_cfg`

Optional:

- `min_retry_gap` (Number) Optional minimum gap between 2 UDP packets for spoof-detect pass, unit is specified by min-retry-gap-interval
- `min_retry_gap_interval` (String) '100ms': 100ms; '1sec': 1sec;
- `spoof_detect` (Number) Force client to retry on udp
- `spoof_detect_retry_timeout` (Number) timeout in seconds
- `spoof_detect_retry_timeout_val_only` (Number) timeout in seconds


<a id="nestedblock--tunnel_encap"></a>
### Nested Schema for `tunnel_encap`

Optional:

- `always` (Block List, Max: 1) (see [below for nested schema](#nestedblock--tunnel_encap--always))
- `gre_always` (Block List, Max: 1) (see [below for nested schema](#nestedblock--tunnel_encap--gre_always))
- `gre_encap` (Number) Enable Tunnel encapsulation using GRE
- `ip_encap` (Number) Enable Tunnel encapsulation using IP in IP

<a id="nestedblock--tunnel_encap--always"></a>
### Nested Schema for `tunnel_encap.always`

Optional:

- `ipv4_addr` (String) IPv4 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)
- `ipv6_addr` (String) IPv6 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)
- `preserve_src_ipv4` (Number) Use original source ip for encapsulation
- `preserve_src_ipv6` (Number) Use original source ip for encapsulation


<a id="nestedblock--tunnel_encap--gre_always"></a>
### Nested Schema for `tunnel_encap.gre_always`

Optional:

- `gre_ipv4` (String) IPv4 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)
- `gre_ipv6` (String) IPv6 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)
- `key_ipv4` (String) Encapsulate with key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)
- `key_ipv6` (String) Encapsulate with key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)
- `preserve_src_ipv4_gre` (Number) Use original source ip for encapsulation
- `preserve_src_ipv6_gre` (Number) Use original source ip for encapsulation


