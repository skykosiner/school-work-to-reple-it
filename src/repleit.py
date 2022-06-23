from dataclasses import dataclass
from typing import List

@dataclass
class FileInfo:
    fileName: str
    fileContents: List[str]

@dataclass
class RepleMeDaddy:
    files: List[FileInfo]
