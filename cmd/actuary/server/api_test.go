package server

import (
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/diogomonica/actuary/cmd/actuary/server/handlers"
	"github.com/diogomonica/actuary/cmd/actuary/server/services"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"strings"
)

func testNewAPI(t *testing.T) {
	api := NewAPI("certPath", "keyPath")
	if api.encryptionKey == nil {
		t.Errorf("Nil encryption key")
	}
	if api.Tokens == nil {
		t.Errorf("Nil tokens")
	}
}
