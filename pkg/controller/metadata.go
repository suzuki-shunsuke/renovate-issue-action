package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

type Metadata struct {
	PackageFileDir string `json:"packageFileDir"`
	PackageName    string `json:"packageName"`
	GroupName      string `json:"groupName"`
	DepName        string `json:"depName"`
	UpdateType     string `json:"updateType"`
	Manager        string `json:"manager"`
}

var errMetadataNotFound = errors.New("metadata isn't found")

func readMetadata(body string, metadata *Metadata) error {
	// <!-- renovate-issue-action: {"packageFileDir": ".github/workflows", "packageName": "", "groupName": "", "depName": "aquaproj/aqua", "manager": "regex", "updateType": "major"} -->
	arr := regexp.MustCompile(`<!-- renovate-issue-action:(.*?)-->`).FindStringSubmatch(body)
	if arr == nil {
		return errMetadataNotFound
	}
	if err := json.Unmarshal([]byte(arr[1]), metadata); err != nil {
		return fmt.Errorf("unmarshal metadata as JSON: %w", err)
	}
	return nil
}
