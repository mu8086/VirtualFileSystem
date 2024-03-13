package common

import (
	"VirtualFileSystem/constants"
	"regexp"
)

func ValidName(name string) bool {
	pattern := `^[A-Za-z0-9_-]+$`

	matched, err := regexp.MatchString(pattern, name)
	if err != nil {
		return false
	}

	return matched
}

func ValidUserName(name string) bool {
	return ValidName(name)
}

func ValidFileName(name string) bool {
	return ValidName(name)
}

func ValidFolderName(name string) bool {
	return ValidName(name)
}

func ValidSortOption(sortOption string) bool {
	return sortOption == constants.OptionSortByName || sortOption == constants.OptionSortByCreated
}

func ValidSortFlag(sortFlag string) bool {
	return sortFlag == constants.FlagSortAsc || sortFlag == constants.FlagSortDesc
}
