package television

import (
	"encoding/json"
	"strconv"

	"github.com/valyala/fasthttp"
)

// Television struct to store credentials and client required for making requests to JioTV API
type Television struct {
	AccessToken string
	SsoToken    string
	Crm         string
	UniqueID    string
	Headers     map[string]string
	Client      *fasthttp.Client
}

// Channel represents Individual channel details from JioTV API
type Channel struct {
	ID       string `json:"channel_id"`
	Name     string `json:"channel_name"`
	URL      string `json:"channel_url"`
	LogoURL  string `json:"logoUrl"`
	Category int    `json:"channelCategoryId"`
	Language int    `json:"channelLanguageId"`
	IsHD     bool   `json:"isHD"`
}

// UnmarshalJSON to Override Channel.ID to convert int from json to string
func (c *Channel) UnmarshalJSON(b []byte) error {
	type Alias Channel
	aux := &struct {
		ID int `json:"channel_id"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	c.ID = strconv.Itoa(aux.ID)
	return nil
}

// ChannelsResponse is the response body for channels from JioTV API
type ChannelsResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Result  []Channel `json:"result"`
}

// Bitrates represents Quality levels for live streams for JioTV API
type Bitrates struct {
	Auto   string `json:"auto"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Medium string `json:"medium"`
}

type MPD struct {
	Result   string   `json:"result"`
	Key      string   `json:"key"`
	Bitrates Bitrates `json:"bitrates"`
}

// LiveURLOutput represents Response of live stream URL request to JioTV API
type LiveURLOutput struct {
	SonyVodStitchAdsCpCustomerID struct {
		Midroll  string `json:"midroll"`
		Postroll string `json:"postroll"`
		Preroll  string `json:"preroll"`
	} `json:"sonyVodStitchAdsCpCustomerID"`
	VmapURL     string   `json:"vmapUrl"`
	Bitrates    Bitrates `json:"bitrates"`
	Code        int      `json:"code"`
	ContentID   float64  `json:"contentId"`
	CurrentTime float64  `json:"currentTime"`
	EndTime     float64  `json:"endTime"`
	Message     string   `json:"message"`
	Result      string   `json:"result"`
	StartTime   float64  `json:"startTime"`
	VodStitch   bool     `json:"vodStitch"`
	Mpd         MPD      `json:"mpd"`
}

// CategoryMap represents Categories for channels
var CategoryMap = map[int]string{
	0:  "All Categories",
	5:  "Entertainment",
	6:  "Movies",
	7:  "Kids",
	8:  "Sports",
	9:  "Lifestyle",
	10: "Infotainment",
	12: "News",
	13: "Music",
	15: "Devotional",
	16: "Business",
	17: "Educational",
	18: "Shopping",
	19: "JioDarshan",
}

// LanguageMap represents Languages for channels
var LanguageMap = map[int]string{
	0:  "All Languages",
	1:  "Hindi",
	2:  "Marathi",
	3:  "Punjabi",
	4:  "Urdu",
	5:  "Bengali",
	6:  "English",
	7:  "Malayalam",
	8:  "Tamil",
	9:  "Gujarati",
	10: "Odia",
	11: "Telugu",
	12: "Bhojpuri",
	13: "Kannada",
	14: "Assamese",
	15: "Nepali",
	16: "French",
	18: "Other",
}

var SONY_CHANNELS = map[string]string{
	"sonyhd":         "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L2RCZHdPaUdhUXZ5MFRBMXpPc2pWNncvbWFzdGVyLm0zdTg=",
	"sonysabhd":      "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L0NyVGl2a0RFU1dxd3ZVajN6RkVZRUEvbWFzdGVyLm0zdTg=",
	"sonypal":        "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L2RoUHJHUndEUnZ1TVF0bWx6cHB6UVEvbWFzdGVyLm0zdTg=",
	"sonypixhd":      "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L3g3clhXZDJFUloydHZ5UVdQbU8xSEEvbWFzdGVyLm0zdTg=",
	"sonymaxhd":      "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L1VjakhOSm1DUTFXUmxHS2xabTczUUEvbWFzdGVyLm0zdTg=",
	"sonymax2":       "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L01kUTVaeS1QU3JhT2NjWHU4amZsQ2cvbWFzdGVyLm0zdTg=",
	"sonywah":        "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L2dYNXJDQmY2UTctRDVBV1ktc292elEvbWFzdGVyLm0zdTg=",
	"sonyten1hd":     "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L3dHNzVuNVU4UnJPS2lGemFXT2JYYkEvbWFzdGVyLm0zdTg=",
	"sonyten2hd":     "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L1Y5aC1peU94UmlHcDQxcHBRU2NEU1EvbWFzdGVyLm0zdTg=",
	"sonyten3hd":     "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L2x0c0NHN1RCU0NTRG15cTByUXR2U0EvbWFzdGVyLm0zdTg=",
	"sonyten4hd":     "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L3NtWXliSV9KVG9XYUh6d294U0U5cUEvbWFzdGVyLm0zdTg=",
	"sonyten5hd":     "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L1NsZV9UUjhyUUl1WkhXenNoRVhZalEvbWFzdGVyLm0zdTg=",
	"sonybbcearthhd": "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50LzZiVldZSUtHUzBDSWEtY09wWlpKUFEvbWFzdGVyLm0zdTg=",
}

var SONY_JIO_MAP = map[string]string{
	"sl291":  "sonyhd",
	"sl154":  "sonysabhd",
	"sl474":  "sonypal",
	"sl762":  "sonypixhd",
	"sl476":  "sonymaxhd",
	"sl483":  "sonymax2",
	"sl1393": "sonywah",
	"sl162":  "sonyten1hd",
	"sl891":  "sonyten2hd",
	"sl892":  "sonyten3hd",
	"sl1772": "sonyten4hd",
	"sl155":  "sonyten5hd",
	"sl852":  "sonybbcearthhd",
}

var SONY_CHANNELS_API = []Channel{
	{
		ID:       "sl291",
		Name:     "SL Sony HD",
		Language: 1,
		Category: 5,
		IsHD:     true,
		LogoURL:  "Sony_HD.png",
	},
	{
		ID:       "sl154",
		Name:     "SL Sony SAB HD",
		Language: 1,
		Category: 5,
		IsHD:     true,
		LogoURL:  "Sony_SAB_HD.png",
	},
	{
		ID:       "sl474",
		Name:     "SL Sony PAL",
		Language: 1,
		Category: 5,
		IsHD:     false,
		LogoURL:  "Sony_Pal.png",
	},
	{
		ID:       "sl762",
		Name:     "SL Sony PIX HD",
		Language: 6,
		Category: 6,
		IsHD:     true,
		LogoURL:  "Sony_Pix_HD.png",
	},
	{
		ID:       "sl476",
		Name:     "SL Sony MAX HD",
		Language: 1,
		Category: 6,
		IsHD:     true,
		LogoURL:  "Sony_Max_HD.png",
	},
	{
		ID:       "sl483",
		Name:     "SL Sony MAX 2",
		Language: 1,
		Category: 6,
		IsHD:     false,
		LogoURL:  "Sony_MAX2.png",
	},
	// Disabled as it requires CORS bypass
	// {
	// 	ID:       "sl1393",
	// 	Name:     "SL Sony WAH",
	// 	Language: 1,
	// 	Category: 5,
	// 	IsHD:     false,
	// 	LogoURL:  "Sony_Wah.png",
	// },
	{
		ID:       "sl162",
		Name:     "SL Sony TEN 1 HD",
		Language: 6,
		Category: 8,
		IsHD:     true,
		LogoURL:  "Ten_HD.png",
	},
	{
		ID:       "sl891",
		Name:     "SL Sony TEN 2 HD",
		Language: 6,
		Category: 8,
		IsHD:     true,
		LogoURL:  "Ten2_HD.png",
	},
	{
		ID:       "sl892",
		Name:     "SL Sony TEN 3 HD",
		Language: 1,
		Category: 8,
		IsHD:     true,
		LogoURL:  "Ten3_HD.png",
	},
	{
		ID:       "sl1772",
		Name:     "SL Sony TEN 4 HD",
		Language: 8,
		Category: 8,
		IsHD:     true,
		LogoURL:  "Ten_4_HD_Tamil.png",
	},
	{
		ID:       "sl155",
		Name:     "SL Sony TEN 5 HD",
		Language: 6,
		Category: 8,
		IsHD:     true,
		LogoURL:  "Six_HD.png",
	},
	{
		ID:       "sl852",
		Name:     "SL Sony BBC Earth HD",
		Language: 6,
		Category: 10,
		IsHD:     true,
		LogoURL:  "Sony_BBC_Earth_HD_English.png",
	},
}

// var ZEE_CHANNELS = map[string]string{
// 	"zeetv":             "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L2RCZHdPaUdhUXZ5MFRBMXpPc2pWNncvbWFzdGVyLm0zdTg=",
// 	"zeetvhd":           "aHR0cHM6Ly9kYWkuZ29vZ2xlLmNvbS9saW5lYXIvaGxzL2V2ZW50L0NyVGl2a0RFU1dxd3ZVajN6RkVZRUEvbWFzdGVyLm0zdTg=",
// 	"zee_anmol":         "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlYW5tb2wxL2RlZmF1bHQvbWFzdGVyLm0zdTg",
// 	"zee_cinema":        "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlc2FsYWFtMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_cinema_hd":     "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlc2FsYWFtMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_bollywood":     "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlYm9sbHl3b29kMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_action":        "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlYWN0aW9uMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_anmol_cinema":  "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlYW5tb2xjaW5lbWExL2RlZmF1bHQvbWFzdGVyLm0zdTg",
// 	"zee_bharat":        "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlaGluZHVzdGFuMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_business":      "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlYnVzaW5lc3MxL2RlZmF1bHQvbWFzdGVyLm0zdTg",
// 	"zee_salaam":        "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlc2FsYWFtMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_marathi":       "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlbWFyYXRoaTEvZGVmYXVsdC9tYXN0ZXIubTN1OA",
// 	"zee_marathi_hd":    "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlbWFyYXRoaWhkMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_talkies":       "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVldGFsa2llczEvZGVmYXVsdC9tYXN0ZXIubTN1OA==",
// 	"zee_talkies_hd":    "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVldGFsa2llc2hkMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_bangla":        "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlYmFuZ2xhMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_bangla_hd":     "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlYmFuZ2xhaGQxL2RlZmF1bHQvbWFzdGVyLm0zdTg",
// 	"zee_bangla_cinema": "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlYmFuZ2xhY2luZW1hMS9kZWZhdWx0L21hc3Rlci5tM3U4",
// 	"zee_tamil":         "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVldGFtaWwxL2RlZmF1bHQvbWFzdGVyLm0zdTg",
// 	"zee_cinemalu":      "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVlY2luZW1hbHUxL2RlZmF1bHQvbWFzdGVyLm0zdTg",
// 	"zee_kannada":       "aHR0cHM6Ly9kMWc4d2dqdXJ6OHZpYS5jbG91ZGZyb250Lm5ldC9icGstdHYvWmVla2FubmFkYTEvZGVmYXVsdC9tYXN0ZXIubTN1OA",
// }

var ZEE_CHANNELS = map[string]string{
	"andtv":             "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Andtv1/default/master.m3u8",
	"andtvhd":           "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Andtvhd1/default/master.m3u8",
	"andpictures":       "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Andpictures1/default/master.m3u8",
	"andpictureshd":     "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Andpictureshd1/default/master.m3u8",
	"zeetv":             "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeetv1/default/master.m3u8",
	"zeetvhd":           "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeetvhd1/default/master.m3u8",
	"zee_anmol":         "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeeanmol1/default/master.m3u8",
	"zee_cinema":        "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeecinema1/default/master.m3u8",
	"zee_cinema_hd":     "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeecinemahd1/default/master.m3u8",
	"zee_bollywood":     "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeebollywood1/default/master.m3u8",
	"zee_action":        "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeeaction1/default/master.m3u8",
	"zee_anmol_cinema":  "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeeanmolcinema1/default/master.m3u8",
	"zee_bharat":        "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeehindustan1/default/master.m3u8",
	"zee_business":      "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeebusiness1/default/master.m3u8",
	"zee_salaam":        "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeesalaam1/default/master.m3u8",
	"zee_marathi":       "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeemarathi1/default/master.m3u8",
	"zee_marathi_hd":    "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeemarathihd1/default/master.m3u8",
	"zee_talkies":       "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeetalkies1/default/master.m3u8",
	"zee_talkies_hd":    "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeetalkieshd1/default/master.m3u8",
	"zee_bangla":        "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeebangla1/default/master.m3u8",
	"zee_bangla_hd":     "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeebanglahd1/default/master.m3u8",
	"zee_bangla_cinema": "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeebanglacinema1/default/master.m3u8",
	"zee_tamil":         "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeetamil1/default/master.m3u8",
	"zee_cinemalu":      "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeecinemalu1/default/master.m3u8",
	"zee_kannada":       "https://d1g8wgjurz8via.cloudfront.net/bpk-tv/Zeekannada1/default/master.m3u8",
}

var ZEE_JIO_MAP = map[string]string{
	"zl2024": "andtv",
	"zl472":  "andtvhd",
	"zl1839": "andpictures",
	"zl185":  "andpictureshd",
	"zl1351": "zeetv",
	"zl167":  "zeetvhd",
	"zl473":  "zee_anmol",
	"zl484":  "zee_cinema",
	"zl165":  "zee_cinema_hd",
	"zl487":  "zee_bollywood",
	"zl488":  "zee_action",
	"zl415":  "zee_anmol_cinema",
	"zl652":  "zee_bharat",
	"zl657":  "zee_business",
	"zl728":  "zee_salaam",
	"zl445":  "zee_marathi",
	"zl1360": "zee_marathi_hd",
	"zl153":  "zee_talkies",
	"zl1358": "zee_talkies_hd",
	"zl625":  "zee_bangla",
	"zl1977": "zee_bangla_hd",
	"zl685":  "zee_bangla_cinema",
	"zl628":  "zee_tamil",
	"zl413":  "zee_cinemalu",
	"zl689":  "zee_kannada",
}

var ZEE_CHANNELS_API = []Channel{
	{
		ID:       "zl2024",
		Name:     "ZL And TV",
		Language: 1, // Hindi
		Category: 5, // Entertainment
		IsHD:     true,
		LogoURL:  "And_TV.png",
	},
	{
		ID:       "zl472",
		Name:     "ZL And TV HD",
		Language: 1, // Hindi
		Category: 5, // Entertainment
		IsHD:     true,
		LogoURL:  "And_TV_HD.png",
	},
	{
		ID:       "zl1839",
		Name:     "ZL And Pictures",
		Language: 1, // Hindi
		Category: 6, // Movies
		IsHD:     true,
		LogoURL:  "And_Pictures.png",
	},
	{
		ID:       "zl185",
		Name:     "ZL And Pictures HD",
		Language: 1, // Hindi
		Category: 6, // Movies
		IsHD:     true,
		LogoURL:  "And_Pictures_HD.png",
	},
	{
		ID:       "zl1351",
		Name:     "ZL Zee TV",
		Language: 1, // Hindi
		Category: 5, // Entertainment
		IsHD:     true,
		LogoURL:  "Zee_TV.png",
	},
	{
		ID:       "zl167",
		Name:     "ZL Zee TV HD",
		Language: 1, // Hindi
		Category: 5, // Entertainment
		IsHD:     true,
		LogoURL:  "Zee_TV_HD.png",
	},
	{
		ID:       "zl473",
		Name:     "ZL Zee Anmol",
		Language: 1, // Hindi
		Category: 5, // Entertainment
		IsHD:     false,
		LogoURL:  "Zee_Anmol.png",
	},
	{
		ID:       "zl484",
		Name:     "ZL Zee Cinema",
		Language: 1, // Hindi
		Category: 6, // Movies
		IsHD:     false,
		LogoURL:  "Zee_Cinema.png",
	},
	{
		ID:       "zl165",
		Name:     "ZL Zee Cinema HD",
		Language: 1, // Hindi
		Category: 6, // Movies
		IsHD:     true,
		LogoURL:  "Zee_Cinema_HD.png",
	},
	{
		ID:       "zl487",
		Name:     "ZL Zee Bollywood",
		Language: 1, // Hindi
		Category: 6, // Movies
		IsHD:     false,
		LogoURL:  "Zee_Classic.png",
	},
	{
		ID:       "zl488",
		Name:     "ZL Zee Action",
		Language: 1, // Hindi
		Category: 6, // Movies
		IsHD:     false,
		LogoURL:  "Zee_Action.png",
	},
	{
		ID:       "zl415",
		Name:     "ZL Zee Anmol Cinema",
		Language: 1, // Hindi
		Category: 6, // Movies
		IsHD:     false,
		LogoURL:  "Zee_Anmol_Cinema.png",
	},
	{
		ID:       "zl652",
		Name:     "ZL Zee Bharat",
		Language: 1,  // Hindi
		Category: 12, // News
		IsHD:     false,
		LogoURL:  "Zee_SANGAM.png",
	},
	{
		ID:       "zl657",
		Name:     "ZL Zee Business",
		Language: 1,  // Hindi
		Category: 16, // Business
		IsHD:     false,
		LogoURL:  "Zee_Business.png",
	},
	{
		ID:       "zl728",
		Name:     "ZL Zee Salaam",
		Language: 4,  // Urdu
		Category: 12, // News
		IsHD:     false,
		LogoURL:  "Zee_Salaam.png",
	},
	{
		ID:       "zl445",
		Name:     "ZL Zee Marathi",
		Language: 2, // Marathi
		Category: 5, // Entertainment
		IsHD:     false,
		LogoURL:  "Zee_Marathi.png",
	},
	{
		ID:       "zl1360",
		Name:     "ZL Zee Marathi HD",
		Language: 2, // Marathi
		Category: 5, // Entertainment
		IsHD:     true,
		LogoURL:  "Zee_Marathi_HD.png",
	},
	{
		ID:       "zl153",
		Name:     "ZL Zee Talkies",
		Language: 2, // Marathi
		Category: 6, // Movies
		IsHD:     false,
		LogoURL:  "Zee_Talkies.png",
	},
	{
		ID:       "zl1358",
		Name:     "ZL Zee Talkies HD",
		Language: 2, // Marathi
		Category: 6, // Movies
		IsHD:     true,
		LogoURL:  "Zee_Talkies_HD.png",
	},
	{
		ID:       "zl625",
		Name:     "ZL Zee Bangla",
		Language: 5, // Bengali
		Category: 5, // Entertainment
		IsHD:     false,
		LogoURL:  "Zee_Bangla.png",
	},
	{
		ID:       "zl1977",
		Name:     "ZL Zee Bangla HD",
		Language: 5, // Bengali
		Category: 5, // Entertainment
		IsHD:     true,
		LogoURL:  "Zee_Bangla_HD.png",
	},
	{
		ID:       "zl685",
		Name:     "ZL Zee Bangla Cinema",
		Language: 5, // Bengali
		Category: 6, // Movies
		IsHD:     false,
		LogoURL:  "Zee_Bangla_Cinema.png",
	},
	{
		ID:       "zl628",
		Name:     "ZL Zee Tamil",
		Language: 8, // Tamil
		Category: 5, // Entertainment
		IsHD:     false,
		LogoURL:  "Zee_Tamil.png",
	},
	{
		ID:       "zl413",
		Name:     "ZL Zee Cinemalu",
		Language: 11, // Telugu
		Category: 6,  // Movies
		IsHD:     false,
		LogoURL:  "Zee_Cinemalu.png",
	},
	{
		ID:       "zl689",
		Name:     "ZL Zee Kannada",
		Language: 13, // Kannada
		Category: 5,  // Entertainment
		IsHD:     false,
		LogoURL:  "Zee_Kannada.png",
	},
}
