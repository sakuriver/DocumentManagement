package define

const ErrorCodeMemoryFailed = "KFSB05200-E"
const ErrorCodeTimerTransactionFailed = "KFSB05300-E"
const ErrorCodeRollBackFailed = "KFSB05400-E"
const ErrorCodeDataBaseQueueWrite = "KFSB06000-E"

var ErrorInfoMap = map[string]string{
	ErrorCodeMemoryFailed:           "SERVER:領域確保に失敗しました。",
	ErrorCodeTimerTransactionFailed: "SERVER:タイマトランザクション起動に失敗しました。",
	ErrorCodeRollBackFailed:         "SERVER:ROLLBACKに失敗しました。",
	ErrorCodeDataBaseQueueWrite:     "SERVER:DBのキューメッセージ書き込みに失敗しました。",
}
