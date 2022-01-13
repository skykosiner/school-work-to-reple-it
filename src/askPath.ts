const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
})


export default function askPath(question: string): Promise<string> {
    return new Promise((resolve, reject) => {
        readline.question(question, (answer: string) => {
            resolve(answer)
        });
    });
};
