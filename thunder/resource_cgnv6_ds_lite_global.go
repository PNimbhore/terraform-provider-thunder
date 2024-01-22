package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DsLiteGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ds_lite_global`: Configure Dual-Stack Lite (DS-Lite) config parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6DsLiteGlobalCreate,
		UpdateContext: resourceCgnv6DsLiteGlobalUpdate,
		ReadContext:   resourceCgnv6DsLiteGlobalRead,
		DeleteContext: resourceCgnv6DsLiteGlobalDelete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send_on_port_unavailable": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'host-unreachable': Send ICMP destination host unreachable; 'admin-filtered': Send ICMP admin filtered; 'disable': Disable ICMP port unavailable message (default);",
						},
						"send_on_user_quota_exceeded": {
							Type: schema.TypeString, Optional: true, Default: "admin-filtered", Description: "'host-unreachable': Send ICMP destination host unreachable; 'admin-filtered': Send ICMP admin filtered (default); 'disable': Disable ICMP quota exceeded message;",
						},
					},
				},
			},
			"inside": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list": {
										Type: schema.TypeString, Optional: true, Description: "Class-list to match for DS-Lite",
									},
								},
							},
						},
					},
				},
			},
			"ip_checksum_error": {
				Type: schema.TypeString, Optional: true, Default: "fix", Description: "'fix': Fix the bad checksum (default); 'drop': Drop packets with a bad checksum;",
			},
			"l4_checksum_error": {
				Type: schema.TypeString, Optional: true, Default: "propagate", Description: "'propagate': Propagate the bad checksum (default); 'fix': Fix the bad checksum; 'drop': Drop packets with a bad checksum;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_tcp_allocated': Total TCP Ports Allocated; 'total_tcp_freed': Total TCP Ports Freed; 'total_udp_allocated': Total UDP Ports Allocated; 'total_udp_freed': Total UDP Ports Freed; 'total_icmp_allocated': Total ICMP Ports Allocated; 'total_icmp_freed': Total ICMP Ports Freed; 'data_session_created': Data Session Created; 'data_session_freed': Data Session Freed; 'user_quota_created': User-Quota Created; 'user_quota_put_in_del_q': User-Quota Freed; 'user_quota_failure': User-Quota Creation Failed; 'nat_port_unavailable_tcp': TCP NAT Port Unavailable; 'nat_port_unavailable_udp': UDP NAT Port Unavailable; 'nat_port_unavailable_icmp': ICMP NAT Port Unavailable; 'new_user_resource_unavailable': New User NAT Resource Unavailable; 'tcp_user_quota_exceeded': TCP User-Quota Exceeded; 'udp_user_quota_exceeded': UDP User-Quota Exceeded; 'icmp_user_quota_exceeded': ICMP User-Quota Exceeded; 'extended_quota_matched': Extended User-Quota Matched; 'extended_quota_exceeded': Extended User-Quota Exceeded; 'data_sesn_user_quota_exceeded': Data Session User-Quota Exceeded; 'data_sesn_rate_user_quota_exceeded': Conn Rate User-Quota Exceeded; 'tcp_fullcone_created': TCP Full-cone Session Created; 'tcp_fullcone_freed': TCP Full-cone Session Freed; 'udp_fullcone_created': UDP Full-cone Session Created; 'udp_fullcone_freed': UDP Full-cone Session Freed; 'fullcone_failure': Full-cone Session Creation Failed; 'hairpin': Hairpin Session Created; 'fullcone_self_hairpinning_drop': Self-Hairpinning Drop; 'endpoint_indep_map_match': Endpoint-Independent Mapping Matched; 'endpoint_indep_filter_match': Endpoint-Independent Filtering Matched; 'inbound_filtered': Endpoint-Dependent Filtering Drop; 'eif_limit_exceeded': Endpoint-Independent Filtering Inbound Limit Exceeded; 'total_tcp_overloaded': TCP Port Overloaded; 'total_udp_overloaded': UDP Port Overloaded; 'port_overloading_smp_inserted_tcp': TCP Port Overloading Session Created; 'port_overloading_smp_inserted_udp': UDP Port Overloading Session Created; 'port_overloading_smp_free_tcp': TCP Port Overloading Session Freed; 'port_overloading_smp_free_udp': UDP Port Overloading Session Freed; 'nat_pool_unusable': NAT Pool Unusable; 'ha_nat_pool_unusable': HA NAT Pool Unusable; 'ha_nat_pool_batch_type_mismatch': HA NAT Pool Batch Type Mismatch; 'no_radius_profile_match': No RADIUS Profile Match; 'truncated_packet': Truncated Packet; 'lid_drop': LSN LID Drop; 'lid_pass_through': LSN LID Pass-through; 'no_class_list_match': No Class-List Match; 'class_list_permit_drop': Permit Class-List Drop; 'user_quota_mem_allocated': User-Quota Memory Allocated; 'user_quota_mem_freed': User-Quota Memory Freed; 'user_quota_created_shadow': Total User-Quota Created; 'quota_marked_deleted': User-Quota Marked Deleted; 'quota_delete_not_in_bucket': User-Quota Delete Not In Bucket; 'user_quota_put_in_del_q_shadow': Total User-Quota Put In Del Q; 'prefix_quota_created': Total Prefix Quota Created; 'prefix_quota_put_in_del_q': Total Prefix Quota Freed; 'prefix_quota_failure': Total Prefix Quota Failure; 'total_user_quota_created': Total Quota Structure Created; 'total_user_quota_put_in_del_q': Total Quota Structure Freed; 'total_user_quota_failure': Total Quota Structure Failure; 'tcp_out_of_state_rst_sent': Total Out of State TCP RST sent; 'tcp_out_of_state_rst_dropped': Total Out of State TCP RST dropped; 'icmp_out_of_state_uqe_admin_filtered_sent': Total User Quota Exceeded ICMP admin filtered sent; 'icmp_out_of_state_uqe_host_unreachable_sent': Total User Quota Exceeded ICMP host unreachable sent; 'icmp_out_of_state_uqe_dropped': Total User Queue Exceeded ICMP notification dropped; 'user_quota_not_found': User-Quota Not Found; 'tcp_fullcone_created_shadow': Total TCP Full-cone sessions created; 'tcp_fullcone_freed_shadow': Total TCP Full-cone sessions freed; 'udp_fullcone_created_shadow': Total UDP Full-cone sessions created; 'udp_fullcone_freed_shadow': Total UDP Full-cone sessions freed; 'udp_alg_fullcone_created': Total UDP ALG Full-cone sessions created; 'udp_alg_fullcone_freed': Total UDP ALG Full-cone sessions freed; 'fullcone_created': Total Full-cone sessions created; 'fullcone_freed': Total Full-cone sessions freed; 'data_session_created_shadow': Total Data Sessions Created; 'data_session_freed_shadow': Total Data Sessions Freed; 'data_session_user_quota_mismatch': Data Session Counter Per User Mismatch; 'extended_quota_mismatched': Extended User-Quota Mismatched; 'nat_port_unavailable_other': Other NAT Port Unavailable; 'nat_port_unavailable': Total NAT Port Unavailable; 'new_user_resource_unavailable_tcp': TCP New User NAT Resource Unavailable; 'new_user_resource_unavailable_udp': UDP New User NAT Resource Unavailable; 'new_user_resource_unavailable_icmp': ICMP New User NAT Resource Unavailable; 'new_user_resource_unavailable_other': Other New User NAT Resource Unavailable; 'total_tcp_allocated_shadow': Total TCP Ports Allocated; 'total_tcp_freed_shadow': Total TCP Ports Freed; 'total_udp_allocated_shadow': Total UDP Ports Allocated; 'total_udp_freed_shadow': Total UDP Ports Freed; 'total_icmp_allocated_shadow': Total ICMP Ports Allocated; 'total_icmp_freed_shadow': Total ICMP Ports Freed; 'udp_alg_no_quota': UDP ALG User-Quota Not Found; 'udp_alg_eim_mismatch': UDP ALG Endpoint-Independent Mapping Mismatch; 'udp_alg_no_nat_ip': UDP ALG User-Quota Unknown NAT IP; 'udp_alg_alloc_failure': UDP ALG Port Allocation Failure; 'sip_alg_no_quota': SIP ALG User-Quota Not Found; 'sip_alg_quota_inc_failure': SIP ALG User-Quota Exceeded; 'sip_alg_no_nat_ip': SIP ALG Unknown NAT IP; 'sip_alg_reuse_contact_fullcone': SIP ALG Reuse Contact Full-cone Session; 'sip_alg_contact_fullcone_mismatch': SIP ALG Contact Full-cone Session Mismatch; 'sip_alg_alloc_contact_port_failure': SIP ALG Alloc Contact NAT Port Failure; 'sip_alg_create_contact_fullcone_failure': SIP ALG Create Contact Full-cone Session Failure; 'sip_alg_release_contact_port_failure': SIP ALG Release Contact NAT Port Failure; 'sip_alg_single_rtp_fullcone': SIP ALG Single RTP Full-cone Found; 'sip_alg_single_rtcp_fullcone': SIP ALG Single RTCP Full-cone Found; 'sip_alg_rtcp_fullcone_mismatch': SIP ALG RTCP Full-cone NAT Port Mismatch; 'sip_alg_reuse_rtp_rtcp_fullcone': SIP ALG Reuse RTP/RTCP Full-cone Session; 'sip_alg_alloc_rtp_rtcp_port_failure': SIP ALG Alloc RTP/RTCP NAT Ports Failure; 'sip_alg_alloc_single_port_failure': SIP ALG Alloc Single RTP or RTCP NAT Port Failure; 'sip_alg_create_single_fullcone_failure': SIP ALG Create Single RTP or RTCP Full-cone Session Failure; 'sip_alg_create_rtp_fullcone_failure': SIP ALG Create RTP Full-cone Session Failure; 'sip_alg_create_rtcp_fullcone_failure': SIP ALG Create RTCP Full-cone Session Failure; 'sip_alg_port_pair_alloc_from_consecutive': SIP ALG Port Pair Allocated From Consecutive; 'sip_alg_port_pair_alloc_from_partition': SIP ALG Port Pair Allocated From Partition; 'sip_alg_port_pair_alloc_from_pool_port_batch': SIP ALG Port Pair Allocated From Pool Port Batch; 'sip_alg_port_pair_alloc_from_quota_consecutive': SIP ALG Port Pair Allocated From Quota Consecutive; 'sip_alg_port_pair_alloc_from_quota_partition': SIP ALG Port Pair Allocated From Quota Partition; 'sip_alg_port_pair_alloc_from_quota_partition_error': SIP ALG Port Pair Allocated From Quota Partition Error; 'sip_alg_port_pair_alloc_from_quota_pool_port_batch': SIP ALG Port Pair Allocated From Quota Pool Port Batch; 'sip_alg_port_pair_alloc_from_quota_pool_port_batch_with_frag': SIP ALG Port Pair Allocated From Quota Port Batch Version 2 with Frag Free Ports;",
						},
						"counters2": {
							Type: schema.TypeString, Optional: true, Description: "'h323_alg_no_quota': H323 ALG User-Quota Not Found; 'h323_alg_quota_inc_failure': H323 ALG User-Quota Exceeded; 'h323_alg_no_nat_ip': H323 ALG Unknown NAT IP; 'h323_alg_reuse_fullcone': H323 ALG Reuse Full-cone Session; 'h323_alg_fullcone_mismatch': H323 ALG Full-cone Session Mismatch; 'h323_alg_alloc_port_failure': H323 ALG Alloc NAT Port Failure; 'h323_alg_create_fullcone_failure': H323 ALG Create Full-cone Session Failure; 'h323_alg_release_port_failure': H323 ALG Release NAT Port Failure; 'h323_alg_single_rtp_fullcone': H323 ALG Single RTP Full-cone Found; 'h323_alg_single_rtcp_fullcone': H323 ALG Single RTCP Full-cone Found; 'h323_alg_rtcp_fullcone_mismatch': H323 ALG RTCP Full-cone NAT Port Mismatch; 'h323_alg_reuse_rtp_rtcp_fullcone': H323 ALG Reuse RTP/RTCP Full-cone Session; 'h323_alg_alloc_rtp_rtcp_port_failure': H323 ALG Alloc RTP/RTCP NAT Ports Failure; 'h323_alg_alloc_single_port_failure': H323 ALG Alloc Single RTP or RTCP NAT Port Failure; 'h323_alg_create_single_fullcone_failure': H323 ALG Create Single RTP or RTCP Full-cone Session Failure; 'h323_alg_create_rtp_fullcone_failure': H323 ALG Create RTP Full-cone Session Failure; 'h323_alg_create_rtcp_fullcone_failure': H323 ALG Create RTCP Full-cone Session Failure; 'h323_alg_port_pair_alloc_from_consecutive': H323 ALG Port Pair Allocated From Consecutive; 'h323_alg_port_pair_alloc_from_partition': H323 ALG Port Pair Allocated From Partition; 'h323_alg_port_pair_alloc_from_pool_port_batch': H323 ALG Port Pair Allocated From Pool Port Batch; 'h323_alg_port_pair_alloc_from_quota_consecutive': H323 ALG Port Pair Allocated From Quota Consecutive; 'h323_alg_port_pair_alloc_from_quota_partition': H323 ALG Port Pair Allocated From Quota Partition; 'h323_alg_port_pair_alloc_from_quota_partition_error': H323 ALG Port Pair Allocated From Quota Partition Error; 'h323_alg_port_pair_alloc_from_quota_pool_port_batch': H323 ALG Port Pair Allocated From Quota Pool Port Batch; 'port_batch_quota_extension_alloc_failure': Port Batch Extension Alloc Failure (No memory); 'port_batch_free_quota_not_found': Port Batch Quota Not Found on Free; 'port_batch_free_port_not_found': Port Batch Port Not Found on Free; 'port_batch_free_wrong_partition': Port Batch Free Wrong Partition; 'radius_query_quota_ext_alloc_failure': RADIUS Query Container Alloc (No Memoty); 'radius_query_quota_ext_alloc_race_free': RADIUS Query Container Alloc Race Free; 'quota_extension_added': Quota Extension Added to Quota; 'quota_extension_removed': Quota Extension Removed from Quota; 'quota_extension_remove_not_found': Quota Extension Not Found on Remove; 'ha_sync_port_batch_sent': HA Port Batch Extension Sync Sent; 'ha_sync_port_batch_rcv': HA Port Batch Extension Sync Received; 'ha_send_port_batch_not_found': HA Port Batch Not Found on Sync Send; 'ha_rcv_port_not_in_port_batch': HA Port Not in Port Batch on Sync Rcv; 'bad_port_to_free': Freeing Bad Port; 'consecutive_port_free': Port Freed from Consecutive Pool; 'partition_port_free': Port Freed from Partition; 'pool_port_batch_port_free': Port Freed from Pool Port Batch; 'port_allocated_from_quota_consecutive': Port Allocated from Quota Consecutive; 'port_allocated_from_quota_partition': Port Allocated from Quota Partition; 'port_allocated_from_quota_pool_port_batch': Port Allocated from Quota Pool Port Batch; 'port_freed_from_quota_consecutive': Port Freed from Quota Consecutive; 'port_freed_from_quota_partition': Port Freed from Quota Partition; 'port_freed_from_quota_pool_port_batch': Port Freed from Quota Pool Port Batch; 'port_batch_allocated_to_quota': Port Batch Allocated to Quota; 'port_batch_freed_from_quota': Port Batch Freed From Quota; 'specific_port_allocated_from_quota_consecutive': Specific Port Allocated from Quota Consecutive; 'specific_port_allocated_from_quota_partition': Specific Port Allocated from Quota Partition; 'specific_port_allocated_from_quota_pool_port_batch': Specific Port Allocated from Quota Pool Port Batch; 'port_batch_container_alloc_failure': Port Batch Container Alloc Out of Memory; 'port_batch_container_alloc_race_free': Port Batch Container Race Free; 'port_overloading_destination_conflict': Port Overloading Destination Conflict; 'port_overloading_out_of_memory': Port Overloading Out of Memory; 'port_overloading_assign_user': Port Overloading Port Assign User; 'port_overloading_assign_user_port_batch': Port Overloading Port Assign User Port Batch; 'port_overloading_inc': Port Overloading Port Increment; 'port_overloading_dec_on_conflict': Port Overloading Port Decrement on Conflict; 'port_overloading_dec_on_free': Port Overloading Port Decrement on Free; 'port_overloading_free_port_on_conflict': Port Overloading Free Port on Conflict; 'port_overloading_free_port_batch_on_conflict': Port Overloading Free Port Batch on Conflict; 'port_overloading_inc_overflow': Port Overloading Inc Overflow; 'port_overloading_attempt_consecutive_ports': Port Overloading on Consecutive Ports; 'port_overloading_attempt_same_partition': Port Overloading on Same Partition; 'port_overloading_attempt_diff_partition': Port Overloading on Different Partition; 'port_overloading_attempt_failed': Port Overloading Attempt Failed; 'port_overloading_conn_free_retry_lookup': Port Overloading Conn Free Retry Lookup; 'port_overloading_conn_free_not_found': Port Overloading Conn Free Not Found; 'port_overloading_smp_mem_allocated': Port Overloading SMP Session Allocated; 'port_overloading_smp_mem_freed': Port Overloading SMP Session Freed; 'port_overloading_smp_inserted': Port Overloading SMP Inserted; 'port_overloading_smp_inserted_tcp_shadow': Total Port Overloading SMP TCP Inserted; 'port_overloading_smp_inserted_udp_shadow': Total Port Overloading SMP UDP Inserted; 'port_overloading_smp_free_tcp_shadow': Total Port Overloading SMP TCP Freed; 'port_overloading_smp_free_udp_shadow': Total Port Overloading SMP UDP Freed; 'port_overloading_smp_put_in_del_q_from_conn': Port Overloading SMP Free From Conn; 'port_overloading_smp_free_dec_failure': Port Overloading SMP Free Dec Failure; 'port_overloading_smp_free_no_quota': Port Overloading SMP Free No Quota; 'port_overloading_smp_free_port': Port Overloading SMP Free Port; 'port_overloading_smp_free_port_from_quota': Port Overloading SMP Free Port From Quota; 'port_overloading_for_no_ports': Port Overloading for No Ports; 'port_overloading_for_no_ports_success': Port Overloading for No Ports Success; 'port_overloading_for_quota_exceeded': Port Overloading for Quota Exceeded; 'port_overloading_for_quota_exceeded_success': Port Overloading for Quota Exceeded Success; 'port_overloading_for_quota_exceeded_race': Port Overloading for Quota Exceeded Race; 'port_overloading_for_quota_exceeded_race_success': Port Overloading for Quota Exceeded Race Success; 'port_overloading_for_new_user': Port Overloading for New User; 'port_overloading_for_new_user_success': Port Overloading for New User Success; 'ha_port_overloading_attempt_failed': HA Port Overloading Attempt Failed; 'ha_port_overloading_for_no_ports': HA Port Overloading for No Ports; 'ha_port_overloading_for_no_ports_success': HA Port Overloading for No Ports Success; 'ha_port_overloading_for_quota_exceeded': HA Port Overloading for Quota Exceeded; 'ha_port_overloading_for_quota_exceeded_success': HA Port Overloading for Quota Exceeded Success; 'ha_port_overloading_for_quota_exceeded_race': HA Port Overloading for Quota Exceeded Race; 'ha_port_overloading_for_quota_exceeded_race_success': HA Port Overloading for Quota Exceeded Race Success; 'ha_port_overloading_for_new_user': HA Port Overloading for New User; 'ha_port_overloading_for_new_user_success': HA Port Overloading for New User Success;",
						},
						"counters3": {
							Type: schema.TypeString, Optional: true, Description: "'nat_pool_force_delete': NAT Pool Force Delete; 'quota_ext_too_many': Quota Ext Too Many; 'nat_pool_not_found_on_free': NAT Pool Not Found on Free; 'standby_class_list_drop': Standby Class-List Drop; 'fullcone_inbound_nat_pool_mismatch': Full-cone Session NAT Pool Mismatch; 'bad_ip_tot_len': Bad IPv4 Total Length; 'ip_checksum_verified': IP Checksum Verified; 'ip_checksum_fixed': IP Checksum Fixed; 'ip_checksum_bad_drop': IP Checksum Bad Drop; 'ip_frag_checksum_fixed': IP Frag Checksum Fixed; 'l4_checksum_verified': L4 Checksum Verified; 'l4_checksum_fixed': L4 Checksum Fixed; 'l4_checksum_bad_drop': L4 Checksum Bad Drop; 'jumbo_list_bad_l4_len': Jumbo List Bad L4 Len; 'frag_list_bad_l4_len': Frag List Bad L4 Len; 'nat_pool_attempt_adding_multiple_free_batches': Attempt Adding Multiple Free Batches to Quota; 'nat_pool_add_free_batch_failed': Add Batch to Quota Failed; 'mgcp_alg_no_quota': MGCP ALG User-Quota Not Found; 'mgcp_alg_quota_inc_failure': MGCP ALG User-Quota Exceeded; 'mgcp_alg_no_nat_ip': MGCP ALG Unknown NAT IP; 'mgcp_alg_reuse_fullcone': MGCP ALG Reuse Full-cone Session; 'mgcp_alg_fullcone_mismatch': MGCP ALG Full-cone Session Mismatch; 'mgcp_alg_alloc_port_failure': MGCP ALG Alloc NAT Port Failure; 'mgcp_alg_create_fullcone_failure': MGCP ALG Create Full-cone Session Failure; 'mgcp_alg_release_port_failure': MGCP ALG Release NAT Port Failure; 'mgcp_alg_single_rtp_fullcone': MGCP ALG Single RTP Full-cone Found; 'mgcp_alg_single_rtcp_fullcone': MGCP ALG Single RTCP Full-cone Found; 'mgcp_alg_rtcp_fullcone_mismatch': MGCP ALG RTCP Full-cone NAT Port Mismatch; 'mgcp_alg_reuse_rtp_rtcp_fullcone': MGCP ALG Reuse RTP/RTCP Full-cone Session; 'mgcp_alg_alloc_rtp_rtcp_port_failure': MGCP ALG Alloc RTP/RTCP NAT Ports Failure; 'mgcp_alg_alloc_single_port_failure': MGCP ALG Alloc Single RTP or RTCP NAT Port Failure; 'mgcp_alg_create_single_fullcone_failure': MGCP ALG Create Single RTP or RTCP Full-cone Session Failure; 'mgcp_alg_create_rtp_fullcone_failure': MGCP ALG Create RTP Full-cone Session Failure; 'mgcp_alg_create_rtcp_fullcone_failure': MGCP ALG Create RTCP Full-cone Session Failure; 'mgcp_alg_port_pair_alloc_from_consecutive': MGCP ALG Port Pair Allocated From Consecutive; 'mgcp_alg_port_pair_alloc_from_partition': MGCP ALG Port Pair Allocated From Partition; 'mgcp_alg_port_pair_alloc_from_pool_port_batch': MGCP ALG Port Pair Allocated From Pool Port Batch; 'mgcp_alg_port_pair_alloc_from_quota_consecutive': MGCP ALG Port Pair Allocated From Quota Consecutive; 'mgcp_alg_port_pair_alloc_from_quota_partition': MGCP ALG Port Pair Allocated From Quota Partition; 'mgcp_alg_port_pair_alloc_from_quota_partition_error': MGCP ALG Port Pair Allocated From Quota Partition Error; 'mgcp_alg_port_pair_alloc_from_quota_pool_port_batch': MGCP ALG Port Pair Allocated From Quota Pool Port Batch; 'user_quota_unusable_drop': User-Quota Unusable Drop; 'user_quota_unusable': User-Quota Marked Unusable; 'nat_pool_same_port_batch_udp_failed': Simultaneous Batch Allocation UDP Port Allocation Failed; 'fwd_ingress_packets_tcp': Forward Ingress Packets TCP; 'fwd_egress_packets_tcp': Forward Egress Packets TCP; 'rev_ingress_packets_tcp': Reverse Ingress Packets TCP; 'rev_egress_packets_tcp': Reverse Egress Packets TCP; 'fwd_ingress_bytes_tcp': Forward Ingress Bytes TCP; 'fwd_egress_bytes_tcp': Forward Egress Bytes TCP; 'rev_ingress_bytes_tcp': Reverse Ingress Bytes TCP; 'rev_egress_bytes_tcp': Reverse Egress Bytes TCP; 'fwd_ingress_packets_udp': Forward Ingress Packets UDP; 'fwd_egress_packets_udp': Forward Egress Packets UDP; 'rev_ingress_packets_udp': Reverse Ingress Packets UDP; 'rev_egress_packets_udp': Reverse Egress Packets UDP; 'fwd_ingress_bytes_udp': Forward Ingress Bytes UDP; 'fwd_egress_bytes_udp': Forward Egress Bytes UDP; 'rev_ingress_bytes_udp': Reverse Ingress Bytes UDP; 'rev_egress_bytes_udp': Reverse Egress Bytes UDP; 'fwd_ingress_packets_icmp': Forward Ingress Packets ICMP; 'fwd_egress_packets_icmp': Forward Egress Packets ICMP; 'rev_ingress_packets_icmp': Reverse Ingress Packets ICMP; 'rev_egress_packets_icmp': Reverse Egress Packets ICMP; 'fwd_ingress_bytes_icmp': Forward Ingress Bytes ICMP; 'fwd_egress_bytes_icmp': Forward Egress Bytes ICMP; 'rev_ingress_bytes_icmp': Reverse Ingress Bytes ICMP; 'rev_egress_bytes_icmp': Reverse Egress Bytes ICMP; 'fwd_ingress_packets_others': Forward Ingress Packets OTHERS; 'fwd_egress_packets_others': Forward Egress Packets OTHERS; 'rev_ingress_packets_others': Reverse Ingress Packets OTHERS; 'rev_egress_packets_others': Reverse Egress Packets OTHERS; 'fwd_ingress_bytes_others': Forward Ingress Bytes OTHERS; 'fwd_egress_bytes_others': Forward Egress Bytes OTHERS; 'rev_ingress_bytes_others': Reverse Ingress Bytes OTHERS; 'rev_egress_bytes_others': Reverse Egress Bytes OTHERS; 'fwd_ingress_pkt_size_range1': Forward Ingress Packet size between 0 and 200; 'fwd_ingress_pkt_size_range2': Forward Ingress Packet size between 201 and 800; 'fwd_ingress_pkt_size_range3': Forward Ingress Packet size between 801 and 1550; 'fwd_ingress_pkt_size_range4': Forward Ingress Packet size between 1551 and 9000; 'fwd_egress_pkt_size_range1': Forward Egress Packet size between 0 and 200; 'fwd_egress_pkt_size_range2': Forward Egress Packet size between 201 and 800; 'fwd_egress_pkt_size_range3': Forward Egress Packet size between 801 and 1550; 'fwd_egress_pkt_size_range4': Forward Egress Packet size between 1551 and 9000; 'rev_ingress_pkt_size_range1': Reverse Ingress Packet size between 0 and 200; 'rev_ingress_pkt_size_range2': Reverse Ingress Packet size between 201 and 800; 'rev_ingress_pkt_size_range3': Reverse Ingress Packet size between 801 and 1550; 'rev_ingress_pkt_size_range4': Reverse Ingress Packet size between 1551 and 9000; 'rev_egress_pkt_size_range1': Reverse Egress Packet size between 0 and 200; 'rev_egress_pkt_size_range2': Reverse Egress Packet size between 201 and 800; 'rev_egress_pkt_size_range3': Reverse Egress Packet size between 801 and 1550; 'rev_egress_pkt_size_range4': Reverse Egress Packet size between 1551 and 9000; 'prefix_quota_mismatch': Prefix Quota NAT IP Mismatch; 'port_overloading_port_tcp_inserted': Port Overloading NAT Port TCP Inserted; 'port_overloading_port_udp_inserted': Port Overloading NAT Port UDP Inserted; 'port_overloading_port_free_tcp': TCP Port Overloading NAT Port Freed; 'port_overloading_port_free_udp': UDP Port Overloading NAT Port Freed;",
						},
					},
				},
			},
			"tcp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mss_clamp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mss_clamp_type": {
										Type: schema.TypeString, Optional: true, Default: "subtract", Description: "'fixed': Specify a fixed max value for the TCP MSS; 'none': No TCP MSS clamping; 'subtract': Specify the value to subtract from the TCP MSS (default: 40);",
									},
									"mss_value": {
										Type: schema.TypeInt, Optional: true, Description: "The max value allowed for the TCP MSS (default: not configured)",
									},
									"mss_subtract": {
										Type: schema.TypeInt, Optional: true, Default: 40, Description: "Specify the value to subtract from the TCP MSS (default: 40)",
									},
									"min": {
										Type: schema.TypeInt, Optional: true, Default: 416, Description: "Specify the min value allowed for the TCP MSS (Specify the min value allowed for the TCP MSS (default: ((576 - 60 - 60 - 40))))",
									},
								},
							},
						},
						"reset_on_error": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"outbound": {
										Type: schema.TypeString, Optional: true, Description: "'disable': Disable send TCP reset on error;",
									},
								},
							},
						},
					},
				},
			},
			"user_quota_prefix_length": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "User Quota Prefix Length (Default: 128)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6DsLiteGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6DsLiteGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6DsLiteGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6DsLiteGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6DsLiteGlobalIcmp(d []interface{}) edpt.Cgnv6DsLiteGlobalIcmp {

	count1 := len(d)
	var ret edpt.Cgnv6DsLiteGlobalIcmp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnPortUnavailable = in["send_on_port_unavailable"].(string)
		ret.SendOnUserQuotaExceeded = in["send_on_user_quota_exceeded"].(string)
	}
	return ret
}

func getObjectCgnv6DsLiteGlobalInside(d []interface{}) edpt.Cgnv6DsLiteGlobalInside {

	count1 := len(d)
	var ret edpt.Cgnv6DsLiteGlobalInside
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Source = getObjectCgnv6DsLiteGlobalInsideSource(in["source"].([]interface{}))
	}
	return ret
}

func getObjectCgnv6DsLiteGlobalInsideSource(d []interface{}) edpt.Cgnv6DsLiteGlobalInsideSource {

	count1 := len(d)
	var ret edpt.Cgnv6DsLiteGlobalInsideSource
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassList = in["class_list"].(string)
	}
	return ret
}

func getSliceCgnv6DsLiteGlobalSamplingEnable(d []interface{}) []edpt.Cgnv6DsLiteGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6DsLiteGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6DsLiteGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		oi.Counters3 = in["counters3"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6DsLiteGlobalTcp(d []interface{}) edpt.Cgnv6DsLiteGlobalTcp {

	count1 := len(d)
	var ret edpt.Cgnv6DsLiteGlobalTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MssClamp = getObjectCgnv6DsLiteGlobalTcpMssClamp(in["mss_clamp"].([]interface{}))
		ret.ResetOnError = getObjectCgnv6DsLiteGlobalTcpResetOnError(in["reset_on_error"].([]interface{}))
	}
	return ret
}

func getObjectCgnv6DsLiteGlobalTcpMssClamp(d []interface{}) edpt.Cgnv6DsLiteGlobalTcpMssClamp {

	count1 := len(d)
	var ret edpt.Cgnv6DsLiteGlobalTcpMssClamp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MssClampType = in["mss_clamp_type"].(string)
		ret.MssValue = in["mss_value"].(int)
		ret.MssSubtract = in["mss_subtract"].(int)
		ret.Min = in["min"].(int)
	}
	return ret
}

func getObjectCgnv6DsLiteGlobalTcpResetOnError(d []interface{}) edpt.Cgnv6DsLiteGlobalTcpResetOnError {

	count1 := len(d)
	var ret edpt.Cgnv6DsLiteGlobalTcpResetOnError
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Outbound = in["outbound"].(string)
	}
	return ret
}

func dataToEndpointCgnv6DsLiteGlobal(d *schema.ResourceData) edpt.Cgnv6DsLiteGlobal {
	var ret edpt.Cgnv6DsLiteGlobal
	ret.Inst.Icmp = getObjectCgnv6DsLiteGlobalIcmp(d.Get("icmp").([]interface{}))
	ret.Inst.Inside = getObjectCgnv6DsLiteGlobalInside(d.Get("inside").([]interface{}))
	ret.Inst.IpChecksumError = d.Get("ip_checksum_error").(string)
	ret.Inst.L4ChecksumError = d.Get("l4_checksum_error").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6DsLiteGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Tcp = getObjectCgnv6DsLiteGlobalTcp(d.Get("tcp").([]interface{}))
	ret.Inst.UserQuotaPrefixLength = d.Get("user_quota_prefix_length").(int)
	//omit uuid
	return ret
}
