package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstDefaultIpProto() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_default_ip_proto`: DDOS IP protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstDefaultIpProtoCreate,
		UpdateContext: resourceDdosDstDefaultIpProtoUpdate,
		ReadContext:   resourceDdosDstDefaultIpProtoRead,
		DeleteContext: resourceDdosDstDefaultIpProtoDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Protocol Number",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"other": {
							Type: schema.TypeString, Optional: true, Description: "DDOS other template",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"default_address_type": {
				Type: schema.TypeString, Required: true, Description: "DefaultAddressType",
			},
		},
	}
}
func resourceDdosDstDefaultIpProtoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultIpProtoCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultIpProto(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDefaultIpProtoRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstDefaultIpProtoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultIpProtoUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultIpProto(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDefaultIpProtoRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstDefaultIpProtoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultIpProtoDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultIpProto(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstDefaultIpProtoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultIpProtoRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultIpProto(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstDefaultIpProtoTemplate(d []interface{}) edpt.DdosDstDefaultIpProtoTemplate {

	count1 := len(d)
	var ret edpt.DdosDstDefaultIpProtoTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Other = in["other"].(string)
	}
	return ret
}

func dataToEndpointDdosDstDefaultIpProto(d *schema.ResourceData) edpt.DdosDstDefaultIpProto {
	var ret edpt.DdosDstDefaultIpProto
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.Template = getObjectDdosDstDefaultIpProtoTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DefaultAddressType = d.Get("default_address_type").(string)
	return ret
}
