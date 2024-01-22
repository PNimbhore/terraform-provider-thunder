package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_nat64_global_trigger_stats_inc`: Configure stats to trigger packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"eif_limit_exceeded": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Endpoint-Independent Filtering Inbound Limit Exceeded",
			},
			"fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
			},
			"fullcone_self_hairpinning_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
			},
			"ha_nat_pool_batch_type_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Batch Type Mismatch",
			},
			"ha_nat_pool_unusable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Unusable",
			},
			"nat_pool_unusable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT Pool Unusable",
			},
			"nat_port_unavailable_icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP NAT Port Unavailable",
			},
			"nat_port_unavailable_tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP NAT Port Unavailable",
			},
			"nat_port_unavailable_udp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP NAT Port Unavailable",
			},
			"new_user_resource_unavailable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for New User NAT Resource Unavailable",
			},
			"no_class_list_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Class-List Match",
			},
			"no_radius_profile_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No RADIUS Profile Match",
			},
			"user_quota_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
			},
			"user_quota_unusable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Marked Unusable",
			},
			"user_quota_unusable_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Unusable Drop",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc
	ret.Inst.Eif_limit_exceeded = d.Get("eif_limit_exceeded").(int)
	ret.Inst.Fullcone_failure = d.Get("fullcone_failure").(int)
	ret.Inst.Fullcone_self_hairpinning_drop = d.Get("fullcone_self_hairpinning_drop").(int)
	ret.Inst.Ha_nat_pool_batch_type_mismatch = d.Get("ha_nat_pool_batch_type_mismatch").(int)
	ret.Inst.Ha_nat_pool_unusable = d.Get("ha_nat_pool_unusable").(int)
	ret.Inst.Nat_pool_unusable = d.Get("nat_pool_unusable").(int)
	ret.Inst.Nat_port_unavailable_icmp = d.Get("nat_port_unavailable_icmp").(int)
	ret.Inst.Nat_port_unavailable_tcp = d.Get("nat_port_unavailable_tcp").(int)
	ret.Inst.Nat_port_unavailable_udp = d.Get("nat_port_unavailable_udp").(int)
	ret.Inst.New_user_resource_unavailable = d.Get("new_user_resource_unavailable").(int)
	ret.Inst.No_class_list_match = d.Get("no_class_list_match").(int)
	ret.Inst.No_radius_profile_match = d.Get("no_radius_profile_match").(int)
	ret.Inst.User_quota_failure = d.Get("user_quota_failure").(int)
	ret.Inst.User_quota_unusable = d.Get("user_quota_unusable").(int)
	ret.Inst.User_quota_unusable_drop = d.Get("user_quota_unusable_drop").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
