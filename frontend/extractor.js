var through = require('through2')
var gutil = require('gulp-util')
var PluginError = gutil.PluginError

// 插件名称
const PLUGIN_NAME = 'gulp-extractor'

// f = function(buildName:string,pathArray:string[])
module.exports = function (f) {
    f = f || function (name, pathArray) {
        gutil.log(PLUGIN_NAME, name, pathArray)
    }
    return through.obj(function (file, enc, cb) {
        if (file.isBuffer()) {
            var buildReg = /<!-- *?build *?: *?(\S*?) *?-->([\s\S]*?)<!-- *?endbuild *?-->/g
            var srcReg = /src *?= *"(.*?)"/g
            var hrefReg = /href *?= *?"(.*?)"/g
            var str = file.contents.toString()
            var result
            while (result = buildReg.exec(str)) {
                var buildName = result[1]
                var subsrc = result[2]
                var pathArray = []
                while (path = hrefReg.exec(subsrc)) {
                    pathArray.push(path[1])
                }
                while (path = srcReg.exec(subsrc)) {
                    pathArray.push(path[1])
                }
                f(buildName, pathArray)
            }
        }
        cb(null, file)
        return
    })
}