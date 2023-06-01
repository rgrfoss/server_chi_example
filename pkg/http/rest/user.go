package rest

import (
	"fmt"
	"net/http"
)

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	// get user
	userId := r.URL.Query().Get("user_id")
	fmt.Println("userID", userId)
}
