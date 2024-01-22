package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtStateNextStateCase() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_state_next_state_case`: Case number\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtStateNextStateCaseCreate,
		UpdateContext: resourceSysUtStateNextStateCaseUpdate,
		ReadContext:   resourceSysUtStateNextStateCaseRead,
		DeleteContext: resourceSysUtStateNextStateCaseDelete,

		Schema: map[string]*schema.Schema{
			"action_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type: schema.TypeString, Required: true, Description: "'send': Test event; 'expect': Expected result; 'wait': Introduce a delay;",
						},
						"template": {
							Type: schema.TypeString, Optional: true, Description: "Packet template",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet drop. Only allowed for output spec",
						},
						"delay": {
							Type: schema.TypeInt, Optional: true, Description: "Delay in seconds",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"l1": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eth_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ethernet_start": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
												},
												"ethernet_end": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
												},
											},
										},
									},
									"trunk_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trunk_start": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
												},
												"trunk_end": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
												},
											},
										},
									},
									"length": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "packet length",
									},
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at L2 header",
									},
									"auto": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auto calculate pkt len",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"l2": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethertype": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "L2 frame type",
									},
									"protocol": {
										Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'arp': arp; 'ipv4': ipv4; 'ipv6': ipv6;",
									},
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "ethertype number",
									},
									"vlan": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Vlan ID on the packet. 0 is untagged",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"mac_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_dst": {
													Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
												},
												"address_type": {
													Type: schema.TypeString, Optional: true, Description: "'broadcast': broadcast; 'multicast': multicast;",
												},
												"virtual_server": {
													Type: schema.TypeString, Optional: true, Description: "vip",
												},
												"nat_pool": {
													Type: schema.TypeString, Optional: true, Description: "Nat pool",
												},
												"ethernet": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
												},
												"ve": {
													Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
												},
												"trunk": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk number",
												},
												"value": {
													Type: schema.TypeString, Optional: true, Description: "Mac Address",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"l3": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "L4 Protocol",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp;",
									},
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "protocol number",
									},
									"checksum": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"ip_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_dst": {
													Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
												},
												"ipv4_address": {
													Type: schema.TypeString, Optional: true, Description: "IP address",
												},
												"ipv6_address": {
													Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
												},
												"virtual_server": {
													Type: schema.TypeString, Optional: true, Description: "vip",
												},
												"nat_pool": {
													Type: schema.TypeString, Optional: true, Description: "Nat pool",
												},
												"ethernet": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
												},
												"ve": {
													Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
												},
												"trunk": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk number",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"tcp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_port": {
										Type: schema.TypeInt, Optional: true, Description: "Source port value",
									},
									"dest_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
									},
									"dest_port_value": {
										Type: schema.TypeInt, Optional: true, Description: "Dest port value",
									},
									"nat_pool": {
										Type: schema.TypeString, Optional: true, Description: "Nat pool port",
									},
									"seq_number": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"ack_seq_number": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"checksum": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"urgent": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"window": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"flags": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"syn": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Syn",
												},
												"ack": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ack",
												},
												"fin": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fin",
												},
												"rst": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Rst",
												},
												"psh": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Psh",
												},
												"ece": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ece",
												},
												"urg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Urg",
												},
												"cwr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cwr",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"options": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mss": {
													Type: schema.TypeInt, Optional: true, Description: "TCP MSS",
												},
												"wscale": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sack_type": {
													Type: schema.TypeString, Optional: true, Description: "'permitted': permitted; 'block': block;",
												},
												"time_stamp_enable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "adds Time Stamp to options",
												},
												"nop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "No Op",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"udp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_port": {
										Type: schema.TypeInt, Optional: true, Description: "Source port value",
									},
									"dest_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
									},
									"dest_port_value": {
										Type: schema.TypeInt, Optional: true, Description: "Dest port value",
									},
									"nat_pool": {
										Type: schema.TypeString, Optional: true, Description: "Nat pool port",
									},
									"length": {
										Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at UDP header",
									},
									"checksum": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"case_number": {
				Type: schema.TypeInt, Required: true, Description: "",
			},
			"repeat": {
				Type: schema.TypeInt, Optional: true, Description: "number of times case should be repeated before state transition",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtStateNextStateCaseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCase(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtStateNextStateCaseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCase(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtStateNextStateCaseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCase(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtStateNextStateCaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCase(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSysUtStateNextStateCaseActionList(d []interface{}) []edpt.SysUtStateNextStateCaseActionList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionList
		oi.Direction = in["direction"].(string)
		oi.Template = in["template"].(string)
		oi.Drop = in["drop"].(int)
		oi.Delay = in["delay"].(int)
		//omit uuid
		oi.L1 = getObjectSysUtStateNextStateCaseActionListL1(in["l1"].([]interface{}))
		oi.L2 = getObjectSysUtStateNextStateCaseActionListL2(in["l2"].([]interface{}))
		oi.L3 = getObjectSysUtStateNextStateCaseActionListL3(in["l3"].([]interface{}))
		oi.Tcp = getObjectSysUtStateNextStateCaseActionListTcp(in["tcp"].([]interface{}))
		oi.Udp = getObjectSysUtStateNextStateCaseActionListUdp(in["udp"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionListL1(d []interface{}) edpt.SysUtStateNextStateCaseActionListL1 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionListL1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthList = getSliceSysUtStateNextStateCaseActionListL1EthList(in["eth_list"].([]interface{}))
		ret.Trunk_list = getSliceSysUtStateNextStateCaseActionListL1Trunk_list(in["trunk_list"].([]interface{}))
		ret.Length = in["length"].(int)
		ret.Value = in["value"].(int)
		ret.Auto = in["auto"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSysUtStateNextStateCaseActionListL1EthList(d []interface{}) []edpt.SysUtStateNextStateCaseActionListL1EthList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionListL1EthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionListL1EthList
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtStateNextStateCaseActionListL1Trunk_list(d []interface{}) []edpt.SysUtStateNextStateCaseActionListL1Trunk_list {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionListL1Trunk_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionListL1Trunk_list
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionListL2(d []interface{}) edpt.SysUtStateNextStateCaseActionListL2 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionListL2
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethertype = in["ethertype"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.Value = in["value"].(int)
		ret.Vlan = in["vlan"].(int)
		//omit uuid
		ret.MacList = getSliceSysUtStateNextStateCaseActionListL2MacList(in["mac_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtStateNextStateCaseActionListL2MacList(d []interface{}) []edpt.SysUtStateNextStateCaseActionListL2MacList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionListL2MacList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionListL2MacList
		oi.SrcDst = in["src_dst"].(string)
		oi.AddressType = in["address_type"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Value = in["value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionListL3(d []interface{}) edpt.SysUtStateNextStateCaseActionListL3 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionListL3
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Protocol = in["protocol"].(int)
		ret.Type = in["type"].(string)
		ret.Value = in["value"].(int)
		ret.Checksum = in["checksum"].(string)
		ret.Ttl = in["ttl"].(int)
		//omit uuid
		ret.IpList = getSliceSysUtStateNextStateCaseActionListL3IpList(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtStateNextStateCaseActionListL3IpList(d []interface{}) []edpt.SysUtStateNextStateCaseActionListL3IpList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionListL3IpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionListL3IpList
		oi.SrcDst = in["src_dst"].(string)
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionListTcp(d []interface{}) edpt.SysUtStateNextStateCaseActionListTcp {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionListTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPort = in["src_port"].(int)
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.SeqNumber = in["seq_number"].(string)
		ret.AckSeqNumber = in["ack_seq_number"].(string)
		ret.Checksum = in["checksum"].(string)
		ret.Urgent = in["urgent"].(string)
		ret.Window = in["window"].(string)
		//omit uuid
		ret.Flags = getObjectSysUtStateNextStateCaseActionListTcpFlags(in["flags"].([]interface{}))
		ret.Options = getObjectSysUtStateNextStateCaseActionListTcpOptions(in["options"].([]interface{}))
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionListTcpFlags(d []interface{}) edpt.SysUtStateNextStateCaseActionListTcpFlags {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionListTcpFlags
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syn = in["syn"].(int)
		ret.Ack = in["ack"].(int)
		ret.Fin = in["fin"].(int)
		ret.Rst = in["rst"].(int)
		ret.Psh = in["psh"].(int)
		ret.Ece = in["ece"].(int)
		ret.Urg = in["urg"].(int)
		ret.Cwr = in["cwr"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionListTcpOptions(d []interface{}) edpt.SysUtStateNextStateCaseActionListTcpOptions {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionListTcpOptions
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mss = in["mss"].(int)
		ret.Wscale = in["wscale"].(int)
		ret.SackType = in["sack_type"].(string)
		ret.TimeStampEnable = in["time_stamp_enable"].(int)
		ret.Nop = in["nop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionListUdp(d []interface{}) edpt.SysUtStateNextStateCaseActionListUdp {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionListUdp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPort = in["src_port"].(int)
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.Length = in["length"].(int)
		ret.Checksum = in["checksum"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointSysUtStateNextStateCase(d *schema.ResourceData) edpt.SysUtStateNextStateCase {
	var ret edpt.SysUtStateNextStateCase
	ret.Inst.ActionList = getSliceSysUtStateNextStateCaseActionList(d.Get("action_list").([]interface{}))
	ret.Inst.CaseNumber = d.Get("case_number").(int)
	ret.Inst.Repeat = d.Get("repeat").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
