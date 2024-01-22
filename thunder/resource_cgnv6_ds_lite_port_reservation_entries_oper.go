package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DsLitePortReservationEntriesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_ds_lite_port_reservation_entries_oper`: Operational Status for the object port-reservation-entries\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6DsLitePortReservationEntriesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tunnel_src_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"tunnel_dest_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_start_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inside_end_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_start_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_end_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6DsLitePortReservationEntriesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLitePortReservationEntriesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLitePortReservationEntriesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6DsLitePortReservationEntriesOperOper := setObjectCgnv6DsLitePortReservationEntriesOperOper(res)
		d.Set("oper", Cgnv6DsLitePortReservationEntriesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6DsLitePortReservationEntriesOperOper(ret edpt.DataCgnv6DsLitePortReservationEntriesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list":  setSliceCgnv6DsLitePortReservationEntriesOperOperEntryList(ret.DtCgnv6DsLitePortReservationEntriesOper.Oper.EntryList),
			"entry_count": ret.DtCgnv6DsLitePortReservationEntriesOper.Oper.EntryCount,
		},
	}
}

func setSliceCgnv6DsLitePortReservationEntriesOperOperEntryList(d []edpt.Cgnv6DsLitePortReservationEntriesOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["tunnel_src_address"] = item.TunnelSrcAddress
		in["tunnel_dest_address"] = item.TunnelDestAddress
		in["inside_address"] = item.InsideAddress
		in["inside_start_port"] = item.InsideStartPort
		in["inside_end_port"] = item.InsideEndPort
		in["nat_address"] = item.NatAddress
		in["nat_start_port"] = item.NatStartPort
		in["nat_end_port"] = item.NatEndPort
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6DsLitePortReservationEntriesOperOper(d []interface{}) edpt.Cgnv6DsLitePortReservationEntriesOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6DsLitePortReservationEntriesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceCgnv6DsLitePortReservationEntriesOperOperEntryList(in["entry_list"].([]interface{}))
		ret.EntryCount = in["entry_count"].(int)
	}
	return ret
}

func getSliceCgnv6DsLitePortReservationEntriesOperOperEntryList(d []interface{}) []edpt.Cgnv6DsLitePortReservationEntriesOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6DsLitePortReservationEntriesOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6DsLitePortReservationEntriesOperOperEntryList
		oi.TunnelSrcAddress = in["tunnel_src_address"].(string)
		oi.TunnelDestAddress = in["tunnel_dest_address"].(string)
		oi.InsideAddress = in["inside_address"].(string)
		oi.InsideStartPort = in["inside_start_port"].(int)
		oi.InsideEndPort = in["inside_end_port"].(int)
		oi.NatAddress = in["nat_address"].(string)
		oi.NatStartPort = in["nat_start_port"].(int)
		oi.NatEndPort = in["nat_end_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6DsLitePortReservationEntriesOper(d *schema.ResourceData) edpt.Cgnv6DsLitePortReservationEntriesOper {
	var ret edpt.Cgnv6DsLitePortReservationEntriesOper

	ret.Oper = getObjectCgnv6DsLitePortReservationEntriesOperOper(d.Get("oper").([]interface{}))
	return ret
}
