package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tursynkhan/todo-app"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok:= c.Get(userCtx)
	if !ok{
		newErrorResponse(c,http.StatusInternalServerError,"user id not found")
	}
	var input todo.Todolist
	if err:=c.BindJSON(&input);err!=nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	
}

func (h *Handler) getAlllists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
