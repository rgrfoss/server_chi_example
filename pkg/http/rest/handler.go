package rest

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"

	db "github.com/rgrfoss/webserver/pkg/db/sqlc"
	m "github.com/rgrfoss/webserver/pkg/middleware"
)

type Handler struct {
	conn *sql.DB
	q    *db.Queries
}

type HandlerArgs struct {
	Conn *sql.DB
}

func newHandler(args *HandlerArgs) Handler {
	q := db.New(args.Conn)
	return Handler{conn: args.Conn, q: q}
}

func InitHandlers(args *HandlerArgs) http.Handler {
	// init validator single instance.

	h := newHandler(&HandlerArgs{Conn: args.Conn})
	r := chi.NewRouter()

	// Basic CORS
	m.CommonMiddlewares(r)

	// example
	r.Get("/user", h.GetUser)

	return r
}
