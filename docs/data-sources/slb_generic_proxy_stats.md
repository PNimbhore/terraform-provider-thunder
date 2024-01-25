---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_generic_proxy_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_generic_proxy_stats: Statistics for the object generic-proxy
  PLACEHOLDER
---

# thunder_slb_generic_proxy_stats (Data Source)

`thunder_slb_generic_proxy_stats`: Statistics for the object generic-proxy

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_generic_proxy_stats" "thunder_slb_generic_proxy_stats" {
}
output "get_slb_generic_proxy_stats" {
  value = ["${data.thunder_slb_generic_proxy_stats.thunder_slb_generic_proxy_stats}"]
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

- `aca_in` (Number) Number of ACAs in
- `aca_out` (Number) Number of ACAs out
- `acr_in` (Number) Number of ACRs in
- `acr_out` (Number) Number of ACRs out
- `asa_in` (Number) Number of ASAs in
- `asa_out` (Number) Number of ASAs out
- `asr_in` (Number) Number of ASRs in
- `asr_out` (Number) Number of ASRs out
- `cca_in` (Number) Number of CCAs in
- `cca_out` (Number) Number of CCAs out
- `cca_t` (Number) Number of CCAs terminate
- `ccr_i` (Number) Number of CCRs initial
- `ccr_in` (Number) Number of CCRs in
- `ccr_out` (Number) Number of CCRs out
- `ccr_t` (Number) Number of CCRs terminate
- `ccr_u` (Number) Number of CCRs update
- `cea_in` (Number) Number of CEAs in
- `cea_out` (Number) Number of CEAs out
- `cer_in` (Number) Number of CERs in
- `cer_out` (Number) Number of CERs out
- `client_fail` (Number) Number of client failures
- `client_select_fail` (Number) Fail to select client
- `close_conn_when_vport_down` (Number) Close client conn when virtual port is down
- `conn_closed_by_client` (Number) Client initiates TCP close/reset
- `conn_closed_by_server` (Number) Server initiates TCP close/reset
- `curr` (Number) Current
- `dcmsg_error` (Number) Diameter cross cpu error
- `dcmsg_fwd_in` (Number) Diameter cross cpu fwd in
- `dcmsg_fwd_out` (Number) Diameter cross cpu fwd out
- `dcmsg_rev_in` (Number) Diameter cross cpu rev in
- `dcmsg_rev_out` (Number) Diameter cross cpu rev out
- `dpa_in` (Number) Number of DPAs in
- `dpa_out` (Number) Number of DPAs out
- `dpr_in` (Number) Number of DPRs in
- `dpr_out` (Number) Number of DPRs out
- `dwa_in` (Number) Number of DWAs in
- `dwa_out` (Number) Number of DWAs out
- `dwr_in` (Number) Number of DWRs in
- `dwr_out` (Number) Number of DWRs out
- `forward_unknown_session_id` (Number) Forward server side message with unknown session id
- `invalid_avp` (Number) AVP value contains illegal chars
- `mismatch_fwd_id` (Number) Diameter mismatch fwd session id
- `mismatch_rev_id` (Number) Diameter mismatch rev session id
- `no_fwd_tuple` (Number) Diameter no fwd tuple matched
- `no_rev_tuple` (Number) Diameter no rev tuple matched
- `no_route` (Number) Number of no routes
- `no_sess` (Number) Number of no sessions
- `no_session_id` (Number) Diameter no session id avp
- `num` (Number) Number
- `other_in` (Number) Number of other messages in
- `other_out` (Number) Number of other messages out
- `reply_error_info_fail` (Number) Fail to reply error info to peer
- `reply_invalid_avp_value` (Number) Reply with invalid AVP error info
- `reply_unable_to_deliver` (Number) Reply with unable to deliver error info
- `reply_unknown_session_id` (Number) Reply with unknown session ID error info
- `reselect_fwd_tuple` (Number) Original client tuple does not exist so reselect another one
- `reselect_fwd_tuple_other_cpu` (Number) Original client tuple does not exist so reselect another one on other CPUs
- `reselect_rev_tuple` (Number) Original server tuple does not exist so reselect another one
- `retry_client_request` (Number) Diameter retry client request
- `retry_client_request_fail` (Number) Diameter retry client request fail
- `server_fail` (Number) Number of server failures
- `snat_fail` (Number) Number of snat failures
- `sta_in` (Number) Number of STAs in
- `sta_out` (Number) Number of STAs out
- `str_in` (Number) Number of STRs in
- `str_out` (Number) Number of STRs out
- `svrsel_fail` (Number) Number of server selection failed
- `terminate_on_cca_t` (Number) Diameter terminate on cca_t
- `total` (Number) Total
- `total_http_req_enter_gen` (Number) Total number of HTTP requests enter generic proxy
- `unkwn_cmd_code` (Number) Diameter unkown cmd code
- `update_latest_server` (Number) Update to the latest server that used a session id
- `user_session` (Number) Number of user sessions

