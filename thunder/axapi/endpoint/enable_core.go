

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableCore struct {
	Inst struct {

    CoreLevel string `json:"core-level" dval:"a10"`
    Full int `json:"full"`
    Uuid string `json:"uuid"`

	} `json:"enable-core"`
}

func (p *EnableCore) GetId() string{
    return "1"
}

func (p *EnableCore) getPath() string{
    return "enable-core"
}

func (p *EnableCore) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableCore::Post")
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

func (p *EnableCore) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableCore::Get")
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
func (p *EnableCore) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableCore::Put")
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

func (p *EnableCore) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableCore::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
