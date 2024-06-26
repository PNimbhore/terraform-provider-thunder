---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_template_tcp_progression_tracking_time_window_tracking Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_template_tcp_progression_tracking_time_window_tracking: Configure and enable TCP Progression Tracking per Time Window
  PLACEHOLDER
---

# thunder_ddos_template_tcp_progression_tracking_time_window_tracking (Resource)

`thunder_ddos_template_tcp_progression_tracking_time_window_tracking`: Configure and enable TCP Progression Tracking per Time Window

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_tcp_progression_tracking_time_window_tracking" "thunder_ddos_template_tcp_progression_tracking_time_window_tracking" {
  name                                = "test"
  progression_tracking_win_enabled    = "enable-check"
  window_sent_max                     = 21795
  window_sent_min                     = 485
  window_rcvd_max                     = 53441
  window_rcvd_min                     = 5219
  window_rcvd_sent_ratio_min          = 624
  window_rcvd_sent_ratio_max          = 42396
  window_violation                    = 233
  progression_tracking_windows_action = "drop"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name
- `progression_tracking_win_enabled` (String) 'enable-check': Enable Progression Tracking per Time Window;

### Optional

- `progression_tracking_windows_action` (String) 'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;
- `progression_tracking_windows_action_list_name` (String) Configure action-list to take when progression tracking violation exceed
- `uuid` (String) uuid of the object
- `window_rcvd_max` (Number) Set the maximum total received byte
- `window_rcvd_min` (Number) Set the minimum total received byte
- `window_rcvd_sent_ratio_max` (Number) Set the maximum received to sent ratio (in unit of 0.1% [1:1000])
- `window_rcvd_sent_ratio_min` (Number) Set the minimum received to sent ratio (in unit of 0.1% [1:1000])
- `window_sent_max` (Number) Set the maximum total sent byte
- `window_sent_min` (Number) Set the minimum total sent byte
- `window_violation` (Number) Set the violation threshold

### Read-Only

- `id` (String) The ID of this resource.


