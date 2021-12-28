package models

type WalletEntity struct {
	ID        uint64
	Title     string
	IsDeleted bool
	OwnerID   uint64
}

type WalletCreateReq struct {
	Title string
}

type WalletUpdareReq struct {
	ID    uint64
	Title string
}

type WalletRes struct {
	ID    uint64
	Title string
}
