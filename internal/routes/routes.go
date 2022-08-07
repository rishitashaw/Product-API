package routes

import (
	"chi"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	// routes
	return &Router{
		config: NewConfig().SetTimeout(serverConfig.GetConfig().GetTimeout()),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouter() *chi.Mux{

}

func (r *Router) SetConfigsRouter() {

}

func RouterHealth(){

}

func RouterProduct(){

}

func EnableTimeout(){

}

func EnableCors(){
	
}

func EnableRecovery(){

}

func EnableRequestID(){

}

func EnableRealIP(){
	
}