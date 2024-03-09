package common

import "VirtualFileSystem/constants"

func ValidUserName(name string) bool {
	return true // TODO
}

func ValidFolderName(name string) bool {
	return true // TODO
}

func ValidSortOption(sortOption string) bool {
	return sortOption == constants.OptionSortByName || sortOption == constants.OptionSortByCreated
}

func ValidSortFlag(sortFlag string) bool {
	return sortFlag == constants.FlagSortAsc || sortFlag == constants.FlagSortDesc
}
