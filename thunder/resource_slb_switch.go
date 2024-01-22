package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSwitch() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_switch`: Configure slb switch\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSwitchCreate,
		UpdateContext: resourceSlbSwitchUpdate,
		ReadContext:   resourceSlbSwitchRead,
		DeleteContext: resourceSlbSwitchDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'fwlb': FWLB; 'licexpire_drop': License Expire Drop; 'bwl_drop': BW Limit Drop; 'rx_kernel': Received kernel; 'rx_arp_req': ARP REQ Rcvd; 'rx_arp_resp': ARP RESP Rcvd; 'vlan_flood': VLAN Flood; 'l2_def_vlan_drop': L2 Default Vlan FWD Drop; 'ipv4_noroute_drop': IPv4 No Route Drop; 'ipv6_noroute_drop': IPv6 No Route Drop; 'prot_down_drop': Prot Down Drop; 'l2_forward': L2 Forward; 'l3_forward_ip': L3 IP Forward; 'l3_forward_ipv6': L3 IPv6 Forward; 'l4_process': L4 Process; 'unknown_prot_drop': Unknown Prot Drop; 'ttl_exceeded_drop': TTL Exceeded Drop; 'linkdown_drop': Link Down Drop; 'sport_drop': SPORT Drop; 'incorrect_len_drop': Incorrect Length Drop; 'ip_defrag': IP Defrag; 'acl_deny': ACL Denys; 'ipfrag_tcp': IP(TCP) Fragment Rcvd; 'ipfrag_overlap': IP Fragment Overlap; 'ipfrag_timeout': IP Fragment Timeout; 'ipfrag_overload': IP Frag Overload Drops; 'ipfrag_reasmoks': IP Fragment Reasm OKs; 'ipfrag_reasmfails': IP Fragment Reasm Fails; 'land_drop': Anomaly Land Attack Drop; 'ipoptions_drop': Anomaly IP OPT Drops; 'badpkt_drop': Bad Pkt Drop; 'pingofdeath_drop': Anomaly PingDeath Drop; 'allfrag_drop': Anomaly All Frag Drop; 'tcpnoflag_drop': Anomaly TCP noFlag Drop; 'tcpsynfrag_drop': Anomaly SYN Frag Drop; 'tcpsynfin_drop': Anomaly TCP SYNFIN Drop; 'ipsec_drop': IPSec Drop; 'bpdu_rcvd': BPDUs Received; 'bpdu_sent': BPDUs Sent; 'ctrl_syn_rate_drop': SYN rate exceeded Drop; 'ip_defrag_invalid_len': IP Invalid Length Frag; 'ipv4_frag_6rd_ok': IPv4 Frag 6RD OK; 'ipv4_frag_6rd_drop': IPv4 Frag 6RD Dropped; 'no_ip_drop': No IP Drop; 'ipv6frag_udp': IPv6 Frag UDP; 'ipv6frag_udp_dropped': IPv6 Frag UDP Dropped; 'ipv6frag_tcp_dropped': IPv6 Frag TCP Dropped; 'ipv6frag_ipip_ok': IPv6 Frag IPIP OKs; 'ipv6frag_ipip_dropped': IPv6 Frag IPIP Drop; 'ip_frag_oversize': IP Fragment oversize; 'ip_frag_too_many': IP Fragment too many; 'ipv4_novlanfwd_drop': IPv4 No L3 VLAN FWD Drop; 'ipv6_novlanfwd_drop': IPv6 No L3 VLAN FWD Drop; 'fpga_error_pkt1': FPGA Error PKT1; 'fpga_error_pkt2': FPGA Error PKT2; 'max_arp_drop': Max ARP Drop; 'ipv6frag_tcp': IPv6 Frag TCP; 'ipv6frag_icmp': IPv6 Frag ICMP; 'ipv6frag_ospf': IPv6 Frag OSPF; 'ipv6frag_esp': IPv6 Frag ESP; 'l4_in_ctrl_cpu': L4 In Ctrl CPU; 'mgmt_svc_drop': Management Service Drop; 'jumbo_frag_drop': Jumbo Frag Drop; 'ipv6_jumbo_frag_drop': IPv6 Jumbo Frag Drop; 'ipipv6_jumbo_frag_drop': IPIPv6 Jumbo Frag Drop; 'ipv6_ndisc_dad_solicits': IPv6 DAD on Solicits; 'ipv6_ndisc_dad_adverts': IPv6 DAD on Adverts; 'ipv6_ndisc_mac_changes': IPv6 DAD MAC Changed; 'ipv6_ndisc_out_of_memory': IPv6 DAD Out-of-memory; 'sp_non_ctrl_pkt_drop': Shared IP mode non ctrl packet to linux drop; 'urpf_pkt_drop': URPF check packet drop; 'fw_smp_zone_mismatch': FW SMP Zone Mismatch; 'ipfrag_udp': IP(UDP) Fragment Rcvd; 'ipfrag_icmp': IP(ICMP) Fragment Rcvd; 'ipfrag_ospf': IP(OSPF) Fragment Rcvd; 'ipfrag_esp': IP(ESP) Fragment Rcvd; 'ipfrag_tcp_dropped': IP Frag TCP Dropped; 'ipfrag_udp_dropped': IP Frag UDP Dropped; 'ipfrag_ipip_dropped': IP Frag IPIP Drop; 'redirect_fwd_fail': Redirect failed in the fwd direction; 'redirect_fwd_sent': Redirect succeeded in the fwd direction; 'redirect_rev_fail': Redirect failed in the rev direction; 'redirect_rev_sent': Redirect succeeded in the rev direction; 'redirect_setup_fail': Redirect connection setup failed; 'ip_frag_sent': IP frag sent; 'invalid_rx_arp_pkt': Invalid ARP PKT Rcvd; 'invalid_sender_mac_arp_drop': ARP PKT dropped due to invalid sender MAC; 'dev_based_arp_drop': ARP PKT dropped due to interface state checks; 'scaleout_arp_drop': ARP PKT dropped due to scaleout checks; 'virtual_ip_not_found_arp_drop': ARP PKT dropped due to virtual IP not found; 'inactive_static_nat_pool_arp_drop': ARP PKT dropped due to inactive static nat pool; 'inactive_nat_pool_arp_drop': ARP PKT dropped due to inactive nat pool; 'scaleout_hairpin_arp_drop': ARP PKT dropped due to scaleout hairpin checks; 'self_grat_arp_drop': Self generated grat ARP PKT dropped; 'self_grat_nat_ip_arp_drop': Self generated grat ARP PKT dropped for NAT IP; 'ip_not_found_arp_drop': ARP PKT dropped due to IP not found; 'dev_link_down_arp_drop': ARP PKT dropped due to interface is down; 'lacp_tx_intf_err_drop': LACP interface error corrected; 'service_chain_sent': Service Chain Packets Sent; 'service_chain_rcvd': Service Chain Packets Rcvd; 'unnumbered_nat_error': Unnumbered NAT error; 'unnumbered_unsupported_drop': Unsupported protocol for unnumbered; 'ipv6frag_gre_dropped': IPv6 Frag gre Drop; 'ipv6_ndisc_dad_prefix_mismatch_drop': IPv6 DAD on Advertise drop for prefix mismatch; 'bw_ignore_limit': BW Limit ignored packets count; 'ppsl_drop_egr': Packet-Per-Sec Limit Drop at egress; 'ppsl_drop_ing': Packet-Per-Sec Limit Drop at ingress; 'ppsl_ignore_limit': Packet-Per-Sec Limit ignored packets count; 'closed_port_syn_drop': Linux Closed Port SYN Drop; 'tls13_drop_req': TLS13-Request-Per-Sec Limit Drop at ingress; 'tls13_ignore_req': TLS13-Request-Per-Sec Limit ignored packets count; 'tls12_drop_req': TLS12-Request-Per-Sec Limit Drop at ingress; 'tls12_ignore_req': TLS12-Request-Per-Sec Limit ignored packets count; 'tls12_tls13_drop_req': TLS12--TLS13-Request-Per-Sec Limit Drop at ingress; 'tls12_tls13_ignore_req': TLS12-TLS13-Request-Per-Sec Limit ignored packets count;",
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
func resourceSlbSwitchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSwitchCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSwitch(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSwitchRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSwitchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSwitchUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSwitch(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSwitchRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSwitchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSwitchDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSwitch(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSwitchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSwitchRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSwitch(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbSwitchSamplingEnable(d []interface{}) []edpt.SlbSwitchSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbSwitchSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSwitchSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSwitch(d *schema.ResourceData) edpt.SlbSwitch {
	var ret edpt.SlbSwitch
	ret.Inst.SamplingEnable = getSliceSlbSwitchSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
