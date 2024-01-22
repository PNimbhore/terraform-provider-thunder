package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVeIpv6StatefulFirewall() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ve_ipv6_stateful_firewall`: Configure Stateful Firewall direction\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceVeIpv6StatefulFirewallCreate,
		UpdateContext: resourceInterfaceVeIpv6StatefulFirewallUpdate,
		ReadContext:   resourceInterfaceVeIpv6StatefulFirewallRead,
		DeleteContext: resourceInterfaceVeIpv6StatefulFirewallDelete,

		Schema: map[string]*schema.Schema{
			"access_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Access-list for traffic from the outside",
			},
			"acl_name": {
				Type: schema.TypeString, Optional: true, Description: "Access-list Name",
			},
			"class_list": {
				Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
			},
			"inside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inside (private) interface for stateful firewall",
			},
			"outside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outside (public) interface for stateful firewall",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceVeIpv6StatefulFirewallCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpv6StatefulFirewallCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpv6StatefulFirewall(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeIpv6StatefulFirewallRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceVeIpv6StatefulFirewallUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpv6StatefulFirewallUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpv6StatefulFirewall(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeIpv6StatefulFirewallRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceVeIpv6StatefulFirewallDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpv6StatefulFirewallDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpv6StatefulFirewall(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceVeIpv6StatefulFirewallRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpv6StatefulFirewallRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpv6StatefulFirewall(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceVeIpv6StatefulFirewall(d *schema.ResourceData) edpt.InterfaceVeIpv6StatefulFirewall {
	var ret edpt.InterfaceVeIpv6StatefulFirewall
	ret.Inst.AccessList = d.Get("access_list").(int)
	ret.Inst.AclName = d.Get("acl_name").(string)
	ret.Inst.ClassList = d.Get("class_list").(string)
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.Outside = d.Get("outside").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
