---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_service_group_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_service_group_stats: Statistics for the object service-group
  PLACEHOLDER
---

# thunder_cgnv6_service_group_stats (Data Source)

`thunder_cgnv6_service_group_stats`: Statistics for the object service-group

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_service_group_stats" "thunder_cgnv6_service_group_stats" {

  name = "testing"
}
output "get_cgnv6_service_group_stats" {
  value = ["${data.thunder_cgnv6_service_group_stats.thunder_cgnv6_service_group_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) CGNV6 Service Name

### Optional

- `member_list` (Block List) (see [below for nested schema](#nestedblock--member_list))
- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--member_list"></a>
### Nested Schema for `member_list`

Required:

- `name` (String) Member name
- `port` (Number) Port number

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--member_list--stats))

<a id="nestedblock--member_list--stats"></a>
### Nested Schema for `member_list.stats`

Optional:

- `curr_conn` (Number) Current connections
- `curr_conn_overflow` (Number) Current connection counter overflow count
- `curr_req` (Number) Current requests
- `curr_ssl_conn` (Number) Current SSL connections
- `fastest_rsp_time` (Number) Fastest response time
- `peak_conn` (Number) Peak connections
- `response_time` (Number) Response time
- `slowest_rsp_time` (Number) Slowest response time
- `state_flaps` (Number) State flaps count
- `total_conn` (Number) Total connections
- `total_fwd_bytes` (Number) Total forward bytes
- `total_fwd_pkts` (Number) Total forward packets
- `total_req` (Number) Total requests
- `total_req_succ` (Number) Total requests success
- `total_rev_bytes` (Number) Total reverse bytes
- `total_rev_pkts` (Number) Total reverse packets
- `total_rev_pkts_inspected` (Number) Total reverse packets inspected
- `total_rev_pkts_inspected_status_code_2xx` (Number) Total reverse packets inspected status code 2xx
- `total_rev_pkts_inspected_status_code_non_5xx` (Number) Total reverse packets inspected status code non 5xx
- `total_ssl_conn` (Number) Total SSL connections



<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `server_selection_fail_drop` (Number) Service selection fail drop
- `server_selection_fail_reset` (Number) Service selection fail reset
- `service_peak_conn` (Number) Service peak connection


