package service

import (
	"fmt"
	"time"
)

type RateError struct {
	serviceId string
	dur       time.Duration
}

func NewRateError(serviceId string, dur time.Duration) error {
	return &RateError{serviceId: serviceId, dur: dur}
}

func (e RateError) Error() string {
	return fmt.Sprintf("limited for %f secs by %s", e.dur.Seconds(), e.serviceId)
}

func (e RateError) Wait() {
	time.Sleep(e.dur)
}

func (e RateError) GetRestrictedTill() time.Time {
	return time.Now().Add(e.dur)
}

type ReCaptchaInvalidError struct {
	token  string
	reason string
}

func NewReCaptchaInvalidError(token string, reason string) error {
	return &ReCaptchaInvalidError{token: token, reason: reason}
}

func (e ReCaptchaInvalidError) Error() string {
	return fmt.Sprintf("invalid recaptcha '%s...%s': %s", e.token[:6], e.token[len(e.token)-6:], e.reason)
}
