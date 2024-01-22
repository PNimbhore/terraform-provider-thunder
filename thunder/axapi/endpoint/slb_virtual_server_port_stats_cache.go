

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServerPortStats struct {
	Inst struct {

    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats SlbVirtualServerPortStatsStats1475 `json:"stats"`
    Name string 

	} `json:"stats"`
}


type SlbVirtualServerPortStatsStats1475 struct {
    Cache SlbVirtualServerPortStatsStatsCache `json:"cache"`
}


type SlbVirtualServerPortStatsStatsCache struct {
    Hits int `json:"hits"`
    Miss int `json:"miss"`
    Bytes_served int `json:"bytes_served"`
    Total_req int `json:"total_req"`
    Caching_req int `json:"caching_req"`
    Nc_req_header int `json:"nc_req_header"`
    Nc_res_header int `json:"nc_res_header"`
    Rv_success int `json:"rv_success"`
    Rv_failure int `json:"rv_failure"`
    Ims_request int `json:"ims_request"`
    Nm_response int `json:"nm_response"`
    Rsp_type_cl int `json:"rsp_type_CL"`
    Rsp_type_ce int `json:"rsp_type_CE"`
    Rsp_type_304 int `json:"rsp_type_304"`
    Rsp_type_other int `json:"rsp_type_other"`
    Rsp_no_compress int `json:"rsp_no_compress"`
    Rsp_gzip int `json:"rsp_gzip"`
    Rsp_deflate int `json:"rsp_deflate"`
    Rsp_other int `json:"rsp_other"`
    Nocache_match int `json:"nocache_match"`
    Match int `json:"match"`
    Invalidate_match int `json:"invalidate_match"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Mem_size int `json:"mem_size"`
    Entry_num int `json:"entry_num"`
    Replaced_entry int `json:"replaced_entry"`
    Aging_entry int `json:"aging_entry"`
    Cleaned_entry int `json:"cleaned_entry"`
    Rsp_type_stream int `json:"rsp_type_stream"`
    Header_save_error int `json:"header_save_error"`
    Rsp_br int `json:"rsp_br"`
}

func (p *SlbVirtualServerPortStats) GetId() string{
    return "1"
}

func (p *SlbVirtualServerPortStats) getPath() string{
    return "slb/virtual-server/" +p.Inst.Name + "/port/" +strconv.Itoa(p.Inst.PortNumber)+"+"+p.Inst.Protocol+"/stats?cache=true"
}

func (p *SlbVirtualServerPortStats) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats::Post")
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

func (p *SlbVirtualServerPortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats::Get")
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
func (p *SlbVirtualServerPortStats) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats::Put")
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

func (p *SlbVirtualServerPortStats) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
