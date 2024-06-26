---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_service_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_service_group: Authentication service group
  PLACEHOLDER
---

# thunder_aam_authentication_service_group (Resource)

`thunder_aam_authentication_service_group`: Authentication service group

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_service_group" "thunder_aam_authentication_service_group" {

  lb_method = "round-robin"
  member_list {
    name            = "test"
    port            = 1812
    member_state    = "enable"
    member_priority = 5
    user_tag        = "6"
    sampling_enable {
      counters1 = "all"
    }
  }
  name     = "test"
  protocol = "udp"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify AAM service group name

### Optional

- `health_check` (String) Health Check (Monitor Name)
- `health_check_disable` (Number) Disable health check
- `lb_method` (String) 'round-robin': Round robin on server level;
- `member_list` (Block List) (see [below for nested schema](#nestedblock--member_list))
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `protocol` (String) 'tcp': TCP AAM service; 'udp': UDP AAM service;
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--member_list"></a>
### Nested Schema for `member_list`

Required:

- `name` (String) Member name
- `port` (Number) Port number

Optional:

- `member_priority` (Number) Priority of Port in the Group
- `member_state` (String) 'enable': Enable member service port; 'disable': Disable member service port;
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--member_list--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--member_list--sampling_enable"></a>
### Nested Schema for `member_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_conn': Total established connections; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_status_code_2xx': Total reverse packets inspected status code 2xx; 'total_rev_pkts_inspected_status_code_non_5xx': Total reverse packets inspected status code non 5xx; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total requests successful; 'peak_conn': peak_conn; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'curr_conn_overflow': Current connection counter overflow count;



<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'server_selection_fail_drop': Drops due to Service selection failure; 'server_selection_fail_reset': Resets sent out for Service selection failure; 'service_peak_conn': Peak connection count for the Service Group; 'service_healthy_host': Service Group healthy host count; 'service_unhealthy_host': Service Group unhealthy host count; 'service_req_count': Service Group request count; 'service_resp_count': Service Group response count; 'service_resp_2xx': Service Group response 2xx count; 'service_resp_3xx': Service Group response 3xx count; 'service_resp_4xx': Service Group response 4xx count; 'service_resp_5xx': Service Group response 5xx count; 'service_curr_conn_overflow': Current connection counter overflow count;


