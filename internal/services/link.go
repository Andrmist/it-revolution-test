package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Andrmist/it-revolution-test-mine/internal/domain"
	"github.com/Andrmist/it-revolution-test-mine/internal/types"
	"github.com/Andrmist/it-revolution-test-mine/internal/utils"
	"github.com/go-chi/chi/v5"
)

type GetOriginalLinkRequest struct {
	OriginalLink string `validate:"required" json:"original_link"`
}

type GetStatisticsResponse struct {
	ShortLink string `json:"short_link"`
	Count     int    `json:"count"`
}

func TransformLink(w http.ResponseWriter, r *http.Request) {
	serverCtx := r.Context().Value("server").(types.ServerContext)
	var data GetOriginalLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var result domain.Link
	serverCtx.DB.Where(domain.Link{OriginalLink: data.OriginalLink}).FirstOrCreate(&result)

	w.Write([]byte(fmt.Sprintf("%s/%s", serverCtx.Config.BaseURL, result.ID)))
}

func GetOriginalLink(w http.ResponseWriter, r *http.Request) {
	serverCtx := r.Context().Value("server").(types.ServerContext)
	id := chi.URLParam(r, "id")

	var result domain.Link
	serverCtx.DB.Where(domain.Link{ID: id}).First(&result)
	result.Count++
	if err := serverCtx.DB.Save(&result).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error", err)
	}

	w.Write([]byte(result.OriginalLink))
}

func GetStatistics(w http.ResponseWriter, r *http.Request) {
	serverCtx := r.Context().Value("server").(types.ServerContext)
	var links []domain.Link
	serverCtx.DB.Find(&links)
	var response []GetStatisticsResponse
	for _, link := range links {
		response = append(response, GetStatisticsResponse{
			ShortLink: fmt.Sprintf("%s/%s", serverCtx.Config.BaseURL, link.ID),
			Count:     link.Count,
		})
	}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error", err)
	}
}
