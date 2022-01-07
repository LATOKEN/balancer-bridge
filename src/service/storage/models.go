package storage

// BlockLog ...
type BlockLog struct {
	Chain      string    `gorm:"type:TEXT"`
	BlockHash  string    `gorm:"type:TEXT"`
	ParentHash string    `gorm:"type:TEXT"`
	Height     int64     `gorm:"type:BIGINT"`
	BlockTime  int64     `gorm:"type:BIGINT"`
	Type       BlockType `gorm:"block_type"`
	CreateTime int64     `gorm:"type:BIGINT"`
}

// TxLog ...
type TxLog struct {
	ID                 int64
	Chain              string `gorm:"type:TEXT"`
	EventID            string
	TxType             TxType      `gorm:"type:tx_types"`
	TxHash             string      `gorm:"type:TEXT"`
	OriginСhainID      string      `gorm:"type:TEXT"`
	DestinationChainID string      `gorm:"type:TEXT"`
	ReceiverAddr       string      `gorm:"type:TEXT"`
	BlockHash          string      `gorm:"type:TEXT"`
	Height             int64       `gorm:"type:BIGINT"`
	Status             TxLogStatus `gorm:"type:tx_log_statuses"`
	EventStatus        EventStatus
	CreateTime         int64  `gorm:"type:BIGINT"`
	UpdateTime         int64  `gorm:"type:BIGINT"`
	DepositNonce       uint64 `gorm:"type:BIGINT"`
	InAmount           string `gorm:"type:TEXT"`
	ConfirmedNum       int64  `gorm:"type:BIGINT"`
}

// Event ...
type Event struct {
	ID                 int64
	EventID            string
	Type               EventStatus
	ChainID            string
	DestinationChainID string
	OriginChainID      string
	ReceiverAddr       string
	InAmount           string
	OutAmount          uint64
	Height             int64
	Status             EventStatus
	CreateTime         int64
	UpdateTime         int64
	TxType             string
}

// TxSent ...
type TxSent struct {
	ID         int64    `json:"id"`
	Chain      string   `json:"chain" gorm:"type:TEXT"`
	Type       TxType   `json:"type" gorm:"type:tx_types"`
	TxHash     string   `json:"tx_hash" gorm:"type:TEXT"`
	ErrMsg     string   `json:"err_msg" gorm:"type:TEXT"`
	Status     TxStatus `json:"status" gorm:"type:tx_statuses"`
	CreateTime int64    `json:"create_time" gorm:"type:BIGINT"`
	UpdateTime int64    `json:"update_time" gorm:"type:BIGINT"`
}

//PriceLog...
type PriceLog struct {
	Name       string `gorm:"primaryKey"`
	Price      string `gorm:"type:TEXT"`
	UpdateTime int64  `json:"update_time" gorm:"type:BIGINT"`
}
