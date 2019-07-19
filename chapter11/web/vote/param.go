package vote

type CreateVoteParam struct {
	Data struct {
		Title       string `json:"title" form:"title" validate:"required"`
		Description string `json:"description" form:"description" validate:"required"`
	} `json:"data" validate:"required"`
}

type PatchVoteParam struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}

type Register struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"min=8,max=16"`
}

func (r Register) ValidAccount() error {
	// 具体的 r.Account 校验
	return nil
}

func (r Register) ValidPassword() error {
	// 具体的 r.Password 校验
	return nil
}
