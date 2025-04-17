package github

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kurochkinivan/commit_history/internal/entity"
)

type Storage struct {
	accessToken string
	owner       string
	repo        string
}

func New(accessToken, owner, repo string) *Storage {
	return &Storage{
		accessToken: accessToken,
		owner:       owner,
		repo:        repo,
	}
}

type (
	saveRequest struct {
		Message string `json:"message"`
		Content string `json:"content"`
		Branch  string `json:"branch"`
	}

	saveResponse struct {
		Content struct {
			HTMLURL string `json:"html_url"`
		} `json:"content"`
	}
)

func (s *Storage) Save(ctx context.Context, file []byte) (string, error) {
	const op = "storage.github.Save"

	filename := fmt.Sprintf("report_%d.xml", time.Now().UnixNano())

	body, err := json.Marshal(&saveRequest{
		Message: "update report",
		Content: base64.StdEncoding.EncodeToString(file),
		Branch:  "main",
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", s.owner, s.repo, filename)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		url,
		bytes.NewReader(body),
	)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("%s: unexpected status %d: %s", op, resp.StatusCode, string(body))
	}

	var result saveResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return result.Content.HTMLURL, nil
}

func (s *Storage) GetAll(ctx context.Context) ([]entity.HistoryContent, error) {
	const op = "storage.github.GetAll"

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents", s.owner, s.repo)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s: unexpected status %d: %s", op, resp.StatusCode, string(body))
	}

	var contents []entity.HistoryContent
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return contents, nil
}

type (
	GetByNameResponse struct {
		Content  string `json:"content"`
		Encoding string `json:"encoding"`
	}
)

func (s *Storage) GetByName(ctx context.Context, filename string) ([]byte, error) {
	const op = "storage.github.GetByName"

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", s.owner, s.repo, filename)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s: unexpected status %d: %s", op, resp.StatusCode, string(body))
	}

	var fileContent GetByNameResponse
	if err := json.NewDecoder(resp.Body).Decode(&fileContent); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if fileContent.Encoding != "base64" {
		return nil, fmt.Errorf("%s: unexpected encoding %s", op, fileContent.Encoding)
	}

	content, err := base64.StdEncoding.DecodeString(fileContent.Content)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to decode base64 content: %w", op, err)
	}

	return content, nil
}
