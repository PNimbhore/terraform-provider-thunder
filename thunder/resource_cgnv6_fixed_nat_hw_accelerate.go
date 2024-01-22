package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatHwAccelerate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_fixed_nat_hw_accelerate`: Fixed-NAT HW accelerate counters\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6FixedNatHwAccelerateCreate,
		UpdateContext: resourceCgnv6FixedNatHwAccelerateUpdate,
		ReadContext:   resourceCgnv6FixedNatHwAccelerateRead,
		DeleteContext: resourceCgnv6FixedNatHwAccelerateDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'entry-create': HW Entries Created; 'entry-create-failure': HW Entry Creation Failed; 'entry-create-fail-server-down': HW Entry Creation Failed - server down; 'entry-create-fail-max-entry': HW Entry Creation Failed - max entries exceeded; 'entry-free': HW Entries Freed; 'entry-free-opp-entry': HW Entries Freed - opposite tuple entry aged-out; 'entry-free-no-hw-prog': HW Entry Freed - no HW prog; 'entry-free-no-conn': HW Entry Freed - no matched conn; 'entry-free-no-sw-entry': HW Entry Freed - no software entry; 'entry-counter': HW Entries Count; 'entry-age-out': HW Entries Aged Out; 'entry-age-out-idle': HW Entries Aged Out - idle timeout; 'entry-age-out-tcp-fin': HW Entries Aged Out - TCP FIN; 'entry-age-out-tcp-rst': HW Entries Aged Out - TCP RST; 'entry-age-out-invalid-dst': HW Entries Aged Out - invalid dst; 'entry-force-hw-invalidate': HW Entries Force HW Invalidate; 'entry-invalidate-server-down': HW Entries Invalidate due to server down; 'tcam-create': TCAM Flows Created; 'tcam-free': TCAM Flows Freed; 'tcam-counter': TCAM Flow Count;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6FixedNatHwAccelerateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatHwAccelerateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatHwAccelerate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatHwAccelerateRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6FixedNatHwAccelerateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatHwAccelerateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatHwAccelerate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatHwAccelerateRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6FixedNatHwAccelerateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatHwAccelerateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatHwAccelerate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6FixedNatHwAccelerateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatHwAccelerateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatHwAccelerate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6FixedNatHwAccelerateSamplingEnable(d []interface{}) []edpt.Cgnv6FixedNatHwAccelerateSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatHwAccelerateSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatHwAccelerateSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatHwAccelerate(d *schema.ResourceData) edpt.Cgnv6FixedNatHwAccelerate {
	var ret edpt.Cgnv6FixedNatHwAccelerate
	ret.Inst.SamplingEnable = getSliceCgnv6FixedNatHwAccelerateSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
