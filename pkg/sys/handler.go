package sys

import (
	"encoding/json"
	"net/http"
)

func (s *Sys) VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(HeaderContentType, ApplicationJSON)
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	if err := enc.Encode(&s.appInfo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Sys) MetricsHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(HeaderContentType, ApplicationJSON)
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	if err := enc.Encode(&s.appInfo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
