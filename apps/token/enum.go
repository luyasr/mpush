package token

type QueryBy int

const (
	QueryByUserId QueryBy = iota
	QueryByAccessToken
	QueryByRefreshToken
	QueryByARToken
)
