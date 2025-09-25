package notification

type NotificationType int64

const (
	Unknown NotificationType = iota
	Status
	News
	Marketing
)

func (nt NotificationType) String() string {
	switch nt {
	case Status:
		return "Status"
	case News:
		return "News"
	case Marketing:
		return "Marketing"
	}
	return "Unknown"
}
