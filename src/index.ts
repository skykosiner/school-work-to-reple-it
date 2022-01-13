import { exit } from "process";
import askPath from "./askPath";
import getFile from "./getFile";

async function main() {
    let path = await askPath("What is the path to the file? ");
    path = "/home/yoni/school-work/" + path;

    console.log(await getFile(path));

    exit(0);
};

main();
