package membership

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/data-portal-server/validation"
	"github.com/go-chi/chi"
)

// delete - Delete membership by id
// @Summary Delete a membership
// @Description Delete membership by ID
// @Tags Membership
// @ID delete-membership-by-id
// @Consume  json
// @Param id path string true "Membership ID"
// @Success 200 {object} model.Membership
// @Failure 400 {array} string
// @Router /memberships/{id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

	membershipID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(membershipID)
	if err != nil {
		validation.InvalidID(w, r)
		return
	}

	membership := &model.Membership{
		ID: uint(id),
	}

	// check record exists or not
	err = model.DB.First(&membership).Error
	if err != nil {
		validation.RecordNotFound(w, r)
		return
	}
	model.DB.Delete(&membership)

	json.NewEncoder(w).Encode(membership)
}