---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_web_category Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_web_category: Web-Category Commands
  PLACEHOLDER
---

# thunder_web_category (Resource)

`thunder_web_category`: Web-Category Commands

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_web_category" "thunder_web_category" {

  category_list_list {
    name                           = "test"
    uncategorized                  = 1
    real_estate                    = 1
    computer_and_internet_security = 1
    financial_services             = 1
    business_and_economy           = 1
    computer_and_internet_info     = 1
    auctions                       = 1
    shopping                       = 1
    cult_and_occult                = 1
    travel                         = 1
    drugs                          = 1
    adult_and_pornography          = 1
    home_and_garden                = 1
    military                       = 1
    social_network                 = 1
    dead_sites                     = 1
    stock_advice_and_tools         = 1
    training_and_tools             = 1
    dating                         = 1
    sex_education                  = 1
    religion                       = 1
    entertainment_and_arts         = 1
    personal_sites_and_blogs       = 1
    legal                          = 1
    local_information              = 1
    streaming_media                = 1
    job_search                     = 1
    gambling                       = 1
    translation                    = 1
    reference_and_research         = 1
    shareware_and_freeware         = 1
    peer_to_peer                   = 1
    marijuana                      = 1
    hacking                        = 1
    games                          = 1
    philosophy_and_politics        = 1
    weapons                        = 1
    pay_to_surf                    = 1
    hunting_and_fishing            = 1
    society                        = 1
    educational_institutions       = 1
    online_greeting_cards          = 1
    sports                         = 1
    swimsuits_and_intimate_apparel = 1
    questionable                   = 1
    kids                           = 1
    hate_and_racism                = 1
    personal_storage               = 1
    violence                       = 1
    keyloggers_and_monitoring      = 1
    search_engines                 = 1
    internet_portals               = 1
    web_advertisements             = 1
    cheating                       = 1
    gross                          = 1
    web_based_email                = 1
    malware_sites                  = 1
    phishing_and_other_fraud       = 1
    proxy_avoid_and_anonymizers    = 1
    spyware_and_adware             = 1
    music                          = 1
    government                     = 1
    nudity                         = 1
    news_and_media                 = 1
    illegal                        = 1
    cdns                           = 1
    internet_communications        = 1
    bot_nets                       = 1
    abortion                       = 1
    health_and_medicine            = 1
    spam_urls                      = 1
    dynamically_generated_content  = 1
    parked_domains                 = 1
    alcohol_and_tobacco            = 1
    image_and_video_search         = 1
    fashion_and_beauty             = 1
    recreation_and_hobbies         = 1
    motor_vehicles                 = 1
    web_hosting_sites              = 1
    nudity_artistic                = 1
    illegal_pornography            = 1
    user_tag                       = "97"
    sampling_enable {
      counters1 = "all"
    }
  }
  cloud_query_cache_size = 222
  cloud_query_disable    = 1
  database_server        = "133"
  online_check_disable   = 1
  port                   = 80
  proxy_server {
    proxy_host    = "16"
    http_port     = 37683
    https_port    = 46027
    auth_type     = "ntlm"
    username      = "99"
    password      = 1
    secret_string = "47"
  }
  remote_syslog_enable = 1
  reputation_scope_list {
    name = "test"
    greater_than {
      greater_trustworthy = 1
    }
    less_than {
      less_trustworthy = 1
    }
    user_tag = "9"
    sampling_enable {
      counters1 = "all"
    }
  }
  rtu_cache_size      = 222
  rtu_update_disable  = 1
  rtu_update_interval = 60
  server_timeout      = 15
  ssl_port            = 443
  statistics {
    sampling_enable {
      counters1 = "all"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `bypassed_urls` (Block List, Max: 1) (see [below for nested schema](#nestedblock--bypassed_urls))
- `category_list_list` (Block List) (see [below for nested schema](#nestedblock--category_list_list))
- `cloud_query_cache_size` (Number) Maximum cache size for storing cloud query results
- `cloud_query_disable` (Number) Disables cloud queries for URL's not present in local database(default enable)
- `database_server` (String) BrightCloud Database Server
- `db_update_time` (String) Time of day to update database (default: 00:00)
- `enable` (Number) Enable BrightCloud SDK
- `intercepted_urls` (Block List, Max: 1) (see [below for nested schema](#nestedblock--intercepted_urls))
- `license` (Block List, Max: 1) (see [below for nested schema](#nestedblock--license))
- `online_check_disable` (Number) Disables online queries for license. By default it is enabled.
- `port` (Number) BrightCloud Query Server Listening Port(default 80)
- `proxy_server` (Block List, Max: 1) (see [below for nested schema](#nestedblock--proxy_server))
- `remote_syslog_enable` (Number) Enable data plane logging to a remote syslog server
- `reputation_scope_list` (Block List) (see [below for nested schema](#nestedblock--reputation_scope_list))
- `rtu_cache_size` (Number) Maximum cache size for storing RTU updates
- `rtu_update_disable` (Number) Disables real time updates(default enable)
- `rtu_update_interval` (Number) Interval to check for real time updates if enabled in mins(default 60)
- `server` (String) BrightCloud Query Server
- `server_timeout` (Number) BrightCloud Servers Timeout in seconds (default: 15s)
- `ssl_port` (Number) BrightCloud Servers SSL Port(default 443)
- `statistics` (Block List, Max: 1) (see [below for nested schema](#nestedblock--statistics))
- `url` (Block List, Max: 1) (see [below for nested schema](#nestedblock--url))
- `use_mgmt_port` (Number) Use management interface for all communication with BrightCloud
- `uuid` (String) uuid of the object
- `web_reputation` (Block List, Max: 1) (see [below for nested schema](#nestedblock--web_reputation))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--bypassed_urls"></a>
### Nested Schema for `bypassed_urls`

Optional:

- `uuid` (String) uuid of the object


<a id="nestedblock--category_list_list"></a>
### Nested Schema for `category_list_list`

Required:

- `name` (String) Web Category List name

Optional:

- `abortion` (Number) Category Abortion
- `adult_and_pornography` (Number) Category Adult and Pornography
- `alcohol_and_tobacco` (Number) Category Alcohol and Tobacco
- `auctions` (Number) Category Auctions
- `bot_nets` (Number) Category Bot Nets
- `business_and_economy` (Number) Category Business and Economy
- `cdns` (Number) Category CDNs
- `cheating` (Number) Category Cheating
- `computer_and_internet_info` (Number) Category Computer and Internet Info
- `computer_and_internet_security` (Number) Category Computer and Internet Security
- `cult_and_occult` (Number) Category Cult and Occult
- `dating` (Number) Category Dating
- `dead_sites` (Number) Category Dead Sites (db Ops only)
- `drugs` (Number) Category Abused Drugs
- `dynamically_generated_content` (Number) Dynamically Generated Content
- `educational_institutions` (Number) Category Educational Institutions
- `entertainment_and_arts` (Number) Category Entertainment and Arts
- `fashion_and_beauty` (Number) Category Fashion and Beauty
- `financial_services` (Number) Category Financial Services
- `gambling` (Number) Category Gambling
- `games` (Number) Category Games
- `government` (Number) Category Government
- `gross` (Number) Category Gross
- `hacking` (Number) Category Hacking
- `hate_and_racism` (Number) Category Hate and Racism
- `health_and_medicine` (Number) Category Health and Medicine
- `home_and_garden` (Number) Category Home and Garden
- `hunting_and_fishing` (Number) Category Hunting and Fishing
- `illegal` (Number) Category Illegal
- `illegal_pornography` (Number) Category Illegal join Adult and Pornography
- `image_and_video_search` (Number) Category Image and Video Search
- `internet_communications` (Number) Category Internet Communications
- `internet_portals` (Number) Category Internet Portals
- `job_search` (Number) Category Job Search
- `keyloggers_and_monitoring` (Number) Category Keyloggers and Monitoring
- `kids` (Number) Category Kids
- `legal` (Number) Category Legal
- `local_information` (Number) Category Local Information
- `malware_sites` (Number) Category Malware Sites
- `marijuana` (Number) Category Marijuana
- `military` (Number) Category Military
- `motor_vehicles` (Number) Category Motor Vehicles
- `music` (Number) Category Music
- `news_and_media` (Number) Category News and Media
- `nudity` (Number) Category Nudity
- `nudity_artistic` (Number) Category Nudity join Entertainment and Arts
- `online_greeting_cards` (Number) Category Online Greeting cards
- `parked_domains` (Number) Category Parked Domains
- `pay_to_surf` (Number) Category Pay to Surf
- `peer_to_peer` (Number) Category Peer to Peer
- `personal_sites_and_blogs` (Number) Category Personal sites and Blogs
- `personal_storage` (Number) Category Personal Storage
- `philosophy_and_politics` (Number) Category Philosophy and Political Advocacy
- `phishing_and_other_fraud` (Number) Category Phishing and Other Frauds
- `proxy_avoid_and_anonymizers` (Number) Category Proxy Avoid and Anonymizers
- `questionable` (Number) Category Questionable
- `real_estate` (Number) Category Real Estate
- `recreation_and_hobbies` (Number) Category Recreation and Hobbies
- `reference_and_research` (Number) Category Reference and Research
- `religion` (Number) Category Religion
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--category_list_list--sampling_enable))
- `search_engines` (Number) Category Search Engines
- `sex_education` (Number) Category Sex Education
- `shareware_and_freeware` (Number) Category Shareware and Freeware
- `shopping` (Number) Category Shopping
- `social_network` (Number) Category Social Network
- `society` (Number) Category Society
- `spam_urls` (Number) Category SPAM URLs
- `sports` (Number) Category Sports
- `spyware_and_adware` (Number) Category Spyware and Adware
- `stock_advice_and_tools` (Number) Category Stock Advice and Tools
- `streaming_media` (Number) Category Streaming Media
- `swimsuits_and_intimate_apparel` (Number) Category Swimsuits and Intimate Apparel
- `training_and_tools` (Number) Category Training and Tools
- `translation` (Number) Category Translation
- `travel` (Number) Category Travel
- `uncategorized` (Number) Uncategorized URLs
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `violence` (Number) Category Violence
- `weapons` (Number) Category Weapons
- `web_advertisements` (Number) Category Web Advertisements
- `web_based_email` (Number) Category Web based email
- `web_hosting_sites` (Number) Category Web Hosting Sites

<a id="nestedblock--category_list_list--sampling_enable"></a>
### Nested Schema for `category_list_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'uncategorized': uncategorized category; 'real-estate': real estate category; 'computer-and-internet-security': computer and internet security category; 'financial-services': financial services category; 'business-and-economy': business and economy category; 'computer-and-internet-info': computer and internet info category; 'auctions': auctions category; 'shopping': shopping category; 'cult-and-occult': cult and occult category; 'travel': travel category; 'drugs': drugs category; 'adult-and-pornography': adult and pornography category; 'home-and-garden': home and garden category; 'military': military category; 'social-network': social network category; 'dead-sites': dead sites category; 'stock-advice-and-tools': stock advice and tools category; 'training-and-tools': training and tools category; 'dating': dating category; 'sex-education': sex education category; 'religion': religion category; 'entertainment-and-arts': entertainment and arts category; 'personal-sites-and-blogs': personal sites and blogs category; 'legal': legal category; 'local-information': local information category; 'streaming-media': streaming media category; 'job-search': job search category; 'gambling': gambling category; 'translation': translation category; 'reference-and-research': reference and research category; 'shareware-and-freeware': shareware and freeware category; 'peer-to-peer': peer to peer category; 'marijuana': marijuana category; 'hacking': hacking category; 'games': games category; 'philosophy-and-politics': philosophy and politics category; 'weapons': weapons category; 'pay-to-surf': pay to surf category; 'hunting-and-fishing': hunting and fishing category; 'society': society category; 'educational-institutions': educational institutions category; 'online-greeting-cards': online greeting cards category; 'sports': sports category; 'swimsuits-and-intimate-apparel': swimsuits and intimate apparel category; 'questionable': questionable category; 'kids': kids category; 'hate-and-racism': hate and racism category; 'personal-storage': personal storage category; 'violence': violence category; 'keyloggers-and-monitoring': keyloggers and monitoring category; 'search-engines': search engines category; 'internet-portals': internet portals category; 'web-advertisements': web advertisements category; 'cheating': cheating category; 'gross': gross category; 'web-based-email': web based email category; 'malware-sites': malware sites category; 'phishing-and-other-fraud': phishing and other fraud category; 'proxy-avoid-and-anonymizers': proxy avoid and anonymizers category; 'spyware-and-adware': spyware and adware category; 'music': music category; 'government': government category; 'nudity': nudity category; 'news-and-media': news and media category; 'illegal': illegal category; 'CDNs': content delivery networks category; 'internet-communications': internet communications category; 'bot-nets': bot nets category; 'abortion': abortion category; 'health-and-medicine': health and medicine category; 'confirmed-SPAM-sources': confirmed SPAM sources category; 'SPAM-URLs': SPAM URLs category; 'unconfirmed-SPAM-sources': unconfirmed SPAM sources category; 'open-HTTP-proxies': open HTTP proxies category; 'dynamically-generated-content': dynamically generated content category; 'parked-domains': parked domains category; 'alcohol-and-tobacco': alcohol and tobacco category; 'private-IP-addresses': private IP addresses category; 'image-and-video-search': image and video search category; 'fashion-and-beauty': fashion and beauty category; 'recreation-and-hobbies': recreation and hobbies category; 'motor-vehicles': motor vehicles category; 'web-hosting-sites': web hosting sites category; 'food-and-dining': food and dining category; 'nudity-artistic': nudity join entertainment and arts; 'illegal-pornography': illegal join adult and pornography;



<a id="nestedblock--intercepted_urls"></a>
### Nested Schema for `intercepted_urls`

Optional:

- `uuid` (String) uuid of the object


<a id="nestedblock--license"></a>
### Nested Schema for `license`

Optional:

- `uuid` (String) uuid of the object


<a id="nestedblock--proxy_server"></a>
### Nested Schema for `proxy_server`

Optional:

- `auth_type` (String) 'ntlm': NTLM authentication(default); 'basic': Basic authentication;
- `domain` (String) Realm for NTLM authentication
- `http_port` (Number) Proxy server HTTP port
- `https_port` (Number) Proxy server HTTPS port(HTTP port will be used if not configured)
- `password` (Number) Password for proxy authentication
- `proxy_host` (String) Proxy server hostname or IP address
- `secret_string` (String) password value
- `username` (String) Username for proxy authentication
- `uuid` (String) uuid of the object


<a id="nestedblock--reputation_scope_list"></a>
### Nested Schema for `reputation_scope_list`

Required:

- `name` (String) Reputation Scope name

Optional:

- `greater_than` (Block List, Max: 1) (see [below for nested schema](#nestedblock--reputation_scope_list--greater_than))
- `less_than` (Block List, Max: 1) (see [below for nested schema](#nestedblock--reputation_scope_list--less_than))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--reputation_scope_list--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--reputation_scope_list--greater_than"></a>
### Nested Schema for `reputation_scope_list.greater_than`

Optional:

- `greater_low_risk` (Number) Reputation score is greater than or equal to 61
- `greater_malicious` (Number) Reputation score is greater than or equal to 1
- `greater_moderate_risk` (Number) Reputation score is greater than or equal to 41
- `greater_suspicious` (Number) Reputation score is greater than or equal to 21
- `greater_threshold` (Number) Reputation score is greater than or equal to the customized score (1-100)
- `greater_trustworthy` (Number) Reputation score is greater than or equal to 81


<a id="nestedblock--reputation_scope_list--less_than"></a>
### Nested Schema for `reputation_scope_list.less_than`

Optional:

- `less_low_risk` (Number) Reputation score is less than or equal to 80
- `less_malicious` (Number) Reputation score is less than or equal to 20
- `less_moderate_risk` (Number) Reputation score is less than or equal to 60
- `less_suspicious` (Number) Reputation score is less than or equal to 40
- `less_threshold` (Number) Reputation score is less than or equal to a customized value (1-100)
- `less_trustworthy` (Number) Reputation score is less than or equal to 100


<a id="nestedblock--reputation_scope_list--sampling_enable"></a>
### Nested Schema for `reputation_scope_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'trustworthy': Trustworthy level(81-100); 'low-risk': Low-risk level(61-80); 'moderate-risk': Moderate-risk level(41-60); 'suspicious': Suspicious level(21-40); 'malicious': Malicious level(1-20);



<a id="nestedblock--statistics"></a>
### Nested Schema for `statistics`

Optional:

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--statistics--sampling_enable))
- `uuid` (String) uuid of the object

<a id="nestedblock--statistics--sampling_enable"></a>
### Nested Schema for `statistics.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'db-lookup': db-lookup; 'cloud-cache-lookup': cloud-cache-lookup; 'cloud-lookup': cloud-lookup; 'rtu-lookup': rtu-lookup; 'lookup-latency': lookup-latency; 'db-mem': db-mem; 'rtu-cache-mem': rtu-cache-mem; 'lookup-cache-mem': lookup-cache-mem;



<a id="nestedblock--url"></a>
### Nested Schema for `url`

Optional:

- `uuid` (String) uuid of the object


<a id="nestedblock--web_reputation"></a>
### Nested Schema for `web_reputation`

Optional:

- `bypassed_urls` (Block List, Max: 1) (see [below for nested schema](#nestedblock--web_reputation--bypassed_urls))
- `intercepted_urls` (Block List, Max: 1) (see [below for nested schema](#nestedblock--web_reputation--intercepted_urls))
- `url` (Block List, Max: 1) (see [below for nested schema](#nestedblock--web_reputation--url))
- `uuid` (String) uuid of the object

<a id="nestedblock--web_reputation--bypassed_urls"></a>
### Nested Schema for `web_reputation.bypassed_urls`

Optional:

- `uuid` (String) uuid of the object


<a id="nestedblock--web_reputation--intercepted_urls"></a>
### Nested Schema for `web_reputation.intercepted_urls`

Optional:

- `uuid` (String) uuid of the object


<a id="nestedblock--web_reputation--url"></a>
### Nested Schema for `web_reputation.url`

Optional:

- `uuid` (String) uuid of the object


