//Options:
//omitExtraWLInCodeBlocks: (boolean) [default false] Omit the trailing newline in a code block. 
//noHeaderId: (boolean) [default false] Disable the automatic generation of header ids. Setting to true overrides prefixHeaderId
//prefixHeaderId: (string/boolean) [default false] Add a prefix to the generated header ids. Passing a string will prefix that string to the header id. Setting to true will add a generic 'section' prefix.
//parseImgDimensions: (boolean) [default false] Enable support for setting image dimensions from within markdown syntax. 
//headerLevelStart: (integer) [default 1] Set the header starting level. 
//simplifiedAutoLink: (boolean) [default false] Turning this on will enable GFM autolink style. 
//literalMidWordUnderscores: (boolean) [default false] Turning this on will stop showdown from interpreting underscores in the middle of words as <em> and <strong> and instead treat them as literal underscores. 
//strikethrough: (boolean) [default false] Enable support for strikethrough syntax.
//tables: (boolean) [default false] Enable support for tables syntax. 
//tablesHeaderId: (boolean) [default false] If enabled adds an id property to table headers tags.
//ghCodeBlocks: (boolean) [default true] Enable support for GFM code block style.
//tasklists:(boolean) [default false] Enable support for GFM takslists.
//smoothLivePreview: (boolean) [default false] Prevents weird effects in live previews due to incomplete input
//smartIndentationFix: (boolean) [default false] Tries to smartly fix indentation problems related to es6 template strings in the midst of indented code.
declare namespace showdown {
    class Converter {
        constructor();
        constructor(options: Object);
        makeHtml(src: string): string;
        setOption(key: string, value: any):void;
        getOption(key: string): any;
    }
    function setOption(key: string, value: any):void;
    function getOption(key: string): any;
    function getDefaultOptions(): Object;
}