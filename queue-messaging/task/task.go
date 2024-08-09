package task

const (
	SendEmailTask   = "SendEmail"
	UploadImageTask = "UploadImage"
)

type (
	SendEmailPayload struct {
		From    string
		To      string
		Subject string
		Message string
	}

	UploadImagePayload struct {
		ImageSource string
		TargetPath  string
	}
)
