from src import AskPath, GetContentsOfFile, GetFilesFromPath

path = AskPath()

fileList = GetFilesFromPath(path)

for files in fileList:
    lines = GetContentsOfFile(files)
