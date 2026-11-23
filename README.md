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
| Set variable `speed` to `5` | `local SCROLL_SPEED = 400` | Stores numbers that control game behavior. |

---

## Phase 1: Experimenting with Game Feel (Physics & Tuning)

**Goal:** Modify numerical variables to see how code parameters instantly change gameplay feel.

1. In the left-side **Project Explorer**, expand the `player` folder and double-click `player.script`.
2. Locate the **`VARIABLE BLOCK`** at the top of the file:

```lua
local GRAVITY       = -2400
local JUMP_FORCE    = 850
local MAX_JUMPS     = 2
```

* **Experiment #1 (Moon Physics):** Change `GRAVITY` to `-1400`. Press `Cmd+B` / `Ctrl+B` to test. Notice how floaty the jump feels!
* **Experiment #2 (Snappy Jump):** Increase `JUMP_FORCE` to `1250` for a higher apex.

---

## Phase 2: Code Tracing and Unlocking the Double Jump

**Goal:** Understand variables and conditional checks (`if` statements) by adding a double-jump feature.

1. In `player/player.script`, look at the variable block:

```lua
local MAX_JUMPS = 2
```

2. **Experiment:** Modify `MAX_JUMPS = 3`. Run the game (`Cmd+B` / `Ctrl+B`). Can you jump three times before hitting the ground?
    * **Trace the Logic:** Every jump subtracts `1` from `jumps_remaining`. When `jumps_remaining` reaches `0`, conditional checks prevent further jumps until the player lands.

---

## Phase 3: Coding a New Feature (Turbo Boost!)

**Goal:** Add a high-speed turbo boost mechanic by editing the shared state file, mapping a new key in the input binding file, updating player input handling, ticking the boost timer once per frame in the game manager, and applying a consistent speed multiplier to scrolling obstacles.

### Step 3.1: Update `main/state.lua` to Track Boost Frames and Speed Multiplier
1. In the left-side **Project Explorer**, expand the `main` folder and double-click `state.lua`.
2. Add `M.boost_frames = 0` and `M.speed_multiplier = 1.0` so the game can track boost duration and keep every scrolling piece at the same speed during each frame:

```lua
local M = {}
M.game_started = false
M.reset_requested = false
M.boost_frames = 0
M.speed_multiplier = 1.0
return M
```

### Step 3.2: Add the `boost` Input Binding
1. In the left-side **Project Explorer**, expand the `input` folder and double-click `game.input_binding`.
2. Add a new key trigger for Input `B` Action `boost`.
3. Verify the code has been added:

```
key_trigger {
  input: KEY_B
  action: "boost"
}
```

### Step 3.3: Handle Input in `player/player.script`
1. Double-click `player/player.script`.
2. Find the `function on_input(self, action_id, action)` function and add the check for `action_id == hash("boost")` to set `state.boost_frames = 15`:

```lua
function on_input(self, action_id, action)
    if action_id == hash("jump") and action.pressed then
      ...
    end

    if action_id == hash("boost") and action.pressed then
        state.boost_frames = 15
    end
end
```

### Step 3.4: Tick `boost_frames` Once Per Frame in `main/main.script`
1. Expand the `main` folder and double-click `main.script`.
2. We will now introduce a **NEW BLOCK** to handle full frame updates.
    * **Why this matters in code:** In Scratch, if you want behavior that doesn't have an exact built-in block, you have to work around the blocks you're given. In text code, you aren't limited to a pre-made set of blocks—you can create new functions and logic whenever your game needs a new capability. In this case, we will create a central frame timer that coordinates every obstacle at once.
    * Add an `update(self, dt)` function so `state.boost_frames` counts down once per frame and sets `state.speed_multiplier` consistently for all obstacles in that frame:

```lua
function update(self, dt)
    if not state.game_started then
        state.speed_multiplier = 1.0
        return
    end

    if state.boost_frames > 0 then
        state.boost_frames = state.boost_frames - 1
        state.speed_multiplier = 2.0
    else
        state.speed_multiplier = 1.0
    end
end
```

### Step 3.5: Apply the Speed Multiplier in `obstacles/scroller.lua`
1. Expand the `obstacles` folder and double-click `scroller.lua`.
2. Modify `M.update` to read `state.speed_multiplier` when moving objects so floors, platforms, and spikes stay aligned:

```lua
function M.update(dt, delete_threshold)
    ...
    pos.x = pos.x - (SCROLL_SPEED * state.speed_multiplier * dt)
    ...
end
```

---

## Phase 4: Messaging & Cross-Object Communication

**Goal:** Understand how separate game objects talk to each other without wires, just like Scratch’s `broadcast [message]` blocks.

1. Open `player/player.script` and find where fall pits delegate to the manager:

```lua
if pos.y <= 0 then
    msg.post("/manager#script", "player_died")
    return
end
```

2. Open `main/main.script` BLOCK 5 to see how the manager listens for that exact message and handles level resetting:

```lua
function on_message(self, message_id, message, sender)
    if message_id == hash("player_died") then
        reset_game_state(self)
    end
end
```

* **The Concept:** Instead of everything living in one massive script, Defold uses **messages** (`msg.post`) to let components coordinate safely. When the player dies, they send a text telegram to the manager, which triggers a clean reset sequence!

---

## Phase 5: Visual Level Design & Hazard Placement

**Goal:** Modify level design via plain text editing.
  * It is the opinion of this tutorial that Defold’s visual layout editor is effective as a viewer but not as an editor for structural layout.

1. Compare `levels/level_1.collection` and `levels/level_2.collection` as text.
    * Defold has an "Open As > Text" option, but it can be deceiving. VS Code or another IDE / text editor is recommended!
2. Challenge: Can you add `level_3` to the game?
    * This WIP Tutorial should provide `level_2` to support this challenge!
```