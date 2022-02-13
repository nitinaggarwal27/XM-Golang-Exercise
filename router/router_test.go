package router

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	routeHandler := Routes()
	if routeHandler != nil {
		t.Logf("Route Initialization Success")
	} else {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	// test case 1 -> when string available in list
	res := contains([]string{"a", "b", "c"}, "a")
	if !res {
		t.Fail()
	}

	res = contains([]string{"a", "b", "c"}, "d")
	if res {
		t.Fail()
	}
}
