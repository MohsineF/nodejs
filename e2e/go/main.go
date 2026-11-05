// A small program over a wide slice of the standard library, so a cold build
// compiles a few hundred packages and a warm one restores them.
package main

import (
	"archive/zip"
	"compress/gzip"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"image/png"
	"net/http"
	"net/http/httptest"
	"net/smtp"
	"regexp"
	"text/tabwriter"
)

func main() {
	_ = zip.Store
	_ = gzip.BestSpeed
	_ = tls.VersionTLS13
	_ = sql.ErrNoRows
	_ = xml.Header
	_ = template.HTMLEscapeString
	_ = png.DefaultCompression
	_ = httptest.DefaultRemoteAddr
	_ = smtp.PlainAuth
	_ = regexp.MustCompile
	_ = tabwriter.AlignRight
	b, _ := json.Marshal(http.StatusOK)
	fmt.Println(string(b))
}
