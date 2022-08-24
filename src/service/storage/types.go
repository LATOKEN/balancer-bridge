package storage

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

type TxType string

const (
	TxTypeFeeTransfer        TxType = "FEE_TRANSFER"
	TxTypeFeeTransferConfirm TxType = "FEE_TRANSFER_CONFIRM"
	TxTypeFeeReversal        TxType = "FEE_REVERSAL"
)

type EventStatus string

const (
	// FEE_TRANSFER
	EventStatusFeeTransferInit          EventStatus = "FEE_TRANSFER_INIT"
	EventStatusFeeTransferInitConfrimed EventStatus = "FEE_TRANSFER_INIT_CONFIRMED"
	EventStatusFeeTransferSent          EventStatus = "FEE_TRANSFER_SENT"
	EventStatusFeeTransferSentConfirmed EventStatus = "FEE_TRANSFER_SENT_CONFIRMED"
	EventStatusFeeTransferConfirmed     EventStatus = "FEE_TRANSFER_CONFIRMED"
	EventStatusFeeTransferFailed        EventStatus = "FEE_TRANSFER_FAILED"
	EventStatusFeeTransferSentFailed    EventStatus = "FEE_TRANSFER_SENT_FAILED"
	EventStatusFeeTransferReversed      EventStatus = "FEE_TRANSFER_REVERSED"

	EventStatusFeeReversalInit       EventStatus = "FEE_REVESAL_INIT"
	EventStatusFeeReversalSent       EventStatus = "FEE_REVERSAL_SENT"
	EventStatusFeeReversalSentFailed EventStatus = "FEE_REVERSAL_SENT_FAILED"
	EventStatusFeeReversalFailed     EventStatus = "FEE_REVERSAL_FAILED"
	EventStatusFeeReversalConfirmed  EventStatus = "FEE_REVERSAL_CONFIRMED"
)

// TxLogStatus ...
type TxLogStatus string

const (
	TxStatusInit      TxLogStatus = "INIT"
	TxStatusConfirmed TxLogStatus = "CONFIRMED"
)
