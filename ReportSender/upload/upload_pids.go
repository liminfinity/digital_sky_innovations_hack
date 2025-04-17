package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// GitHubReference представляет ссылку (branch, tag, etc)
type GitHubReference struct {
	Ref    string `json:"ref"`
	NodeID string `json:"node_id"`
	URL    string `json:"url"`
	Object struct {
		Type string `json:"type"`
		SHA  string `json:"sha"`
		URL  string `json:"url"`
	} `json:"object"`
}

// GitHubTreeRequest запрос на создание дерева
type GitHubTreeRequest struct {
	Base string                  `json:"base_tree,omitempty"`
	Tree []GitHubTreeRequestItem `json:"tree"`
}

// GitHubTreeRequestItem элемент дерева
type GitHubTreeRequestItem struct {
	Path    string `json:"path"`
	Mode    string `json:"mode"`
	Type    string `json:"type"`
	Content string `json:"content,omitempty"`
}

// GitHubTreeResponse ответ на создание дерева
type GitHubTreeResponse struct {
	SHA  string `json:"sha"`
	URL  string `json:"url"`
	Tree []struct {
		Path string `json:"path"`
		Mode string `json:"mode"`
		Type string `json:"type"`
		SHA  string `json:"sha"`
		Size int    `json:"size,omitempty"`
		URL  string `json:"url"`
	} `json:"tree"`
}

// GitHubCommitRequest запрос на создание коммита
type GitHubCommitRequest struct {
	Message string   `json:"message"`
	Tree    string   `json:"tree"`
	Parents []string `json:"parents"`
}

// GitHubCommitResponse ответ на создание коммита
type GitHubCommitResponse struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

// GitHubUpdateRefRequest запрос на обновление ссылки
type GitHubUpdateRefRequest struct {
	SHA   string `json:"sha"`
	Force bool   `json:"force"`
}

func UploadFilesInSingleCommit(dirPath, owner, repo, accessToken string) error {

	refURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/main", owner, repo)
	ref, err := getReference(refURL, accessToken)
	if err != nil {
		return fmt.Errorf("error getting branch reference: %w", err)
	}

	log.Printf("Current branch HEAD commit: %s", ref.Object.SHA)

	date := time.Now().Format(time.DateOnly)
	treeItems, err := prepareTreeItems(dirPath, date)
	if err != nil {
		return fmt.Errorf("error preparing tree items: %w", err)
	}

	log.Printf("Prepared %d files for commit", len(treeItems))

	treeURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees", owner, repo)
	treeRequest := GitHubTreeRequest{
		Base: ref.Object.SHA,
		Tree: treeItems,
	}

	newTree, err := createTree(treeURL, accessToken, treeRequest)
	if err != nil {
		return fmt.Errorf("error creating tree: %w", err)
	}

	log.Printf("Created new tree with SHA: %s", newTree.SHA)

	commitURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits", owner, repo)
	commitRequest := GitHubCommitRequest{
		Message: fmt.Sprintf("Upload files batch %s", time.Now().Format(time.RFC3339)),
		Tree:    newTree.SHA,
		Parents: []string{ref.Object.SHA},
	}

	newCommit, err := createCommit(commitURL, accessToken, commitRequest)
	if err != nil {
		return fmt.Errorf("error creating commit: %w", err)
	}

	log.Printf("Created new commit with SHA: %s", newCommit.SHA)

	updateRefRequest := GitHubUpdateRefRequest{
		SHA:   newCommit.SHA,
		Force: false,
	}

	err = updateReference(refURL, accessToken, updateRefRequest)
	if err != nil {
		return fmt.Errorf("error updating reference: %w", err)
	}

	log.Printf("Successfully updated branch to point to new commit")
	return nil
}

func getReference(url, token string) (*GitHubReference, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error (status %d): %s", resp.StatusCode, string(body))
	}

	var reference GitHubReference
	if err := json.NewDecoder(resp.Body).Decode(&reference); err != nil {
		return nil, err
	}

	return &reference, nil
}

func prepareTreeItems(dirPath, datePrefix string) ([]GitHubTreeRequestItem, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	treeItems := make([]GitHubTreeRequestItem, 0, len(files))
	for _, file := range files {
		if file.IsDir() || file.Name() == "pid.xml" {
			continue
		}

		path := filepath.Join(dirPath, file.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", file.Name(), err)
		}

		remotePath := fmt.Sprintf("%s/%s", datePrefix, file.Name())

		treeItems = append(treeItems, GitHubTreeRequestItem{
			Path:    remotePath,
			Mode:    "100644",
			Type:    "blob",
			Content: string(content),
		})

		err = os.Remove(path)
		if err != nil {
			return nil, fmt.Errorf("failed to remove xml file: %w", err)
		}
		log.Printf("remove file %s", path)
	}

	return treeItems, nil
}

func createTree(url, token string, request GitHubTreeRequest) (*GitHubTreeResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error (status %d): %s", resp.StatusCode, string(body))
	}

	var treeResponse GitHubTreeResponse
	if err := json.NewDecoder(resp.Body).Decode(&treeResponse); err != nil {
		return nil, err
	}

	return &treeResponse, nil
}

func createCommit(url, token string, request GitHubCommitRequest) (*GitHubCommitResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error (status %d): %s", resp.StatusCode, string(body))
	}

	var commitResponse GitHubCommitResponse
	if err := json.NewDecoder(resp.Body).Decode(&commitResponse); err != nil {
		return nil, err
	}

	return &commitResponse, nil
}

func updateReference(url, token string, request GitHubUpdateRefRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GitHub API error (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}
