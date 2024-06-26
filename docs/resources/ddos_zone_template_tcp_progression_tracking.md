---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_zone_template_tcp_progression_tracking Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_zone_template_tcp_progression_tracking: Configure and enable TCP Progression Tracking
  PLACEHOLDER
---

# thunder_ddos_zone_template_tcp_progression_tracking (Resource)

`thunder_ddos_zone_template_tcp_progression_tracking`: Configure and enable TCP Progression Tracking

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_tcp_progression_tracking" "thunder_ddos_zone_template_tcp_progression_tracking" {
  name = "test"
  connection_tracking {
    progression_tracking_conn_enabled = "enable-check"
    conn_sent_max                     = 20
    conn_sent_min                     = 10
    conn_rcvd_max                     = 20
    conn_rcvd_min                     = 10
    conn_rcvd_sent_ratio_min          = 15
    conn_rcvd_sent_ratio_max          = 30
    conn_duration_max                 = 44
    conn_duration_min                 = 22
    conn_violation                    = 5
    progression_tracking_conn_action  = "drop"
  }
  progression_tracking_enabled     = "enable-check"
  request_response_model           = "enable"
  violation                        = 33
  ignore_tls_handshake             = 1
  response_length_max              = 20
  request_length_min               = 20
  request_length_max               = 40
  response_request_min_ratio       = 20
  response_request_max_ratio       = 40
  first_request_max_time           = 3
  request_to_response_max_time     = 30
  response_to_request_max_time     = 3
  profiling_request_response_model = 1
  profiling_connection_life_model  = 1
  progression_tracking_action      = "drop"
  time_window_tracking {
    progression_tracking_win_enabled    = "enable-check"
    window_sent_max                     = 20
    window_sent_min                     = 10
    window_rcvd_max                     = 20
    window_rcvd_min                     = 10
    window_rcvd_sent_ratio_min          = 10
    window_rcvd_sent_ratio_max          = 20
    window_violation                    = 2
    progression_tracking_windows_action = "drop"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name
- `progression_tracking_enabled` (String) 'enable-check': Enable Progression Tracking Check;

### Optional

- `connection_tracking` (Block List, Max: 1) (see [below for nested schema](#nestedblock--connection_tracking))
- `first_request_max_time` (Number) Set the maximum wait time from connection creation until the first data is transmitted over the connection (100 ms)
- `ignore_tls_handshake` (Number) Ignore TLS handshake
- `profiling_connection_life_model` (Number) Enable auto-config progression tracking learning for connection model
- `profiling_request_response_model` (Number) Enable auto-config progression tracking learning for Request Response model
- `profiling_time_window_model` (Number) Enable auto-config progression tracking learning for time window model
- `progression_tracking_action` (String) 'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;
- `progression_tracking_action_list_name` (String) Configure action-list to take when progression tracking violation exceed
- `request_length_max` (Number) Set the maximum request length
- `request_length_min` (Number) Set the minimum request length
- `request_response_model` (String) 'enable': Enable Request Response Model; 'disable': Disable Request Response Model;
- `request_to_response_max_time` (Number) Set the maximum request to response time (100 ms)
- `response_length_max` (Number) Set the maximum response length
- `response_length_min` (Number) Set the minimum response length
- `response_request_max_ratio` (Number) Set the maximum response to request ratio (in unit of 0.1% [1:1000])
- `response_request_min_ratio` (Number) Set the minimum response to request ratio (in unit of 0.1% [1:1000])
- `response_to_request_max_time` (Number) Set the maximum response to request time (100 ms)
- `time_window_tracking` (Block List, Max: 1) (see [below for nested schema](#nestedblock--time_window_tracking))
- `uuid` (String) uuid of the object
- `violation` (Number) Set the violation threshold

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--connection_tracking"></a>
### Nested Schema for `connection_tracking`

Optional:

- `conn_duration_max` (Number) Set the maximum duration time (in unit of 100ms, up to 24 hours)
- `conn_duration_min` (Number) Set the minimum duration time (in unit of 100ms, up to 24 hours)
- `conn_rcvd_max` (Number) Set the maximum total received byte
- `conn_rcvd_min` (Number) Set the minimum total received byte
- `conn_rcvd_sent_ratio_max` (Number) Set the maximum received to sent ratio (in unit of milli-, 0.001)
- `conn_rcvd_sent_ratio_min` (Number) Set the minimum received to sent ratio (in unit of milli-, 0.001)
- `conn_sent_max` (Number) Set the maximum total sent byte
- `conn_sent_min` (Number) Set the minimum total sent byte
- `conn_violation` (Number) Set the violation threshold
- `progression_tracking_conn_action` (String) 'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;
- `progression_tracking_conn_action_list_name` (String) Configure action-list to take when progression tracking violation exceed
- `progression_tracking_conn_enabled` (String) 'enable-check': Enable General Progression Tracking per Connection;
- `uuid` (String) uuid of the object


<a id="nestedblock--time_window_tracking"></a>
### Nested Schema for `time_window_tracking`

Optional:

- `progression_tracking_win_enabled` (String) 'enable-check': Enable Progression Tracking per Time Window;
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


