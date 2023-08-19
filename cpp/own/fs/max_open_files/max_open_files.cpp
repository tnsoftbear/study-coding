#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <unistd.h>
#include <sys/resource.h>
#include <filesystem>

/**
 * This code will create and open huge amount of files in ./tmp directory: "file_0.txt", "file_1.txt", "file_2.txt" and so on.
 * The count of created files is limited by soft limit.
 * To delete created files call: rm -rf ./tmp
 * Alternatively: find ./tmp -maxdepth 1 -name "*.txt" -exec rm {} \;
 */

int create_tmp_dir() {
    std::string dirName = "tmp";

    if (std::filesystem::exists(dirName) && std::filesystem::is_directory(dirName)) {
        return 0;
    }

    if (std::filesystem::create_directory(dirName)) {
        std::cout << "Directory \"" << dirName << "\" created." << std::endl;
        return 0;
    }

    std::cerr << "Cannot create directory \"" << dirName << "\"." << std::endl;
    return 1;
}

int main() {
    struct rlimit limit;
    if (getrlimit(RLIMIT_NOFILE, &limit) == -1) {
        std::cerr << "Failed to get resource limit" << std::endl;
        return 1;
    }

    int numFilesToOpen = limit.rlim_cur; // Открываем столько файлов, сколько разрешает Soft Limit
    std::cout << "We plan to create " << numFilesToOpen << " files. Are you sure? [y/N]";
    char answer;
    std::cin >> answer;
    if (answer != 'y') {
        std::cout << "Exiting..." << std::endl;
        return 0;
    }

    if (create_tmp_dir() != 0) {
        std::cout << "Exiting..." << std::endl;
        return 1;
    };

    std::vector<std::ofstream> files;
    for (int i = 0; i < numFilesToOpen; ++i) {
        std::string filename = "tmp/file_" + std::to_string(i) + ".txt";
        std::ofstream file(filename);
        if (!file.is_open()) {
            std::cerr << "Could not open file: " << filename << std::endl;
            break;
        }
        files.push_back(std::move(file));
        std::cout << "Created and opened file: " << filename << std::endl;
    }

    return 0;
}