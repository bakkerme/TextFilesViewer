package main

import (
	"regexp"
	"errors"
)

var ErrEndOfStack = errors.New("End of Stack")

type DirStack struct {
	dirList []string
}

func (dirStack *DirStack) getCurrentDirPath () (string) {
	tempPath := ""
	for _, value := range dirStack.dirList {
		tempPath += value + "/"
	}

	return tempPath
}

func (dirStack *DirStack) pushToWorkingDirStack (path string) {
	processedPath := path

	testForEndSlash := `/$`
	matchEndSlash, _ := regexp.MatchString(testForEndSlash, path)
	if matchEndSlash {
		processedPath = path[0:len(path) - 1]
	}

	testForStartSlash := `^/`
	matchStartSlash, _ := regexp.MatchString(testForStartSlash, processedPath)
	if matchStartSlash {
		processedPath = processedPath[1:len(processedPath)]
	}

	testForRemainingSlash := `/`
	matchRemainingSlash, _ := regexp.MatchString(testForRemainingSlash, processedPath)
	if matchRemainingSlash { // This probably is multiple paths. Lets use them all
		processedPaths := regexp.MustCompile("/").Split(processedPath, -1)
		for _, value := range processedPaths {
			dirStack.dirList = append(dirStack.dirList, value)
		}
	} else {
		dirStack.dirList = append(dirStack.dirList, processedPath)
	}
}

func (dirStack *DirStack) popFromWorkingDirStack () (string, error) {
	if(len(dirStack.dirList) == 0) {
		return "", ErrEndOfStack
	}

	popItem := dirStack.dirList[len(dirStack.dirList)-1]
	dirStack.dirList = dirStack.dirList[:len(dirStack.dirList)-1]

	return popItem, nil
}

func (dirStack *DirStack) getFinalDirEntry () (string, error) {
	if(len(dirStack.dirList) == 0) {
		return "", ErrEndOfStack
	}

	return dirStack.dirList[len(dirStack.dirList) - 1], nil
}
