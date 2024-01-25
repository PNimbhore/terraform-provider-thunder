---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_virtual_port Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_template_virtual_port: Virtual port template
  PLACEHOLDER
---

# thunder_slb_template_virtual_port (Resource)

`thunder_slb_template_virtual_port`: Virtual port template

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_virtual_port" "virtual-port" {
  name                       = "testvirtualport"
  user_tag                   = "virtualport1"
  reset_unknown_conn         = 1
  ignore_tcp_msl             = 1
  rate                       = 500
  snat_msl                   = 10
  allow_syn_otherflags       = 1
  aflow                      = 1
  conn_limit                 = 100
  conn_limit_reset           = 1
  conn_limit_no_logging      = 1
  drop_unknown_conn          = 1
  reset_l7_on_failover       = 1
  pkt_rate_type              = "src-ip-port"
  pkt_rate_interval          = "100ms"
  pkt_rate_limit_reset       = 9000
  snat_port_preserve         = 1
  conn_rate_limit            = 50
  rate_interval              = "100ms"
  conn_rate_limit_reset      = 1
  conn_rate_limit_no_logging = 1
  non_syn_initiation         = 1
  dscp                       = 50
  log_options                = "no-logging"
  when_rr_enable             = 1
  allow_vip_to_rport_mapping = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Virtual port template name

### Optional

- `aflow` (Number) Use aFlow to eliminate the traffic surge
- `allow_syn_otherflags` (Number) Allow initial SYN packet with other flags
- `allow_vip_to_rport_mapping` (Number) Allow mapping of VIP to real port
- `conn_limit` (Number) Connection limit
- `conn_limit_no_logging` (Number) Do not log connection over limit event
- `conn_limit_reset` (Number) Send client reset when connection over limit
- `conn_rate_limit` (Number) Connection rate limit
- `conn_rate_limit_no_logging` (Number) Do not log connection over limit event
- `conn_rate_limit_reset` (Number) Send client reset when connection rate over limit
- `drop_unknown_conn` (Number) Drop conection if receives TCP packet without SYN or RST flag and it does not belong to any existing connections
- `dscp` (Number) Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)
- `ignore_tcp_msl` (Number) reclaim TCP resource immediately without MSL
- `log_options` (String) 'no-logging': Do not log over limit event; 'no-repeat-logging': log once for over limit event. Default is log once per minute;
- `non_syn_initiation` (Number) Allow initial TCP packet to be non-SYN
- `pkt_rate_interval` (String) '100ms': Source IP and port rate limit per 100ms; 'second': Source IP and port rate limit per second (default);
- `pkt_rate_limit_reset` (Number) send client-side reset (reset after packet limit)
- `pkt_rate_type` (String) 'src-ip-port': Source IP and port rate limit; 'src-port': Source port rate limit;
- `rate` (Number) Source IP and port rate limit (Packet rate limit)
- `rate_interval` (String) '100ms': Use 100 ms as sampling interval; 'second': Use 1 second as sampling interval;
- `reset_l7_on_failover` (Number) Send reset to L7 client and server connection upon a failover
- `reset_unknown_conn` (Number) Send reset back if receives TCP packet without SYN or RST flag and it does not belong to any existing connections
- `snat_msl` (Number) Source NAT MSL (Source NAT MSL value (seconds))
- `snat_port_preserve` (Number) Source NAT Port Preservation
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `when_rr_enable` (Number) Only do rate limit if CPU RR triggered

### Read-Only

- `id` (String) The ID of this resource.

