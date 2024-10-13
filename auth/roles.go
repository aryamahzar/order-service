package auth

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

var rolePermissions = map[Role][]string{
	Admin: {
		"CreateOrder",
		"ListOrders",
		"GetOrder",
		"UpdateOrder",
		"CancelOrder",
	},
	User: {
		"CreateOrder",
		"ListOrders",
		"GetOrder",
		"UpdateOrder",
		"CancelOrder",
	},
}

func HasPermission(role Role, permission string) bool {
	for _, p := range rolePermissions[role] {
		if p == permission {
			return true
		}
	}
	return false
}
