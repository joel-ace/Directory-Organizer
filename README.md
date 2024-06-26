# Directory Organizer

## Overview

This Go application is designed to help organize your cluttered Desktop or Downloads folder by arranging files into specific directories based on their extensions. It categorizes files into various types such as Images, Audio, Video, Documents, Programming files, Applications, Zip archives, Folders, and Others.

## Background
As part of my ongoing journey in learning the Go programming language, I took the initiative to apply what I have learned so far to solve a practical problem I have. Despite being aware of existing solutions, I saw this as an opportunity to reinforce my understanding of Go and contribute a solution tailored to my needs.

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
cd Directory-Organizer
```

## Run the application:
You can either build an executable or run the app from the command line

### Building an executable
To create an executable for the application, use the following command:
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

#### Before
<img src="https://github.com/joel-ace/Directory-Organizer/assets/23579556/7a088ce1-dda9-4876-9960-1f6e3158e337" alt="Screenshot of Desktop before running application" />

#### After
<img src="https://github.com/joel-ace/Directory-Organizer/assets/23579556/9407a178-a88f-42a2-a4fa-6092fc2cd76e" alt="Screenshot of Desktop after running application" />
<img src="https://github.com/joel-ace/Directory-Organizer/assets/23579556/b77869ef-c13c-415d-92d9-5ea0db7bfed2" alt="Screenshot of arranged file directories" />

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
