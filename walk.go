//Package walk - Walker is a recursive file walker
//This is different from your regular os.Walk walker because this is
//intended to be faster
//It does not call an os.Stat on any of the files instead returns the Full File Path
//And leaves the determination of the next action to the callback
//In addition this makes heavy use of the go subroutines
//there by not guaranteeing the order of returned elements
package walk

import (
	"fmt"
	"os"
	"path"
	"sync"
)

//Scan - Method called to scan the contents of a folder.
// Takes the below arguments :
// scanFolder : string - The locaiton of the starting folder from which to scan recursively
// callback : func(string) - the function to be called when a file is found
// The actual implementation of the file processing will be put in the callback
//This method relies heavily on go-subroutines
func Scan(startFolder string, callBack func(string, string)) {
	_, err := os.Open(startFolder)
	if err != nil {
		fmt.Println("Location does not exist")
		os.Exit(-1)
	}

	walkerWaitGroup.Add(1)
	go recursiveFolder(startFolder, "", callBack)
	walkerWaitGroup.Wait()
}

//walkerWaitGroup - Instance of sync.WaitGroup used to
//make the control wait until all folders and files have been processed
var walkerWaitGroup sync.WaitGroup

//readAllFiles - Constant used to tell system to read all files
const readAllFiles = -1

//recursiveFolder - This loops through all the contents of the folder
// and for every file it calls the callback function passing the file path
// and for each folder it recursively calls the method again in a
//go subroutine
func recursiveFolder(dirpath string, fileName string, callBack func(string, string)) {

	//release the waitgroup once the method is processed
	defer walkerWaitGroup.Done()
	//open a folder
	folderPtr, err := os.Open(dirpath)
	if err != nil {
		//if unable to open the folder log and return
		fmt.Println("Unable to open folder ", err)
		return
	}

	//read all files
	fileInfoArr, _ := folderPtr.Readdir(readAllFiles)

	//loop over the fileInfoArr
	for _, fileInfo := range fileInfoArr {

		if fileInfo.IsDir() {
			//add 1 to the waitgroup
			walkerWaitGroup.Add(1)

			go recursiveFolder(path.Join(dirpath, fileInfo.Name()), "", callBack)
		} else {
			//call the callback function with the file name
			callBack(dirpath, fileInfo.Name())
		}

	}

}
