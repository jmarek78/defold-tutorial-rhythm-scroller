components {
  id: "spawner_script"
  component: "/main/spawner.script"
}
embedded_components {
  id: "spike_factory"
  type: "factory"
  data: "prototype: \"/obstacles/spike.go\"\n"
  "load_dynamically: false\n"
  ""
}