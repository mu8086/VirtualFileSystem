package common

import (
	"VirtualFileSystem/constants"
	"regexp"

	"github.com/spf13/viper"
)

var pattern string

func ValidName(name, pattern string) bool {
	if len(pattern) == 0 {
		pattern = viper.GetString("name.pattern")
		if len(pattern) == 0 {
			pattern = `^[A-Za-z0-9_-]+$`
		}
	}

	matched, err := regexp.MatchString(pattern, name)
	if err != nil {
		return false
	}

	return matched
}

func ValidUserName(name string) bool {
	return ValidName(name, pattern)
}

func ValidFileName(name string) bool {
	return ValidName(name, pattern)
}

func ValidFolderName(name string) bool {
	return ValidName(name, pattern)
}

func ValidSortOption(sortOption string) bool {
	return sortOption == constants.OptionSortByName || sortOption == constants.OptionSortByCreated
}

func ValidSortFlag(sortFlag string) bool {
	return sortFlag == constants.FlagSortAsc || sortFlag == constants.FlagSortDesc
}
