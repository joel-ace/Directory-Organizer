# Directory Organizer

## Overview

This Go application is designed to help organize your cluttered Desktop or Downloads folder by arranging files into specific directories based on their extensions. It categorizes files into various types such as Images, Audio, Video, Documents, Programming files, Applications, Zip archives, Folders, and Others.

## Background
As part of my ongoing journey in learning the Go programming language, I've taken the initiative to apply what I have learned so far to solve a practical problem I have. Despite being aware of existing solutions, I saw this as an opportunity to reinforce my understanding of Go and contribute a solution tailored to my specific needs.

#### Before


#### After


## Features
- **Automatic Organization**: The application automatically sorts files into appropriate directories based on their extensions.
- **Intuitive Usage**: Simple command-line interface for easy interaction.

## Installation
- Make sure you have Go installed on your system. If not, you can download and install it from [here](https://go.dev/doc/install).
- Clone this repository to your local machine:
```bash
git clone https://github.com/joel-ace/Directory-Organizer.git
```
-  Navigate to the directory:
```bash
cd <directory_name>
```

## Run the application:
You can either build an executable or run the app from the command line

### Building an executable
To create an executable for the application, you can use the following command
```bash
go build
```

### Running the Application from Command Line
- Alternatively, you can run the application directly from the command line without building an executable by using:
```bash
go run main.go
```

## Usage
Upon running the application, you will be prompted to choose the directory you want to clean:

- Enter 1 for Desktop.
- Enter 2 for Downloads.

The application will then organize the files in the selected directory into subdirectories based on their file types.
Once the organization process is completed, you will see a confirmation message.

## Customization
- You can customize the directory names and supported file extensions by modifying the respective constants in the `organizer.go` file.
- To add new file categories or extensions, update the `FileDirectories` struct and extension maps accordingly in the `GetFileCategories` and `GetSupportedFileExtensionsMap` functions.

## Note
This application is designed to organize files in your Desktop or Downloads directory. Make sure to use it cautiously

Author
--------------
- Joel Akwevagbe

## License
This project is licensed under the MIT License.
