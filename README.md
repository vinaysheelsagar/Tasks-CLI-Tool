# Tasks-CLI-Tool
**tacl: A Versatile Task Management CLI**

**Overview**

tacl is a powerful command-line interface designed to simplify task management. Built with Go, Cobra, and Viper, tacl offers a flexible and efficient way to create, track, and prioritize your tasks.

**Key Features:**

* **Intuitive Command-Line Interface:** Easily interact with your tasks using simple commands.
* **Task Organization:** Categorize tasks for better organization and focus.
* **Priority Management:** Assign priorities to tasks to effectively manage your workload.
* **Task Completion Tracking:** Mark tasks as completed and review your progress.
* **Flexible Configuration:** Customize tacl to your preferences using configuration files.

**Getting Started**

**Installation:**

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/vinaysheelsagar/Tasks-CLI.git
   cd Tasks-CLI
   ```
2. **Build and Install:**
   ```bash
   go build
   sudo mv tacl /usr/local/bin/
   ```

**Basic Usage:**

```bash
tacl help
```

This will display the available commands and their usage.

**Key Commands:**

* **Create a task:**
  ```bash
  tacl create "Write a report"
  ```
* **List all tasks:**
  ```bash
  tacl list
  ```
* **View a specific task:**
  ```bash
  tacl view <task_name>
  ```
* **Mark a task as complete:**
  ```bash
  tacl check <task_name>
  ```
* **Update a task:**
  ```bash
  tacl update <task_name> "<new description>"
  ```
* **Delete a task:**
  ```bash
  tacl delete <task_name>
  ```

**Configuration**

tacl supports configuration through environment variables and configuration files (e.g., `config.yaml`). You can customize various settings, such as default category, priority, and output format.

**Contributing**

We welcome contributions! Please refer to the `CONTRIBUTING.md` file for guidelines.

**License**

tacl is licensed under the MIT License. See the `LICENSE` file for details.

**Additional Notes:**

* For more advanced usage and customization, refer to the `help` command or the documentation.
* We encourage you to explore the full potential of tacl and provide feedback.

By leveraging the power of Go, Cobra, and Viper, tacl offers a robust and efficient solution for managing your tasks.
