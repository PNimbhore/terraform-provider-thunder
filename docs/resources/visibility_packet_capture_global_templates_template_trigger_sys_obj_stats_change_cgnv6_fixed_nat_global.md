---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_global Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_global: Configure triggers for cgnv6.fixed-nat.global object
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_global (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_global`: Configure triggers for cgnv6.fixed-nat.global object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_global" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_global" {

  name = "test"
  trigger_stats_inc {
    nat_port_unavailable_tcp                = 1
    nat_port_unavailable_udp                = 1
    nat_port_unavailable_icmp               = 1
    session_user_quota_exceeded             = 1
    fullcone_failure                        = 1
    nat44_inbound_filtered                  = 1
    nat64_inbound_filtered                  = 1
    dslite_inbound_filtered                 = 1
    nat44_eif_limit_exceeded                = 1
    nat64_eif_limit_exceeded                = 1
    dslite_eif_limit_exceeded               = 1
    standby_drop                            = 1
    fixed_nat_fullcone_self_hairpinning_dro = 1
    sixrd_drop                              = 1
    dest_rlist_drop                         = 1
    dest_rlist_pass_through                 = 1
    dest_rlist_snat_drop                    = 1
    config_not_found                        = 1
    port_overload_failed                    = 1
    ha_session_user_quota_exceeded          = 1
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `trigger_stats_inc` (Block List, Max: 1) (see [below for nested schema](#nestedblock--trigger_stats_inc))
- `trigger_stats_rate` (Block List, Max: 1) (see [below for nested schema](#nestedblock--trigger_stats_rate))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--trigger_stats_inc"></a>
### Nested Schema for `trigger_stats_inc`

Optional:

- `config_not_found` (Number) Enable automatic packet-capture for Fixed NAT Config not Found
- `dest_rlist_drop` (Number) Enable automatic packet-capture for Fixed NAT Dest Rule List Drop
- `dest_rlist_pass_through` (Number) Enable automatic packet-capture for Fixed NAT Dest Rule List Pass-Through
- `dest_rlist_snat_drop` (Number) Enable automatic packet-capture for Fixed NAT Dest Rules List Source NAT Drop
- `dslite_eif_limit_exceeded` (Number) Enable automatic packet-capture for DS-Lite Endpoint-Independent-Filtering Limit Exceeded
- `dslite_inbound_filtered` (Number) Enable automatic packet-capture for DS-Lite Endpoint-Dependent Filtering Drop
- `fixed_nat_fullcone_self_hairpinning_dro` (Number) Enable automatic packet-capture for Self-Hairpinning Drop
- `fullcone_failure` (Number) Enable automatic packet-capture for Full-Cone Session Creation Failed
- `ha_session_user_quota_exceeded` (Number) Enable automatic packet-capture for HA Sessions User Quota Exceeded
- `nat44_eif_limit_exceeded` (Number) Enable automatic packet-capture for NAT44 Endpoint-Independent-Filtering Limit Exceeded
- `nat44_inbound_filtered` (Number) Enable automatic packet-capture for NAT44 Endpoint-Dependent Filtering Drop
- `nat64_eif_limit_exceeded` (Number) Enable automatic packet-capture for NAT64 Endpoint-Independent-Filtering Limit Exceeded
- `nat64_inbound_filtered` (Number) Enable automatic packet-capture for NAT64 Endpoint-Dependent Filtering Drop
- `nat_port_unavailable_icmp` (Number) Enable automatic packet-capture for ICMP NAT Port Unavailable
- `nat_port_unavailable_tcp` (Number) Enable automatic packet-capture for TCP NAT Port Unavailable
- `nat_port_unavailable_udp` (Number) Enable automatic packet-capture for UDP NAT Port Unavailable
- `port_overload_failed` (Number) Enable automatic packet-capture for Port overload failed
- `session_user_quota_exceeded` (Number) Enable automatic packet-capture for Sessions User Quota Exceeded
- `sixrd_drop` (Number) Enable automatic packet-capture for Fixed NAT IPv6 in IPv4 Packet Drop
- `standby_drop` (Number) Enable automatic packet-capture for Fixed NAT LID Standby Drop
- `uuid` (String) uuid of the object


<a id="nestedblock--trigger_stats_rate"></a>
### Nested Schema for `trigger_stats_rate`

Optional:

- `config_not_found` (Number) Enable automatic packet-capture for Fixed NAT Config not Found
- `dest_rlist_drop` (Number) Enable automatic packet-capture for Fixed NAT Dest Rule List Drop
- `dest_rlist_pass_through` (Number) Enable automatic packet-capture for Fixed NAT Dest Rule List Pass-Through
- `dest_rlist_snat_drop` (Number) Enable automatic packet-capture for Fixed NAT Dest Rules List Source NAT Drop
- `dslite_eif_limit_exceeded` (Number) Enable automatic packet-capture for DS-Lite Endpoint-Independent-Filtering Limit Exceeded
- `dslite_inbound_filtered` (Number) Enable automatic packet-capture for DS-Lite Endpoint-Dependent Filtering Drop
- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `fixed_nat_fullcone_self_hairpinning_dro` (Number) Enable automatic packet-capture for Self-Hairpinning Drop
- `fullcone_failure` (Number) Enable automatic packet-capture for Full-Cone Session Creation Failed
- `ha_session_user_quota_exceeded` (Number) Enable automatic packet-capture for HA Sessions User Quota Exceeded
- `nat44_eif_limit_exceeded` (Number) Enable automatic packet-capture for NAT44 Endpoint-Independent-Filtering Limit Exceeded
- `nat44_inbound_filtered` (Number) Enable automatic packet-capture for NAT44 Endpoint-Dependent Filtering Drop
- `nat64_eif_limit_exceeded` (Number) Enable automatic packet-capture for NAT64 Endpoint-Independent-Filtering Limit Exceeded
- `nat64_inbound_filtered` (Number) Enable automatic packet-capture for NAT64 Endpoint-Dependent Filtering Drop
- `nat_port_unavailable_icmp` (Number) Enable automatic packet-capture for ICMP NAT Port Unavailable
- `nat_port_unavailable_tcp` (Number) Enable automatic packet-capture for TCP NAT Port Unavailable
- `nat_port_unavailable_udp` (Number) Enable automatic packet-capture for UDP NAT Port Unavailable
- `port_overload_failed` (Number) Enable automatic packet-capture for Port overload failed
- `session_user_quota_exceeded` (Number) Enable automatic packet-capture for Sessions User Quota Exceeded
- `sixrd_drop` (Number) Enable automatic packet-capture for Fixed NAT IPv6 in IPv4 Packet Drop
- `standby_drop` (Number) Enable automatic packet-capture for Fixed NAT LID Standby Drop
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

