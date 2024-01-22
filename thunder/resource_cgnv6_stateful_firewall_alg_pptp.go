package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgPptp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_alg_pptp`: Configure PPTP ALG for NAT stateful firewall (default: enabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallAlgPptpCreate,
		UpdateContext: resourceCgnv6StatefulFirewallAlgPptpUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallAlgPptpRead,
		DeleteContext: resourceCgnv6StatefulFirewallAlgPptpDelete,

		Schema: map[string]*schema.Schema{
			"pptp_value": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable ALG;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'calls-established': Calls Established; 'call-req-pns-call-id-mismatch': Call ID Mismatch on Call Request; 'call-reply-pns-call-id-mismatch': Call ID Mismatch on Call Reply; 'gre-session-created': GRE Session Created; 'gre-session-freed': GRE Session Freed; 'call-req-retransmit': Call Request Retransmit; 'call-req-new': Call Request New; 'call-req-ext-alloc-failure': Call Request Ext Alloc Failure; 'call-reply-call-id-unknown': Call Reply Unknown Client Call ID; 'call-reply-retransmit': Call Reply Retransmit; 'call-reply-ext-ext-alloc-failure': Call Request Ext Alloc Failure; 'smp-app-type-mismatch': SMP App Type Mismatch; 'smp-client-call-id-mismatch': SMP Client Call ID Mismatch; 'smp-sessions-created': SMP Session Created; 'smp-sessions-freed': SMP Session Freed; 'smp-alloc-failure': SMP Session Alloc Failure; 'gre-conn-creation-failure': GRE Conn Alloc Failure; 'gre-conn-ext-creation-failure': GRE Conn Ext Alloc Failure; 'gre-no-fwd-route': GRE No Fwd Route; 'gre-no-rev-route': GRE No Rev Route; 'gre-no-control-conn': GRE No Control Conn; 'gre-conn-already-exists': GRE Conn Already Exists; 'gre-free-no-ext': GRE Free No Ext; 'gre-free-no-smp': GRE Free No SMP; 'gre-free-smp-app-type-mismatch': GRE Free SMP App Type Mismatch; 'control-freed': Control Session Freed; 'control-free-no-ext': Control Free No Ext; 'control-free-no-smp': Control Free No SMP; 'control-free-smp-app-type-mismatch': Control Free SMP App Type Mismatch;",
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
func resourceCgnv6StatefulFirewallAlgPptpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgPptpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgPptp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgPptpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgPptpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgPptpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgPptp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgPptpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallAlgPptpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgPptpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgPptp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgPptpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgPptpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgPptp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6StatefulFirewallAlgPptpSamplingEnable(d []interface{}) []edpt.Cgnv6StatefulFirewallAlgPptpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6StatefulFirewallAlgPptpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6StatefulFirewallAlgPptpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallAlgPptp(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgPptp {
	var ret edpt.Cgnv6StatefulFirewallAlgPptp
	ret.Inst.PptpValue = d.Get("pptp_value").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6StatefulFirewallAlgPptpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
