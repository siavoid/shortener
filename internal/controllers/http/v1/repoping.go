package v1

import "net/http"

func (s *Server) repoPing(w http.ResponseWriter, r *http.Request) {
	err := s.u.RepoPing(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
