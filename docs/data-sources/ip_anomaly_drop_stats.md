---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_anomaly_drop_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_anomaly_drop_stats: Statistics for the object anomaly-drop
  PLACEHOLDER
---

# thunder_ip_anomaly_drop_stats (Data Source)

`thunder_ip_anomaly_drop_stats`: Statistics for the object anomaly-drop

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_anomaly_drop_stats" "thunder_ip_anomaly_drop_stats" {
}
output "get_ip_anomaly_drop_stats" {
  value = ["${data.thunder_ip_anomaly_drop_stats.thunder_ip_anomaly_drop_stats}"]
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

- `bad_ip_flg` (Number) Bad IP Flags Drop
- `bad_ip_frg_offset` (Number) Bad IP Fragment Offset Drop
- `bad_ip_hdrlen` (Number) Bad IP Header Len Drop
- `bad_ip_payload_len` (Number) Bad IP Payload Len Drop
- `bad_ip_ttl` (Number) Bad IP TTL Drop
- `bad_tcp_urg_offset` (Number) TCP Bad Urgent Offset Drop
- `csum` (Number) Bad IP Checksum Drop
- `emp_frg` (Number) Empty Fragment Drop
- `emp_mic_frg` (Number) Micro Fragment Drop
- `frg` (Number) IPv4 Fragment Drop
- `gre_pptp_err` (Number) GRE PPTP Error Drop
- `ipip_tnl_err` (Number) IP-over-IP Tunnel Error Drop
- `ipip_tnl_msmtch` (Number) IP-over-IP Tunnel Mismatch Drop
- `ipv6_eh_ah` (Number) IPv6 Authentication Header Drop
- `ipv6_eh_dest` (Number) IPv6 Destination Header Drop
- `ipv6_eh_esp` (Number) IPv6 ESP Header Drop
- `ipv6_eh_frag` (Number) IPv6 Fragmentation Header Drop
- `ipv6_eh_hbh` (Number) IPv6 Hop by Hop Header Drop
- `ipv6_eh_malformed` (Number) IPv6 Malformed Extension Header Drop
- `ipv6_eh_mobility` (Number) IPv6 Mobility Header Drop
- `ipv6_eh_none` (Number) IPv6 No Next Header Drop
- `ipv6_eh_other` (Number) IPv6 Unknown Extension Header Drop
- `ipv6_eh_routing` (Number) IPv6 Routing Header Drop
- `land` (Number) Land Attack Drop
- `no_ip_payload` (Number) No IP Payload drop
- `nvgre_err` (Number) GRE Tunnel Error Drop
- `opt` (Number) IPv4 Options Drop
- `over_ip_payload` (Number) Oversize IP Payload Drop
- `pod` (Number) ICMP Ping of Death Drop
- `runt_ip_hdr` (Number) Runt IP Header Drop
- `runt_tcp_udp_hdr` (Number) Runt TCP/UDP Header Drop
- `tcp_bad_csum` (Number) TCP Bad Checksum Drop
- `tcp_bad_iplen` (Number) TCP Bad IP Length Drop
- `tcp_frg_hdr` (Number) TCP Fragmented Header Drop
- `tcp_null_frg` (Number) TCP Null Flags Drop
- `tcp_null_scan` (Number) TCP Null Scan Drop
- `tcp_opt_err` (Number) TCP Option Error Drop
- `tcp_sht_hdr` (Number) TCP Short Header Drop
- `tcp_syn_fin` (Number) TCP Syn and Fin Drop
- `tcp_syn_frg` (Number) TCP Syn Fragment Drop
- `tcp_xmas` (Number) TCP XMAS Flags Drop
- `tcp_xmas_scan` (Number) TCP XMAS Scan Drop
- `udp_bad_csum` (Number) UDP Bad Checksum Drop
- `udp_bad_len` (Number) UDP Bad Length Drop
- `udp_kerb_frg` (Number) UDP Kerberos Fragment Drop
- `udp_port_lb` (Number) UDP Port Loopback Drop
- `udp_srt_hdr` (Number) UDP Short Header Drop
- `vxlan_err` (Number) VXLAN Tunnel Error Drop


