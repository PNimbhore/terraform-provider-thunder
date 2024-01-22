package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHttpProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_http_proxy`: Configure HTTP Proxy global\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbHttpProxyCreate,
		UpdateContext: resourceSlbHttpProxyUpdate,
		ReadContext:   resourceSlbHttpProxyRead,
		DeleteContext: resourceSlbHttpProxyDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num': Num; 'curr_proxy': Curr Proxy Conns; 'total_proxy': Total Proxy Conns; 'req': HTTP requests; 'req_succ': HTTP requests(succ); 'noproxy': No proxy error; 'client_rst': Client RST; 'server_rst': Server RST; 'notuple': No tuple error; 'parsereq_fail': Parse req fail; 'svrsel_fail': Server selection fail; 'fwdreq_fail': Fwd req fail; 'fwdreq_fail_buff': Fwd req fail - buff; 'fwdreq_fail_rport': Fwd req fail - rport; 'fwdreq_fail_route': Fwd req fail - route; 'fwdreq_fail_persist': Fwd req fail - persist; 'fwdreq_fail_server': Fwd req fail - server; 'fwdreq_fail_tuple': Fwd req fail - tuple; 'fwdreqdata_fail': fwdreqdata_fail; 'req_retran': Packets retrans; 'req_ofo': Packets ofo; 'server_resel': Server reselection; 'svr_prem_close': Server premature close; 'new_svrconn': Server conn made; 'snat_fail': Source NAT failure; 'tcpoutrst': Out RSTs; 'full_proxy': Full proxy tot; 'full_proxy_post': Full proxy POST; 'full_proxy_pipeline': Full proxy pipeline; 'full_proxy_fpga_err': Full proxy fpga err; 'req_over_limit': Request over limit; 'req_rate_over_limit': Request rate over limit; 'l4_switching': L4 switching; 'cookie_switching': Cookie switching; 'aflex_switching': aFleX switching; 'http_policy_switching': HTTP Policy switching; 'url_switching': URL switching; 'host_switching': Host switching; 'lb_switching': Normal LB switching; 'l4_switching_ok': L4 switching (succ); 'cookie_switching_ok': Cookie switching (succ); 'aflex_switching_ok': aFleX switching (succ); 'http_policy_switching_ok': HTTP Policy switching (succ); 'url_switching_ok': URL switching (succ); 'host_switching_ok': Host switching (succ); 'lb_switching_ok': Normal LB switch. (succ); 'l4_switching_enqueue': L4 switching (enQ); 'cookie_switching_enqueue': Cookie switching (enQ); 'aflex_switching_enqueue': aFleX switching (enQ); 'http_policy_switching_enqueue': HTTP Policy switching (enQ); 'url_switching_enqueue': URL switching (enQ); 'host_switching_enqueue': Host switching (enQ); 'lb_switching_enqueue': Normal LB switch. (enQ); 'retry_503': Retry on 503; 'aflex_retry': aFleX http retry; 'aflex_lb_reselect': aFleX lb reselect; 'aflex_lb_reselect_ok': aFleX lb reselect (succ); 'client_rst_request': Client RST - request; 'client_rst_connecting': Client RST - connecting; 'client_rst_connected': Client RST - connected; 'client_rst_response': Client RST - response; 'server_rst_request': Server RST - request; 'server_rst_connecting': Server RST - connecting; 'server_rst_connected': Server RST - connected; 'server_rst_response': Server RST - response; 'invalid_header': Invalid header; 'too_many_headers': Too many headers; 'line_too_long': Line too long; 'header_name_too_long': Header name too long; 'wrong_resp_header': Wrong response header; 'header_insert': Header insert; 'header_delete': Header delete; 'insert_client_ip': Insert client IP; 'negative_req_remain': Negative request remain; 'negative_resp_remain': Negative response remain; 'large_cookie': Large cookies; 'large_cookie_header': Large cookie headers; 'huge_cookie': Huge cookies; 'huge_cookie_header': Huge cookie headers; 'parse_cookie_fail': Parse cookie fail; 'parse_setcookie_fail': Parse set-cookie fail; 'asm_cookie_fail': Assemble cookie fail; 'asm_cookie_header_fail': Asm cookie header fail; 'asm_setcookie_fail': Assemble set-cookie fail; 'asm_setcookie_header_fail': Asm set-cookie hdr fail; 'client_req_unexp_flag': Client req unexp flags; 'connecting_fin': Connecting FIN; 'connecting_fin_retrans': Connecting FIN retran; 'connecting_fin_ofo': Connecting FIN ofo; 'connecting_rst': Connecting RST; 'connecting_rst_retrans': Connecting RST retran; 'connecting_rst_ofo': Connecting RST ofo; 'connecting_ack': Connecting ACK; 'pkts_ofo': Packets ofo; 'pkts_retrans': Packets retrans; 'pkts_retrans_ack_finwait': retrans ACK FWAIT; 'pkts_retrans_fin': retrans FIN; 'pkts_retrans_rst': retrans RST; 'pkts_retrans_push': retrans PSH; 'stale_sess': Stale sess; 'server_resel_failed': Server re-select failed; 'compression_before': Tot data before compress; 'compression_after': Tot data after compress; 'response_1xx': Status code 1XX; 'response_100': Status code 100; 'response_101': Status code 101; 'response_102': Status code 102; 'response_2xx': Status code 2XX; 'response_200': Status code 200; 'response_201': Status code 201; 'response_202': Status code 202; 'response_203': Status code 203; 'response_204': Status code 204; 'response_205': Status code 205; 'response_206': Status code 206; 'response_207': Status code 207; 'response_3xx': Status code 3XX; 'response_300': Status code 300; 'response_301': Status code 301; 'response_302': Status code 302; 'response_303': Status code 303; 'response_304': Status code 304; 'response_305': Status code 305; 'response_306': Status code 306; 'response_307': Status code 307; 'response_4xx': Status code 4XX; 'response_400': Status code 400; 'response_401': Status code 401; 'response_402': Status code 402; 'response_403': Status code 403; 'response_404': Status code 404; 'response_405': Status code 405; 'response_406': Status code 406; 'response_407': Status code 407; 'response_408': Status code 408; 'response_409': Status code 409; 'response_410': Status code 410; 'response_411': Status code 411; 'response_412': Status code 412; 'response_413': Status code 413; 'response_414': Status code 414; 'response_415': Status code 415; 'response_416': Status code 416; 'response_417': Status code 417; 'response_418': Status code 418; 'response_422': Status code 422; 'response_423': Status code 423; 'response_424': Status code 424; 'response_425': Status code 425; 'response_426': Status code 426; 'response_449': Status code 449; 'response_450': Status code 450; 'response_5xx': Status code 5XX; 'response_500': Status code 500; 'response_501': Status code 501; 'response_502': Status code 502; 'response_503': Status code 503; 'response_504': Status code 504; 'response_505': Status code 505; 'response_506': Status code 506; 'response_507': Status code 507; 'response_508': Status code 508; 'response_509': Status code 509; 'response_510': Status code 510; 'response_6xx': Status code 6XX; 'response_unknown': Status code unknown; 'req_http10': Request 1.0; 'req_http11': Request 1.1; 'response_http10': Resp 1.0; 'response_http11': Resp 1.1; 'req_get': Method GET; 'req_head': Method HEAD; 'req_put': Method PUT; 'req_post': Method POST; 'req_trace': Method TRACE; 'req_options': Method OPTIONS; 'req_connect': Method CONNECT; 'req_delete': Method DELETE; 'req_unknown': Method UNKNOWN; 'req_content_len': Req content len; 'rsp_content_len': Resp content len; 'rsp_chunk': Resp chunk encoding; 'req_chunk': Req chunk encoding; 'compress_rsp': Compress req; 'compress_del_accept_enc': Compress del accept enc; 'compress_resp_already_compressed': Resp already compressed; 'compress_content_type_excluded': Compress cont type excl; 'compress_no_content_type': Compress no cont type; 'compress_resp_lt_min': Compress resp less than min; 'compress_resp_no_cl_or_ce': Compress resp no CL/CE; 'compress_ratio_too_high': Compress ratio too high; 'cache_rsp': HTTP req (cache succ); 'close_on_ddos': Close on DDoS; 'req_http10_keepalive': 1.0 Keepalive; 'req_sz_1k': Req less than equal to 1K; 'req_sz_2k': Req less than equal to 2K; 'req_sz_4k': Req less than equal to 4K;",
						},
						"counters2": {
							Type: schema.TypeString, Optional: true, Description: "'req_sz_8k': Req less than equal to 8K; 'req_sz_16k': Req less than equal to 16K; 'req_sz_32k': Req less than equal to 32K; 'req_sz_64k': Req less than equal to 64K; 'req_sz_256k': Req less than equal to 256K; 'req_sz_gt_256k': Req greater than 256K; 'rsp_sz_1k': Resp less than equal to 1K; 'rsp_sz_2k': Resp less than equal to 2K; 'rsp_sz_4k': Resp less than equal to 4K; 'rsp_sz_8k': Resp less than equal to 8K; 'rsp_sz_16k': Resp less than equal to 16K; 'rsp_sz_32k': Resp less than equal to 32K; 'rsp_sz_64k': Resp less than equal to 64K; 'rsp_sz_256k': Resp less than equal to 256K; 'rsp_sz_gt_256k': Resp greater than 256K; 'chunk_sz_512': Chunk less than equal to 512; 'chunk_sz_1k': Chunk less than equal to 1K; 'chunk_sz_2k': Chunk less than equal to 2K; 'chunk_sz_4k': Chunk less than equal to 4K; 'chunk_sz_gt_4k': Chunk greater than 4K; 'pconn_connecting': pconn connecting; 'pconn_connected': pconn connected; 'pconn_connecting_failed': pconn conn failed; 'chunk_bad': Bad Chunk; 'req_10u': Rsp time less than 10u; 'req_20u': Rsp time less than 20u; 'req_50u': Rsp time less than 50u; 'req_100u': Rsp time less than 100u; 'req_200u': Rsp time less than 200u; 'req_500u': Rsp time less than 500u; 'req_1m': Rsp time less than 1m; 'req_2m': Rsp time less than 2m; 'req_5m': Rsp time less than 5m; 'req_10m': Rsp time less than 10m; 'req_20m': Rsp time less than 20m; 'req_50m': Rsp time less than 50m; 'req_100m': Rsp time less than 100m; 'req_200m': Rsp time less than 200m; 'req_500m': Rsp time less than 500m; 'req_1s': Rsp time less than 1s; 'req_2s': Rsp time less than 2s; 'req_5s': Rsp time less than 5s; 'req_over_5s': Rsp time greater than equal to 5s; 'insert_client_port': Insert client Port; 'req_track': Method TRACK; 'connect_req': Total HTTP CONNECT requests; 'req_enter_ssli': Total HTTP requests enter SSLi; 'non_http_bypass': Non-HTTP bypass; 'decompression_before': Tot data before decompress; 'decompression_after': Tot data after decompress; 'req_http2': Request 2.0; 'response_http2': Resp 2.0; 'req_timeout_retry': Retry on Req Timeout; 'req_timeout_close': Close on Req Timeout; 'doh_req': DoH Requests; 'doh_req_get': DoH GET Requests; 'doh_req_post': DoH POST Requests; 'doh_non_doh_req': DoH non DoH Requests; 'doh_non_doh_req_get': DoH non DoH GET Requests; 'doh_non_doh_req_post': DoH non DoH POST Requests; 'doh_resp': DoH Responses; 'doh_tc_resp': DoH TC Responses; 'doh_udp_dns_req': DoH UDP DNS Requests; 'doh_udp_dns_resp': DoH UDP DNS Responses; 'doh_tcp_dns_req': DoH TCP DNS Requests; 'doh_tcp_dns_resp': DoH TCP DNS Responses; 'doh_req_send_failed': DoH Request Send Failed; 'doh_resp_send_failed': DoH Response Send Failed; 'doh_malloc_fail': DoH Memory alloc failed; 'doh_req_udp_retry': DoH UDP Retry; 'doh_req_udp_retry_fail': DoH UDP Retry failed; 'doh_req_tcp_retry': DoH TCP Retry; 'doh_req_tcp_retry_fail': DoH TCP Retry failed; 'doh_snat_failed': DoH Source NAT failed; 'doh_path_not_found': DoH URI Path not found; 'doh_get_dns_arg_failed': DoH GET dns arg not found in uri; 'doh_get_base64_decode_failed': DoH GET base64url decode failed; 'doh_post_content_type_mismatch': DoH POST content-type not found; 'doh_post_payload_not_found': DoH POST payload not found; 'doh_post_payload_extract_failed': DoH POST payload extract failed; 'doh_non_doh_method': DoH Non DoH HTTP request method rcvd; 'doh_tcp_send_failed': DoH serv TCP DNS send failed; 'doh_udp_send_failed': DoH serv UDP DNS send failed; 'doh_query_time_out': DoH serv Query timed out; 'doh_dns_query_type_a': DoH Query type A; 'doh_dns_query_type_aaaa': DoH Query type AAAA; 'doh_dns_query_type_ns': DoH Query type NS; 'doh_dns_query_type_cname': DoH Query type CNAME; 'doh_dns_query_type_any': DoH Query type ANY; 'doh_dns_query_type_srv': DoH Query type SRV; 'doh_dns_query_type_mx': DoH Query type MX; 'doh_dns_query_type_soa': DoH Query type SOA; 'doh_dns_query_type_others': DoH Query type Others; 'doh_resp_setup_failed': DoH Response setup failed; 'doh_resp_header_alloc_failed': DoH Resp hdr alloc failed; 'doh_resp_que_failed': DoH Resp queue failed; 'doh_resp_udp_frags': DoH UDP Frags Rcvd; 'doh_resp_tcp_frags': DoH TCP Frags Rcvd; 'doh_serv_sel_failed': DoH Server Select Failed; 'doh_retry_w_tcp': DoH Retry with TCP SG; 'doh_get_uri_too_long': DoH GET URI too long; 'doh_post_payload_too_large': DoH POST Payload too large; 'doh_dns_malformed_query': DoH DNS Malformed Query; 'doh_dns_resp_rcode_err_format': DoH DNS Response rcode ERR_FORMAT; 'doh_dns_resp_rcode_err_server': DoH DNS Response rcode ERR_SERVER; 'doh_dns_resp_rcode_err_name': DoH DNS Response rcode ERR_NAME; 'doh_dns_resp_rcode_err_type': DoH DNS Response rcode ERR_TYPE; 'doh_dns_resp_rcode_refuse': DoH DNS Response rcode REFUSE; 'doh_dns_resp_rcode_yxdomain': DoH DNS Response rcode YXDOMAIN; 'doh_dns_resp_rcode_yxrrset': DoH DNS Response rcode YXRRSET; 'doh_dns_resp_rcode_nxrrset': DoH DNS Response rcode NXRRSET; 'doh_dns_resp_rcode_notauth': DoH DNS Response rcode NOTAUTH; 'doh_dns_resp_rcode_notzone': DoH DNS Response rcode NOTZONE; 'doh_dns_resp_rcode_other': DoH DNS Response rcode OTHER; 'compression_before_br': Tot data before brotli compress; 'compression_after_br': Tot data after brotli compress; 'compression_before_total': Tot data before compress; 'compression_after_total': Tot data after compress; 'decompression_before_br': Tot data before brotli decompress; 'decompression_after_br': Tot data after brotli decompress; 'decompression_before_total': Tot data before decompress; 'decompression_after_total': Tot data after decompress; 'compress_rsp_br': Compress req with brotli; 'compress_rsp_total': Compress req; 'h2up_content_length_alias': HTTP2 content length alias; 'malformed_h2up_header_value': Malformed HTTP2 header value; 'malformed_h2up_scheme_value': Malformed HTTP2 scheme value; 'h2up_with_transfer_encoding': HTTP2 with transfer-encoding header; 'multiple_content_length': Multiple content-length headers; 'multiple_transfer_encoding': Multiple transfer-encoding headers; 'transfer_encoding_and_content_length': Transfer-encoding header with Content-Length header; 'get_and_payload': GET method with content-length header or transfer-encoding header; 'h2up_with_host_and_auth': HTTP2 with host header and authority header; 'req_http3': Request 3.0; 'response_http3': Resp 3.0; 'header_filter_rule_hit': Hit header filter rule; 'http1_client_idle_timeout': HTTP1 client idle timeout; 'http2_client_idle_timeout': HTTP2 client idle timeout; 'http_disallowed_methods': HTTP disallowed methods; 'http_allowed_methods': HTTP allowed methods; 'req_http11_new_proxy': Request 1.1 (new proxy);",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbHttpProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttpProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttpProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHttpProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbHttpProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttpProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttpProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHttpProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbHttpProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttpProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttpProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbHttpProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttpProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttpProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbHttpProxySamplingEnable(d []interface{}) []edpt.SlbHttpProxySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbHttpProxySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHttpProxySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHttpProxy(d *schema.ResourceData) edpt.SlbHttpProxy {
	var ret edpt.SlbHttpProxy
	ret.Inst.SamplingEnable = getSliceSlbHttpProxySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
