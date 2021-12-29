package models

type TransactionType int

const (
	Earn  TransactionType = 0
	Spend TransactionType = 1
)

type TransactionEntity struct {
	ID              uint64
	Amount          float64
	TransactionType TransactionType
	Description     string
	TagIDs          []uint64
	Owner           uint64
	WalletID        uint64
	IsDeleted       bool
}

type TransactionCreateReq struct {
	Amount          float64         `json:"amount"`
	TransactionType TransactionType `json:"type"`
	Description     string          `json:"description"`
	WalletID        uint64          `json:"wallet_id"`
	TagIDs          []uint64        `json:"tags"`
}

type TransactionUpdateReq struct {
	ID              uint64          `json:"id"`
	Amount          float64         `json:"amount"`
	TransactionType TransactionType `json:"type"`
	Description     string          `json:"description"`
	WalletID        uint64          `json:"wallet_id"`
	TagIDs          []uint64        `json:"tags"`
}

type TransactionRes struct {
	ID              uint64          `json:"id"`
	Amount          float64         `json:"amount"`
	TransactionType TransactionType `json:"type"`
	Description     string          `json:"description"`
	WalletID        uint64          `json:"wallet_id"`
	TagIDs          []uint64        `json:"tags"`
}
