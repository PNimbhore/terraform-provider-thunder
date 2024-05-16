package thunder

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"log"
	"reflect"
	"strings"
)

func Provider() *schema.Provider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name/IP of the THUNDER",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username with API access to the THUNDER",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user's password",
			},
			"partition": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "partition name",
				Default:      "shared",
				ValidateFunc: validation.StringLenBetween(0, 14),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"thunder_aam_aaa_policy":      resourceAamAaaPolicy(), 
			
		},
		DataSourcesMap: map[string]*schema.Resource{
			"thunder_aam_aaa_policy_aaa_rule_stats":      resourceAamAaaPolicyAaaRuleStats(), 
					},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	device := Thunder{
		Host:      d.Get("address").(string),
		User:      d.Get("username").(string),
		Password:  d.Get("password").(string),
		Partition: d.Get("partition").(string),
	}
	var diags diag.Diagnostics

	err := device.Connect()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create thunder client",
			Detail:   "Unable to authenticate user for authenticated thunder client",
		})
		return device, diags
	}
	err1 := device.callPartition()
	if err1 != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to switch partition",
			Detail:   "Unable to switch to partition " + device.Partition,
		})
		return device, diags
	}
	return device, diags
}

func mapEntity(d map[string]interface{}, obj interface{}) {
	val := reflect.ValueOf(obj).Elem()
	for field := range d {
		f := val.FieldByName(strings.Title(field))
		if f.IsValid() {
			if f.Kind() == reflect.Slice {
				incoming := d[field].([]interface{})
				s := reflect.MakeSlice(f.Type(), len(incoming), len(incoming))
				for i := 0; i < len(incoming); i++ {
					s.Index(i).Set(reflect.ValueOf(incoming[i]))
				}
				f.Set(s)
			} else {
				f.Set(reflect.ValueOf(d[field]))
			}
		} else {
			log.Printf("[WARN] You probably weren't expecting %s to be an invalid field", field)
		}
	}
}
