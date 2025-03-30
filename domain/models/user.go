package models

type User struct {
	Id           uint
	ProviderId   string
	ProviderName string
	UserName     string
}

func NewUser(id uint, providerId, providerName, UserName string) *User {
	return &User{
		Id:           id,
		ProviderId:   providerId,
		ProviderName: providerName,
		UserName:     UserName,
	}
}
