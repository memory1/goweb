package controller
import (
	"log"
	"net/http"
	"testing"
	"html/template"
	"io/ioutil"
	"net/http/httptest"
	
)

func TestLoginExecutesCorrectTemplate(t *testing.T){
	h :=new(login)
	expected := "login template"
	h.loginTemplate, _ = template.New("").Parse(expected)
	r := httptest.NewRequest(http.MethodGet, "/login", nil)
	w := httptest.NewRecorder()

	h.handleLogin(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)
	log.Printf("actual: %v", string(actual))
	
	if string(actual) != expected {
		t.Errorf("Failed execute correct template")
	}

}