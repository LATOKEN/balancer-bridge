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
	SwapID             string
	TxType             TxType      `gorm:"type:tx_types"`
	TxHash             string      `gorm:"type:TEXT"`
	OriginСhainID      string      `gorm:"type:TEXT"`
	DestinationChainID string      `gorm:"type:TEXT"`
	ReceiverAddr       string      `gorm:"type:TEXT"`
	ResourceID         string      `gorm:"type:TEXT"`
	BlockHash          string      `gorm:"type:TEXT"`
	Height             int64       `gorm:"type:BIGINT"`
	Status             TxLogStatus `gorm:"type:tx_log_statuses"`
	EventStatus        EventStatus `gorm:"type:TEXT"`
	ConfirmedNum       uint32      `gorm:"type:BIGINT"`
	CreateTime         int64       `gorm:"type:BIGINT"`
	UpdateTime         int64       `gorm:"type:BIGINT"`
	DepositNonce       uint64      `gorm:"type:BIGINT"`
	InAmount           string      `gorm:"type:TEXT"`
	Data               string      `gorm:"type:TEXT"`
}

// Event ...
type Event struct {
	SwapID             string      `json:"swap_id" gorm:"primaryKey"`
	ChainID            string      `json:"chain_id" gorm:"primaryKey"`
	DestinationChainID string      `json:"destination_chain_id" gorm:"TEXT"`
	OriginChainID      string      `json:"origin_chain_id" gorm:"TEXT"`
	ReceiverAddr       string      `json:"receiver_addr" gorm:"TEXT"`
	InAmount           string      `json:"in_amount" gorm:"TEXT"`
	ResourceID         string      `json:"resource_id" gorm:"TEXT"`
	Height             int64       `json:"height" gorm:"BIGINT"`
	Status             EventStatus `json:"status" gorm:"TEXT"`
	DepositNonce       uint64      `json:"deposit_nonce" gorm:"BIGINT"`
	CreateTime         int64       `json:"create_time" gorm:"BIGINT"`
	UpdateTime         int64       `json:"update_time" gorm:"BIGINT"`
	Data               string      `json:"data" gorm:"TEXT"`
}

// TxSent ...
type TxSent struct {
	ID         int64    `json:"id"`
	Chain      string   `json:"chain" gorm:"type:TEXT"`
	SwapID     string   `json:"swap_id" gorm:"type:TEXT"`
	Type       TxType   `json:"type" gorm:"type:tx_types"`
	TxHash     string   `json:"tx_hash" gorm:"type:TEXT"`
	ErrMsg     string   `json:"err_msg" gorm:"type:TEXT"`
	Status     TxStatus `json:"status" gorm:"type:tx_statuses"`
	CreateTime int64    `json:"create_time" gorm:"type:BIGINT"`
	UpdateTime int64    `json:"update_time" gorm:"type:BIGINT"`
}

// PriceLog...
type PriceLog struct {
	Name       string `gorm:"primaryKey"`
	Price      string `gorm:"type:TEXT"`
	UpdateTime int64  `json:"update_time" gorm:"type:BIGINT"`
}

type ResourceId struct {
	Name string `gorm:"primaryKey"`
	ID   string `gorm:"type:TEXT"`
}
