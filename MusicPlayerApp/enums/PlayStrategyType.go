package enums

type PlayStrategyType string

const (
	SEQUENTIAL   PlayStrategyType = "SEQUENTIAL"
	RANDOM       PlayStrategyType = "RANDOM"
	CUSTOM_QUEUE PlayStrategyType = "CUSTOM_QUEUE"
)
