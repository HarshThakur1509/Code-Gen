package edit

import (
	"fmt"
	"os"
	"strings"
)

// CodeEditRequest defines the structure for code modification requests
type CodeEditRequest struct {
	FilePath      string `json:"filePath"`
	CommentMarker string `json:"commentMarker"`
	NewCode       string `json:"newCode"`
}

// EditFile allows modifying a file by inserting code at a specific comment marker
func EditFile(request CodeEditRequest) error {
	// Validate input
	if request.FilePath == "" {
		return fmt.Errorf("file path cannot be empty")
	}

	// Read the file
	content, err := os.ReadFile(request.FilePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Convert content to a string
	fileContent := string(content)

	// Check if the comment marker exists
	if request.CommentMarker != "" && !strings.Contains(fileContent, request.CommentMarker) {
		return fmt.Errorf("comment marker not found in the file")
	}

	// Split the content and insert new code
	var updatedContent string
	if request.CommentMarker != "" {
		parts := strings.Split(fileContent, request.CommentMarker)
		updatedContent = parts[0] + request.CommentMarker + "\n" + request.NewCode + "\n" + parts[1]
	} else {
		// If no comment marker, append to the end of the file
		updatedContent = fileContent + "\n" + request.NewCode
	}

	// Write the updated content back to the file
	err = os.WriteFile(request.FilePath, []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}
