from abc import ABC, abstractmethod

class FileHandler(ABC):
    @abstractmethod
    def read_file(self, file_path: str) -> list:
        """Read a file and return its contents as a list of lines."""
        pass

    @abstractmethod
    def write_file(self, file_path: str, content: str):
        """Write content to a file."""
        pass