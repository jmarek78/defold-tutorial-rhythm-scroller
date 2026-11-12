components {
  id: "script"
  component: "/obstacles/ground.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"idle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/ground.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 10.0
    y: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.2\n"
  "restitution: 0.0\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
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
  "   data: 320.0\n"
  "   data: 32.0\n"
  "   data: 10.0\n"
  "}\n"
  ""
}