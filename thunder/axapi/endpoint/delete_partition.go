

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DeletePartition struct {
	Inst struct {

    Id1 int `json:"id1"`
    PartitionName string `json:"partition-name"`

	} `json:"partition"`
}

func (p *DeletePartition) GetId() string{
    return "1"
}

func (p *DeletePartition) getPath() string{
    return "delete/partition"
}

func (p *DeletePartition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DeletePartition::Post")
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

func (p *DeletePartition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DeletePartition::Get")
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
func (p *DeletePartition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DeletePartition::Put")
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

func (p *DeletePartition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DeletePartition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
