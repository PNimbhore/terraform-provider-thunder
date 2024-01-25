---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_acos_events_local_logging_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_acos_events_local_logging_stats: Statistics for the object local-logging
  PLACEHOLDER
---

# thunder_acos_events_local_logging_stats (Data Source)

`thunder_acos_events_local_logging_stats`: Statistics for the object local-logging

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_acos_events_local_logging_stats" "thunder_acos_events_local_logging_stats" {
}
output "get_acos_events_local_logging_stats" {
  value = ["${data.thunder_acos_events_local_logging_stats.thunder_acos_events_local_logging_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `disk_over_thres` (Number) Number of logs Dropped, Disk reached threshold
- `freed` (Number) Local logging Stopped
- `in_bytes` (Number) Number of bytes successfully stored
- `in_bytes_backlog` (Number) Number of backlog bytes loaded from disk
- `in_discard_bytes` (Number) Number of old bytes discarded to fit in new logs
- `in_discard_logs` (Number) Number of old logs discarded to fit in new logs
- `in_logs` (Number) Number of logs successfully stored
- `in_logs_backlog` (Number) Number of backlogs loaded from disk
- `in_store_fail_no_space` (Number) Number of logs Dropped, failed without disk space
- `init_fail` (Number) Local logging Init Fail
- `init_pass` (Number) Local logging Init Successful
- `not_inited` (Number) Number of logs Dropped, Local logging not inited
- `out_bytes` (Number) Number of bytes sent to log-servers
- `out_error` (Number) Number of errors during send
- `out_logs` (Number) Number of logs sent to log servers
- `rate_limited` (Number) Number of logs Dropped, Rate limited
- `remaining_bytes` (Number) Total number of remaining bytes yet to be sent
- `remaining_logs` (Number) Total number of remaining logs yet to be sent
- `sent_to_store` (Number) Number of logs sent to be stored
- `sent_to_store_fail` (Number) Number of Logs sent to be stored Failed
- `store_fail` (Number) Number of logs failed to be stored

