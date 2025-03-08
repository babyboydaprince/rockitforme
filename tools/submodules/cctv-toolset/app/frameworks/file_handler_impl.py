from interfaces.file_handler import FileHandler

class FileHandlerImpl(FileHandler):
    def read_file(self, file_path: str) -> list:
        """Read a file and return its contents as a list of lines."""
        try:
            with open(file_path, "r") as file:
                return file.readlines()
        except FileNotFoundError:
            print(f"Error: File not found: {file_path}")
            return []
        except Exception as e:
            print(f"Error reading file {file_path}: {e}")
            return []

    def write_file(self, file_path: str, content: str):
        """Write content to a file."""
        try:
            with open(file_path, "w") as file:
                file.write(content)
        except Exception as e:
            print(f"Error writing to file {file_path}: {e}")