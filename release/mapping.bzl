def gen_mappings(os, arch):
  return {
    "v2ray_study/release/doc": "doc",
    "v2ray_study/release/config": "",
    "v2ray_study/main/" + os + "/" + arch: "",
    "v2ray_ext/tools/control/main/" + os + "/" + arch: "",
  }
