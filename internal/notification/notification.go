package notification

type NotificationService interface {
	Send(t NotificationType, userId string, message string)
}

type NotificationServiceImpl struct {
	gateway *Gateway
	rl      *RateLimiter
}

func NewNotificationServiceImpl(gateway *Gateway, rl *RateLimiter) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		gateway: gateway,
		rl:      rl,
	}
}

func (n *NotificationServiceImpl) Send(notificationType NotificationType, userId string, message string) bool {
	ok := n.rl.Allow(userId, notificationType)

	if !ok {
		return false
	}

	n.gateway.send(userId, message)
	return true
}
