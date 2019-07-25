package http

import (
	"encoding/json"
	"go-rest-chat/src/api/domain/user/entities"
	"net/http"
)

func (handler *CheckHandler) Check(w http.ResponseWriter, r *http.Request) {
	ok, err := handler.usecases.Check()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}

	if ok {
		if err = json.NewEncoder(w).Encode(map[string]string{"health": "ok"}); err != nil {
			http.Error(w, "Write error", http.StatusInternalServerError)
		}
	}
}
