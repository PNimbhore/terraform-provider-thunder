package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAState() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_state`: HA VRRP-A Global Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAStateCreate,
		UpdateContext: resourceVrrpAStateUpdate,
		ReadContext:   resourceVrrpAStateRead,
		DeleteContext: resourceVrrpAStateDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'sync_pkt_tx_counter': Conn Sync Sent counter; 'sync_pkt_rcv_counter': Conn Sync Received counter; 'sync_rx_create_counter': Conn Sync Create Session Received counter; 'sync_rx_del_counter': Conn Sync Del Session Received counter; 'sync_rx_update_age_counter': Conn Sync Update Age Received counter; 'sync_tx_create_counter': Conn Sync Create Session Sent counter; 'sync_tx_del_counter': Conn Sync Del Session Sent counter; 'sync_tx_update_age_counter': Conn Sync Update Age Sent counter; 'sync_rx_persist_create_counter': Conn Sync Create Persist Session Pkts Received counter; 'sync_rx_persist_del_counter': Conn Sync Delete Persist Session Pkts Received counter; 'sync_rx_persist_update_age_counter': Conn Sync Update Persist Age Pkts Received counter; 'sync_tx_persist_create_counter': Conn Sync Create Persist Session Pkts Sent counter; 'sync_tx_persist_del_counter': Conn Sync Delete Persist Session Pkts Sent counter; 'sync_tx_persist_update_age_counter': Conn Sync Update Persist Age Pkts Sent counter; 'query_pkt_tx_counter': Conn Query sent counter; 'query_pkt_rcv_counter': Conn Query Received counter; 'sync_tx_smp_radius_table_counter': Conn Sync Update LSN RADIUS Sent counter; 'sync_rx_smp_radius_table_counter': Conn Sync Update LSN RADIUS Received counter; 'query_tx_max_packed': Max Query Msg Per Packet; 'query_tx_min_packed': Min Query Msg Per Packet; 'query_pkt_invalid_idx_counter': Conn Query Invalid Interface; 'query_tx_get_buff_failed': Conn Query Get Buff Failure; 'query_rx_zero_info_counter': Conn Query Packet Empty; 'query_rx_full_info_counter': Conn Query Packet Full; 'query_rx_unk_counter': Conn Query Unknown Type; 'sync_pkt_invalid_idx_counter': Conn Sync Invalid Interface; 'sync_tx_get_buff_failed': Conn Sync Get Buff Failure; 'sync_tx_total_info_counter': Conn Sync Total Info Pkts Sent counter; 'sync_tx_create_ext_bit_counter': Conn Sync Create with Ext Sent counter; 'sync_tx_update_seqnos_counter': Conn Sync Update Seq Num Sent counter; 'sync_tx_min_packed': Max Sync Msg Per Packet; 'sync_tx_max_packed': Min Sync Msg Per Packet; 'sync_rx_len_invalid': Conn Sync Length Invalid; 'sync_persist_rx_len_invalid': Persist Conn Sync Length Invalid; 'sync_persist_rx_proto_not_supported': Persist Conn Sync Protocol Invalid; 'sync_persist_rx_type_invalid': Persist Conn Sync Type Invalid; 'sync_persist_rx_cannot_process_mandatory': Persist Conn Sync Process Mandatory Invalid; 'sync_persist_rx_ext_bit_process_error': Persist Conn Sync Proc Ext Bit Failure; 'sync_persist_rx_no_such_vport': Persist Conn Sync Virt Port Not Found; 'sync_persist_rx_vporttype_not_supported': Persist Conn Sync Virt Port Type Invalid; 'sync_persist_rx_no_such_rport': Persist Conn Sync Real Port Not Found; 'sync_persist_rx_no_such_sg_group': Persist Conn Sync No Service Group Found; 'sync_persist_rx_no_sg_group_info': Persist Conn Sync No Service Group Info Found; 'sync_persist_rx_conn_get_failed': Persist Conn Sync Get Conn Failure; 'sync_rx_no_such_vport': Conn Sync Virt Port Not Found; 'sync_rx_no_such_rport': Conn Sync Real Port Not Found; 'sync_rx_cannot_process_mandatory': Conn Sync Process Mandatory Invalid; 'sync_rx_ext_bit_process_error': Conn Sync Proc Ext Bit Failure; 'sync_rx_create_ext_bit_counter': Conn Sync Create with Ext Received counter; 'sync_rx_conn_exists': Conn Sync Create Conn Exists; 'sync_rx_conn_get_failed': Conn Sync Get Conn Failure; 'sync_rx_proto_not_supported': Conn Sync Protocol Invalid; 'sync_rx_no_dst_for_vport_inline': Conn Sync 'dst' not found for vport inline; 'sync_rx_no_such_nat_pool': Conn Sync NAT Pool Error; 'sync_rx_no_such_sg_node': Conn Sync no SG node found; 'sync_rx_del_no_such_session': Conn Sync Del Conn not Found; 'sync_rx_type_invalid': Conn Sync Type Invalid; 'sync_rx_zero_info_counter': Conn Sync Packet Empty; 'sync_rx_dcmsg_counter': Conn Sync forward CPU; 'sync_rx_total_info_counter': Conn Sync Total Info Pkts Received counter; 'sync_rx_update_seqnos_counter': Conn Sync Update Seq Num Received counter; 'sync_rx_unk_counter': Conn Sync Unknown Type; 'sync_rx_apptype_not_supported': Conn Sync App Type Invalid; 'sync_query_dcmsg_counter': Conn Sync query forward CPU; 'sync_get_buff_failed_rt': Conn Sync Get Buff Failure No Route; 'sync_get_buff_failed_port': Conn Sync Get Buff Failure Wrong Port; 'sync_rx_lsn_create_sby': Conn Sync LSN Create Standby; 'sync_rx_nat_create_sby': Conn Sync NAT Create Standby; 'sync_rx_nat_alloc_sby': Conn Sync NAT Alloc Standby; 'sync_rx_insert_tuple': Conn Sync Insert Tuple; 'sync_rx_sfw': Conn Sync SFW; 'sync_rx_create_static_sby': Conn Sync Create Static Standby; 'sync_rx_ext_pptp': Conn Sync Ext PPTP; 'sync_rx_ext_rtsp': Conn Sync Ext RTSP; 'sync_rx_reserve_ha': Conn Sync Reserve HA Conn; 'sync_rx_seq_deltas': Conn Sync Seq Deltas Failure; 'sync_rx_ftp_control': Conn Sync FTP Control Failure; 'sync_rx_ext_lsn_acl': Conn Sync LSN ACL Failure; 'sync_rx_ext_lsn_ac_idle_timeout': Conn Sync LSN ACL Idle Timeout Failure; 'sync_rx_ext_sip_alg': Conn Sync SIP TCP ALG Failure; 'sync_rx_ext_h323_alg': Conn Sync H323 TCP ALG Failure; 'sync_rx_ext_nat_mac': Conn Sync NAT MAC Failure; 'sync_tx_lsn_fullcone': Conn Sync Update LSN Fullcone Sent counter; 'sync_rx_lsn_fullcone': Conn Sync Update LSN Fullcone Received counter; 'sync_err_lsn_fullcone': Conn Sync LSN Fullcone Failure; 'sync_tx_update_sctp_conn_addr': Update SCTP Addresses Sent; 'sync_rx_update_sctp_conn_addr': Update SCTP Addresses Received; 'sync_rx_ext_nat_alg_tcp_info': Conn Sync NAT ALG TCP Information; 'sync_rx_ext_dcfw_rule_id': Conn Sync FIREWALL session rule ID information Failure; 'sync_rx_ext_dcfw_log': Conn Sync FIREWALL session logging information Failure; 'sync_rx_estab_counter': Conn Sync rcv established state; 'sync_tx_estab_counter': Conn Sync send established state; 'sync_rx_zone_failure_counter': Conn Sync Zone Failure; 'sync_rx_ext_fw_http_logging': FW HTTP Logging Sync Failures; 'sync_rx_ext_dcfw_rule_idle_timeout': Conn Sync FIREWALL session rule idle timeout information Failure; 'sync_rx_ext_fw_gtp_info': FW GTP Info Received; 'sync_rx_not_expect_sync_pkt': unexpected session sync packets; 'sync_rx_ext_fw_apps': Conn Sync FIREWALL application information Failure; 'sync_tx_mon_entity': Acos Monitoring Entities Sync Messages Sent; 'sync_rx_mon_entity': Acos monitoring Entities Sync Messages Received; 'sync_rx_ext_fw_gtp_log_info': FW GTP Log Info Received; 'sync_rx_ext_fw_gtp_u_info': FW GTP U Info Received; 'sync_rx_ext_fw_gtp_ext_info': FW GTP Ext Info Received; 'sync_rx_ext_fw_gtp_log_ext_info': FW GTP Ext Log Info Received; 'sync_rx_ddos_drop_counter': Conn Sync receive ddos protect packet; 'sync_rx_invalid_sync_packet_counter': Conn Sync receive invalid packet; 'sync_pkt_empty_buff_counter': Conn Sync drop sending packet for empty buffer; 'sync_pkt_no_sending_vgrp_counter': Conn Sync drop sending packet for invalid sending virtual group; 'sync_pkt_no_receiving_vgrp_counter': Conn Sync drop sending packet for invalid receiving virtual group; 'query_pkt_no_receiving_ip_counter': Conn Sync drop sending packet for invalid receiving ip; 'sync_pkt_failed_buff_copy_counter': Conn Sync drop sending packet for failure in sending buffer copy; 'sync_rx_bad_protocol_counter': Conn Sync receive packet with bad protocol; 'sync_rx_no_vgrp_counter': Conn Sync receive packet with no virtual group; 'sync_rx_by_inactive_peer_counter': Conn Sync receive packet by inactive peer; 'sync_rx_table_entry_update_counter': Conn Sync receive packet with table entry update; 'sync_rx_table_entry_create_counter': Conn Sync receive packet with table entry create;",
						},
						"counters2": {
							Type: schema.TypeString, Optional: true, Description: "'sync_rx_table_entry_del_counter': Conn Sync receive packet with table entry delete; 'sync_rx_aflex_update_counter': Conn Sync receive packet with aflex update; 'sync_rx_aflex_create_counter': Conn Sync receive packet with aflex create; 'sync_rx_aflex_del_counter': Conn Sync receive packet with aflex delete; 'sync_rx_aflex_frag_counter': Conn Sync receive packet with aflex fragment; 'query_rx_invalid_partition_counter': Conn Sync receive query packet with invalid partition; 'query_rx_invalid_ha_group_counter': Conn Sync receive query packet with invalid ha group; 'query_rx_invalid_sync_version_counter': Conn Sync receive query packet with invalid sync version; 'query_rx_invalid_msg_dir_counter': Conn Sync receive query packet with invalid message dir; 'sync_rx_out_of_order_pkt_counter': total number of out of order packets received; 'sync_rx_unreached_pkt_counter': total number of unreached packets; 'sync_rx_ext_fw_gtp_echo_ext_info': FW GTP Echo Ext Info Received; 'sync_rx_smp_create_counter': Sync Create SMP Session Pkts Received counter; 'sync_rx_smp_delete_counter': Sync Delete SMP Session Pkts Received counter; 'sync_rx_smp_update_counter': Sync Update SMP Session Pkts Received counter; 'sync_tx_smp_create_counter': Sync Create SMP Session Pkts Sent counter; 'sync_tx_smp_delete_counter': Sync Delete SMP Session Pkts Sent counter; 'sync_tx_smp_update_counter': Sync Update SMP Session Pkts Sent counter; 'sync_rx_smp_clear_counter': Sync Clear SMP Session Pkts Received counter; 'sync_tx_smp_clear_counter': Sync Clear SMP Session Pkts Sent counter; 'sync_rx_ext_fw_so_shadow_ext_info': FW Scaleout Shadow Ext Info Received; 'sync_tx_aflex_table_entry_add_counter': Sync send packet with aflex table entry add; 'sync_rx_aflex_table_entry_add_counter': Sync receive packet with aflex table entry add; 'sync_tx_aflex_table_entry_append_counter': Sync send packet with aflex table entry append; 'sync_rx_aflex_table_entry_append_counter': Sync receive packet with aflex table entry append; 'sync_tx_aflex_table_entry_delete_counter': Sync send packet with aflex table entry delete; 'sync_rx_aflex_table_entry_delete_counter': Sync receive packet with aflex table entry delete; 'sync_tx_aflex_table_entry_incr_counter': Sync send packet with aflex table entry incr; 'sync_rx_aflex_table_entry_incr_counter': Sync receive packet with aflex table entry incr; 'sync_tx_aflex_table_entry_lookup_counter': Sync send packet with aflex table entry lookup; 'sync_rx_aflex_table_entry_lookup_counter': Sync receive packet with aflex table entry lookup; 'sync_tx_aflex_table_entry_lifetime_counter': Sync send packet with aflex table entry lifetime; 'sync_rx_aflex_table_entry_lifetime_counter': Sync receive packet with aflex table entry lifetime; 'sync_tx_aflex_table_entry_replace_counter': Sync send packet with aflex table entry replace; 'sync_rx_aflex_table_entry_replace_counter': Sync receive packet with aflex table entry replace; 'sync_tx_aflex_table_entry_set_counter': Sync send packet with aflex table entry set; 'sync_rx_aflex_table_entry_set_counter': Sync receive packet with aflex table entry set; 'sync_tx_aflex_table_entry_timeout_counter': Sync send packet with aflex table entry timeout; 'sync_rx_aflex_table_entry_timeout_counter': Sync receive packet with aflex table entry timeout; 'sync_tx_aflex_table_entry_fastsync_counter': Sync send packet with aflex table entry fast sync; 'sync_rx_aflex_table_entry_fastsync_counter': Sync receive packet with aflex table entry fast sync; 'sync_tx_aflex_table_entry_error_counter': Error on send packet with aflex table entry; 'sync_tx_aflex_table_entry_not_eligible_counter': send of aflex table entry not eligible; 'sync_rx_ext_fw_limit_entry': Sync FW Limit Entry Info Failure; 'sync_tx_fw_set_dscp_counter': Conn Sync send fw set dscp counter; 'sync_rx_fw_set_dscp_counter': Conn Sync receive fw set dscp counter; 'dns_cache_sync_tx_create_counter': DNS Cache Sync Create Sent counter; 'dns_cache_sync_rx_create_counter': DNS Cache Sync Create Received counter; 'dns_cache_sync_tx_del_counter': DNS Cache Sync Del Sent counter; 'dns_cache_sync_rx_del_counter': DNS Cache Sync Del Received counter; 'dns_cache_sync_tx_frag_counter': DNS Cache Sync Frag Sent counter; 'dns_cache_sync_rx_frag_counter': DNS Cache Sync Frag Received counter; 'dns_cache_sync_tx_error_counter': DNS Cache Sync Error Sent counter; 'dns_cache_sync_rx_error_counter': DNS Cache Sync Error Received counter;",
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
func resourceVrrpAStateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAStateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAState(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAStateRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAStateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAStateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAState(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAStateRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAStateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAStateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAState(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAStateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAStateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAState(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVrrpAStateSamplingEnable(d []interface{}) []edpt.VrrpAStateSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VrrpAStateSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAStateSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAState(d *schema.ResourceData) edpt.VrrpAState {
	var ret edpt.VrrpAState
	ret.Inst.SamplingEnable = getSliceVrrpAStateSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
