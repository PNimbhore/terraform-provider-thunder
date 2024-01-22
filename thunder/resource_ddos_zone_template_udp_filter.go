package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateUdpFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_udp_filter`: UDP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateUdpFilterCreate,
		UpdateContext: resourceDdosZoneTemplateUdpFilterUpdate,
		ReadContext:   resourceDdosZoneTemplateUdpFilterRead,
		DeleteContext: resourceDdosZoneTemplateUdpFilterDelete,

		Schema: map[string]*schema.Schema{
			"byte_offset_filter": {
				Type: schema.TypeString, Optional: true, Description: "Filter using Berkeley Packet Filter syntax",
			},
			"udp_filter_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src;",
			},
			"udp_filter_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"udp_filter_inverse_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of the matching",
			},
			"udp_filter_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"udp_filter_regex": {
				Type: schema.TypeString, Optional: true, Description: "Regex Expression",
			},
			"udp_filter_seq": {
				Type: schema.TypeInt, Optional: true, Description: "Sequence number",
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
func resourceDdosZoneTemplateUdpFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateUdpFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateUdpFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateUdpFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateUdpFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateUdpFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateUdpFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateUdpFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateUdpFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateUdpFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateUdpFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateUdpFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateUdpFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateUdpFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateUdpFilter(d *schema.ResourceData) edpt.DdosZoneTemplateUdpFilter {
	var ret edpt.DdosZoneTemplateUdpFilter
	ret.Inst.ByteOffsetFilter = d.Get("byte_offset_filter").(string)
	ret.Inst.UdpFilterAction = d.Get("udp_filter_action").(string)
	ret.Inst.UdpFilterActionListName = d.Get("udp_filter_action_list_name").(string)
	ret.Inst.UdpFilterInverseMatch = d.Get("udp_filter_inverse_match").(int)
	ret.Inst.UdpFilterName = d.Get("udp_filter_name").(string)
	ret.Inst.UdpFilterRegex = d.Get("udp_filter_regex").(string)
	ret.Inst.UdpFilterSeq = d.Get("udp_filter_seq").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
