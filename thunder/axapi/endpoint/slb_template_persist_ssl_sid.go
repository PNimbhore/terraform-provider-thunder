

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePersistSslSid struct {
	Inst struct {

    DontHonorConnRules int `json:"dont-honor-conn-rules"`
    Name string `json:"name"`
    Timeout int `json:"timeout" dval:"5"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ssl-sid"`
}

func (p *SlbTemplatePersistSslSid) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplatePersistSslSid) getPath() string{
    return "slb/template/persist/ssl-sid"
}

func (p *SlbTemplatePersistSslSid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistSslSid::Post")
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

func (p *SlbTemplatePersistSslSid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistSslSid::Get")
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
func (p *SlbTemplatePersistSslSid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistSslSid::Put")
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

func (p *SlbTemplatePersistSslSid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistSslSid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
