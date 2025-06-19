# Deltatree (?) – hierarchical folder-based task manager

## Principles
- every task is a folder
- you can organize tasks hierarchically, just put a folder inside a folder
- you can attach files to a task, just put them in the task's folder
- folder naming follows [todo.txt rules](https://github.com/todotxt/todo.txt)
- child tasks inherit all properties from parent task but can override some of them

## Example structure
```
- my tasks
    - deltatree task manager +deltatree @programming
        - x 2025-06-18 2025-06-19 write readme file
        - 2025-06-15 write unit tests
    - work +work
        - x 2025-06-16 2025-06-15 take vacation for 26 and 27 of July
        - ...
```
