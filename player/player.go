components {
  id: "script"
  component: "/player/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"idle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/player.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 1.0
    y: 1.0
  }
}
embedded_components {
  id: "label"
  type: "label"
  data: "text: \"Jump to start gameplay!\"\n"
  "font: \"/fonts/game.font\"\n"
  "material: \"/builtins/fonts/font.material\"\n"
  "size {\n"
  "  x: 500.0\n"
  "  y: 60.0\n"
  "}\n"
  "color {\n"
  "  x: 1.0\n"
  "  y: 1.0\n"
  "  z: 1.0\n"
  "  w: 1.0\n"
  "}\n"
  ""
  position {
    x: 0.0
    y: 70.0
    z: 0.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "mask: \"hazard\"\n"
  "embedded_collision_shape {\n"
  "   shapes {\n"
  "     shape_type: TYPE_BOX\n"
  "     position {\n"
  "     }\n"
  "     rotation {\n"
  "     }\n"
  "     index: 0\n"
  "     count: 3\n"
  "   }\n"
  "   data: 30.0\n"
  "   data: 18.0\n"
  "   data: 10.0\n"
  "}\n"
  ""
}