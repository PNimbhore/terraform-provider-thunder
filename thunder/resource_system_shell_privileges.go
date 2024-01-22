package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemShellPrivileges() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_shell_privileges`: Configuration for special shell privileges\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemShellPrivilegesCreate,
		UpdateContext: resourceSystemShellPrivilegesUpdate,
		ReadContext:   resourceSystemShellPrivilegesRead,
		DeleteContext: resourceSystemShellPrivilegesDelete,

		Schema: map[string]*schema.Schema{
			"enable_shell_privileges": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable the shell privileges for a given customer",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemShellPrivilegesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemShellPrivilegesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemShellPrivileges(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemShellPrivilegesRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemShellPrivilegesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemShellPrivilegesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemShellPrivileges(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemShellPrivilegesRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemShellPrivilegesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemShellPrivilegesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemShellPrivileges(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemShellPrivilegesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemShellPrivilegesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemShellPrivileges(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemShellPrivileges(d *schema.ResourceData) edpt.SystemShellPrivileges {
	var ret edpt.SystemShellPrivileges
	ret.Inst.EnableShellPrivileges = d.Get("enable_shell_privileges").(int)
	//omit uuid
	return ret
}
