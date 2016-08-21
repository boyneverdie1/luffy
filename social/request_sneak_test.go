package luffy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestSneak(t *testing.T) {
	reqResult := RequestSneak("asdf123", "JKL456")
	assert.NoError(t, reqResult)

	isRequested := IsRequested("asdf123", "JKL456")
	assert.True(t, isRequested)
}

func TestRequestToNoExistUser(t *testing.T) {
	reqSneak := RequestSneak("asdf123", ":!#$%^&*()_+-=;")
        assert.Error(t, reqSneak, "This function must be error because requested sneak to no user exist")
}

func TestRequestToYourOwn(t *testing.T){
	reqSneak := RequestSneak("asdf123", "asdf123")
	assert.Error(t, reqSneak, "This function must be error because requested sneak to your own nick")
}

func TestRequestFromUnknown(t *testing.T){
	reqSneak := RequestSneak(":!#$%^&*()_+-=;", "asdf123")
	assert.Error(t, reqSenak, "This function must be error because requested sneak from an unknown user")
}
