package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgMgcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_alg_mgcp`: Change LSN MGCP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnAlgMgcpCreate,
		UpdateContext: resourceCgnv6LsnAlgMgcpUpdate,
		ReadContext:   resourceCgnv6LsnAlgMgcpRead,
		DeleteContext: resourceCgnv6LsnAlgMgcpDelete,

		Schema: map[string]*schema.Schema{
			"mgcp_value": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable MGCP ALG for LSN;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'auep': MGCP AUEP; 'aucx': MGCP AUCX; 'crcx': MGCP CRCX; 'dlcx': MGCP DLCX; 'epcf': MGCP EPCF; 'mdcx': MGCP MDCX; 'ntfy': MGCP NTFY; 'rqnt': MGCP RQNT; 'rsip': MGCP RSIP; 'parse-error': MGCP Message Parse Error; 'conn-ext-creation-failure': MGCP Create Connection Extension Failure; 'third-party-sdp': MGCP Third-Party SDP; 'sdp-process-candidate-failure': MGCP Operate SDP Media Candidate Attribute Failure; 'sdp-op-failure': MGCP Operate SDP Failure; 'sdp-alloc-port-map-success': MGCP Alloc SDP Port Map Success; 'sdp-alloc-port-map-failure': MGCP Alloc SDP Port Map Failure; 'modify-failure': MGCP Message Modify Failure; 'rewrite-failure': MGCP Message Rewrite Failure; 'tcp-out-of-order-drop': TCP Out-of-Order Drop;",
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
func resourceCgnv6LsnAlgMgcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgMgcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgMgcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgMgcpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnAlgMgcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgMgcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgMgcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgMgcpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnAlgMgcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgMgcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgMgcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnAlgMgcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgMgcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgMgcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnAlgMgcpSamplingEnable(d []interface{}) []edpt.Cgnv6LsnAlgMgcpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnAlgMgcpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnAlgMgcpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgMgcp(d *schema.ResourceData) edpt.Cgnv6LsnAlgMgcp {
	var ret edpt.Cgnv6LsnAlgMgcp
	ret.Inst.MgcpValue = d.Get("mgcp_value").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6LsnAlgMgcpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
