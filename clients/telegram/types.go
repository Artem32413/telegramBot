package telegram

type UpdateResponse struct{
	Ok bool `json:"ok"`
	Result []Updates `json:"result"`
}

type Updates struct{
	ID int64 `json:"uppdate_id"`
	Message string `json:"message"`
}