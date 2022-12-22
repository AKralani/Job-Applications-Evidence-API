package evidence

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/akralani/jobapplications/cors"
	"golang.org/x/net/websocket"
)

const evidencesPath = "evidences"

func handleEvidences(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		evidenceList, err := getEvidenceList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(evidenceList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var evidence Evidence
		err := json.NewDecoder(r.Body).Decode(&evidence)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		evidenceID, err := insertEvidence(evidence)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"evidenceId":%d}`, evidenceID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleEvidence(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", evidencesPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	evidenceID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		evidence, err := getEvidence(evidenceID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if evidence == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(evidence)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPut:
		var evidence Evidence
		err := json.NewDecoder(r.Body).Decode(&evidence)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if *evidence.EvidenceID != evidenceID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateEvidence(evidence)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		err := removeEvidence(evidenceID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// SetupRoutes:
func SetupRoutes(apiBasePath string) {
	evidencesHandler := http.HandlerFunc(handleEvidences)
	evidenceHandler := http.HandlerFunc(handleEvidence)
	http.Handle("/websocket", websocket.Handler(evidenceSocket))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, evidencesPath), cors.Middleware(evidencesHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, evidencesPath), cors.Middleware(evidenceHandler))
}
