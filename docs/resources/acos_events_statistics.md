---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_acos_events_statistics Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_acos_events_statistics: acos events global statistics
  PLACEHOLDER
---

# thunder_acos_events_statistics (Resource)

`thunder_acos_events_statistics`: acos events global statistics

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_statistics" "thunder_acos_events_statistics" {
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'msg_sent': Messages sent, to Remote; 'msg_sent_logdb': Messages sent, to LogDB; 'msg_dropped_format_not_defined': Messages Dropped, format not defined; 'msg_dropped_malloc_failure': Messages Dropped, malloc failure; 'msg_dropped_no_template': Messages Dropped, no active template; 'msg_dropped_selector': Messages Dropped, selector does not enable msg; 'msg_dropped_too_long': Messages Dropped, invalid length; 'msg_dropped_craft_fail': Messages Dropped, msg crafting failed; 'msg_dropped_local_log_ratelimit': Messages Dropped, local log ratelimited; 'msg_dropped_remote_log_ratelimit': Messages Dropped, remote log ratelimited; 'msg_dropped_send_failed': Messages Dropped, send failed; 'msg_dropped_no_active_member': Messages Dropped, no active member in collector grp; 'msg_dropped_route_fail': Messages Dropped, Route lookup failed; 'msg_dropped_other': Messages Dropped, unexpected error; 'no_template': Message API called, with no active template; 'msg_dropped_lost_during_config_change': Messages Dropped, lost during config change; 'local_enqueue_pass': Messages enqueue to Logd passed; 'msg_sent_to_logd': Messages sent to Logd via IPC; 'msg_retry_after_socket_fail': Messages retried to be sent to Logd via IPC; 'msg_sent_direct_syslog': Messages sent to syslog directly from axlog; 'msg_dropped_send_to_logd_fail': Messages Dropped, send to Logd via IPC failed; 'msg_dropped_trylock_fail': Messages Dropped, Trylock failed in axlog; 'msg_dropped_remote_cplane_log_ratelimit': Messages Dropped, Remote cplane log ratelimited; 'msg_dropped_remote_dplane_log_ratelimit': Messages Dropped, Remote dplane log ratelimited; 'msg_dropped_local_enqueue_failed': Messages Dropped, Enqueue to Logd failed; 'msg_dropped_grp_not_used': Messages Dropped, Collector group not used; 'msg_sent_remote_cplane': Messages Sent, to remote in logd; 'msg_dropped_no_template_logd': Messages Dropped, no active template in Logd; 'msg_dropped_lost_during_config_change_logd': Messages Dropped, lost during config change in Logd; 'msg_dropped_craft_fail_logd': Messages Dropped, msg crafting failed in Logd; 'msg_dropped_send_failed_logd': Messages Dropped, send failed in Logd; 'msg_dropped_no_active_member_logd': Messages Dropped, no active member in collector grp in Logd; 'msg_dropped_other_logd': Messages Dropped, unexpected error in Logd; 'msg_dropped_invalid_part': Messages Dropped, Invalid partition Id; 'acos_evt_test_logs_ticks': Number of ticks when running ACOS Event Test Logs; 'param_msg_sent_to_hc': Parameterized log sent to HC; 'param_msg_sent_fail': Parameterized log send to HC failed; 'param_msg_encode_fail': Parameterized log AVRO encoding failed;


