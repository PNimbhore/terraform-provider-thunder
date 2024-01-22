

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SystemHealthCheck struct {
	Inst struct {

    L2bfdMultiplier int `json:"l2bfd-multiplier"`
    L2bfdRxInterval int `json:"l2bfd-rx-interval"`
    L2bfdTxInterval int `json:"l2bfd-tx-interval"`
    L2hmHcName string `json:"l2hm-hc-name"`
    MethodL2bfd int `json:"method-l2bfd"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"health-check"`
}

func (p *SystemHealthCheck) GetId() string{
    return url.QueryEscape(p.Inst.L2hmHcName)
}

func (p *SystemHealthCheck) getPath() string{
    return "system/health-check"
}

func (p *SystemHealthCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemHealthCheck::Post")
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

func (p *SystemHealthCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemHealthCheck::Get")
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
func (p *SystemHealthCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemHealthCheck::Put")
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

func (p *SystemHealthCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemHealthCheck::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
