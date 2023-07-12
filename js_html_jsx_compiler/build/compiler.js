const lg = console.log;
function Compiler() {
    this.file = '<div style="background:salmon">hi</div> ';
    this.Result = function () {
        const tokenizer = new Tokenizer(this.file);
        const tokens = tokenizer.GetTokens();
        lg(tokens);
    };
}
function Tokenizer(file) {
    this.tokens = [];
    this.Content = file;
    this.Index = 0;
    this.Length = this.Content.length;
    this.timeToGetAttribute = false;
    this.timeToGetInnerText = false;
    this.GetTokens = function () {
        return this.tokens;
    };
    while (true) {
        if (this.Index >= this.Length) {
            lg(this.Content);
            return;
        }
        lg(this.Content[this.Index]);
        if (this.timeToGetAttribute) {
            this.Index += 1;
            if (isChar(this.Content[this.Index])) {
                let str = "";
                while (true) {
                    if (this.Content[this.Index] === ">") {
                        break;
                    }
                    str += this.Content[this.Index];
                    this.Index += 1;
                }
                this.timeToGetAttribute = false;
                let t = new Tok("ATTR", str);
                this.tokens = this.tokens.concat(t);
            }
        }
        if (isLeft(this.Content[this.Index])) {
            let t = new Tok("LEFT", "<");
            this.tokens = this.tokens.concat(t);
            this.Index += 1;
            continue;
        }
        if (isRight(this.Content[this.Index])) {
            let t = new Tok("RIGHT", ">");
            this.tokens = this.tokens.concat(t);
            this.Index += 1;
            this.timeToGetInnerText = true;
            continue;
        }
        if (isChar(this.Content[this.Index])) {
            let str = "";
            while (true) {
                if (!isChar(this.Content[this.Index])) {
                    break;
                }
                str += this.Content[this.Index];
                this.Index += 1;
            }
            if (str === "div") {
                let t = new Tok("TAG", str);
                this.tokens = this.tokens.concat(t);
                this.timeToGetAttribute = true;
                lg(this.tokens);
            }
            continue;
        }
        this.Index += 1;
    }
    this.tokAttribute = function () {
        let str = "";
    };
}
function Tok(symbol, value) {
    this.symbol = symbol;
    this.value = value;
}
function isChar(c) {
    let n = c.charCodeAt(0);
    return (n >= 65 && n < 91) || (n >= 97 && n < 123);
}
function isLeft(c) {
    return c === "<";
}
function isRight(c) {
    return c === ">";
}
function Noder(tokens) {
}
export default Compiler;
