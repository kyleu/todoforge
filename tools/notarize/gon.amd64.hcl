# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_amd64_v1/todoforge"]
bundle_id = "dev.kyleu.todoforge"

notarize {
  path = "./build/dist/todoforge_0.0.1_darwin_amd64_desktop.dmg"
  bundle_id = "dev.kyleu.todoforge"
}

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/todoforge_0.0.1_darwin_amd64.dmg"
  volume_name = "TODO Forge"
}

zip {
  output_path = "./build/dist/todoforge_0.0.1_darwin_amd64_notarized.zip"
}
