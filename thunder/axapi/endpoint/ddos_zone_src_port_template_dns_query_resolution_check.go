

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneSrcPortTemplateDnsQueryResolutionCheck struct {
	Inst struct {

    BigResponseAction string `json:"big-response-action" dval:"default"`
    BigResponseSize int `json:"big-response-size"`
    DomainLockupAction string `json:"domain-lockup-action" dval:"default"`
    SessionTimeoutValue int `json:"session-timeout-value"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"query-resolution-check"`
}

func (p *DdosZoneSrcPortTemplateDnsQueryResolutionCheck) GetId() string{
    return "1"
}

func (p *DdosZoneSrcPortTemplateDnsQueryResolutionCheck) getPath() string{
    return "ddos/zone-src-port-template/dns/" +p.Inst.Name + "/query-resolution-check"
}

func (p *DdosZoneSrcPortTemplateDnsQueryResolutionCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateDnsQueryResolutionCheck::Post")
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

func (p *DdosZoneSrcPortTemplateDnsQueryResolutionCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateDnsQueryResolutionCheck::Get")
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
func (p *DdosZoneSrcPortTemplateDnsQueryResolutionCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateDnsQueryResolutionCheck::Put")
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

func (p *DdosZoneSrcPortTemplateDnsQueryResolutionCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateDnsQueryResolutionCheck::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
