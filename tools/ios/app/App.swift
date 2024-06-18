import SwiftUI
import TodoforgeServer

@main
struct Project: App {
    init() {
        print("starting TODO Forge...")
        let path = NSSearchPathForDirectoriesInDomains(.libraryDirectory, .userDomainMask, true)
        let port = TodoforgeServer.CmdLib(path[0])
        print("TODO Forge started on port [\(port)]")
        let url = URL.init(string: "http://localhost:\(port)/")!
        self.cv = ContentView(url: URLRequest(url: url))
    }

    var cv: ContentView

    var body: some Scene {
        WindowGroup {
            cv
        }
    }
}
