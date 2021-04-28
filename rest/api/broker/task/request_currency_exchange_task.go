package task

type CurrencyExchangeTask struct {
	ResultQueueName     string  `json:"result_queue_name"`
	Value               float64 `json:"value"`
	CurrentCurrencyName string  `json:"currentCurrencyName"`
	NewCurrencyName     string  `json:"newCurrencyName"`
	Result              float64 `json:"result"`
}

func BuildCurrencyExchangeTaskForUser(
	resultQueueName string,
	value float64,
	currentCurrencyName string,
	newCurrencyName string) CurrencyExchangeTask {

	return CurrencyExchangeTask{
		ResultQueueName:     resultQueueName,
		Value:               value,
		CurrentCurrencyName: currentCurrencyName,
		NewCurrencyName:     newCurrencyName,
		Result:              0,
	}
}
