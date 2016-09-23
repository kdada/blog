
// 用户详细信息
export class UserDetail {
    Id: number          //用户id
    Email: string       //用户邮箱
    Name: string        //用户昵称
    CreateTime: string  //创建时间
    Status: number      //状态码:1-正常,2-禁止登录
    Reason: string      //处于当前状态的原因
}