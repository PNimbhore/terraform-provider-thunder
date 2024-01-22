

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugA10support struct {
	Inst struct {

    Duration int `json:"duration" dval:"2"`
    Password string `json:"password"`
    Uuid string `json:"uuid"`

	} `json:"a10support"`
}

func (p *DebugA10support) GetId() string{
    return "1"
}

func (p *DebugA10support) getPath() string{
    return "debug/a10support"
}

func (p *DebugA10support) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugA10support::Post")
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

func (p *DebugA10support) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugA10support::Get")
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
func (p *DebugA10support) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugA10support::Put")
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

func (p *DebugA10support) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugA10support::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
