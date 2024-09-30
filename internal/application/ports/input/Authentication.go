package inputports

type Authentication struct {
	UserId string
}

func UserAuthentication(userId string) *Authentication {
	return &Authentication{
		UserId: userId,
	}
}
