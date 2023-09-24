# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_all/todoforge"]
bundle_id = "com.kyleu.todoforge"

notarize {
  path = "./build/dist/todoforge_0.0.1_darwin_all_desktop.dmg"
  bundle_id = "com.kyleu.todoforge"
}

apple_id {
  username = "todoforge@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/todoforge_0.0.1_darwin_all.dmg"
  volume_name = "TODO Forge"
}

zip {
  output_path = "./build/dist/todoforge_0.0.1_darwin_all_notarized.zip"
}
