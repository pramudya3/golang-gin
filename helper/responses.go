package helper

type (
	ResponseSuccess struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta,omitempty"`
	}

	ResponseFailed struct {
		Message interface{} `json:"message"`
	}
)
