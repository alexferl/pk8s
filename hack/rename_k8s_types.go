package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	files, err := filepath.Glob("k8s/*.go")
	if err != nil {
		log.Fatalf("failed reading files: %v", err)
	}

	re := regexp.MustCompile(`model_io_k8s_.*_(v\d+(?:alpha\d+|beta\d+)?)_(.+)\.go`)

	excludePatterns := []string{
		`_list\.go$`,
		`_watch_event\.go$`,
		`_pkg_version_info.go$`,
	}

	for _, file := range files {
		shouldExclude := false
		for _, pattern := range excludePatterns {
			if matched, _ := regexp.MatchString(pattern, file); matched {
				shouldExclude = true
				break
			}
		}

		if shouldExclude {
			err := os.Remove(file)
			if err != nil {
				log.Fatalf("failed removing file %s: %v\n", file, err)
			} else {
				fmt.Printf("removed file: %s\n", file)
			}
			continue
		}

		matches := re.FindStringSubmatch(file)
		if len(matches) == 3 {
			version := matches[1]
			resource := matches[2]
			newName := fmt.Sprintf("k8s/%s_%s.go", resource, version)

			if file != newName {
				err := os.Rename(file, newName)
				if err != nil {
					log.Fatalf("failed renaming %s to %s: %v\n", file, newName, err)
				} else {
					fmt.Printf("renamed %s to %s\n", file, newName)
				}
			}

			content, err := os.ReadFile(newName)
			if err != nil {
				log.Fatalf("failed reading file %s: %v\n", newName, err)
			}

			stringRe := regexp.MustCompile(`IoK8s\w*?(V\d+(?:alpha\d+|beta\d+)?)(\w+)`)
			newContent := stringRe.ReplaceAllString(string(content), "${2}${1}")

			err = os.WriteFile(newName, []byte(newContent), 0o644)
			if err != nil {
				log.Fatalf("failed writing file %s: %v\n", newName, err)
			} else {
				fmt.Printf("updated strings in %s\n", newName)
			}
		}
	}

	directories := []string{"k8s/.openapi-generator", "k8s/api"}
	for _, dir := range directories {
		err = os.RemoveAll(dir)
		if err != nil {
			log.Fatalf("failed deleting folder %s: %v\n", dir, err)
		} else {
			fmt.Printf("successfully deleted folder: %s\n", dir)
		}
	}
}
