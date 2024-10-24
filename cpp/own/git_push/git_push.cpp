#include <iostream>
#include <cstdlib>
#include <ctime>
#include <string>
#include <unistd.h>

// g++ -o git_push git_push.cpp -static
// sudo mv git_push /usr/bin
// crontab -e
// 5 0,12 * * * /usr/bin/git_push /home/manjaro/zettelkaster

int main(int argc, char* argv[]) {
    if (argc != 2) {
        std::cerr << "Usage: " << argv[0] << " <directory_path>" << std::endl;
        return 1;
    }

    std::string directoryPath = argv[1];

    if (chdir(directoryPath.c_str()) != 0) {
        std::cerr << "Error changing directory to " << directoryPath << std::endl;
        return 1;
    }    

    // Проверка наличия измененных файлов
    std::string gitStatusCmd = "git status --porcelain " + directoryPath;
    std::cout << gitStatusCmd << std::endl;
    if (system(gitStatusCmd.c_str()) != 0) {
        std::cerr << "Error checking git status" << std::endl;
        return 1;
    }

    // Добавление измененных файлов в staging area
    std::string gitAddCmd = "git add " + directoryPath + "/*";
    std::cout << gitAddCmd << std::endl;
    auto gitAddCmdExitStatus = system(gitAddCmd.c_str());
    // if (gitAddCmdExitStatus != 0) {
    //     std::cerr << "Error adding files to staging area " << gitAddCmdExitStatus << std::endl;
    //     return 1;
    // }
    std::cout << "Added files to staging area" << std::endl;

    // Формирование комментария
    std::time_t now = std::time(nullptr);
    char timestamp[20];
    std::strftime(timestamp, sizeof(timestamp), "%Y-%m-%d %H:%M", std::localtime(&now));
    std::string commitComment = std::string(timestamp);

    // Создание коммита
    std::string gitCommitCmd = "git commit -m \"" + commitComment + "\"";
    std::cout << gitCommitCmd << std::endl;
    if (system(gitCommitCmd.c_str()) != 0) {
        std::cerr << "Error committing changes" << std::endl;
        return 1;
    }
    std::cout << "Committed changes with comment: " << commitComment << std::endl;

    std::string gitPullRebaseCmd = "git pull --rebase";
    std::cout << gitPullRebaseCmd << std::endl;
    if (system(gitPullRebaseCmd.c_str()) != 0) {
        std::cerr << "Error pulling changes with rebase" << std::endl;
        return 1;
    }

    // Пуш в удаленный репозиторий
    std::string gitPushCmd = "git push origin main";
    std::cout << gitPushCmd << std::endl;
    if (system(gitPushCmd.c_str()) != 0) {
        std::cerr << "Error pushing changes" << std::endl;
        return 1;
    }
    std::cout << "Pushed changes to remote repository" << std::endl;

    return 0;
}
