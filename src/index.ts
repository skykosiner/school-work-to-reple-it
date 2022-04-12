import askPath from "./askPath";
import getFile from "./getFile";

async function main() {
    let path = await askPath("What is the path to the file? ");
    path = "/home/yoni/school-python/" + path;

    await getFile(path);
};

main();
