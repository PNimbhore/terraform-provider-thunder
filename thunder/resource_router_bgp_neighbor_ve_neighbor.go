package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNeighborVeNeighbor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_neighbor_ve_neighbor`: Specify a VE unnumbered neighbor\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpNeighborVeNeighborCreate,
		UpdateContext: resourceRouterBgpNeighborVeNeighborUpdate,
		ReadContext:   resourceRouterBgpNeighborVeNeighborRead,
		DeleteContext: resourceRouterBgpNeighborVeNeighborDelete,

		Schema: map[string]*schema.Schema{
			"peer_group_name": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"unnumbered": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve": {
				Type: schema.TypeInt, Required: true, Description: "Virtual ethernet interface number",
			},
			"as_number": {
				Type: schema.TypeString, Required: true, Description: "AsNumber",
			},
		},
	}
}
func resourceRouterBgpNeighborVeNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborVeNeighborCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborVeNeighbor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborVeNeighborRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpNeighborVeNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborVeNeighborUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborVeNeighbor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborVeNeighborRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpNeighborVeNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborVeNeighborDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborVeNeighbor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpNeighborVeNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborVeNeighborRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborVeNeighbor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpNeighborVeNeighbor(d *schema.ResourceData) edpt.RouterBgpNeighborVeNeighbor {
	var ret edpt.RouterBgpNeighborVeNeighbor
	ret.Inst.PeerGroupName = d.Get("peer_group_name").(string)
	ret.Inst.Unnumbered = d.Get("unnumbered").(int)
	//omit uuid
	ret.Inst.Ve = d.Get("ve").(int)
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
