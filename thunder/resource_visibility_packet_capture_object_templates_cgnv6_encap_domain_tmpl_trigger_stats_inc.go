package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"inbound_addr_port_validation_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Destination Address Port Validation Failed",
			},
			"inbound_dest_unreachable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv6 Destination Address Unreachable",
			},
			"inbound_rev_lookup_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Reverse Route Lookup Failed",
			},
			"interface_not_configured": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Interfaces not Configured Dropped",
			},
			"outbound_addr_validation_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Source Address Validation Failed",
			},
			"outbound_dest_unreachable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv4 Destination Address Unreachable",
			},
			"outbound_rev_lookup_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Reverse Route Lookup Failed",
			},
			"packet_mtu_exceeded": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Exceeded MTU",
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
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc
	ret.Inst.Inbound_addr_port_validation_failed = d.Get("inbound_addr_port_validation_failed").(int)
	ret.Inst.Inbound_dest_unreachable = d.Get("inbound_dest_unreachable").(int)
	ret.Inst.Inbound_rev_lookup_failed = d.Get("inbound_rev_lookup_failed").(int)
	ret.Inst.Interface_not_configured = d.Get("interface_not_configured").(int)
	ret.Inst.Outbound_addr_validation_failed = d.Get("outbound_addr_validation_failed").(int)
	ret.Inst.Outbound_dest_unreachable = d.Get("outbound_dest_unreachable").(int)
	ret.Inst.Outbound_rev_lookup_failed = d.Get("outbound_rev_lookup_failed").(int)
	ret.Inst.Packet_mtu_exceeded = d.Get("packet_mtu_exceeded").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
