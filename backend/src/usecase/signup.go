package UseCases

type (
	SignupUseCase interface {
		Execute()
	}
	SignupInput struct {
	}
	SignupPresenter interface {
		Output()
	}
	SignupOutput struct {
	}
	SignupInteractor struct {
	}
)

func NewSignupInteractor() {

}

func (interactor *SignupInteractor) Execute() {

}
