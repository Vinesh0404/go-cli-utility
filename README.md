# 🧮 Versatile Go Command-Line Utility

A simple, all-in-one Go program that provides a menu-driven interface for common utility tasks.

## ✨ Features

  * `➕` **Arithmetic Calculator:** Perform basic addition (`+`), subtraction (`-`), multiplication (`*`), and division (`/`) on two integer inputs.
  * `🔢` **Factorial Calculator:** Recursively calculate the factorial of a given non-negative integer.
  * `⭐` **Pattern Drawer:** Draw a right-angled triangle pattern of stars (`*`) based on a specified height.

## 📋 Prerequisites

To run this project, you only need to have **Go** (version 1.18 or newer) installed on your system.

  * [Download and Install Go](https://golang.org/doc/install)

## 🚀 Getting Started

1.  **Save the Code:**
    Save the project code into a file named `main.go`.

2.  **Run the Program:**
    Open your terminal, navigate to the directory where you saved the file, and execute the following command:

    ```bash
    go run main.go
    ```

3.  **Interact with the Menu:**
    Once running, the program will display the main menu. Just type the number corresponding to the task you want to perform and press `Enter`.

## 🖥️ Usage Examples

### Main Menu

When you start the program, you will see this menu:

```text
--- Versatile Command-Line Utility ---

Main Menu: 
1. Arithmetic
2. Factorial
3. Pattern
4. Exit

Select an Option: 
```

### Example 1: Arithmetic

```text
Select an Option: 1

--- Arithmetic Calculator ---
Enter first number: 10
Enter operator (+, -, *, /): +
Enter second number: 5

Result: 10 + 5 = 15

Press 'Enter' to return to the main menu...
```

### Example 2: Factorial

```text
Select an Option: 2

--- Factorial Calculator ---
Enter a non-negative number: 5

Result: Factorial of 5 is 120

Press 'Enter' to return to the main menu...
```

### Example 3: Pattern

```text
Select an Option: 3

--- Geometric Pattern Drawer ---
Enter the height of the triangle: 4

*
**
***
****

Press 'Enter' to return to the main menu...
```

### Example 4: Exit

```text
Select an Option: 4

Thank you for using the utility. Goodbye!
```

## 🤝 Contributing

Pull requests are welcome\! For major changes, please open an issue first to discuss what you would like to change.

## 📜 License

This project is licensed under the **MIT License**. See the `LICENSE.md` file for details (or add your preferred license).
