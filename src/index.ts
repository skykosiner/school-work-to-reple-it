import { exit } from "process";
import askPath from "./askPath";

async function main() {
    let path = await askPath("What is the path to the file? ");

    path = "/home/yoni/school-work/" + path;

    exit(0);
};

main();
