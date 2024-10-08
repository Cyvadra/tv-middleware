package structures

// aggressive trade mode
type PositionDiff struct {
	// ensure market position is equal to TradingView
	Ticker             string
	TargetPositionSize float32 // negative -> short position
	PrevPositionSize   float32
	TimestampCreated   float64
}

type Broker interface {
	Connect() error
	Close() error
	Equity() float32
	FreeMargin() float32
	SupportTicker(string) bool
	CurrentPositionOnTicker(string) float32
	EnsurePositionDiff(*PositionDiff) error
}
