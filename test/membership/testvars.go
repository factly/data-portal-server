package membership

import (
	"errors"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var Membership map[string]interface{} = map[string]interface{}{
	"status":     "Test Status",
	"user_id":    1,
	"payment_id": 1,
	"plan_id":    1,
}

var invalidMembership map[string]interface{} = map[string]interface{}{
	"status":    "Test Status",
	"user_id":   0,
	"paymentid": 0,
	"planid":    1,
}

var membershiplist []map[string]interface{} = []map[string]interface{}{
	{
		"status":     "Test Status 1",
		"user_id":    1,
		"payment_id": 1,
		"plan_id":    1,
	},
	{
		"status":     "Test Status 2",
		"user_id":    1,
		"payment_id": 1,
		"plan_id":    1,
	},
}

var MembershipCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "status", "user_id", "payment_id", "plan_id"}

var selectQuery string = regexp.QuoteMeta(`SELECT * FROM "dp_membership"`)
var countQuery string = regexp.QuoteMeta(`SELECT count(*) FROM "dp_membership"`)
var errMembershipUserFK error = errors.New(`pq: insert or update on table "dp_membership" violates foreign key constraint "dp_membership_user_id_dp_user_id_foreign"`)
var errMembershipPlanFK error = errors.New(`pq: insert or update on table "dp_membership" violates foreign key constraint "dp_membership_plan_id_dp_plan_id_foreign"`)
var errMembershipPaymentFK error = errors.New(`pq: insert or update on table "dp_membership" violates foreign key constraint "dp_membership_payment_id_dp_payment_id_foreign"`)

const basePath string = "/memberships"
const path string = "/memberships/{membership_id}"

func MembershipSelectMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(MembershipCols).
			AddRow(1, time.Now(), time.Now(), nil, Membership["status"], Membership["user_id"], Membership["payment_id"], Membership["plan_id"]))
}
