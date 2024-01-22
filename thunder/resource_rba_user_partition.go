package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRbaUserPartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rba_user_partition`: RBA configuration for the access privilege of a user within one partition\n\n__PLACEHOLDER__",
		CreateContext: resourceRbaUserPartitionCreate,
		UpdateContext: resourceRbaUserPartitionUpdate,
		ReadContext:   resourceRbaUserPartitionRead,
		DeleteContext: resourceRbaUserPartitionDelete,

		Schema: map[string]*schema.Schema{
			"partition_name": {
				Type: schema.TypeString, Required: true, Description: "partition name",
			},
			"role_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role": {
							Type: schema.TypeString, Optional: true, Description: "Role in a given partition",
						},
					},
				},
			},
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"object": {
							Type: schema.TypeString, Optional: true, Description: "Lineage of object class for permitted operation",
						},
						"operation": {
							Type: schema.TypeString, Optional: true, Description: "'no-access': no-access; 'read': read; 'oper': oper; 'write': write;",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceRbaUserPartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaUserPartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaUserPartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaUserPartitionRead(ctx, d, meta)
	}
	return diags
}

func resourceRbaUserPartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaUserPartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaUserPartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRbaUserPartitionRead(ctx, d, meta)
	}
	return diags
}
func resourceRbaUserPartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaUserPartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaUserPartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRbaUserPartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRbaUserPartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRbaUserPartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRbaUserPartitionRoleList(d []interface{}) []edpt.RbaUserPartitionRoleList {

	count1 := len(d)
	ret := make([]edpt.RbaUserPartitionRoleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaUserPartitionRoleList
		oi.Role = in["role"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRbaUserPartitionRuleList(d []interface{}) []edpt.RbaUserPartitionRuleList {

	count1 := len(d)
	ret := make([]edpt.RbaUserPartitionRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RbaUserPartitionRuleList
		oi.Object = in["object"].(string)
		oi.Operation = in["operation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRbaUserPartition(d *schema.ResourceData) edpt.RbaUserPartition {
	var ret edpt.RbaUserPartition
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	ret.Inst.RoleList = getSliceRbaUserPartitionRoleList(d.Get("role_list").([]interface{}))
	ret.Inst.RuleList = getSliceRbaUserPartitionRuleList(d.Get("rule_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
