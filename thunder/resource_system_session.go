package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSession() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_session`: Session Entries\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSessionCreate,
		UpdateContext: resourceSystemSessionUpdate,
		ReadContext:   resourceSystemSessionRead,
		DeleteContext: resourceSystemSessionDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_l4_conn': Total L4 Count; 'conn_counter': Conn Count; 'conn_freed_counter': Conn Freed; 'total_l4_packet_count': Total L4 Packet Count; 'total_l7_packet_count': Total L7 Packet Count; 'total_l4_conn_proxy': Total L4 Conn Proxy Count; 'total_l7_conn': Total L7 Conn; 'total_tcp_conn': Total TCP Conn; 'curr_free_conn': Curr Free Conn; 'tcp_est_counter': TCP Established; 'tcp_half_open_counter': TCP Half Open; 'tcp_half_close_counter': TCP Half Closed; 'udp_counter': UDP Count; 'ip_counter': IP Count; 'other_counter': Non TCP/UDP IP sessions; 'reverse_nat_tcp_counter': Reverse NAT TCP; 'reverse_nat_udp_counter': Reverse NAT UDP; 'tcp_syn_half_open_counter': TCP SYN Half Open; 'conn_smp_alloc_counter': Conn SMP Alloc; 'conn_smp_free_counter': Conn SMP Free; 'conn_smp_aged_counter': Conn SMP Aged; 'ssl_count_curr': Curr SSL Count; 'ssl_count_total': Total SSL Count; 'server_ssl_count_curr': Current SSL Server Count; 'server_ssl_count_total': Total SSL Server Count; 'client_ssl_reuse_total': Total SSL Client Reuse; 'server_ssl_reuse_total': Total SSL Server Reuse; 'ssl_failed_total': Total SSL Failures; 'ssl_failed_ca_verification': SSL Cert Auth Verification Errors; 'ssl_server_cert_error': SSL Server Cert Errors; 'ssl_client_cert_auth_fail': SSL Client Cert Auth Failures; 'total_ip_nat_conn': Total IP Nat Conn; 'total_l2l3_conn': Totl L2/L3 Connections; 'client_ssl_ctx_malloc_failure': Client SSL Ctx malloc Failures; 'conn_type_0_available': Conn Type 0 Available; 'conn_type_1_available': Conn Type 1 Available; 'conn_type_2_available': Conn Type 2 Available; 'conn_type_3_available': Conn Type 3 Available; 'conn_type_4_available': Conn Type 4 Available; 'conn_smp_type_0_available': Conn SMP Type 0 Available; 'conn_smp_type_1_available': Conn SMP Type 1 Available; 'conn_smp_type_2_available': Conn SMP Type 2 Available; 'conn_smp_type_3_available': Conn SMP Type 3 Available; 'conn_smp_type_4_available': Conn SMP Type 4 Available; 'sctp-half-open-counter': SCTP Half Open; 'sctp-est-counter': SCTP Established; 'nonssl_bypass': NON SSL Bypass Count; 'ssl_failsafe_total': Total SSL Failsafe Count; 'ssl_forward_proxy_failed_handshake_total': Total SSL Forward Proxy Failed Handshake Count; 'ssl_forward_proxy_failed_tcp_total': Total SSL Forward Proxy Failed TCP Count; 'ssl_forward_proxy_failed_crypto_total': Total SSL Forward Proxy Failed Crypto Count; 'ssl_forward_proxy_failed_cert_verify_total': Total SSL Forward Proxy Failed Certificate Verification Count; 'ssl_forward_proxy_invalid_ocsp_stapling_total': Total SSL Forward Proxy Invalid OCSP Stapling Count; 'ssl_forward_proxy_revoked_ocsp_total': Total SSL Forward Proxy Revoked OCSP Response Count; 'ssl_forward_proxy_failed_cert_signing_total': Total SSL Forward Proxy Failed Certificate Signing Count; 'ssl_forward_proxy_failed_ssl_version_total': Total SSL Forward Proxy Unsupported version Count; 'ssl_forward_proxy_sni_bypass_total': Total SSL Forward Proxy SNI Bypass Count; 'ssl_forward_proxy_client_auth_bypass_total': Total SSL Forward Proxy Client Auth Bypass Count; 'conn_app_smp_alloc_counter': Conn APP SMP Alloc; 'diameter_conn_counter': Diameter Conn Count; 'diameter_conn_freed_counter': Diameter Conn Freed; 'debug_tcp_counter': Hidden TCP sessions for CGNv6 Stateless Technologies; 'debug_udp_counter': Hidden UDP sessions for CGNv6 Stateless Technologies; 'total_fw_conn': Total Firewall Conn; 'total_local_conn': Total Local Conn; 'total_curr_conn': Total Curr Conn; 'client_ssl_fatal_alert': client ssl fatal alert; 'client_ssl_fin_rst': client ssl fin rst; 'fp_session_fin_rst': FP Session FIN/RST; 'server_ssl_fatal_alert': server ssl fatal alert; 'server_ssl_fin_rst': server ssl fin rst; 'client_template_int_err': client template internal error; 'client_template_unknown_err': client template unknown error; 'server_template_int_err': server template int error; 'server_template_unknown_err': server template unknown error; 'total_debug_conn': Total Debug Conn; 'ssl_forward_proxy_failed_aflex_total': Total SSL Forward Proxy AFLEX Count; 'ssl_forward_proxy_cert_subject_bypass_total': Total SSL Forward Proxy Certificate Subject Bypass Count; 'ssl_forward_proxy_cert_issuer_bypass_total': Total SSL Forward Proxy Certificate Issuer Bypass Count; 'ssl_forward_proxy_cert_san_bypass_total': Total SSL Forward Proxy Certificate SAN Bypass Count; 'ssl_forward_proxy_no_sni_bypass_total': Total SSL Forward Proxy No SNI Bypass Count; 'ssl_forward_proxy_no_sni_reset_total': Total SSL Forward Proxy No SNI reset Count; 'ssl_forward_proxy_username_bypass_total': Total SSL Forward Proxy Username Bypass Count; 'ssl_forward_proxy_ad_grpup_bypass_total': Total SSL Forward Proxy AD-Group Bypass Count; 'diameter_concurrent_user_sessions_counter': Diameter Concurrent User-Sessions; 'client_ssl_session_ticket_reuse_total': Total SSL Client Session Ticket Reuse; 'server_ssl_session_ticket_reuse_total': Total SSL Server Session Ticket Reuse; 'total_clientside_early_data_connections': Total clientside early data connections; 'total_serverside_early_data_connections': Total serverside early data connections; 'total_clientside_failed_early_data-connections': Total clientside failed early data connections; 'total_serverside_failed_early_data-connections': Total serverside failed early data connections; 'ssl_forward_proxy_esni_bypass_total': Total SSL Forward Proxy ESNI Bypass Count; 'ssl_forward_proxy_esni_reset_total': Total SSL Forward Proxy ESNI Drop Count; 'total_logging_conn': Total Logging Conn; 'gtp_c_est_counter': GTP-C Established; 'gtp_c_half_open_counter': GTP-C Half Open; 'gtp_u_counter': GTP-U Count; 'gtp_c_echo_counter': GTP-C Echo Count; 'gtp_u_echo_counter': GTP-U Echo Count; 'gtp_curr_free_conn': GTP Current Available Conn; 'gtp_cum_conn_counter': GTP cumulative Conn Count; 'gtp_cum_conn_freed_counter': GTP cumulative Conn Freed; 'fw_blacklist_sess': Blacklist Sessions; 'fw_blacklist_sess_created': Blacklist Session Created; 'fw_blacklist_sess_freed': Blacklist Session Freed; 'server_tcp_est_counter': Server TCP Established; 'server_tcp_half_open_counter': Server TCP Half Open; 'sched_conn_with_wrong_next_idx_to_rml': Attempt to Put a Conn to RML Whose next_idx is NOT Invalid; 'free_conn_not_in_sp': Attempt to Free a Conn Whoes Address Not in Session Pool;",
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
func resourceSystemSessionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSession(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSessionRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSessionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSession(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSessionRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSessionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSession(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSessionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSession(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemSessionSamplingEnable(d []interface{}) []edpt.SystemSessionSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemSessionSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemSessionSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemSession(d *schema.ResourceData) edpt.SystemSession {
	var ret edpt.SystemSession
	ret.Inst.SamplingEnable = getSliceSystemSessionSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
