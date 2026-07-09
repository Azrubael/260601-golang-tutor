package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func buildServer(t *testing.T, dir string) string {
	t.Helper()
	exe := filepath.Join(dir, "http_gin_test_server.exe")
	cmd := exec.Command("go", "build", "-o", exe, "server.go")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("build failed: %v\noutput: %s", err, string(out))
	}
	return exe
}

func startServer(t *testing.T, exe string) *exec.Cmd {
	t.Helper()
	cmd := exec.Command(exe)
	// show server logs in test output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start server: %v", err)
	}
	return cmd
}

func waitReady(t *testing.T, base string, timeout time.Duration) {
	t.Helper()
	client := &http.Client{Timeout: 2 * time.Second}
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		resp, err := client.Get(base)
		if err == nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
	t.Fatalf("server not ready at %s after %v", base, timeout)
}

func basicAuthHeader() (string, string, string) {
	user := os.Getenv("BASIC_AUTH_USER")
	pass := os.Getenv("BASIC_AUTH_PASS")
	if user == "" {
		user = "admin"
	}
	if pass == "" {
		pass = "admin"
	}
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+pass))
	return user, pass, auth
}

func doRequestWithAuthFallback(t *testing.T, client *http.Client, base, method, path string, body interface{}) (*http.Response, []byte, error) {
	t.Helper()
	url := base + path
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	// try without auth
	resp, err := client.Do(req)
	if err == nil && resp.StatusCode != http.StatusUnauthorized {
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return resp, b, nil
	}
	if resp != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	// retry with auth
	_, _, auth := basicAuthHeader()
	req2, _ := http.NewRequest(method, url, nil)
	if body != nil {
		b, _ := json.Marshal(body)
		req2.Body = io.NopCloser(bytes.NewReader(b))
		req2.Header.Set("Content-Type", "application/json")
	}
	req2.Header.Set("Authorization", auth)
	resp2, err2 := client.Do(req2)
	if err2 != nil {
		return nil, nil, err2
	}
	b2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	return resp2, b2, nil
}

func TestIntegration(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// ensure we operate in the project folder (server.go is expected here)
	dir := wd

	exe := buildServer(t, dir)
	defer os.Remove(exe)

	cmd := startServer(t, exe)
	// ensure we stop the server
	defer func() {
		_ = cmd.Process.Kill()
		cmd.Wait()
	}()

	base := "http://localhost:8080"
	waitReady(t, base, 15*time.Second)

	client := &http.Client{Timeout: 5 * time.Second}

	// 1) GET /videos
	t.Log("GET /videos")
	resp, body, err := doRequestWithAuthFallback(t, client, base, "GET", "/videos", nil)
	if err != nil {
		t.Fatalf("GET /videos failed: %v", err)
	}
	t.Logf("Status: %s\nBody: %s", resp.Status, string(body))

	// 2) POST /videos (create)
	t.Log("POST /videos")
	newVideo := map[string]interface{}{
		"id":       1,
		"title":    "Go Test Video",
		"author":   "test",
		"playTime": 123,
		"likes":    0,
	}
	resp, body, err = doRequestWithAuthFallback(t, client, base, "POST", "/videos", newVideo)
	if err != nil {
		t.Fatalf("POST /videos failed: %v", err)
	}
	t.Logf("Status: %s\nBody: %s", resp.Status, string(body))

	// 3) GET /videos/1
	t.Log("GET /videos/1")
	resp, body, err = doRequestWithAuthFallback(t, client, base, "GET", "/videos/1", nil)
	if err != nil {
		t.Fatalf("GET /videos/1 failed: %v", err)
	}
	t.Logf("Status: %s\nBody: %s", resp.Status, string(body))

	// 4) PUT /videos/1 (update)
	t.Log("PUT /videos/1")
	upd := map[string]interface{}{
		"id":       1,
		"title":    "Updated Title",
		"author":   "tester",
		"playTime": 321,
		"likes":    10,
	}
	resp, body, err = doRequestWithAuthFallback(t, client, base, "PUT", "/videos/1", upd)
	if err != nil {
		t.Fatalf("PUT /videos/1 failed: %v", err)
	}
	t.Logf("Status: %s\nBody: %s", resp.Status, string(body))

	// 5) DELETE /videos/1
	t.Log("DELETE /videos/1")
	resp, body, err = doRequestWithAuthFallback(t, client, base, "DELETE", "/videos/1", nil)
	if err != nil {
		t.Fatalf("DELETE /videos/1 failed: %v", err)
	}
	t.Logf("Status: %s\nBody: %s", resp.Status, string(body))

	// 6) Confirm deletion (expect non-200 / maybe 404)
	t.Log("GET /videos/1 after delete (expect not found)")
	resp, body, err = doRequestWithAuthFallback(t, client, base, "GET", "/videos/1", nil)
	if err != nil {
		t.Fatalf("GET /videos/1 (after delete) request failed: %v", err)
	}
	t.Logf("Status: %s\nBody: %s", resp.Status, string(body))
	// if server returns 200 after delete, fail
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected video to be deleted, but GET /videos/1 returned 200")
	}
}