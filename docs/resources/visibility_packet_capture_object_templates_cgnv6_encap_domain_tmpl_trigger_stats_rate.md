---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl_trigger_stats_rate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl_trigger_stats_rate: Configure stats to triggers packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl_trigger_stats_rate (Resource)

`thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl_trigger_stats_rate`: Configure stats to triggers packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl_trigger_stats_rate" {

  name                                = "test"
  duration                            = 60
  inbound_addr_port_validation_failed = 1
  inbound_dest_unreachable            = 1
  inbound_rev_lookup_failed           = 1
  interface_not_configured            = 1
  outbound_addr_validation_failed     = 1
  outbound_dest_unreachable           = 1
  outbound_rev_lookup_failed          = 1
  packet_mtu_exceeded                 = 1
  threshold_exceeded_by               = 5
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `inbound_addr_port_validation_failed` (Number) Enable automatic packet-capture for Inbound IPv4 Destination Address Port Validation Failed
- `inbound_dest_unreachable` (Number) Enable automatic packet-capture for Inbound IPv6 Destination Address Unreachable
- `inbound_rev_lookup_failed` (Number) Enable automatic packet-capture for Inbound IPv4 Reverse Route Lookup Failed
- `interface_not_configured` (Number) Enable automatic packet-capture for Interfaces not Configured Dropped
- `outbound_addr_validation_failed` (Number) Enable automatic packet-capture for Outbound IPv6 Source Address Validation Failed
- `outbound_dest_unreachable` (Number) Enable automatic packet-capture for Outbound IPv4 Destination Address Unreachable
- `outbound_rev_lookup_failed` (Number) Enable automatic packet-capture for Outbound IPv6 Reverse Route Lookup Failed
- `packet_mtu_exceeded` (Number) Enable automatic packet-capture for Packet Exceeded MTU
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


