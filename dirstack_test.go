package main

import (
	"testing"
)

func TestEmptyDirStack(t *testing.T) {
	dirStack := DirStack{}
	currDirPath := dirStack.getCurrentDirPath()
	if currDirPath != "" {
		t.Errorf("Default stack is not empty, instead %s", currDirPath)
	}
}

func TestPushDirStack(t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("testPath")

	currDirPath := dirStack.getCurrentDirPath()
	if currDirPath != "testPath/" {
		t.Errorf("TestPath stack is not 'testPath/', instead %s", currDirPath)
	}
}

func TestPushDirStackWithLeadingAndTrailingSlashes(t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("/testPath/")

	currDirPath := dirStack.getCurrentDirPath()
	if currDirPath != "testPath/" {
		t.Errorf("TestPath stack is not 'testPath/', instead %s", currDirPath)
	}
}

func TestPushMultipleDirStack(t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("testPath")
	dirStack.pushToWorkingDirStack("testPath2")

	currDirPath := dirStack.getCurrentDirPath()
	if currDirPath != "testPath/testPath2/" {
		t.Errorf("TestPath stack is not 'testPath/testPath2/', instead %s", currDirPath)
	}
}

func TestPushMultipleDirStackInSingleString(t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("testPath/testPath2")

	currDirPath := dirStack.getCurrentDirPath()
	if currDirPath != "testPath/testPath2/" {
		t.Errorf("TestPath stack is not 'testPath/testPath2/', instead %s", currDirPath)
	}
}

func TestPushMultipleDirStackInSingleStringWithLeadingAndTailingSlashes(t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("/testPath/testPath2/")

	currDirPath := dirStack.getCurrentDirPath()
	if currDirPath != "testPath/testPath2/" {
		t.Errorf("TestPath stack is not 'testPath/testPath2/', instead %s", currDirPath)
	}
}

func TestDirStackPop(t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("testPath")
	dirStack.pushToWorkingDirStack("testPath2")

	currDirPath := dirStack.getCurrentDirPath()
	if currDirPath != "testPath/testPath2/" {
		t.Errorf("TestPath stack is not 'testPath/testPath2/', instead %s", currDirPath)
	}

	popped, err := dirStack.popFromWorkingDirStack()

	if err != nil {
		t.Errorf("popFromWorkingDirStack returned error: %s", err.Error())
	}

	if popped != "testPath2" {
		t.Errorf("TestPath popped item is not testPath2, instead %s", popped)
	}

	currDirPath = dirStack.getCurrentDirPath()
	if currDirPath != "testPath/" {
		t.Errorf("TestPath stack after pop is not 'testPath/', instead %s", currDirPath)
	}
}

func TestDirStackPopToEmpty(t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("testPath")
	dirStack.pushToWorkingDirStack("testPath2")

	_, err := dirStack.popFromWorkingDirStack()
	if err != nil {
		t.Errorf("popFromWorkingDirStack returned error: %s", err.Error())
	}

	_, err2 := dirStack.popFromWorkingDirStack()
	if err2 != nil {
		t.Errorf("popFromWorkingDirStack returned error: %s", err2.Error())
	}

	currDirPath := dirStack.getCurrentDirPath()
	if currDirPath != "" {
		t.Errorf("TestPath stack after pop is not '', instead %s", currDirPath)
	}
}

func TestDirStackPopTooFarError(t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("testPath")
	dirStack.pushToWorkingDirStack("testPath2")

	dirStack.popFromWorkingDirStack()
	dirStack.popFromWorkingDirStack()
	_, err := dirStack.popFromWorkingDirStack()

	if  err != ErrEndOfStack {
		t.Errorf("Stack was popped beyond limit but no error was recieved")
	}
}

func TestDirStackPopThenPush(t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("testPath")
	dirStack.popFromWorkingDirStack()
	dirStack.pushToWorkingDirStack("testPath2")

	currDirPath := dirStack.getCurrentDirPath()
	if currDirPath != "testPath2/" {
		t.Errorf("TestPath stack after pop is not 'testPath2/', instead %s", currDirPath)
	}
}

func TestDirStackFinalEntry (t *testing.T) {
	dirStack := DirStack{}
	dirStack.pushToWorkingDirStack("testPath")
	dirStack.pushToWorkingDirStack("testPath2")

	entry, err := dirStack.getFinalDirEntry()
	if err != nil {
		t.Error(err)
	}

	if entry != "testPath2" {
		t.Errorf("getFinalDirEntry should return 'testPath2', instead %s", entry)
	}
}

func TestDirStackFinalEntryWithNoStackError (t *testing.T) {
	dirStack := DirStack{}

	_, err := dirStack.getFinalDirEntry()
	if err != ErrEndOfStack {
		t.Error("getFinalDirEntry should return an error when getting final from empty stack")
	}
}
