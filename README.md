# Virtual File System

This project aims to implement a virtual file system with user and file management capabilities using GoLang 1.20+.

## Objective

The objective of this virtual file system is to provide users with the ability to manage their folders and files efficiently. Users can register unique usernames and have an arbitrary number of folders and files within their scope.

## Features

### 1. User Registration

Allow users to register a unique, case insensitive username. Users can have an arbitrary number of folders and files.

- **Create User:** `register [username]`

### 2. Folder Management

Users can create, delete, list, and rename folders. Folder names must be unique within the user's scope and are case insensitive. Folders have an optional description field.

- **Create Folder:** `create-folder [username] [foldername] [description]?`
- **Delete Folder:** `delete-folder [username] [foldername]`
- **List Folders:** `list-folders [username] [--sort-name|--sort-created] [asc|desc]`
  - Each field should be separated by whitespace or tab characters.
  - The order of printed folder information is determined by the `--sort-name` or `--sort-created` combined with `asc` or `desc` flags.
- **Rename Folder:** `rename-folder [username] [foldername] [new-folder-name]`

### 3. File Management

Users can create, delete, and list all files within a specified folder. File names must be unique within the same folder and are case insensitive. Files have an optional description field.

- **Create File:** `create-file [username] [foldername] [filename] [description]?`
- **Delete File:** `delete-file [username] [foldername] [filename]`
- **List Files:** `list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]`
  - Each field should be separated by whitespace or tab characters.
  - The order of printed file information is determined by the `--sort-name` or `--sort-created` combined with `asc` or `desc` flags.

## Input Restrictions

- `[username]`, `[foldername]`, `[filename]` only accept these characters: uppercase and lowercase letters, numbers, dash, and underscore.

## Build & Run

### Prerequisites

- GoLang 1.20 or higher installed on your system.

### Build

- To build the virtual file system project, follow these steps:

1. Clone the repository:

    ```bash
    git clone <repository_url>
    ```

2. Navigate to the project directory:

    ```bash
    cd VirtualFileSystem
    ```

3. Build the project:

    ```bash
    go build
    ```

### Run

- To run the virtual file system project, execute the following command:
    ```bash
    ./VirtualFileSystem
    ```

## Example

- Register two users, user1 and user2
    ```
    # register user1
    Add user1 successfully.
    ```
    ```
    # register user2
    Add user2 successfully.
    ```

- Create a folder for user1 and user2 with the same folder name
    ```
    # create-folder user1 folder1
    Create folder1 successfully.
    ```
    ```
    # create-folder user2 folder1
    Create folder1 successfully.
    ```

- Attempt to create a folder with an existing name for user1
    ```
    # create-folder user1 folder1
    Error: folder1 has already existed.
    ```

- Create a folder with a description for user1
    ```
    # create-folder user1 folder2 this-is-folder-2
    Create folder2 successfully.
    ```

- List folders for user1 sorted by name in ascending order
    ```
    # list-folders user1 --sort-name asc
    folder1 2023-01-01 15:00:00 user1
    folder2 this-is-folder-2 2023-01-01 15:00:10 user1
    ```

- List folders for user2 sorted by name in ascending order
    ```
    # list-folders user2
    folder1 2023-01-01 15:05:00 user2
    ```

- Create a file with a description for user1 in folder1
    ```
    # create-file user1 folder1 file1 this-is-file1
    Create file1 in user1/folder1 successfully
    ```

- Create a file named config with a description for user1 in folder1
    ```
    # create-file user1 folder1 config a-config-file
    Create config in user1/folder1 successfully.
    ```

- Attempt to create an existing file.
    ```
    # create-file user1 folder1 config a-config-file
    Error: the config has already existed.
    ```

- Attempt to create a file for an unregistered user.
    ```
    # create-file user-abc folder-abc config a-config-file
    Error: The user-abc doesn't exist.
    ```

- Attempt to type an unsupported command
    ```
    # list data
    Error: Unrecognized command
    ```

- Attempt to list files with incorrect flags
    ```
    # list-files user1 folder1 --sort a
    Usage: list files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
    ```

- List files for user1 within folder1 sorted by name in descending order
    ```
    # list-files user1 folder1 --sort-name desc
    file1 this-is-file1 2023-01-01 15:00:20 folder1 user1
    config a-config-file 2023-01-01 15:00:30 folder1 user1
    ```