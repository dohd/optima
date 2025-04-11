package config

import (
    "log"
    "path/filepath"
    "runtime"
)

func GetProjectRoot() string {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        log.Fatal("unable to get current file path")
    }

    // Adjust based on where your project root is
    projectRoot := filepath.Join(filepath.Dir(filename), "..", "..")

    // Get the absolute path to the project root
    absProjectRoot, err := filepath.Abs(projectRoot)
    if err != nil {
        log.Fatal("failed to get absolute path: ", err)
    }

    return absProjectRoot
}
