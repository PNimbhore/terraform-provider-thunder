package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_cgnv6_dns64_vs_port_tmpl_trigger_stats_severity`: Configure triggers severity for system object stats\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityDelete,

		Schema: map[string]*schema.Schema{
			"drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
			},
			"drop_alert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
			},
			"drop_critical": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
			},
			"drop_warning": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
			},
			"error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
			},
			"error_alert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
			},
			"error_critical": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
			},
			"error_warning": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
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
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity
	ret.Inst.Drop = d.Get("drop").(int)
	ret.Inst.DropAlert = d.Get("drop_alert").(int)
	ret.Inst.DropCritical = d.Get("drop_critical").(int)
	ret.Inst.DropWarning = d.Get("drop_warning").(int)
	ret.Inst.Error = d.Get("error").(int)
	ret.Inst.ErrorAlert = d.Get("error_alert").(int)
	ret.Inst.ErrorCritical = d.Get("error_critical").(int)
	ret.Inst.ErrorWarning = d.Get("error_warning").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
