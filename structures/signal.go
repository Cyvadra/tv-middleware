package structures

// JSON signal from TradingView
type TvcSignal struct {
	Ticker                 string `json:"ticker" eg:"REEFUSDT.P"`
	Ex                     string `json:"ex" eg:"BITGET"`
	Close                  string `json:"close" eg:"0.004699"`
	Open                   string `json:"open" eg:"0.004381"`
	High                   string `json:"high" eg:"0.004699"`
	Low                    string `json:"low" eg:"0.00433"`
	Time                   string `json:"time" eg:"2024-09-26T20:00:00Z"`
	Volume                 string `json:"volume" eg:"95896516"`
	Timenow                string `json:"timenow" eg:"2024-09-26T21:35:31Z"`
	Interval               string `json:"interval" eg:"240"`
	PositionSize           string `json:"position_size" eg:"655780.706"`
	Action                 string `json:"action" eg:"buy"`
	Contracts              string `json:"contracts" eg:"655780.706"`
	Price                  string `json:"price" eg:"0.004699"`
	Id                     string `json:"id" eg:"REEFUSDT_UMCBL"`
	MarketPosition         string `json:"market_position" eg:"long"`
	MarketPositionSize     string `json:"market_position_size" eg:"655780.706"`
	PrevMarketPosition     string `json:"prev_market_position" eg:"flat"`
	PrevMarketPositionSize string `json:"prev_market_position_size" eg:"0"`
	Exchange               string `json:"exchange" eg:"bitget"`
	Lever                  string `json:"lever" eg:"1"`
	TdMode                 string `json:"td_mode" eg:"isolated"`
	Symbol                 string `json:"symbol" eg:"REEFUSDT_UMCBL"`
	OrdType                string `json:"ord_type" eg:"market"`
	OrdBase                string `json:"ord_base" eg:"qty"`
	Amount                 string `json:"amount" eg:""`
	StrategyMethod         string `json:"strategy_method" eg:"auto"`
	Delay                  string `json:"delay" eg:"0"`
	SltpType               string `json:"sltp_type" eg:"no"`
	Aid                    string `json:"aid" eg:"0"`
	ApiSec                 string `json:"api_sec" eg:"*********"`
}

type TvFormatedSignal struct {
	Action       string
	Contracts    float32
	Ticker       string
	PositionSize float32
	ExitCode     int8
}
