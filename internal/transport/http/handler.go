package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) HandleUserByID() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			path := r.URL.Path
			parts := strings.Split(path, "/")
			id, _ := strconv.Atoi(parts[2])
			user, err := s.userService.GetByID(ctx, id)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			resp, _ := json.Marshal(user)
			w.Write(resp)
		},
	)
}
