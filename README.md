
# r2todo 📝

----

![demo](r2todo.gif)


`r2todo` is a simple and efficient command-line TODO application that keeps you focused by managing your tasks in a CSV file. It's designed to be quick to operate and free from internet-based distractions, ensuring that you own your data. 🚀

## Features ✨

- Manage TODO items with ease. ✅
- Store tasks in a CSV file. 📂
- No internet distractions—keep your data private. 🔒

## Installation 🔧

To install `r2todo`, you can build it from source using Go. Make sure you have Go installed on your machine.

1. **Clone the repository:**

    ```sh
    git clone https://github.com/IamMuuo/r2todo.git
    ```

2. **Navigate to the project directory:**

    ```sh
    cd r2todo
    ```

3. **Build the application:**

    ```sh
    go build -o r2todo
    ```

4. **Move the binary to a directory in your PATH:**

    ```sh
    mv r2todo /usr/local/bin/
    ```

    Adjust the path as necessary based on your operating system and environment.

## Usage 🛠️

Run `r2todo [command]` to interact with the application. Here’s a brief overview of available commands:

- **complete**: Toggles the completion status of a TODO item. ✔️
- **completion**: Generate the autocompletion script for the specified shell. 🖥️
- **create**: Creates a new TODO item. ✏️
- **delete**: Deletes a TODO item from the CSV file. ❌
- **help**: Displays help about any command. ❓
- **list**: Lists all currently saved TODO items. 📋

### Flags 🚩

- `-a, --delete-all`: Specifies whether to delete all TODO items. 🗑️
- `-d, --description string`: The description of the TODO item. 📝
- `-h, --help`: Displays help information. ❓
- `-i, --task-id int`: The ID of the TODO item. 🔢

### Examples 💡

- **List all TODO items:**

    ```sh
    r2todo list
    ```

- **Create a new TODO item:**

    ```sh
    r2todo create --description "Finish the report"
    ```

- **Complete a TODO item:**

    ```sh
    r2todo complete --task-id 1
    ```

- **Delete a TODO item:**

    ```sh
    r2todo delete --task-id 1
    ```

- **Delete all TODO items:**

    ```sh
    r2todo delete --delete-all
    ```

## Autocompletion ⌨️

To enable shell autocompletion for `r2todo`, use the `completion` command to generate the script for your shell. For example:

```sh
r2todo completion bash > /etc/bash_completion.d/r2todo
```


## Supported Shells 🌟
- Bash
- Zsh
- Fish

# Contributing 🤝
Contributions are welcome! If you have suggestions, bug reports, or would like to contribute code, please follow these steps:

    Fork the repository. 🍴
    Create a new branch for your feature or fix. 🌿
    Make your changes and test them thoroughly. 🛠️
    Submit a pull request with a clear description of your changes. 📝

# License
- Licensed under the Apache License see License for more info

# Questions? 🤔
If you have any questions or need further clarification, feel free to open an issue on GitHub or contact the project maintainers.

---
Enjoy managing your tasks with r2todo! 🎉
