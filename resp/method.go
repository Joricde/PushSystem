package resp

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Option interface {
	apply(*Response)
}

type optionFunc func(*Response)

func (f optionFunc) apply(r *Response) {
	f(r)
}

func New(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResp(opts ...Option) *Response {
	resp := &Response{
		Code:    SUCCESS,
		Message: Flags[SUCCESS],
		Data:    nil,
	}
	return resp.withOptions(opts...)
}

func NewErrorRResp(opts ...Option) *Response {
	resp := &Response{
		Code:    ERROR,
		Message: Flags[ERROR],
		Data:    nil,
	}
	return resp.withOptions(opts...)
}

func NewInvalidResp(opts ...Option) *Response {
	resp := &Response{
		Code:    InvalidParams,
		Message: Flags[InvalidParams],
		Data:    nil,
	}
	return resp.withOptions(opts...)
}

func NewDefaultSuccessResp(opts ...Option) *Response {
	resp := &Response{
		Code:    SUCCESS,
		Message: Flags[SUCCESS],
		Data:    map[string]string{"result": "ok"},
	}
	return resp.withOptions(opts...)
}

func (r *Response) withOptions(opts ...Option) *Response {
	c := r.clone()
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

func WithCode(code int) Option {
	return optionFunc(func(r *Response) {
		r.Code = code
	})
}

func WithMessage(msg string) Option {
	return optionFunc(func(r *Response) {
		r.Message = msg
	})
}

func WithData(data interface{}) Option {
	return optionFunc(func(r *Response) {
		r.Data = data
	})
}

func (r *Response) clone() *Response {
	response := *r
	return &response
}
