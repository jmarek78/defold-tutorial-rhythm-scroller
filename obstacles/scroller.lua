local state = require "main.state"
local M = {}

local SCROLL_SPEED = 400

function M.update(dt, delete_threshold)
    -- Use the shared state module instead of individual messages
    if not state.game_started then return end

    local pos = go.get_position()
    pos.x = pos.x - (SCROLL_SPEED * dt)
    go.set_position(pos)

    if pos.x < delete_threshold then
        go.delete()
    end
end

return M