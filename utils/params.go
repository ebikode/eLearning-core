package utils

import (
	"net/http"
	"strconv"
)

func PaginationParams(r *http.Request) (int,int) {
	params := r.URL.Query()
	pgParam := params.Get("page")
	lmParam := params.Get("limit")
		var page int = 1
		var limit int = 20


		if len(pgParam) > 0 {
			pg, err := strconv.ParseInt(pgParam,0,64)
			if err == nil {
				page = int(pg)
			}
		}
		if len(lmParam) > 0 {
			lm, err := strconv.ParseUint(lmParam,0,64)
			if err == nil {
				limit = int(lm)
			}
		}

		return page, limit
}