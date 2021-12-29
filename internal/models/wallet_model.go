package models

type WalletEntity struct {
	ID        uint64
	Title     string
	IsDeleted bool
	OwnerID   uint64
}

type WalletCreateReq struct {
	Title string `json:"title"`
}

type WalletUpdareReq struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type WalletRes struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}
