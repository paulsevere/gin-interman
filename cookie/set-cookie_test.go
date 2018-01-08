package cookie

import "testing"

// TestDetermineCookieDomain asdf
func TestDetermineCookieDomain(t *testing.T) {
	domainString := determineCookieDomain("https://test-libertymutual.com")
	if domainString != ".libertymutual.com" {
		t.Fail()
	}
	domainString = determineCookieDomain("http://localhost:8080")
	if domainString != "localhost" {
		t.Fail()
	}
	domainString = determineCookieDomain("http://google.com")
	if domainString != "http://google.com" {
		t.Fail()
	}
}
