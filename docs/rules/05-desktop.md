# 05 Desktop Development Rules

## Rule: Respect platform conventions and user expectations

**When:** Building desktop application UI and interactions

**What:** Follow platform-specific guidelines for menus, shortcuts, dialogs, and window behavior

**Why:** Users expect consistent behavior across applications on their platform; violating conventions creates friction

**How to verify:** Test on target platforms; verify that keyboard shortcuts match platform standards; check that dialogs and menus follow platform patterns

**Example (correct):**
```python
# Platform-appropriate keyboard shortcuts
if sys.platform == 'darwin':
    SAVE_SHORTCUT = 'Cmd+S'
    QUIT_SHORTCUT = 'Cmd+Q'
else:
    SAVE_SHORTCUT = 'Ctrl+S'
    QUIT_SHORTCUT = 'Ctrl+Q'

menu.add_command(label="Save", accelerator=SAVE_SHORTCUT)
```

**Example (incorrect):**
```python
# Hardcoded shortcuts ignore platform conventions
menu.add_command(label="Save", accelerator='Ctrl+S')  # Wrong on macOS
menu.add_command(label="Quit", accelerator='Alt+F4')  # Wrong on macOS
```
