

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate struct {
	Inst struct {

    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Duration int `json:"duration" dval:"60"`
    Entry_create_failures int `json:"entry_create_failures"`
    Header_save_error int `json:"header_save_error"`
    Nc_req_header int `json:"nc_req_header"`
    Nc_res_header int `json:"nc_res_header"`
    Rv_failure int `json:"rv_failure"`
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate) getPath() string{
    return "visibility/packet-capture/object-templates/slb-templ-cache-tmpl/" +p.Inst.Name + "/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
