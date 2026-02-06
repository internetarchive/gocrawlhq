package gocrawlhq

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
)

// gzipTransport wraps an http.RoundTripper and gzip-compresses request bodies.
type gzipTransport struct {
	next http.RoundTripper
}

func (t *gzipTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body == nil || req.Body == http.NoBody {
		return t.next.RoundTrip(req)
	}

	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	if _, err := io.Copy(w, req.Body); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	req.Body.Close()

	req.Body = io.NopCloser(&buf)
	req.ContentLength = int64(buf.Len())
	req.Header.Set("Content-Encoding", "gzip")

	return t.next.RoundTrip(req)
}
