package actions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/factly/data-portal-api/models"
)

// tag request body
type tag struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

// GetTags - Get all tags
// @Summary Show all tags
// @Description Get all tags
// @Tags Tag
// @ID get-all-tags
// @Produce  json
// @Param limit query string false "limt per page"
// @Param page query string false "page number"
// @Success 200 {array} models.Tag
// @Router /tags [get]
func GetTags(w http.ResponseWriter, r *http.Request) {

	var tags []models.Tag
	p := r.URL.Query().Get("page")
	pg, _ := strconv.Atoi(p) // pg contains page number
	l := r.URL.Query().Get("limit")
	li, _ := strconv.Atoi(l) // li contains perPage number

	offset := 0 // no. of records to skip
	limit := 5  // limt

	if li > 0 && li <= 10 {
		limit = li
	}

	if pg > 1 {
		offset = (pg - 1) * limit
	}

	models.DB.Offset(offset).Limit(limit).Model(&models.Tag{}).Find(&tags)

	json.NewEncoder(w).Encode(tags)
}
