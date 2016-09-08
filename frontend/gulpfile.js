//引用
var gulp = require("gulp")
var gutil = require("gulp-util")
var htmlreplace = require("gulp-html-replace")
var concat = require("gulp-concat")
var uglify = require("gulp-uglify")
var md5 = require("gulp-md5-plus")
var cleanCSS = require("gulp-clean-css")
var pack = require("gulp-webpack")
var rename = require("gulp-rename")
var clean = require("gulp-clean")
var webpack = require("webpack")
var Q = require("q")
var path = require("path")
var extractor = require("./extractor")
var ngtemplate = require("./ngtemplate")

//根目录
var root = "./"

//生成目标根目录
var dest = "../web/"

// 需要监视的文件
var files = {
    views: "./views/**/*.html",
    css: "./css/**/*.css",
    module: ["./src/**/*.ts","./tmpls/**/*.html"]
}

//需要移动的文件和目录
var static = ["./files/**", "./fonts/**", "./img/**", "./views/layout.json"]

//生成目标目录
var output = {
    views: dest + "views",
    css: dest + "css",
    js: dest + "js",
    module: dest + "js"
}

//html替换路径
var urlPaths = {
    css: "/css/",
    js: "/js/",
    module: "/js/"
}

//html替换扩展名
var exts = {
    css: ".css",
    js: ".js",
    module: ".js"
}

//html 替换配置
var replaceConfig = {
    // "buildName":"filePath"
}

//编译配置
var config = {
    // "ext":{
    //     "fileName":["filePath"]
    // }
}

//分析html,根据html的定义生成replaceConfig和config
function analyzeHtml() {
    return function () {
        replaceConfig = {}
        config = {}
        return gulp.src(files.views)
            .pipe(extractor(function (buildName, pathArray) {
                var index = buildName.lastIndexOf("-")
                var ext = buildName.substring(index + 1)
                var filename = buildName.substring(0, index) + exts[ext]
                var base = urlPaths[ext]
                var dist = base + filename
                replaceConfig[buildName] = dist
                var obj = config[ext] || {}
                var arr = obj[filename] || []
                for (var i = 0; i < pathArray.length; i++) {
                    //缺少文件去重
                    arr.push(root + pathArray[i])
                }
                obj[filename] = arr
                config[ext] = obj
            }))
    }
}

//根据replaceConfig编译并替换html
function buildHtml() {
    return function () {
        return gulp.src(files.views)
            .pipe(htmlreplace(replaceConfig))
            .pipe(gulp.dest(output.views))
    }
}

////根据config.css合并css,prod决定是否压缩
function buildCSS(prod) {
    return function () {
        var src = config.css
        console.log(config.css)
        if (src) {
            var ds = []
            for (var fileName in src) {
                var d = Q.defer()
                ds.push(d.promise)
                var files = src[fileName]
                var s = gulp.src(files)
                    .pipe(concat(fileName))
                if (prod) {
                    s.pipe(cleanCSS())
                }
                s.pipe(gulp.dest(output.css))
                    .on("end", d.resolve)
            }
            if (ds.length > 0) {
                return Q.all(ds)
            }
        }
    }
}

//根据config.js合并js,prod决定是否压缩
function buildJs(prod) {
    return function () {
        var src = config.js
        console.log(config.js)
        if (src) {
            var ds = []
            for (var fileName in src) {
                var d = Q.defer()
                ds.push(d.promise)
                var files = src[fileName]
                var s = gulp.src(files)
                    .pipe(concat(fileName))
                if (prod) {
                    s.pipe(uglify())
                }
                s.pipe(gulp.dest(output.js))
                    .on("end", d.resolve)
            }
            if (ds.length > 0) {
                return Q.all(ds)
            }
        }
    }
}

//根据config.module编译并合并模块,prod决定是否压缩
function buildModule(prod) {
    return function () {
        var src = config.module
        console.log(config.module)
        if (src) {
            var s = gulp.src(root)
                .pipe(pack({
                    entry: src,
                    output: {
                        filename: "[name]"
                    },
                    resolve: {
                        root: path.resolve("./"),
                        extensions: ["", ".js", ".ts"]
                    },
                    resolveLoader: {
                        alias: {
                            "ngtemplate": path.resolve("./ngtemplate")
                        }
                    },
                    module: {
                        loaders: [
                            {
                                test: /\.ts$/,
                                loaders: ["awesome-typescript-loader", "ngtemplate"]
                            },
                            {
                                test: /\.(html|css)$/,
                                loader: 'raw-loader'
                            }
                        ]
                    },
                    externals: {
                        // "@angular/common": "window.ng.common",
                        // "@angular/compiler": "window.ng.compiler",
                        // "@angular/forms": "window.ng.forms",
                        // "@angular/router": "window.ng.router",
                        // "@angular/platform-browser": "window.ng.platformBrowser",
                        // "@angular/platform-browser-dynamic": "window.ng.platformBrowserDynamic",
                        // "@angular/core": "window.ng.core",
                        // "@angular/http": "window.ng.http",
                        // "rxjs": "window.Rx"
                    },
                    plugins: [
                        new webpack.optimize.CommonsChunkPlugin("vendor_module.js", "vendor_module.js")
                    ]
                }))
            if (prod) {
                s.pipe(uglify())
            }
            return s.pipe(gulp.dest(output.module))
        }
    }
}

//分析html,根据html的定义生成replaceConfig和config
gulp.task("analyze-html", analyzeHtml())

//根据replaceConfig编译并替换html
gulp.task("build-html", ["analyze-html"], buildHtml())

//移动文件
gulp.task("move-files", function () {
    gulp.src(static, {
        base: "./"
    })
        .pipe(gulp.dest(dest))
})

//根据config.css合并css
gulp.task("dev-css", ["analyze-html"], buildCSS(false))

//根据config.js合并js
gulp.task("dev-js", ["analyze-html"], buildJs(false))

//根据config.module编译并合并模块
gulp.task("dev-module", ["analyze-html"], buildModule(false))

//根据config.css合并压缩css
gulp.task("build-css", ["analyze-html"], buildCSS(true))

//根据config.js合并压缩js
gulp.task("build-js", ["analyze-html"], buildJs(true))

//根据config.module编译并合并压缩模块
gulp.task("build-module", ["analyze-html"], buildModule(true))

//资源md5标记
gulp.task("md5", ["build-html", "build-css", "build-js", "build-module", "move-files"], function () {
    return gulp.src([output.css + "/**/*.css", output.js + "/**/*.js"], {
        base: dest
    })
        .pipe(clean({
            force: true
        }))
        .pipe(md5(10, output.views + "/**/*.html"))
        .pipe(gulp.dest(dest))
})

//清理目标文件夹
gulp.task("clean", function () {
    return gulp.src(dest)
        .pipe(clean({
            force: true
        }))
})

//开发环境任务
gulp.task("default", ["build-html", "dev-css", "dev-js", "dev-module", "move-files"])

//生产环境任务
gulp.task("prod", ["md5"])

//监听文件变化
gulp.task("watch", function () {
    console.log(files, static)
    gulp.watch(files.views, ["default"])
    gulp.watch(files.css, ["dev-css"])
    gulp.watch(files.module, ["dev-module"])
    gulp.watch(static, ["move-files"])
})