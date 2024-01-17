package handler

type Handler interface {
	NFTHandler() *NFTHandler
}

type handler struct {
	nftHandler *NFTHandler
}

// nolint:ireturn
func NewHandler(nftHandler *NFTHandler) Handler {
	return &handler{
		nftHandler: nftHandler,
	}
}

func (h handler) NFTHandler() *NFTHandler {
	return h.nftHandler
}
