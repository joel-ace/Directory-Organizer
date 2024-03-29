package organizer

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"strings"
)

const ArrangedFileDirectoryName = "CleanedFiles"

type FileDirectories struct {
	Images       string
	Audio        string
	Video        string
	Folders      string
	Documents    string
	Programming  string
	Others       string
	Applications string
	Zips         string
}

// GetFileCategories returns a map of file categories to their corresponding directory names.
func GetFileCategories() FileDirectories {
	fileDirectories := FileDirectories{
		Images:       "Images",
		Audio:        "Audio",
		Video:        "Video",
		Folders:      "Folders",
		Documents:    "Documents",
		Programming:  "Programming",
		Applications: "Applications",
		Zips:         "Zips",
		Others:       "Others",
	}

	return fileDirectories
}

// GetSupportedFileExtensionsMap returns a map of file extensions to their corresponding directories.
//
// Returns:
// - a map where the key is a file extension and the value is the directory any file with this extension should be saved to.
func GetSupportedFileExtensionsMap() map[string]string {
	fileCategories := GetFileCategories()

	extensions := map[string][]string{
		fileCategories.Images:       ImageExtensions,
		fileCategories.Audio:        AudioExtensions,
		fileCategories.Video:        VideoExtensions,
		fileCategories.Documents:    DocumentExtensions,
		fileCategories.Programming:  ProgrammingExtensions,
		fileCategories.Applications: ApplicationExtensions,
		fileCategories.Zips:         ZipExtensions,
	}

	extensionsToDirectoryMap := make(map[string]string)

	for key := range extensions {
		for i := 0; i < len(extensions[key]); i++ {
			extensionsToDirectoryMap[extensions[key][i]] = key
		}
	}

	return extensionsToDirectoryMap
}

// GetFileInfo retrieves information about a file or directory.
//
// Parameters:
// - path: path to the file or directory.
//
// Returns:
// - os.FileInfo: Information about the file if file exists.
// - bool: indicates whether the file exists
func GetFileInfo(path string) (os.FileInfo, bool) {
	fileInfo, checkDirectoryExistError := os.Stat(path)
	return fileInfo, checkDirectoryExistError == nil
}

// GetHomeDirectory retrieves the path to the user's home directory.
func GetHomeDirectory() string {
	homeDir := os.Getenv("HOME")

	if homeDir == "" {
		currentUserHomeDir, err := os.UserHomeDir()

		if err != nil {
			currentUser, _ := user.Current()
			return currentUser.HomeDir
		}
		homeDir = currentUserHomeDir
	}

	return homeDir
}

// getHomeDirectoriesPathByName retrieves the path to a directory inside the user's home directory using the dir name.
//
// Parameters:
// name - The name of the directory.
//
// Returns:
// - the path to the directory.
func getHomeDirectoriesPathByName(name string) string {
	HomeDirectory := GetHomeDirectory()
	return filepath.Join(HomeDirectory, name)
}

// GetDesktopDirectory retrieves the path to the Desktop directory.
//
// Returns:
// - the path to the Desktop directory.
func GetDesktopDirectory() string {
	return getHomeDirectoriesPathByName("Desktop")
}

// GetArrangedFileDirectory retrieves the path to the directory where the arranged files will be moved to.
//
// Parameters:
// - targetDirectory: parent directory where the arranged file directory will be created.
//
// Returns:
// - path to the arranged file directory.
func GetArrangedFileDirectory(targetDirectory string) string {
	return filepath.Join(targetDirectory, ArrangedFileDirectoryName)
}

// GetDownloadDirectory retrieves the path to the Downloads directory.
//
// Returns:
// - the path to the Downloads directory.
func GetDownloadDirectory() string {
	return getHomeDirectoriesPathByName("Downloads")
}

// CreateDirectory creates a directory if it does not already exist.
//
// Parameters:
// - path: the path of the directory to create.
// - permission: the permission mode for the directory.
func CreateDirectory(path string, permission os.FileMode) {
	if _, exists := GetFileInfo(path); !exists {
		os.Mkdir(path, permission)
	}
}

// ReadDirectory reads the contents of a directory.
//
// Parameters:
// - path: the path to the directory.
//
// Returns:
// - an array of the directory contents.
func ReadDirectory(path string) []os.DirEntry {
	if _, exists := GetFileInfo(path); !exists {
		err := fmt.Sprintf("Directory with path: %v not found", path)
		log.Fatal(err)
		return []os.DirEntry{}
	}

	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
		return []os.DirEntry{}
	}

	return files
}

// CreateDirectories creates all required directories for the supported file extensions.
//
// Parameters:
// - fileCategories: a struct containing the required file categories.
// - path: the path to the parent directory where the directories will be created.
func CreateDirectories(fileCategories FileDirectories, path string) {
	categoryTypes := reflect.TypeOf(fileCategories)

	CreateDirectory(path, 0755)

	for i := 0; i < categoryTypes.NumField(); i++ {
		field := categoryTypes.Field(i)
		CreateDirectory(filepath.Join(path, field.Name), 0755)
	}
}

// GetFileExtensions retrieves the extension of a file.
//
// Parameters:
// - file: the file whose extension is to be retrieved.
//
// Returns:
// - the extension of the file.
func GetFileExtensions(file os.DirEntry) string {
	extension := filepath.Ext(file.Name())
	extension = strings.ToLower(strings.TrimPrefix(extension, "."))
	return extension
}

// incrementPathName adds a number to the end of a file with duplicate name and increments it to avoid name conflicts.
//
// Parameters:
// - originalPath: the original path of the file.
//
// Returns:
// - the new path with an incremented suffixed file name.
func incrementPathName(originalPath string) string {
	_, err := os.Stat(originalPath)
	if err == nil {
		base := filepath.Base(originalPath)
		dir := filepath.Dir(originalPath)
		ext := filepath.Ext(base)
		name := base[:len(base)-len(ext)]

		for i := 1; ; i++ {
			newPath := filepath.Join(dir, fmt.Sprintf("%s_%d%s", name, i, ext))
			_, err := os.Stat(newPath)

			if os.IsNotExist(err) {
				return newPath
			}
		}
	}
	return originalPath
}

// getDirectorySize calculates the size of a directory.
//
// Parameters:
// - dir: the path to the directory.
//
// Returns:
// - the size of the directory.
// - error encountered if any.
func getDirectorySize(dir string) (int64, error) {
	var size int64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// Move moves a file to a new location.
//
// Parameters:
// - sourcePath: the current path of the file.
//   - destinationPath: the new path of the file.
func Move(sourcePath, destinationPath string) {
	sourceStat, _ := os.Stat(sourcePath)
	destinationStat, err := os.Stat(destinationPath)

	if err != nil {
		os.Rename(sourcePath, destinationPath)
		return
	}

	if sourceStat.Size() == destinationStat.Size() {
		destinationPath = incrementPathName(destinationPath)
	}

	os.Rename(sourcePath, destinationPath)
}

// moveDirectory moves a directory to a new location.
//
// Parameters:
// - sourcePath: the current path of the directory.
// - destinationPath: the new path of the directory.
func moveDirectory(sourcePath string, destinationPath string) {
	sourceDirSize, _ := getDirectorySize(sourcePath)
	destinationDirSize, _ := getDirectorySize(destinationPath)

	if sourceDirSize == destinationDirSize {
		destinationPath = incrementPathName(destinationPath)
	}

	renameError := os.Rename(sourcePath, destinationPath)

	if renameError != nil {
		fmt.Printf("Error occured moving directory: %v \n", renameError)
	}
}

// organize organizes the given file according to its type and moves it to the appropriate directory.
//
// Parameters:
// - file: the file to be organized.
// - fileCategories: struct containing the directory names for file categories.
// - parentDirectory: the parent directory of the file to be moved/organized.
// - arrangedFileDirectory: the directory where organized files will be moved.
func organize(
	file os.DirEntry,
	fileCategories FileDirectories,
	parentDirectory string,
	arrangedFileDirectory string,
) {
	var folder string
	sourceDirectory := filepath.Join(parentDirectory, file.Name())

	if file.IsDir() {
		folder = fileCategories.Folders
		moveDirectory(sourceDirectory, filepath.Join(arrangedFileDirectory, folder, file.Name()))
	} else {
		extension := GetFileExtensions(file)
		extensionToFolderMap := GetSupportedFileExtensionsMap()

		value, hasValue := extensionToFolderMap[extension]

		if hasValue {
			folder = value
		} else {
			folder = fileCategories.Others
		}

		Move(sourceDirectory, filepath.Join(arrangedFileDirectory, folder, file.Name()))
	}
}

// OrganizeFiles organizes a list of files using the "organize" function according to their types.
//
// Parameters:
// - files: the list of files to be organized.
// - fileCategories: struct containing the directory names for different file categories.
// - targetDirectory: the directory containing the files to be organized.
// - arrangedFileDirectory: the directory where organized files will be moved.
func OrganizeFiles(
	files []os.DirEntry,
	fileCategories FileDirectories,
	targetDirectory string,
	arrangedFileDirectory string,
) {
	for _, file := range files {
		if file.Name() == ArrangedFileDirectoryName {
			continue
		}

		organize(file, fileCategories, targetDirectory, arrangedFileDirectory)
	}
	fmt.Println("================ COMPLETED ================")
}
