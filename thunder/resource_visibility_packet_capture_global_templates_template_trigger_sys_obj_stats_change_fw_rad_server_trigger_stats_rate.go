package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_rad_server_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"ha_standby_dropped": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Standby Dropped",
			},
			"invalid_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Radius Request has Invalid Key Field",
			},
			"ipv6_prefix_length_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Framed IPV6 Prefix Length Mismatch",
			},
			"radius_request_dropped": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Malformed Packet)",
			},
			"radius_table_full": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Table Full)",
			},
			"request_bad_secret_dropped": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Bad Secret Dropped",
			},
			"request_ignored": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Table Full Dropped",
			},
			"request_malformed_dropped": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Malformed Dropped",
			},
			"request_no_key_vap_dropped": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request No Key Attribute Dropped",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.HaStandbyDropped = d.Get("ha_standby_dropped").(int)
	ret.Inst.InvalidKey = d.Get("invalid_key").(int)
	ret.Inst.Ipv6PrefixLengthMismatch = d.Get("ipv6_prefix_length_mismatch").(int)
	ret.Inst.RadiusRequestDropped = d.Get("radius_request_dropped").(int)
	ret.Inst.RadiusTableFull = d.Get("radius_table_full").(int)
	ret.Inst.RequestBadSecretDropped = d.Get("request_bad_secret_dropped").(int)
	ret.Inst.RequestIgnored = d.Get("request_ignored").(int)
	ret.Inst.RequestMalformedDropped = d.Get("request_malformed_dropped").(int)
	ret.Inst.RequestNoKeyVapDropped = d.Get("request_no_key_vap_dropped").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
