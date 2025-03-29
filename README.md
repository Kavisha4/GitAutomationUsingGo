# Git Automation Using Go

## Overview

`GitAutomationUsingGo` is a command-line tool written in Go that automates common Git workflows, making it easier to manage version control tasks efficiently.

## Features

- Automates `git add`, `git commit`, and `git push`
- Reduces manual intervention and minimizes errors
- Improves developer productivity by streamlining Git workflows
- Can be extended to support additional Git operations

## Installation

### 1. Clone the Repository

```sh
 git clone https://github.com/Kavisha4/GitAutomationUsingGo.git
 cd GitAutomationUsingGo
```

### 2. Build the Executable

```sh
go build -o git-auto
```

### 3. Make It Global (Optional)

```sh
sudo mv git-auto /usr/local/bin/git-auto
sudo chmod +x /usr/local/bin/git-auto
```

## Usage

### Initialize a New Repository

```sh
mkdir MyNewRepo && cd MyNewRepo
git init
echo "# My New Project" > README.md
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin <your-github-repo-url>
git push -u origin main
```

### Automate Git Commit and Push

```sh
git-auto ~/MyNewRepo "Your commit message"
```

## Future Enhancements

- Auto-Pull & Conflict Resolution
- Smart Commit Messages
- Multi-Repo Automation
- CI/CD Integration
- Better Error Handling

## Contributing

Feel free to fork the repository and create pull requests. Suggestions and improvements are welcome!

## License

This project is licensed under the MIT License.

