package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_win_trigger_stats_inc`: Configure stats to trigger packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"kerberos_delete_kdc_keytab_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Deletion Failure",
			},
			"kerberos_generate_kdc_keytab_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Generation Failure",
			},
			"kerberos_job_start_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Job Start Error",
			},
			"kerberos_other_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Other Error",
			},
			"kerberos_polling_control_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Polling Control Error",
			},
			"kerberos_pw_change_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password change failure",
			},
			"kerberos_pw_expiry": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password expiry",
			},
			"kerberos_request_dropped": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Dropped Request",
			},
			"kerberos_response_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Error Response",
			},
			"kerberos_response_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Failure Response",
			},
			"kerberos_response_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout Response",
			},
			"kerberos_timeout_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout",
			},
			"kerberos_validate_kdc_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Validation Failure",
			},
			"ntlm_authentication_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Authentication Failure",
			},
			"ntlm_job_start_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Job Start Error",
			},
			"ntlm_other_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Other Error",
			},
			"ntlm_polling_control_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Polling Control Error",
			},
			"ntlm_prepare_req_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Prepare Request Failed",
			},
			"ntlm_proto_negotiation_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Protocol Negotiation Failure",
			},
			"ntlm_request_dropped": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Dropped Request",
			},
			"ntlm_response_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Error Response",
			},
			"ntlm_response_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Failure Response",
			},
			"ntlm_response_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Timeout Response",
			},
			"ntlm_session_setup_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Session Setup Failure",
			},
			"ntlm_timeout_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total NTLM Timeout",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc
	ret.Inst.KerberosDeleteKdcKeytabFailure = d.Get("kerberos_delete_kdc_keytab_failure").(int)
	ret.Inst.KerberosGenerateKdcKeytabFailure = d.Get("kerberos_generate_kdc_keytab_failure").(int)
	ret.Inst.KerberosJobStartError = d.Get("kerberos_job_start_error").(int)
	ret.Inst.KerberosOtherError = d.Get("kerberos_other_error").(int)
	ret.Inst.KerberosPollingControlError = d.Get("kerberos_polling_control_error").(int)
	ret.Inst.KerberosPwChangeFailure = d.Get("kerberos_pw_change_failure").(int)
	ret.Inst.KerberosPwExpiry = d.Get("kerberos_pw_expiry").(int)
	ret.Inst.KerberosRequestDropped = d.Get("kerberos_request_dropped").(int)
	ret.Inst.KerberosResponseError = d.Get("kerberos_response_error").(int)
	ret.Inst.KerberosResponseFailure = d.Get("kerberos_response_failure").(int)
	ret.Inst.KerberosResponseTimeout = d.Get("kerberos_response_timeout").(int)
	ret.Inst.KerberosTimeoutError = d.Get("kerberos_timeout_error").(int)
	ret.Inst.KerberosValidateKdcFailure = d.Get("kerberos_validate_kdc_failure").(int)
	ret.Inst.NtlmAuthenticationFailure = d.Get("ntlm_authentication_failure").(int)
	ret.Inst.NtlmJobStartError = d.Get("ntlm_job_start_error").(int)
	ret.Inst.NtlmOtherError = d.Get("ntlm_other_error").(int)
	ret.Inst.NtlmPollingControlError = d.Get("ntlm_polling_control_error").(int)
	ret.Inst.NtlmPrepareReqFailed = d.Get("ntlm_prepare_req_failed").(int)
	ret.Inst.NtlmProtoNegotiationFailure = d.Get("ntlm_proto_negotiation_failure").(int)
	ret.Inst.NtlmRequestDropped = d.Get("ntlm_request_dropped").(int)
	ret.Inst.NtlmResponseError = d.Get("ntlm_response_error").(int)
	ret.Inst.NtlmResponseFailure = d.Get("ntlm_response_failure").(int)
	ret.Inst.NtlmResponseTimeout = d.Get("ntlm_response_timeout").(int)
	ret.Inst.NtlmSessionSetupFailed = d.Get("ntlm_session_setup_failed").(int)
	ret.Inst.NtlmTimeoutError = d.Get("ntlm_timeout_error").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
