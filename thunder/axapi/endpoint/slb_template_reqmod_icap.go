

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateReqmodIcap struct {
	Inst struct {

    Action string `json:"action" dval:"continue"`
    AllowedHttpMethods string `json:"allowed-http-methods"`
    BypassIpCfg []SlbTemplateReqmodIcapBypassIpCfg `json:"bypass-ip-cfg"`
    DisableHttpServerReset int `json:"disable-http-server-reset"`
    FailClose int `json:"fail-close"`
    FailureAction string `json:"failure-action" dval:"continue"`
    IncludeProtocolInUri int `json:"include-protocol-in-uri"`
    LogOnlyAllowedMethod int `json:"log-only-allowed-method"`
    Logging string `json:"logging"`
    MinPayloadSize int `json:"min-payload-size"`
    Name string `json:"name"`
    Preview int `json:"preview" dval:"32768"`
    ServerSsl string `json:"server-ssl"`
    ServiceGroup string `json:"service-group"`
    ServiceUrl string `json:"service-url"`
    SharedPartitionPersistSourceIpTemplate int `json:"shared-partition-persist-source-ip-template"`
    SharedPartitionTcpProxyTemplate int `json:"shared-partition-tcp-proxy-template"`
    SourceIp string `json:"source-ip"`
    TcpProxy string `json:"tcp-proxy"`
    TemplatePersistSourceIpShared string `json:"template-persist-source-ip-shared"`
    TemplateTcpProxyShared string `json:"template-tcp-proxy-shared"`
    Timeout int `json:"timeout" dval:"5"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    XAuthUrl int `json:"x-auth-url"`

	} `json:"reqmod-icap"`
}


type SlbTemplateReqmodIcapBypassIpCfg struct {
    BypassIp string `json:"bypass-ip"`
    Mask string `json:"mask"`
}

func (p *SlbTemplateReqmodIcap) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateReqmodIcap) getPath() string{
    return "slb/template/reqmod-icap"
}

func (p *SlbTemplateReqmodIcap) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateReqmodIcap::Post")
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

func (p *SlbTemplateReqmodIcap) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateReqmodIcap::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *SlbTemplateReqmodIcap) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateReqmodIcap::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *SlbTemplateReqmodIcap) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateReqmodIcap::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
