import Foundation

class FileLockResolver {
    private var fileLocks: [String: Bool] = [:]
    private let queue = DispatchQueue(label: "FileLockResolver", attributes: .concurrent)

    func lockFile(_ filePath: String) {
        queue.async(flags: .barrier) {
            self.fileLocks[filePath] = true
        }
    }

    func unlockFile(_ filePath: String) {
        queue.async(flags: .barrier) {
            self.fileLocks[filePath] = false
        }
    }

    func isFileLocked(_ filePath: String) -> Bool {
        var isLocked = false
        queue.sync {
            isLocked = self.fileLocks[filePath] ?? false
        }
        return isLocked
    }
}