package apierror

type GenericErrorCode string

const (
	Generic         GenericErrorCode = "GENERIC"
	UnexpectedError GenericErrorCode = "UNEXPECTED_ERROR"
)

type getPropertyByName func(errorCode GenericErrorCode, propertyName string) string

var GetPropertyByName getPropertyByName
var ErrorCode = "error-code"
var ErrorMessage = "error-message"

func (g GenericErrorCode) GetCodeByEnv() string {
	return GetPropertyByName(g, ErrorCode)
}
func (g GenericErrorCode) GetMessageByEnv() string {
	return GetPropertyByName(g, ErrorMessage)
}
