package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleListDefaultOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_rule_list_default_oper`: Operational Status for the object default\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnRuleListDefaultOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"proto": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"start_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"end_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dscp_match": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"action_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dnat_domain_unresolved_drops": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_list": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_list": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"no_snat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pool": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"pool_shared": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_alg": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"timeout_val": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fast": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dscp_direction": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dscp_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"rule_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceCgnv6LsnRuleListDefaultOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDefaultOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDefaultOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnRuleListDefaultOperOper := setObjectCgnv6LsnRuleListDefaultOperOper(res)
		d.Set("oper", Cgnv6LsnRuleListDefaultOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnRuleListDefaultOperOper(ret edpt.DataCgnv6LsnRuleListDefaultOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list":  setSliceCgnv6LsnRuleListDefaultOperOperRuleList(ret.DtCgnv6LsnRuleListDefaultOper.Oper.RuleList),
			"rule_count": ret.DtCgnv6LsnRuleListDefaultOper.Oper.RuleCount,
		},
	}
}

func setSliceCgnv6LsnRuleListDefaultOperOperRuleList(d []edpt.Cgnv6LsnRuleListDefaultOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["hits"] = item.Hits
		in["proto"] = item.Proto
		in["start_port"] = item.StartPort
		in["end_port"] = item.EndPort
		in["dscp_match"] = item.DscpMatch
		in["action"] = item.Action
		in["action_type"] = item.ActionType
		in["domain_name"] = item.DomainName
		in["dnat_domain_unresolved_drops"] = item.DnatDomainUnresolvedDrops
		in["ipv4_list"] = item.Ipv4List
		in["port_list"] = item.PortList
		in["no_snat"] = item.NoSnat
		in["vrid"] = item.Vrid
		in["pool"] = item.Pool
		in["pool_shared"] = item.PoolShared
		in["http_alg"] = item.HttpAlg
		in["timeout_val"] = item.TimeoutVal
		in["fast"] = item.Fast
		in["dscp_direction"] = item.DscpDirection
		in["dscp_value"] = item.DscpValue
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6LsnRuleListDefaultOperOper(d []interface{}) edpt.Cgnv6LsnRuleListDefaultOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefaultOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceCgnv6LsnRuleListDefaultOperOperRuleList(in["rule_list"].([]interface{}))
		ret.RuleCount = in["rule_count"].(int)
	}
	return ret
}

func getSliceCgnv6LsnRuleListDefaultOperOperRuleList(d []interface{}) []edpt.Cgnv6LsnRuleListDefaultOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDefaultOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDefaultOperOperRuleList
		oi.Hits = in["hits"].(int)
		oi.Proto = in["proto"].(string)
		oi.StartPort = in["start_port"].(int)
		oi.EndPort = in["end_port"].(int)
		oi.DscpMatch = in["dscp_match"].(string)
		oi.Action = in["action"].(string)
		oi.ActionType = in["action_type"].(string)
		oi.DomainName = in["domain_name"].(string)
		oi.DnatDomainUnresolvedDrops = in["dnat_domain_unresolved_drops"].(int)
		oi.Ipv4List = in["ipv4_list"].(string)
		oi.PortList = in["port_list"].(string)
		oi.NoSnat = in["no_snat"].(int)
		oi.Vrid = in["vrid"].(int)
		oi.Pool = in["pool"].(string)
		oi.PoolShared = in["pool_shared"].(int)
		oi.HttpAlg = in["http_alg"].(string)
		oi.TimeoutVal = in["timeout_val"].(int)
		oi.Fast = in["fast"].(int)
		oi.DscpDirection = in["dscp_direction"].(string)
		oi.DscpValue = in["dscp_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnRuleListDefaultOper(d *schema.ResourceData) edpt.Cgnv6LsnRuleListDefaultOper {
	var ret edpt.Cgnv6LsnRuleListDefaultOper

	ret.Oper = getObjectCgnv6LsnRuleListDefaultOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
