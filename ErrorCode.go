package apierror

type GenericErrorCode string

const (
	Generic         GenericErrorCode = "GENERIC"
	UnexpectedError GenericErrorCode = "UNEXPECTED_ERROR"
)

type getPropertyByName func(errorCode GenericErrorCode, propertyName string) string

var GetPropertyByName getPropertyByName

func (g GenericErrorCode) GetCodeByEnv() string {
	return GetPropertyByName(g, "error-code")
}
func (g GenericErrorCode) GetMessageByEnv() string {
	return GetPropertyByName(g, "error-message")
}
