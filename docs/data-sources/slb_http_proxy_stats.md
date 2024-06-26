---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_http_proxy_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_http_proxy_stats: Statistics for the object http-proxy
  PLACEHOLDER
---

# thunder_slb_http_proxy_stats (Data Source)

`thunder_slb_http_proxy_stats`: Statistics for the object http-proxy

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_http_proxy_stats" "thunder_slb_http_proxy_stats" {
}
output "get_slb_http_proxy_stats" {
  value = ["${data.thunder_slb_http_proxy_stats.thunder_slb_http_proxy_stats}"]
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

- `cache_rsp` (Number) HTTP req (cache succ)
- `chunk_sz_1k` (Number) Chunk less than equal to 1K
- `chunk_sz_2k` (Number) Chunk less than equal to 2K
- `chunk_sz_4k` (Number) Chunk less than equal to 4K
- `chunk_sz_512` (Number) Chunk less than equal to 512
- `chunk_sz_gt_4k` (Number) Chunk greater than 4K
- `client_rst` (Number) Client RST
- `close_on_ddos` (Number) Close on DDoS
- `compression_after` (Number) Tot data after compress
- `compression_after_br` (Number) Tot data after brotli compress
- `compression_after_total` (Number) Tot data after compress
- `compression_before` (Number) Tot data before compress
- `compression_before_br` (Number) Tot data before brotli compress
- `compression_before_total` (Number) Tot data before compress
- `connect_req` (Number) Total HTTP CONNECT requests
- `curr_proxy` (Number) Curr Proxy Conns
- `decompression_after` (Number) Tot data after decompress
- `decompression_after_br` (Number) Tot data after brotli decompress
- `decompression_after_total` (Number) Tot data after decompress
- `decompression_before` (Number) Tot data before decompress
- `decompression_before_br` (Number) Tot data before brotli decompress
- `decompression_before_total` (Number) Tot data before decompress
- `doh_dns_malformed_query` (Number) DoH DNS Malformed Query
- `doh_dns_query_type_a` (Number) DoH Query type A
- `doh_dns_query_type_aaaa` (Number) DoH Query type AAAA
- `doh_dns_query_type_any` (Number) DoH Query type ANY
- `doh_dns_query_type_cname` (Number) DoH Query type CNAME
- `doh_dns_query_type_mx` (Number) DoH Query type MX
- `doh_dns_query_type_ns` (Number) DoH Query type NS
- `doh_dns_query_type_others` (Number) DoH Query type Others
- `doh_dns_query_type_soa` (Number) DoH Query type SOA
- `doh_dns_query_type_srv` (Number) DoH Query type SRV
- `doh_dns_resp_rcode_err_format` (Number) DoH DNS Response rcode ERR_FORMAT
- `doh_dns_resp_rcode_err_name` (Number) DoH DNS Response rcode ERR_NAME
- `doh_dns_resp_rcode_err_server` (Number) DoH DNS Response rcode ERR_SERVER
- `doh_dns_resp_rcode_err_type` (Number) DoH DNS Response rcode ERR_TYPE
- `doh_dns_resp_rcode_notauth` (Number) DoH DNS Response rcode NOTAUTH
- `doh_dns_resp_rcode_notzone` (Number) DoH DNS Response rcode NOTZONE
- `doh_dns_resp_rcode_nxrrset` (Number) DoH DNS Response rcode NXRRSET
- `doh_dns_resp_rcode_other` (Number) DoH DNS Response rcode OTHER
- `doh_dns_resp_rcode_refuse` (Number) DoH DNS Response rcode REFUSE
- `doh_dns_resp_rcode_yxdomain` (Number) DoH DNS Response rcode YXDOMAIN
- `doh_dns_resp_rcode_yxrrset` (Number) DoH DNS Response rcode YXRRSET
- `doh_get_base64_decode_failed` (Number) DoH GET base64url decode failed
- `doh_get_dns_arg_failed` (Number) DoH GET dns arg not found in uri
- `doh_get_uri_too_long` (Number) DoH GET URI too long
- `doh_malloc_fail` (Number) DoH Memory alloc failed
- `doh_non_doh_method` (Number) DoH Non DoH HTTP request method rcvd
- `doh_non_doh_req` (Number) DoH non DoH Requests
- `doh_non_doh_req_get` (Number) DoH non DoH GET Requests
- `doh_non_doh_req_post` (Number) DoH non DoH POST Requests
- `doh_path_not_found` (Number) DoH URI Path not found
- `doh_post_content_type_mismatch` (Number) DoH POST content-type not found
- `doh_post_payload_extract_failed` (Number) DoH POST payload extract failed
- `doh_post_payload_not_found` (Number) DoH POST payload not found
- `doh_post_payload_too_large` (Number) DoH POST Payload too large
- `doh_query_time_out` (Number) DoH serv Query timed out
- `doh_req` (Number) DoH Requests
- `doh_req_get` (Number) DoH GET Requests
- `doh_req_post` (Number) DoH POST Requests
- `doh_req_send_failed` (Number) DoH Request Send Failed
- `doh_req_tcp_retry` (Number) DoH TCP Retry
- `doh_req_tcp_retry_fail` (Number) DoH TCP Retry failed
- `doh_req_udp_retry` (Number) DoH UDP Retry
- `doh_req_udp_retry_fail` (Number) DoH UDP Retry failed
- `doh_resp` (Number) DoH Responses
- `doh_resp_header_alloc_failed` (Number) DoH Resp hdr alloc failed
- `doh_resp_que_failed` (Number) DoH Resp queue failed
- `doh_resp_send_failed` (Number) DoH Response Send Failed
- `doh_resp_setup_failed` (Number) DoH Response setup failed
- `doh_resp_tcp_frags` (Number) DoH TCP Frags Rcvd
- `doh_resp_udp_frags` (Number) DoH UDP Frags Rcvd
- `doh_retry_w_tcp` (Number) DoH Retry with TCP SG
- `doh_serv_sel_failed` (Number) DoH Server Select Failed
- `doh_snat_failed` (Number) DoH Source NAT failed
- `doh_tc_resp` (Number) DoH TC Responses
- `doh_tcp_dns_req` (Number) DoH TCP DNS Requests
- `doh_tcp_dns_resp` (Number) DoH TCP DNS Responses
- `doh_tcp_send_failed` (Number) DoH serv TCP DNS send failed
- `doh_udp_dns_req` (Number) DoH UDP DNS Requests
- `doh_udp_dns_resp` (Number) DoH UDP DNS Responses
- `doh_udp_send_failed` (Number) DoH serv UDP DNS send failed
- `fwdreq_fail` (Number) Fwd req fail
- `fwdreqdata_fail` (Number)
- `new_svrconn` (Number) Server conn made
- `noproxy` (Number) No proxy error
- `notuple` (Number) No tuple error
- `parsereq_fail` (Number) Parse req fail
- `req` (Number) HTTP requests
- `req_100m` (Number) Rsp time less than 100m
- `req_100u` (Number) Rsp time less than 100u
- `req_10m` (Number) Rsp time less than 10m
- `req_10u` (Number) Rsp time less than 10u
- `req_1m` (Number) Rsp time less than 1m
- `req_1s` (Number) Rsp time less than 1s
- `req_200m` (Number) Rsp time less than 200m
- `req_200u` (Number) Rsp time less than 200u
- `req_20m` (Number) Rsp time less than 20m
- `req_20u` (Number) Rsp time less than 20u
- `req_2m` (Number) Rsp time less than 2m
- `req_2s` (Number) Rsp time less than 2s
- `req_500m` (Number) Rsp time less than 500m
- `req_500u` (Number) Rsp time less than 500u
- `req_50m` (Number) Rsp time less than 50m
- `req_50u` (Number) Rsp time less than 50u
- `req_5m` (Number) Rsp time less than 5m
- `req_5s` (Number) Rsp time less than 5s
- `req_connect` (Number) Method CONNECT
- `req_content_len` (Number) Req content len
- `req_delete` (Number) Method DELETE
- `req_enter_ssli` (Number) Total HTTP requests enter SSLi
- `req_get` (Number) Method GET
- `req_head` (Number) Method HEAD
- `req_http2` (Number) Request 2.0
- `req_http3` (Number) Request 3.0
- `req_ofo` (Number) Packets ofo
- `req_options` (Number) Method OPTIONS
- `req_over_5s` (Number) Rsp time greater than equal to 5s
- `req_over_limit` (Number) Request over limit
- `req_post` (Number) Method POST
- `req_put` (Number) Method PUT
- `req_rate_over_limit` (Number) Request rate over limit
- `req_retran` (Number) Packets retrans
- `req_succ` (Number) HTTP requests(succ)
- `req_sz_16k` (Number) Req less than equal to 16K
- `req_sz_1k` (Number) Req less than equal to 1K
- `req_sz_256k` (Number) Req less than equal to 256K
- `req_sz_2k` (Number) Req less than equal to 2K
- `req_sz_32k` (Number) Req less than equal to 32K
- `req_sz_4k` (Number) Req less than equal to 4K
- `req_sz_64k` (Number) Req less than equal to 64K
- `req_sz_8k` (Number) Req less than equal to 8K
- `req_sz_gt_256k` (Number) Req greater than 256K
- `req_trace` (Number) Method TRACE
- `req_track` (Number) Method TRACK
- `req_unknown` (Number) Method UNKNOWN
- `response_100` (Number) Status code 100
- `response_101` (Number) Status code 101
- `response_102` (Number) Status code 102
- `response_1xx` (Number) Status code 1XX
- `response_200` (Number) Status code 200
- `response_201` (Number) Status code 201
- `response_202` (Number) Status code 202
- `response_203` (Number) Status code 203
- `response_204` (Number) Status code 204
- `response_205` (Number) Status code 205
- `response_206` (Number) Status code 206
- `response_207` (Number) Status code 207
- `response_2xx` (Number) Status code 2XX
- `response_300` (Number) Status code 300
- `response_301` (Number) Status code 301
- `response_302` (Number) Status code 302
- `response_303` (Number) Status code 303
- `response_304` (Number) Status code 304
- `response_305` (Number) Status code 305
- `response_306` (Number) Status code 306
- `response_307` (Number) Status code 307
- `response_3xx` (Number) Status code 3XX
- `response_400` (Number) Status code 400
- `response_401` (Number) Status code 401
- `response_402` (Number) Status code 402
- `response_403` (Number) Status code 403
- `response_404` (Number) Status code 404
- `response_405` (Number) Status code 405
- `response_406` (Number) Status code 406
- `response_407` (Number) Status code 407
- `response_408` (Number) Status code 408
- `response_409` (Number) Status code 409
- `response_410` (Number) Status code 410
- `response_411` (Number) Status code 411
- `response_412` (Number) Status code 412
- `response_413` (Number) Status code 413
- `response_414` (Number) Status code 414
- `response_415` (Number) Status code 415
- `response_416` (Number) Status code 416
- `response_417` (Number) Status code 417
- `response_418` (Number) Status code 418
- `response_422` (Number) Status code 422
- `response_423` (Number) Status code 423
- `response_424` (Number) Status code 424
- `response_425` (Number) Status code 425
- `response_426` (Number) Status code 426
- `response_449` (Number) Status code 449
- `response_450` (Number) Status code 450
- `response_4xx` (Number) Status code 4XX
- `response_500` (Number) Status code 500
- `response_501` (Number) Status code 501
- `response_502` (Number) Status code 502
- `response_503` (Number) Status code 503
- `response_504` (Number) Status code 504
- `response_505` (Number) Status code 505
- `response_506` (Number) Status code 506
- `response_507` (Number) Status code 507
- `response_508` (Number) Status code 508
- `response_509` (Number) Status code 509
- `response_510` (Number) Status code 510
- `response_5xx` (Number) Status code 5XX
- `response_6xx` (Number) Status code 6XX
- `response_http2` (Number) Resp 2.0
- `response_http3` (Number) Resp 3.0
- `response_unknown` (Number) Status code unknown
- `rsp_chunk` (Number) Resp chunk encoding
- `rsp_content_len` (Number) Resp content len
- `rsp_sz_16k` (Number) Resp less than equal to 16K
- `rsp_sz_1k` (Number) Resp less than equal to 1K
- `rsp_sz_256k` (Number) Resp less than equal to 256K
- `rsp_sz_2k` (Number) Resp less than equal to 2K
- `rsp_sz_32k` (Number) Resp less than equal to 32K
- `rsp_sz_4k` (Number) Resp less than equal to 4K
- `rsp_sz_64k` (Number) Resp less than equal to 64K
- `rsp_sz_8k` (Number) Resp less than equal to 8K
- `rsp_sz_gt_256k` (Number) Resp greater than 256K
- `server_resel` (Number) Server reselection
- `server_rst` (Number) Server RST
- `snat_fail` (Number) Source NAT failure
- `svr_prem_close` (Number) Server premature close
- `svrsel_fail` (Number) Server selection fail
- `total_proxy` (Number) Total Proxy Conns


