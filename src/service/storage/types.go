package storage

var (
	LaChain string = "LA"
	// PosChain - polygon mumbai chain
	PosChain string = "POS"
	// BscChain - BSC testnet chain
	BscChain string = "BSC"
	//EthChain - Ropsten testnet chain
	EthChain string = "ETH"
	// BtcChain - Bitcoin chain
	BtcChain string = "BTC"
	// AvaxChain - Avalanche chain
	AvaxChain string = "AVAX"
	// FtmChain - Fantom chain
	FtmChain string = "FTM"
)

// BlockType ...
type BlockType string

const (
	BlockTypeCurrent BlockType = "CURRENT"
	BlockTypeParent  BlockType = "PARENT"
)

// TxStatus ...
type TxStatus string

const (
	TxSentStatusInit     TxStatus = "INIT"
	TxSentStatusNotFound TxStatus = "NOT_FOUND"
	TxSentStatusPending  TxStatus = "PENDING"
	TxSentStatusFailed   TxStatus = "FAILED"
	TxSentStatusSuccess  TxStatus = "SUCCESS"
	TxSentStatusLost     TxStatus = "LOST"
)

// !!! TODO !!!

//
type TxType string

const (
	TxTypeFeeTransfer TxType = "FEE_TRANSFER"
)

type EventStatus string

const (
	// FEE_TRANSFER
	EventStatusFeeTransferInit          EventStatus = "FEE_TRANSFER_INIT"
	EventStatusFeeTransferInitConfrimed EventStatus = "FEE_TRANSFER_INIT_CONFIRMED"
	EventStatusFeeTransferSent          EventStatus = "FEE_TRANSFER_SENT"
	EventStatusFeeTransferConfirmed     EventStatus = "FEE_TRANSFER_CONFIRMED"
	EventStatusFeeTransferFailed        EventStatus = "FEE_TRANSFER_FAILED"
	EventStatusFeeTransferSentFailed    EventStatus = "FEE_TRANSFER_SENT_FAILED"
)

// TxLogStatus ...
type TxLogStatus string

const (
	TxStatusInit      TxLogStatus = "INIT"
	TxStatusConfirmed TxLogStatus = "CONFIRMED"
)
