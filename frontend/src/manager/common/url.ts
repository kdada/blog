import { URLSearchParams } from '@angular/http';

// SearchParams 将obj对象转换为URLSearchParams
export function SearchParams(obj: any): URLSearchParams {
    var params = new URLSearchParams()
    for (var k in obj) {
        params.append(k,obj[k])
    }
    return params
}