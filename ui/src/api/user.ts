import request from '@/utils/request'

export type UserLoginReq = {
  username: string
  password: string
}

export function userLogin(username: string, password: string) {
  const data: UserLoginReq = { username: username, password: password }
  return request({
    url: '/api/v1/card/login',
    method: 'post',
    data
  })
}
