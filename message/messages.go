package message

import (
	"fmt"
)

var TextColor = struct {
	Error string
	Info  string
	Warn  string
	Debug string
	Reset string
}{
	Error: "\033[0;31m",
	Info:  "\033[0;32m",
	Warn:  "\033[0;33m",
	Debug: "\033[0;34m",
	Reset: "\033[0m",
}

func GetInvalidRuleIdWarning(ruleId string) string {
	return fmt.Sprintf(TextColor.Warn+"Warning "+TextColor.Reset+"Invalid rule ID %s\n", ruleId+TextColor.Reset)
}

func GetMissingFileOrFolderError() string {
	return "Specify a file or folder path to scan\n"
}

func GetCreatingIgnoreListInstanceError(err error) string {
	return fmt.Sprintf("Error while creating instance for ignore list %v", err)
}

func GetReadingCurrentWorkDirectoryError(pathErr error) string {
	return fmt.Sprintf("Error while reading current work directory: %v", pathErr)
}

func GetIgnoreFilesWalkError(ignoreFilesWalkError error) string {
	return fmt.Sprintf("Error while scanning for ignore files %v", ignoreFilesWalkError)
}

func GetPathFetchingError(err error) string {
	return fmt.Sprintf("Error fetching path: %+v\n", err)
}

func GetFilesFetchingError(err error) string {
	return fmt.Sprintf("Error in fetching files from directory: %+v\n", err)
}

func GetMarshallingOutputError(err error) string {
	return fmt.Sprintf("Error marshalling output: %+v\n", err)
}

func GetFileReadError(fileName string, err error) string {
	return fmt.Sprintf("Error reading file %s: %v", fileName, err.Error())
}

func GetInvalidTemplateFileError(fileError error) string {
	return fmt.Sprintf("Invalid template file  %s", fileError)
}

func GetInvalidConfigFileError(path string) string {
	return fmt.Sprintf("Invalid config file. Expecting a \"yaml\" or \"json\" extension %s", path)
}

func GetFileUnmarshalingError(err error) string {
	return fmt.Sprintf("Error during Unmarshaling of file %s", err)
}

func GetMarshalIndentError(err error) string {
	return fmt.Sprintf("%v", err)
}
