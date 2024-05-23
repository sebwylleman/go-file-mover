# goFileMover 🤖

If manually managing files scattered across multiple directories is a time-consuming headache, _goFileMover_ offers a simple, automated solution. This script consolidates files of a specific type into a single location, saving you valuable time and effort.

## About

Born from the desire to streamline file organisation, this script was my way of embracing automation. It turned out to be a fantastic learning experience, diving into Go's filesystem operations and showcasing its cross-platform prowess. I can now compile this script for various operating systems and easily share the executable with friends and colleagues.

## Features

- **Recursive Search:** Dives deep into nested directories, leaving no file unturned.
- **Customisable Extension:** Tailor the search to a specific file type (e.g., `.html`, `.txt`, `.jpg`, `.mp4`).
- **Centralised Collection:** Gathers all matching files into a designated folder.
- **Cross-Platform:** Works like a charm on Windows, macOS, and Linux.
- **Easy to Use:** Simple configuration and execution.

## Usage

1. **Create Target Directory:** Make a new directory named `html_files` in the same location where you have saved the `goFileMover.go` script. This is where the script will move all the files it finds. Feel free to change the directory name to your liking, but make sure to also change the name of the `targeDir` variable in the script to match your new directory name.
2. **Set the Extension:**
   - Open the `goFileMover.go` file in a text editor.
   - Locate the line:
     ```go
     if filepath.Ext(path) == ".html" {
     ```
   - Replace `.html` with the desired file extension (e.g., `".txt"`, `".jpg"`, `".mp4"`, `".jpeg"`):
     ```go
     if filepath.Ext(path) == ".jpeg" {
     ```
3. **Run the Script:**
   - **If you have Go installed:**
     - Open your terminal and navigate to the script's location.
     - Run `go build goFileMover.go`.
     - Execute the compiled file: `./goFileMover` (or `goFileMover.exe` on Windows).
   - **If you don't have Go installed:**
     - Ask me for the executable compiled for your operating system.
     - Place the executable in the same directory as the `goFileMover.go` script and the `html_files` folder.
     - Double-click the executable to run it.

## License

This project is licensed under the [MIT License](https://github.com/sebwylleman/goFileMover/blob/main/license.txt). This means you are free to use, modify, and distribute the software for any purpose, including commercial applications.

## Contribution

Contributions are welcome! Feel free to fork this repository, make improvements, and submit pull requests.
