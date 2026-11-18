--------------------------------------------------------------------------------
-- BLOCK 1: MODULE SETUP & CONFIGURATION (Like Scratch Custom Block Settings)
--------------------------------------------------------------------------------
local state = require "main.state"
local M = {}

-- Constant scrolling speed to match the rhythm runner pace
local SCROLL_SPEED = 400

--------------------------------------------------------------------------------
-- BLOCK 2: SCROLLER UPDATE FUNCTION (Like Scratch "forever" loop motion)
--------------------------------------------------------------------------------
function M.update(dt, delete_threshold)
    -- Do not move obstacles until the player presses jump to start the game
    if not state.game_started then 
        return 
    end

    -- Move the obstacle object to the left over time (dt = delta time / frame speed)
    local pos = go.get_position()
    pos.x = pos.x - (SCROLL_SPEED * dt)
    go.set_position(pos)

    -- Cleanup Check: Delete the obstacle once it scrolls past the left edge of the screen
    if pos.x < delete_threshold then
        go.delete()
    end
end

return M