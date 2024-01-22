package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameLevelIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_name_level_indicator`: DDoS Indicator Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoNameLevelIndicatorCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoNameLevelIndicatorUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoNameLevelIndicatorRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoNameLevelIndicatorDelete,

		Schema: map[string]*schema.Schema{
			"data_packet_size": {
				Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
			},
			"score": {
				Type: schema.TypeInt, Optional: true, Description: "Score corresponding to the indicator",
			},
			"src_threshold_large_num": {
				Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
			},
			"src_threshold_num": {
				Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
			},
			"src_threshold_str": {
				Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
			},
			"src_violation_actions": {
				Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this src indicator threshold reaches",
			},
			"type": {
				Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'frag-rate': rate of incoming fragmented packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_threshold_large_num": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
			},
			"zone_threshold_num": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
			},
			"zone_threshold_str": {
				Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
			},
			"zone_violation_actions": {
				Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this zone indicator threshold reaches",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"level_num": {
				Type: schema.TypeString, Required: true, Description: "LevelNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}
func resourceDdosDstZoneIpProtoProtoNameLevelIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameLevelIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameLevelIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNameLevelIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNameLevelIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameLevelIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameLevelIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNameLevelIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoNameLevelIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameLevelIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameLevelIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNameLevelIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameLevelIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameLevelIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneIpProtoProtoNameLevelIndicator(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameLevelIndicator {
	var ret edpt.DdosDstZoneIpProtoProtoNameLevelIndicator
	ret.Inst.DataPacketSize = d.Get("data_packet_size").(int)
	ret.Inst.Score = d.Get("score").(int)
	ret.Inst.SrcThresholdLargeNum = d.Get("src_threshold_large_num").(int)
	ret.Inst.SrcThresholdNum = d.Get("src_threshold_num").(int)
	ret.Inst.SrcThresholdStr = d.Get("src_threshold_str").(string)
	ret.Inst.SrcViolationActions = d.Get("src_violation_actions").(string)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneThresholdLargeNum = d.Get("zone_threshold_large_num").(int)
	ret.Inst.ZoneThresholdNum = d.Get("zone_threshold_num").(int)
	ret.Inst.ZoneThresholdStr = d.Get("zone_threshold_str").(string)
	ret.Inst.ZoneViolationActions = d.Get("zone_violation_actions").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.LevelNum = d.Get("level_num").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	return ret
}
