

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats struct {
    
    Stats SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats `json:"stats"`
    WebCategoryList string `json:"web-category-list"`
    Name string 

}
type DataSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats struct {
    DtSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats `json:"web-category-list"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStatsStats struct {
    Hits int `json:"hits"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats) getPath() string{
    
    return "slb/template/policy/" +p.Name + "/forward-policy/source/" +p.Name + "/destination/web-category-list/"+p.WebCategoryList+"/stats"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats,error) {
logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
