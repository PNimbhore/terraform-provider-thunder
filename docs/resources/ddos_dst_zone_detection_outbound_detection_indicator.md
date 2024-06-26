---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_zone_detection_outbound_detection_indicator Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_zone_detection_outbound_detection_indicator: Outbound indicator configuration
  PLACEHOLDER
---

# thunder_ddos_dst_zone_detection_outbound_detection_indicator (Resource)

`thunder_ddos_dst_zone_detection_outbound_detection_indicator`: Outbound indicator configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection_outbound_detection_indicator" "thunder_ddos_dst_zone_detection_outbound_detection_indicator" {
  threshold_num = 470644108
  type          = "pkt-rate"
  user_tag      = "86"
  zone_name     = "test"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `type` (String) 'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets;
- `zone_name` (String) ZoneName

### Optional

- `data_packet_size` (Number) Expected minimal data size
- `tcp_window_size` (Number) Expected minimal window size
- `threshold_large_num` (Number) Threshold for each geo-location
- `threshold_num` (Number) Threshold for each geo-location
- `threshold_str` (String) Threshold for each geo-location (Non-zero floating point)
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


