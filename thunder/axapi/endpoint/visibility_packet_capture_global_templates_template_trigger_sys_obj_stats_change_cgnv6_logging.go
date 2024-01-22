

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc1979 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate1980 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"cgnv6-logging"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc1979 struct {
    LogDropped int `json:"log-dropped"`
    ConnTcpDropped int `json:"conn-tcp-dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate1980 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    LogDropped int `json:"log-dropped"`
    ConnTcpDropped int `json:"conn-tcp-dropped"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/cgnv6-logging"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
