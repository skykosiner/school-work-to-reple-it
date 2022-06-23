from typing import List
from pathlib import Path

def AskPath() -> str:
    path = input("What is the path to the file? ")
    path = "/home/yoni/school-python/" + path

    return path

fileStr: List[str] = []

def GetFilesFromPath(path: str) -> List[str]:
    for posix_path in Path(path).glob("*"):
        if posix_path.is_dir():
            if ".git" in posix_path.as_posix():
                continue

            GetFilesFromPath(posix_path.as_posix())
        else:
            fileStr.append(posix_path.as_posix())

    return fileStr

def GetContentsOfFile(filePath: str) -> List[str]:
    lines: List[str] = []

    with open(filePath, "r") as f:
        for line in f.readlines():
            lines.append(line)

    return lines
