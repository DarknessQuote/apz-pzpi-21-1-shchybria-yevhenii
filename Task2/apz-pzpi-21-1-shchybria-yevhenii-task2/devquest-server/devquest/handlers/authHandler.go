package handlers

import (
	"devquest-server/devquest/infrastructure"
	"devquest-server/devquest/utils"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Auth struct {
	*infrastructure.Auth
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var reqPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := utils.ReadJSON(w, r, &reqPayload); err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	log.Println(reqPayload.Username, reqPayload.Password)

	user := infrastructure.JWTUser{
		ID: uuid.New(),
		Username: reqPayload.Username,
		RoleName: "user",
	}

	tokens, err := a.GenerateTokenPairs(&user)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	refreshCookie := a.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	utils.WriteJSON(w, http.StatusAccepted, tokens)
}