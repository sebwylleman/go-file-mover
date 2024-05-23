package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 1. Target directory (where to put HTML files)
	targetDir := "html_files" // Create this directory beforehand

	// 2. Error handling if the directory doesn't exist
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory '%s' not found. Creating...\n", targetDir)
		os.Mkdir(targetDir, 0755) // Create with default permissions
	}

	// 3. Recursive function to walk the file tree
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 4. Skip directories
		if info.IsDir() {
			return nil
		}

		// 5. Check if the file has the ".html" extension
		if filepath.Ext(path) == ".html" {
			fmt.Printf("Found HTML file: %s\n", path)

			// 6. Construct the destination path
			newPath := filepath.Join(targetDir, filepath.Base(path))

			// 7. Move the file (rename in this case)
			if err := os.Rename(path, newPath); err != nil {
				fmt.Printf("Error moving file '%s': %v\n", path, err)
				return err
			}
		}
		return nil
	})

	// 8. Handle any errors during the walk
	if err != nil {
		fmt.Println("Error walking the path:", err)
	} else {
		fmt.Println("HTML files moved successfully!")
	}
}
