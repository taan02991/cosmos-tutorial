package rest

import (
	"github.com/gorilla/mux"

	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// RegisterRoutes registers blog-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding
	r.HandleFunc("/blog/comment", listCommentHandler(cliCtx, "blog")).Methods("GET")
	r.HandleFunc("/blog/comment", createCommentHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/blog/post", listPostHandler(cliCtx, "blog")).Methods("GET")
	r.HandleFunc("/blog/post", createPostHandler(cliCtx)).Methods("POST")

}

func listPostHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-post", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
