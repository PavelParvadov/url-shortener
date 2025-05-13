package stat

import (
	"net/http"
	"time"
	"url/configs"
	"url/pkg/middleware"
	"url/pkg/res"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

type StatHandler struct {
	StatRepository *StatRepository
}
type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))

}

func (h *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != FilterByDay && by != FilterByMonth {
			http.Error(w, "Invalid param", http.StatusBadRequest)
			return
		}

		stats := h.StatRepository.GetStats(by, from, to)
		res.JsonResponse(w, stats, 200)

	}
}
