//go:generate ~/go/bin/mockgen -source ./notify.go -destination ../mock/notify_mock.go -package domain_mock

package domain

type NotifyService interface {
	NotifySlackError(text string)
	NotifyLog(text string) error
}
