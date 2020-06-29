package endpoints

import (
	"net/http"

	al "github.com/ebikode/eLearning-core/domain/activity_log"
	ut "github.com/ebikode/eLearning-core/utils"
)

// GetActivityLogsEndpoint ...
func GetActivityLogsEndpoint(cs al.ActivityLogService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		page, limit := ut.PaginationParams(r)

		// var activityLogs []*ActivityLog
		activityLogs := cs.GetActivityLogs(page, limit)

		var nextPage int

		if len(activityLogs) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["activity_logs"] = activityLogs
		ut.Respond(w, r, resp)
	}
}
