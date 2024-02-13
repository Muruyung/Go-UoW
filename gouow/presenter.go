package gouow

type (
	// TX transaction structure for keep transaction sql
	TX struct {
		Tx    interface{}
		UseTx bool
	}

	key string
)

// TX_KEY key for store and get value from context
const TX_KEY key = "unit_of_work"
