package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const AdminLogCreateNews string = "CreateNews"
const AdminLogPatchNews string = "PatchNews"
const AdminLogDeleteNews string = "DeleteNews"

const AdminLogPatchUserNodeLevel string = "PatchUserNodeLevel"
const AdminLogPatchUserLevel string = "PatchUserLevel"
const AdminLogPatchUserActive string = "PatchUserActive"
const AdminLogPatchUserWithdraw string = "PatchUserWithdraw"
const AdminLogPatchUserEmail string = "PatchUserEmail"

const AdminLogPostNodeApprove string = "PostNodeApprove"
const AdminLogPostReWithdraw string = "PostReWithdraw"

type AdminLog struct {
	ID         uint64
	UUID       uuid.UUID
	Account    string
	LogType    string
	BeforeData string
	AfterData  string
	UpdatedAt  time.Time
	CreatedAt  time.Time
}
