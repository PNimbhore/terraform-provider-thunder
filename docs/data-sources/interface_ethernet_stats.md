---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_ethernet_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_interface_ethernet_stats: Statistics for the object ethernet
  PLACEHOLDER
---

# thunder_interface_ethernet_stats (Data Source)

`thunder_interface_ethernet_stats`: Statistics for the object ethernet

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_ethernet_stats" "thunder_interface_ethernet_stats" {
  ifnum = 2
}
output "get_interface_ethernet_stats" {
  value = ["${data.thunder_interface_ethernet_stats.thunder_interface_ethernet_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ifnum` (Number) Ethernet interface number

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `bytes_input` (Number) Input bytes
- `bytes_output` (Number) Output bytes
- `collisions` (Number) Collisions
- `crc` (Number) CRC
- `drops` (Number) Drops
- `frame` (Number) Frames
- `giants` (Number) Giants
- `giants_output` (Number) Output Giants
- `input_errors` (Number) Input errors
- `input_utilization` (Number) Input Utilization
- `load_interval` (Number) Load Interval
- `output_errors` (Number) Output errors
- `output_utilization` (Number) Output Utilization
- `packets_input` (Number) Input packets
- `packets_output` (Number) Output packets
- `rate_byte_rcvd` (Number) Byte received rate bits/sec
- `rate_byte_sent` (Number) Byte sent rate bits/sec
- `rate_pkt_rcvd` (Number) Packet received rate packets/sec
- `rate_pkt_sent` (Number) Packet sent rate packets/sec
- `received_broadcasts` (Number) Received broadcasts
- `received_multicasts` (Number) Received multicasts
- `received_unicasts` (Number) Received unicasts
- `runts` (Number) Runts
- `transmitted_broadcasts` (Number) Transmitted broadcasts
- `transmitted_multicasts` (Number) Transmitted multicasts
- `transmitted_unicasts` (Number) Transmitted unicasts


