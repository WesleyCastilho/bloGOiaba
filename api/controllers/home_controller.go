package controllers

import (
	"github.com/WesleyCastilho/blog/api/responses"
	"net/http"
)


func (server *Server) Home(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "Bem vindo ao Blog do Goiaba")
}