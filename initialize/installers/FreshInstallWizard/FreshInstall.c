#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <limits.h>

// Function to check distro from /etc/os-release
void detect_distro(char *distro, size_t size) {
    FILE *fp = fopen("/etc/os-release", "r");
    if (!fp) {
        fprintf(stderr, "❌ Could not open /etc/os-release. Unsupported system.\n");
        exit(1);
    }

    char line[256];
    while (fgets(line, sizeof(line), fp)) {
        if (strncmp(line, "ID=", 3) == 0) {
            strncpy(distro, line + 3, size - 1);
            // remove newline/quotes
            distro[strcspn(distro, "\n")] = 0;
            if (distro[0] == '\"')
                memmove(distro, distro + 1, strlen(distro));
            if (distro[strlen(distro) - 1] == '\"')
                distro[strlen(distro) - 1] = 0;
            fclose(fp);
            return;
        }
    }
    fclose(fp);
    fprintf(stderr, "❌ Could not detect Linux distribution.\n");
    exit(1);
}

// Function to run a system command and check return code
void run_cmd(const char *cmd) {
    printf("\n\n➡️ Running: %s\n", cmd, "\n");
    int ret = system(cmd);
    if (ret != 0) {
        fprintf(stderr, "❌ Command failed: %s\n", cmd);
        exit(1);
    }
}

// Run a command and stream stdout/stderr in real time
void run_cmd_stream(const char *cmd) {
    FILE *fp = popen(cmd, "r");
    if (!fp) {
        fprintf(stderr, "❌ Failed to run command: %s\n", cmd);
        exit(1);
    }

    char buffer[256];
    while (fgets(buffer, sizeof(buffer), fp) != NULL) {
        printf("%s", buffer); // forward output to stdout
    }

    int ret = pclose(fp);
    if (ret != 0) {
        fprintf(stderr, "❌ Build command failed: %s\n", cmd);
        exit(1);
    }
}

int main() {
    char distro[64];
    detect_distro(distro, sizeof(distro));

    run_cmd("clear");
    printf("\n🔄 Rock It For Me! Fresh install has just begun.");
    printf("\n📦 Detected distribution: %s\n", distro);
    printf("⚡ Fresh Install requires sudo privileges.\n");

    if (strcmp(distro, "ubuntu") == 0 || strcmp(distro, "debian") == 0) {
        run_cmd("sudo apt update");
        run_cmd("sudo apt install -y golang");
        run_cmd("sudo apt install -y libpcap-dev");
    }
    else if (strcmp(distro, "fedora") == 0) {
        run_cmd("sudo dnf install -y golang");
        run_cmd("sudo dnf install -y libpcap-devel");
    }
    else if (strcmp(distro, "arch") == 0 || strcmp(distro, "archlinux") == 0) {
        run_cmd("sudo pacman -Sy --noconfirm go");
        run_cmd("sudo pacman -Sy --noconfirm libpcap");
    }
    else {
        fprintf(stderr, "❌ Unsupported distribution: %s\n", distro);
        return 1;
    }

    // Get project root = current working directory
    char cwd[PATH_MAX];
    if (getcwd(cwd, sizeof(cwd)) == NULL) {
        perror("❌ Failed to get current working directory");
        exit(1);
    }
    printf("\n📂 Project root detected: %s\n", cwd);

    // Change directory into ./rockitforme inside project root
    if (chdir(cwd) != 0) {
        perror("❌ Failed to change directory to ./rockitforme");
        printf("The given path: \n", cwd);
        exit(1);
    }

    // Run the Go build command and show output
    printf("➡️ Building Rock it For Me...");
    run_cmd_stream("go build rockitforme");
    // run_cmd("clear");
    printf("\n\n✅ Rock It For Me! has been built successfully!\n");

    // Keep the original next-steps block but commented out
    /*
    printf("--------------------------------------------------\n");
    printf("Now you can run Rock It For Me using:\n");
    printf("   ➜ sudo go run main.go\n");
    printf("\nOr build the executable binary using:\n");
    printf("   ➜ go build rockitforme\n");
    printf("   (on project root directory)\n");
    printf("--------------------------------------------------\n");
    */

    printf("\nHit ENTER to exit...");
    getchar();

    return 0;
}
