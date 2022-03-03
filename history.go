package vcapool

type HistoryDate struct {
	From int64 `json:"from" bson:"from"`
	To   int64 `json:"to" bson:"to"`
}

type HistoryDateList []HistoryDate

func (i *HistoryDateList) Add(from int64, to int64) {
	historyDate := HistoryDate{
		From: from,
		To:   to,
	}
	*i = append(*i, historyDate)
}
