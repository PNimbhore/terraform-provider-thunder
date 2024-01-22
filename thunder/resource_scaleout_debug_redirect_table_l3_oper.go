package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugRedirectTableL3Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_redirect_table_l3_oper`: Operational Status for the object l3\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugRedirectTableL3OperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"part_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"device_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"red_follow_shared": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"red_table_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id1": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dir": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dst_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"total_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"mac": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"real_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"count_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dst_index": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"dst_index_valid": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceScaleoutDebugRedirectTableL3OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugRedirectTableL3OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugRedirectTableL3Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugRedirectTableL3OperOper := setObjectScaleoutDebugRedirectTableL3OperOper(res)
		d.Set("oper", ScaleoutDebugRedirectTableL3OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugRedirectTableL3OperOper(ret edpt.DataScaleoutDebugRedirectTableL3Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"part_name":   ret.DtScaleoutDebugRedirectTableL3Oper.Oper.Part_name,
			"device_list": setSliceScaleoutDebugRedirectTableL3OperOperDevice_list(ret.DtScaleoutDebugRedirectTableL3Oper.Oper.Device_list),
		},
	}
}

func setSliceScaleoutDebugRedirectTableL3OperOperDevice_list(d []edpt.ScaleoutDebugRedirectTableL3OperOperDevice_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["device_id"] = item.Device_id
		in["red_follow_shared"] = item.Red_follow_shared
		in["red_table_list"] = setSliceScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list(item.Red_table_list)
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list(d []edpt.ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["id1"] = item.Id1
		in["dir"] = item.Dir
		in["dst_ip"] = item.Dst_ip
		in["total_count"] = item.Total_count
		in["mac"] = item.Mac
		in["real_port"] = item.Real_port
		in["count_list"] = setSliceScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list(item.Count_list)
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list(d []edpt.ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_index"] = item.Dst_index
		in["dst_index_valid"] = item.Dst_index_valid
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugRedirectTableL3OperOper(d []interface{}) edpt.ScaleoutDebugRedirectTableL3OperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugRedirectTableL3OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Part_name = in["part_name"].(string)
		ret.Device_list = getSliceScaleoutDebugRedirectTableL3OperOperDevice_list(in["device_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugRedirectTableL3OperOperDevice_list(d []interface{}) []edpt.ScaleoutDebugRedirectTableL3OperOperDevice_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugRedirectTableL3OperOperDevice_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugRedirectTableL3OperOperDevice_list
		oi.Device_id = in["device_id"].(int)
		oi.Red_follow_shared = in["red_follow_shared"].(int)
		oi.Red_table_list = getSliceScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list(in["red_table_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list(d []interface{}) []edpt.ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list
		oi.Id1 = in["id1"].(int)
		oi.Dir = in["dir"].(string)
		oi.Dst_ip = in["dst_ip"].(string)
		oi.Total_count = in["total_count"].(int)
		oi.Mac = in["mac"].(string)
		oi.Real_port = in["real_port"].(int)
		oi.Count_list = getSliceScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list(in["count_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list(d []interface{}) []edpt.ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list
		oi.Dst_index = in["dst_index"].(int)
		oi.Dst_index_valid = in["dst_index_valid"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugRedirectTableL3Oper(d *schema.ResourceData) edpt.ScaleoutDebugRedirectTableL3Oper {
	var ret edpt.ScaleoutDebugRedirectTableL3Oper

	ret.Oper = getObjectScaleoutDebugRedirectTableL3OperOper(d.Get("oper").([]interface{}))
	return ret
}
