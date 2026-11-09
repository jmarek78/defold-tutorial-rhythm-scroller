# defold-tutorial-rhythm-scroller
Rhythm-scroller Defold tutorial

# Build a Rhythm Scroller Game in Defold: From Blocks to Real Code!

Welcome! If you have built games using Scratch or Code.org, you already know how game logic works (loops, conditions, and triggers). In this tutorial, you will take your first step into **real text-based software development** using **Lua** and **Defold** - the same technology professionals use to build lightweight 2D games.

---

## Step 0: Getting the Starter Project Running

1. Open **Defold Editor**.
2. Select **Open Project** and navigate to your cloned repository folder (`defold-tutorial-rhythm-scroller`).
3. Press **`Ctrl + B`** (Windows/Linux) or **`Cmd + B`** (Mac) to build and run the game immediately!
4. Try pressing `Space` to jump over obstacles.

---

## Scratch-to-Lua Translator Guide

When writing text code, syntax matters, but the logical concepts are identical to block coding:

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

1. In the left-side **Project Explorer**, expand the `game` folder and double-click `player.script`.
2. Locate the **`VARIABLE BLOCK`** at the top of the file:

```lua
local GRAVITY       = -1200
local JUMP_IMPULSE  = 550
local FORWARD_SPEED = 300
```

* **Experiment #1 (Moon Physics):** Change `GRAVITY` to `-400`. Press `Cmd+B` / `Ctrl+B` to test. Notice how floaty the jump feels!
* **Experiment #2 (Super Speed):** Change `FORWARD_SPEED` to `600`. Notice how much faster your reaction times need to be!
* **Experiment #3 (High Jump):** Change `JUMP_IMPULSE` to `850` alongside strong gravity (`-2000`) for a fast, snappy jump arc.

---

## Phase 2: Code Modification (Unlocking the Double Jump)

**Goal:** Understand variables and conditional checks (`if` statements) by adding a double-jump feature.

1. In `player.script`, look at line 6 of the **VARIABLE BLOCK**:

```lua
local MAX_JUMPS = 2
```

2. Look down at the **CONTROLLER BLOCK**:

```lua
if action_id == hash("jump") and action.pressed then
    if jumps_remaining > 0 then
        velocity_y = JUMP_IMPULSE
        is_grounded = false
        jumps_remaining = jumps_remaining - 1
    end
end
```

* **Challenge:** Modify `MAX_JUMPS = 3` in the Variable Block. Run the game (`Cmd+B` / `Ctrl+B`). Can you jump three times before hitting the ground?
* **Trace the Logic:** Every jump subtracts `1` from `jumps_remaining`. When `jumps_remaining` reaches `0`, the condition `if jumps_remaining > 0` prevents further jumps until the player lands on a ground block (which resets `jumps_remaining` back to `MAX_JUMPS`).

---

## Phase 3: Visual Level Design & Hazard Placement

**Goal:** Design custom levels using Defold’s visual layout editor.

1. Open `game/main.collection` from the **Project Explorer** pane.
2. **Adding a Spike:**
   * Right-click the `main` node in the Outline window (top right).
   * Select **Add Game Object File** and choose `game/hazard.go`.
   * Use the **Move Tool** (`W`) to place the spike along the runner track.
3. **Adding Platforms:**
   * Right-click `main` $\rightarrow$ **Add Game Object File** $\rightarrow$ Select `game/ground.go`.
   * Position the block higher in the air to build floating platforms.
4. Press **`Cmd + B`** / **`Ctrl + B`** to playtest your custom level!
