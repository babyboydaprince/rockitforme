#include <stdio.h>
#include <stdlib.h>
#include <string.h>

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
    printf("➡️ Running: %s\n", cmd);
    int ret = system(cmd);
    if (ret != 0) {
        fprintf(stderr, "❌ Command failed: %s\n", cmd);
        exit(1);
    }
}

int main() {
    char distro[64];
    detect_distro(distro, sizeof(distro));

    printf("📦 Detected distribution: %s\n", distro);
    printf("⚡ This script requires sudo privileges.\n");

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

    printf("\n✅ Installation completed successfully!\n");
    printf("--------------------------------------------------\n");
    printf("Now you can run Rock It For Me using:\n");
    printf("   ➜ sudo go run main.go\n");
    printf("\nOr build the executable binary using:\n");
    printf("   ➜ go build rockitforme\n");
    printf("   (on project root directory)\n");
    printf("--------------------------------------------------\n");

    printf("\nHit ENTER to exit...");
    getchar();

    return 0;
}
