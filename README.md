# GitHub Activity

**GitHub Activity** is a command-line tool that retrieves the latest activity of a specified GitHub user.

Project idea from <https://roadmap.sh/projects/github-user-activity>

## Features

- Fetch the latest public activity for a GitHub user.
- Simple and lightweight CLI.

## Installation

1. Clone the repository and navigate to the project root.

   ```bash
   git clone https://github.com/PavelNikoltsev/github-activity.git
   cd github-activity
   ```
2. Build the project:

   ```bash
   make build
   ```

## Usage

To run the program, execute the `github-activity` file from the root directory:

```bash
./github-activity
```

Once the program is running, you will be prompted to enter a GitHub username. Enter the username, and the latest public activity for that user will be displayed.

### Example:

```bash
Enter username to get the latest github activity: PavelNikoltsev
```

```bash
Loading events...
September 28, 2024 at 7:03 PM UTC - pushed 1 commit to repo PavelNikoltsev/github-activity
September 27, 2024 at 3:26 PM UTC - created branch main in repo PavelNikoltsev/github-activity
September 27, 2024 at 3:16 PM UTC - created repository PavelNikoltsev/github-activity
September 27, 2024 at 2:49 PM UTC - pushed 1 commit to repo PavelNikoltsev/tasker
September 27, 2024 at 2:44 PM UTC - pushed 1 commit to repo PavelNikoltsev/tasker
September 27, 2024 at 2:14 PM UTC - pushed 1 commit to repo PavelNikoltsev/tasker
September 26, 2024 at 2:00 PM UTC - created branch main in repo PavelNikoltsev/task-tracker
September 26, 2024 at 1:47 PM UTC - created repository PavelNikoltsev/task-tracker
```

The program will then fetch and display the latest activity for the GitHub user `PavelNikoltsev`.
