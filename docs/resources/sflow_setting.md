---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_sflow_setting Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_sflow_setting: Configure sFlow
  PLACEHOLDER
---

# thunder_sflow_setting (Resource)

`thunder_sflow_setting`: Configure sFlow

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_setting" "thunder_sflow_setting" {
  append_mapping_info         = 1
  counter_polling_interval    = 20
  default_counter_polling_mtu = 1500
  local_collection            = "enable"
  management_link_utilization = 5733
  packet_sampling_rate        = 5733
  source_ip_use_mgmt          = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `append_mapping_info` (Number) Allow TPS to always send mapping ctr block (260, 271, and 272)
- `counter_polling_interval` (Number) sFlow counter polling interval, default is 20
- `default_counter_polling_mtu` (Number) Default MTU for counter-polling packets - DDoS 3.2 format only (Default: 1500)
- `local_collection` (String) 'enable': Enable local sflow collection; 'disable': Disable local sflow collection;
- `local_t1_polling_interval` (Number) Set sFlow local counter polling interval for T1 stats
- `local_t2_polling_interval` (Number) Set sFlow local counter polling interval for T2 stats
- `management_link_utilization` (Number) limit management link speed in (Mbps)
- `management_link_utilization_percentage` (Number) percentage limit of the management link speed (Default is 30%)
- `max_header` (Number) Configure maximum number of bytes that should be copied from a sampled packet (default: 128) (The maximum number of bytes (Default: 128))
- `packet_sampling_rate` (Number) sFlow packet sampling rate, default is 1000
- `port_range_end` (Number) Source port-range end
- `port_range_start` (Number) Source port-range
- `randomize_source_port` (String) 'enable': Randomize source port; 'disable': Fix source port 6343; 'packet-sampling-only': Only randomized source port for packet-sampling (Default);
- `source_ip_use_mgmt` (Number) Use management interface's IP address for source IP of sFlow packets
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


