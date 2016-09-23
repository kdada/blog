
// ReplaceDate 将标准日期字符串转换为yyyy-MM-dd hh:mm:ss
export function ReplaceDate<T>(array: Array<T>, field: string):Array<T> {
    array.forEach(
        (v:any, i:number, a:Array<T>) => {
            var date = <string>v[field]
            v[field] = date.substring(0, 10) + " " + date.substring(11, 19)
        }
    )
    return array
}