package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_sip`: Configure SIP\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSipCreate,
		UpdateContext: resourceSlbSipUpdate,
		ReadContext:   resourceSlbSipRead,
		DeleteContext: resourceSlbSipDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'msg_proxy_current': Number of current sip proxy connections; 'msg_proxy_total': Total number of sip proxy connections; 'msg_proxy_mem_allocd': msg_proxy_mem_allocd; 'msg_proxy_mem_cached': msg_proxy_mem_cached; 'msg_proxy_mem_freed': msg_proxy_mem_freed; 'msg_proxy_client_recv': Number of SIP messages received from client; 'msg_proxy_client_send_success': Number of SIP messages received from client and forwarded to server; 'msg_proxy_client_incomplete': Number of packet which contains incomplete message; 'msg_proxy_client_drop': Number of AX drop; 'msg_proxy_client_connection': Connecting server; 'msg_proxy_client_fail': Number of SIP messages received from client but failed to forward to server; 'msg_proxy_client_fail_parse': msg_proxy_client_fail_parse; 'msg_proxy_client_fail_process': msg_proxy_client_fail_process; 'msg_proxy_client_fail_snat': msg_proxy_client_fail_snat; 'msg_proxy_client_exceed_tmp_buff': msg_proxy_client_exceed_tmp_buff; 'msg_proxy_client_fail_send_pkt': msg_proxy_client_fail_send_pkt; 'msg_proxy_client_fail_start_server_Conn': msg_proxy_client_fail_start_server_Conn; 'msg_proxy_server_recv': Number of SIP messages received from server; 'msg_proxy_server_send_success': Number of SIP messages received from server and forwarded to client; 'msg_proxy_server_incomplete': Number of packet which contains incomplete message; 'msg_proxy_server_drop': Number of AX drop; 'msg_proxy_server_fail': Number of SIP messages received from server but failed to forward to client; 'msg_proxy_server_fail_parse': msg_proxy_server_fail_parse; 'msg_proxy_server_fail_process': msg_proxy_server_fail_process; 'msg_proxy_server_fail_selec_connt': msg_proxy_server_fail_selec_connt; 'msg_proxy_server_fail_snat': msg_proxy_server_fail_snat; 'msg_proxy_server_exceed_tmp_buff': msg_proxy_server_exceed_tmp_buff; 'msg_proxy_server_fail_send_pkt': msg_proxy_server_fail_send_pkt; 'msg_proxy_create_server_conn': Number of server connection system tries to create; 'msg_proxy_start_server_conn': Number of server connection created successfully; 'msg_proxy_fail_start_server_conn': Number of server connection create failed; 'msg_proxy_server_conn_fail_snat': msg_proxy_server_conn_fail_snat; 'msg_proxy_fail_construct_server_conn': msg_proxy_fail_construct_server_conn; 'msg_proxy_fail_reserve_pconn': msg_proxy_fail_reserve_pconn; 'msg_proxy_start_server_conn_failed': msg_proxy_start_server_conn_failed; 'msg_proxy_server_conn_already_exists': msg_proxy_server_conn_already_exists; 'msg_proxy_fail_insert_server_conn': msg_proxy_fail_insert_server_conn; 'msg_proxy_parse_msg_fail': msg_proxy_parse_msg_fail; 'msg_proxy_process_msg_fail': msg_proxy_process_msg_fail; 'msg_proxy_no_vport': msg_proxy_no_vport; 'msg_proxy_fail_select_server': msg_proxy_fail_select_server; 'msg_proxy_fail_alloc_mem': msg_proxy_fail_alloc_mem; 'msg_proxy_unexpected_err': msg_proxy_unexpected_err; 'msg_proxy_l7_cpu_failed': msg_proxy_l7_cpu_failed; 'msg_proxy_l4_to_l7': msg_proxy_l4_to_l7; 'msg_proxy_l4_from_l7': msg_proxy_l4_from_l7; 'msg_proxy_to_l4_send_pkt': msg_proxy_to_l4_send_pkt; 'msg_proxy_l4_from_l4_send': msg_proxy_l4_from_l4_send; 'msg_proxy_l7_to_L4': msg_proxy_l7_to_L4; 'msg_proxy_mag_back': msg_proxy_mag_back; 'msg_proxy_fail_dcmsg': msg_proxy_fail_dcmsg; 'msg_proxy_deprecated_conn': msg_proxy_deprecated_conn; 'msg_proxy_hold_msg': msg_proxy_hold_msg; 'msg_proxy_split_pkt': msg_proxy_split_pkt; 'msg_proxy_pipline_msg': msg_proxy_pipline_msg; 'msg_proxy_client_reset': msg_proxy_client_reset; 'msg_proxy_server_reset': msg_proxy_server_reset; 'session_created': SIP Session created; 'session_freed': SIP Session freed; 'session_in_rml': session_in_rml; 'session_invalid': session_invalid; 'conn_allocd': conn_allocd; 'conn_freed': conn_freed; 'session_callid_allocd': session_callid_allocd; 'session_callid_freed': session_callid_freed; 'line_mem_allocd': line_mem_allocd; 'line_mem_freed': line_mem_freed; 'table_mem_allocd': table_mem_allocd; 'table_mem_freed': table_mem_freed; 'cmsg_no_uri_header': cmsg_no_uri_header; 'cmsg_no_uri_session': cmsg_no_uri_session; 'sg_no_uri_header': sg_no_uri_header; 'smsg_no_uri_session': smsg_no_uri_session; 'line_too_long': line_too_long; 'fail_read_start_line': fail_read_start_line; 'fail_parse_start_line': fail_parse_start_line; 'invalid_start_line': invalid_start_line; 'request_unknown_version': request_unknown_version; 'response_unknown_version': response_unknown_version; 'request_unknown': request_unknown; 'fail_parse_headers': fail_parse_headers; 'too_many_headers': too_many_headers; 'invalid_header': invalid_header; 'header_name_too_long': header_name_too_long; 'body_too_big': body_too_big; 'fail_get_counter': fail_get_counter; 'msg_no_call_id': msg_no_call_id; 'identify_dir_failed': identify_dir_failed; 'no_sip_request': no_sip_request; 'deprecated_msg': deprecated_msg; 'fail_insert_callid_session': fail_insert_callid_session; 'fail_insert_uri_session': fail_insert_uri_session; 'fail_insert_header': fail_insert_header; 'select_server_conn': select_server_conn; 'select_server_conn_by_callid': select_server_conn_by_callid; 'select_server_conn_by_uri': select_server_conn_by_uri; 'select_server_conn_by_rev_tuple': select_server_conn_by_rev_tuple; 'select_server_conn_failed': select_server_conn_failed; 'select_client_conn': select_client_conn; 'X_forward_for_select_client': X_forward_for_select_client; 'call_id_select_client': call_id_select_client; 'uri_select_client': uri_select_client; 'client_select_failed': client_select_failed; 'acl_denied': acl_denied; 'assemble_frag_failed': assemble_frag_failed; 'wrong_ip_version': wrong_ip_version; 'size_too_large': size_too_large; 'fail_split_fragment': fail_split_fragment; 'client_keepalive_received': client_keepalive_received; 'server_keepalive_received': server_keepalive_received; 'client_keepalive_send': client_keepalive_send; 'server_keepalive_send': server_keepalive_send; 'ax_health_check_received': ax_health_check_received; 'client_request': client_request; 'client_request_ok': client_request_ok; 'concatenate_msg': concatenate_msg; 'save_uri': save_uri; 'save_uri_ok': save_uri_ok; 'save_call_id': save_call_id; 'save_call_id_ok': save_call_id_ok; 'msg_translation': msg_translation; 'msg_translation_fail': msg_translation_fail; 'msg_trans_start_line': msg_trans_start_line; 'msg_trans_start_headers': msg_trans_start_headers; 'msg_trans_body': msg_trans_body; 'request_register': request_register; 'request_invite': request_invite; 'request_ack': request_ack; 'request_cancel': request_cancel; 'request_bye': request_bye; 'request_options': request_options; 'request_prack': request_prack; 'request_subscribe': request_subscribe; 'request_notify': request_notify; 'request_publish': request_publish; 'request_info': request_info; 'request_refer': request_refer; 'request_message': request_message; 'request_update': request_update; 'response_unknown': response_unknown; 'response_1XX': response_1XX; 'response_2XX': response_2XX; 'response_3XX': response_3XX; 'response_4XX': response_4XX; 'response_5XX': response_5XX; 'response_6XX': response_6XX; 'ha_send_sip_session': ha_send_sip_session; 'ha_send_sip_session_ok': ha_send_sip_session_ok; 'ha_fail_get_msg_header': ha_fail_get_msg_header; 'ha_recv_sip_session': ha_recv_sip_session; 'ha_insert_sip_session_ok': ha_insert_sip_session_ok; 'ha_update_sip_session_ok': ha_update_sip_session_ok; 'ha_invalid_pkt': ha_invalid_pkt; 'ha_fail_alloc_sip_session': ha_fail_alloc_sip_session; 'ha_fail_alloc_call_id': ha_fail_alloc_call_id; 'ha_fail_clone_sip_session': ha_fail_clone_sip_session; 'save_smp_call_id_rtp': save_smp_call_id_rtp; 'update_smp_call_id_rtp': update_smp_call_id_rtp; 'smp_call_id_rtp_session_match': smp_call_id_rtp_session_match; 'smp_call_id_rtp_session_not_match': smp_call_id_rtp_session_not_match; 'process_error_when_message_switch': process_error_when_message_switch;",
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
func resourceSlbSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSipRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSipRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbSipSamplingEnable(d []interface{}) []edpt.SlbSipSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbSipSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSipSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSip(d *schema.ResourceData) edpt.SlbSip {
	var ret edpt.SlbSip
	ret.Inst.SamplingEnable = getSliceSlbSipSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}