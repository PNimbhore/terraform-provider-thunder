---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_link_probe_entry_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_link_probe_entry_stats: Statistics for the object entry
  PLACEHOLDER
---

# thunder_slb_link_probe_entry_stats (Data Source)

`thunder_slb_link_probe_entry_stats`: Statistics for the object entry

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_link_probe_entry_stats" "thunder_slb_link_probe_entry_stats" {
}
output "get_slb_link_probe_entry_stats" {
  value = ["${data.thunder_slb_link_probe_entry_stats.thunder_slb_link_probe_entry_stats}"]
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

- `curr_entries` (Number) Current Entry Count
- `err_entry_create_failed` (Number) Entry Creation Failure
- `err_entry_create_oom` (Number) Entry Creation Out of Memory
- `err_entry_insert_failed` (Number) Entry Insert Failed
- `err_l4_sess_alloc` (Number) Error allocating L4 session for probe
- `err_probe_tcp_conn_send` (Number) Error in initiating TCP connection for probe
- `err_smart_nat_alloc` (Number) Error creating Smart NAT Instance
- `err_smart_nat_port_alloc` (Number) Error obtaining Smart NAT source port
- `err_tmpl_probe_create_failed` (Number) Probe Template Creation Failure
- `err_tmpl_probe_create_oom` (Number) Probe Template Creation Out of Memory
- `probe_tcp_conn_sent` (Number) TCP connection sent for probe
- `total_created` (Number) Total Entry Created
- `total_freed` (Number) Total Entry Freed
- `total_http_probes_sent` (Number) Total HTTP Probes Sent out
- `total_http_response_bad` (Number) Total HTTP responses not matching probe template config
- `total_http_response_good` (Number) Total HTTP responses matching probe template config
- `total_http_response_received` (Number) Total HTTP responses received
- `total_inserted` (Number) Total Entry Inserted
- `total_ready_to_free` (Number) Total Entry Ready To Free
- `total_tcp_err` (Number) Total TCP errors in probes sent out


