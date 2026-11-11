# Build a Rhythm Scroller Game in Defold: From Blocks to Real Code!

Welcome! If you have built games using Scratch or Code.org, you already know how game logic works (loops, conditions, and triggers). In this tutorial, you will take your first step into **real text-based software development** using **Lua** and **Defold**—the same technology professionals use to build lightweight 2D games.

---

## Step 0: Getting the Starter Project Running

1. Open **Defold Editor**.
2. Select **Open Project** and navigate to your cloned repository folder (`defold-tutorial-rhythm-scroller`).
3. Press **`Ctrl + B`** (Windows/Linux) or **`Cmd + B`** (Mac) to build and run the game immediately!
4. Try pressing `Space` to jump over obstacles.

---

## Scratch-to-Lua Translator Guide

When writing text code, syntax matters, but the logical concepts are similar to block coding:

| Scratch Concept | Lua / Defold Concept | What It Does |
| :--- | :--- | :--- |
| `when green flag clicked` | `function init(self)` | Runs setup code once when the level opens. |
| `forever` | `function update(self, dt)` | Runs 60 times every second to update positions and physics. |
| `when [space] key pressed` | `function on_input(self, action_id, action)` | Listens for keyboard presses or mouse clicks. |
| `if touching [Spike]?` | `function on_message(self, message_id, message)` | Triggers events when objects collide. |
| Set variable `speed` to `5` | `local FORWARD_SPEED = 300` | Stores numbers that control game behavior. |

---

## Phase 1: Experimenting with Game Feel (Physics & Tuning)

**Goal:** Modify numerical variables to see how code parameters instantly change gameplay feel.

1. In the left-side **Project Explorer**, expand the `player` folder and double-click `player.script`.
2. Locate the **`VARIABLE BLOCK`** at the top of the file:

```
local GRAVITY       = -2400
local JUMP_FORCE    = 850
local MAX_JUMPS     = 2
```

* **Experiment #1 (Moon Physics):** Change `GRAVITY` to `-800`. Press `Cmd+B` / `Ctrl+B` to test. Notice how floaty the jump feels!
* **Experiment #2 (Snappy Jump):** Increase `JUMP_FORCE` to `1000` for a higher apex.

---

## Phase 2: Code Modification (Unlocking the Double Jump)

**Goal:** Understand variables and conditional checks (`if` statements) by adding a double-jump feature.

1. In `player/player.script`, look at the variable block:

```
local MAX_JUMPS = 2
```

2. **Challenge:** Modify `MAX_JUMPS = 3`. Run the game (`Cmd+B` / `Ctrl+B`). Can you jump three times before hitting the ground?
* **Trace the Logic:** Every jump subtracts `1` from `jumps_remaining`. When `jumps_remaining` reaches `0`, conditional checks prevent further jumps until the player lands.

---

## Phase 3: Visual Level Design & Hazard Placement

**Goal:** Design custom levels using Defold’s visual layout editor.

1. Open `main/main.collection` from the **Project Explorer** pane.
2. **Adding a Spike:**
   * Right-click the `main` node in the Outline window (top right).
   * Select **Add Game Object File** and choose `obstacles/spike.go`.
   * Use the **Move Tool** (`W`) to place the spike along the runner track.
3. **Adding Platforms:**
   * Right-click `main` $\rightarrow$ **Add Game Object File** $\rightarrow$ Select `obstacles/ground.go`.
   * Position the block higher in the air to build floating platforms.
4. Press **`Cmd + B`** / **`Ctrl + B`** to playtest your custom level!
