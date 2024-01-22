package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4Neighbor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv4_flowspec_neighbor_ipv4_neighbor`: Specify a peer-group neighbor router\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborDelete,

		Schema: map[string]*schema.Schema{
			"activate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
			},
			"neighbor_ipv4": {
				Type: schema.TypeString, Required: true, Description: "Neighbor address",
			},
			"neighbor_route_map_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nbr_route_map": {
							Type: schema.TypeString, Optional: true, Description: "Apply route map to neighbor (Name of route map)",
						},
						"nbr_rmap_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
						},
					},
				},
			},
			"send_community_val": {
				Type: schema.TypeString, Optional: true, Default: "both", Description: "'all': Send Standard, Extended, and Large Community attributes; 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes; 'large': Send Large Community attributes;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"as_number": {
				Type: schema.TypeString, Required: true, Description: "AsNumber",
			},
		},
	}
}
func resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4Neighbor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4Neighbor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4Neighbor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4Neighbor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4Neighbor(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4Neighbor {
	var ret edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4Neighbor
	ret.Inst.Activate = d.Get("activate").(int)
	ret.Inst.NeighborIpv4 = d.Get("neighbor_ipv4").(string)
	ret.Inst.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborNeighborRouteMapLists(d.Get("neighbor_route_map_lists").([]interface{}))
	ret.Inst.SendCommunityVal = d.Get("send_community_val").(string)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
