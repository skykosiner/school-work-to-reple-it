import fs from "fs";

export default function getFile(path: string): Promise<string[]> {
    return new Promise((res, rej) => {
        //@ts-ignore
        fs.readFile(path, (err: Error | undefined, data: Buffer) => {
            if (err) {
                console.log(err);
                rej(err);
            } else {
                const arr = data.toString().split("\n")
                const ret: string[] = [];
                for (const a in arr) {
                    if (arr[a] !== "") ret.push(arr[a]);
                };
                res(ret);
            };
        });
    });
};
