package handler

type Handler interface {
	AuthHandler() *AuthHandler
	NFTHandler() *NFTHandler
}

type handler struct {
	nftHandler  *NFTHandler
	authHandler *AuthHandler
}

// nolint:ireturn
func NewHandler(nftHandler *NFTHandler, authHandler *AuthHandler) Handler {
	return &handler{
		nftHandler:  nftHandler,
		authHandler: authHandler,
	}
}

func (h handler) NFTHandler() *NFTHandler {
	return h.nftHandler
}

func (h handler) AuthHandler() *AuthHandler {
	return h.authHandler
}
