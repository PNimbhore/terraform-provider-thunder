package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDst() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_class_list_app_type_src_dst`: DDOS APP type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstDelete,

		Schema: map[string]*schema.Schema{
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_l4": {
							Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
						},
						"dns": {
							Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
						},
						"http": {
							Type: schema.TypeString, Optional: true, Description: "DDOS http template",
						},
						"sip": {
							Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
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
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "ClassListName",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListAppTypeSrcDst(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListAppTypeSrcDst(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListAppTypeSrcDst(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListAppTypeSrcDstRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListAppTypeSrcDst(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstEntrySrcDstPairClassListAppTypeSrcDstTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListAppTypeSrcDstTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListAppTypeSrcDstTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func dataToEndpointDdosDstEntrySrcDstPairClassListAppTypeSrcDst(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairClassListAppTypeSrcDst {
	var ret edpt.DdosDstEntrySrcDstPairClassListAppTypeSrcDst
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Template = getObjectDdosDstEntrySrcDstPairClassListAppTypeSrcDstTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
