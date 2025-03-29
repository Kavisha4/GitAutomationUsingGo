# Git Automation Using Go

## Overview

`GitAutomationUsingGo` is a command-line tool written in Go that automates common Git workflows, making it easier to manage version control tasks efficiently.

## Features

- Automates `git add`, `git commit`, and `git push`
- Reduces manual intervention and minimizes errors
- Improves developer productivity by streamlining Git workflows
- Can be extended to support additional Git operations
- Uses Cobra for a structured CLI experience
- Supports Git features like stash, checkout, and rebase
- Includes GitHub Actions workflows for CI/CD automation

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
git-auto commit "Your commit message"
```

### Using Additional Git Features

- **Stash Changes:** `git-auto stash`
- **Checkout Branch:** `git-auto checkout <branch>`
- **Rebase Branch:** `git-auto rebase <branch>`

## GitHub Workflows Integration

This project includes GitHub Actions workflows for CI/CD automation.

### **Testing GitHub Workflows**

1. Navigate to `.github/workflows/ci.yml` and ensure it is correctly configured.
2. Push a commit to trigger the workflow:

```sh
git add .
git commit -m "Testing GitHub workflows"
git push origin main
```

3. Check the GitHub Actions tab in your repository to monitor the workflow execution.

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

