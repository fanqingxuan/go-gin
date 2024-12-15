package middleware

import "go-gin/internal/httpx"

func Init(r *httpx.Engine) {

	r.Before(BeforeSampleA(), BeforeSampleB())
	r.After(AfterSampleB())
	// r.Before(TokenCheck())

}
