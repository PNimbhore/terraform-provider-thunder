---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist: Configure triggers for slb.persist object
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist`: Configure triggers for slb.persist object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by      = 5
    duration                   = 60
    hash_tbl_trylock_fail      = 1
    hash_tbl_create_fail       = 1
    hash_tbl_rst_updown        = 1
    hash_tbl_rst_adddel        = 1
    url_hash_fail              = 1
    header_hash_fail           = 1
    src_ip_fail                = 1
    src_ip_new_sess_cache_fail = 1
    src_ip_new_sess_sel_fail   = 1
    src_ip_hash_fail           = 1
    dst_ip_fail                = 1
    dst_ip_new_sess_cache_fail = 1
    dst_ip_new_sess_sel_fail   = 1
    dst_ip_hash_fail           = 1
    cssl_sid_not_found         = 1
    cssl_sid_not_match         = 1
    sssl_sid_not_found         = 1
    sssl_sid_not_match         = 1
    ssl_sid_persist_fail       = 1
    ssl_sid_session_fail       = 1
    cookie_persist_fail        = 1
    cookie_not_found           = 1
    cookie_invalid             = 1
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `trigger_stats_inc` (Block List, Max: 1) (see [below for nested schema](#nestedblock--trigger_stats_inc))
- `trigger_stats_rate` (Block List, Max: 1) (see [below for nested schema](#nestedblock--trigger_stats_rate))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--trigger_stats_inc"></a>
### Nested Schema for `trigger_stats_inc`

Optional:

- `cookie_invalid` (Number) Enable automatic packet-capture for Invalid persist cookie
- `cookie_not_found` (Number) Enable automatic packet-capture for Persist cookie not found
- `cookie_persist_fail` (Number) Enable automatic packet-capture for Cookie persist fail
- `cssl_sid_not_found` (Number) Enable automatic packet-capture for Client SSL SID not found
- `cssl_sid_not_match` (Number) Enable automatic packet-capture for Client SSL SID not match
- `dst_ip_fail` (Number) Enable automatic packet-capture for DST IP persist fail
- `dst_ip_hash_fail` (Number) Enable automatic packet-capture for DST IP hash persist fail
- `dst_ip_new_sess_cache_fail` (Number) Enable automatic packet-capture for DST IP new sess fail (c)
- `dst_ip_new_sess_sel_fail` (Number) Enable automatic packet-capture for DST IP new sess fail (s)
- `hash_tbl_create_fail` (Number) Enable automatic packet-capture for Hash tbl create fail
- `hash_tbl_rst_adddel` (Number) Enable automatic packet-capture for Hash tbl reset (add/del)
- `hash_tbl_rst_updown` (Number) Enable automatic packet-capture for Hash tbl reset (up/down)
- `hash_tbl_trylock_fail` (Number) Enable automatic packet-capture for Hash tbl lock fail
- `header_hash_fail` (Number) Enable automatic packet-capture for Header hash persist fail
- `src_ip_fail` (Number) Enable automatic packet-capture for SRC IP persist fail
- `src_ip_hash_fail` (Number) Enable automatic packet-capture for SRC IP hash persist fail
- `src_ip_new_sess_cache_fail` (Number) Enable automatic packet-capture for SRC IP new sess fail (c)
- `src_ip_new_sess_sel_fail` (Number) Enable automatic packet-capture for SRC IP new sess fail (s)
- `ssl_sid_persist_fail` (Number) Enable automatic packet-capture for SSL SID persist fail
- `ssl_sid_session_fail` (Number) Enable automatic packet-capture for Create SSL SID fail
- `sssl_sid_not_found` (Number) Enable automatic packet-capture for Server SSL SID not found
- `sssl_sid_not_match` (Number) Enable automatic packet-capture for Server SSL SID not match
- `url_hash_fail` (Number) Enable automatic packet-capture for URL hash persist fail
- `uuid` (String) uuid of the object


<a id="nestedblock--trigger_stats_rate"></a>
### Nested Schema for `trigger_stats_rate`

Optional:

- `cookie_invalid` (Number) Enable automatic packet-capture for Invalid persist cookie
- `cookie_not_found` (Number) Enable automatic packet-capture for Persist cookie not found
- `cookie_persist_fail` (Number) Enable automatic packet-capture for Cookie persist fail
- `cssl_sid_not_found` (Number) Enable automatic packet-capture for Client SSL SID not found
- `cssl_sid_not_match` (Number) Enable automatic packet-capture for Client SSL SID not match
- `dst_ip_fail` (Number) Enable automatic packet-capture for DST IP persist fail
- `dst_ip_hash_fail` (Number) Enable automatic packet-capture for DST IP hash persist fail
- `dst_ip_new_sess_cache_fail` (Number) Enable automatic packet-capture for DST IP new sess fail (c)
- `dst_ip_new_sess_sel_fail` (Number) Enable automatic packet-capture for DST IP new sess fail (s)
- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `hash_tbl_create_fail` (Number) Enable automatic packet-capture for Hash tbl create fail
- `hash_tbl_rst_adddel` (Number) Enable automatic packet-capture for Hash tbl reset (add/del)
- `hash_tbl_rst_updown` (Number) Enable automatic packet-capture for Hash tbl reset (up/down)
- `hash_tbl_trylock_fail` (Number) Enable automatic packet-capture for Hash tbl lock fail
- `header_hash_fail` (Number) Enable automatic packet-capture for Header hash persist fail
- `src_ip_fail` (Number) Enable automatic packet-capture for SRC IP persist fail
- `src_ip_hash_fail` (Number) Enable automatic packet-capture for SRC IP hash persist fail
- `src_ip_new_sess_cache_fail` (Number) Enable automatic packet-capture for SRC IP new sess fail (c)
- `src_ip_new_sess_sel_fail` (Number) Enable automatic packet-capture for SRC IP new sess fail (s)
- `ssl_sid_persist_fail` (Number) Enable automatic packet-capture for SSL SID persist fail
- `ssl_sid_session_fail` (Number) Enable automatic packet-capture for Create SSL SID fail
- `sssl_sid_not_found` (Number) Enable automatic packet-capture for Server SSL SID not found
- `sssl_sid_not_match` (Number) Enable automatic packet-capture for Server SSL SID not match
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `url_hash_fail` (Number) Enable automatic packet-capture for URL hash persist fail
- `uuid` (String) uuid of the object


