import request from '@/utils/request'

export type AddCardReq = {
  content: string
}

export function addCard(content : string) {
  const data: AddCardReq = { content: content}
  return request({
    url: '/api/v1/card',
    method: 'post',
    data
  })
}

export type UpdateCardReq = {
  id: number
  content: string
}

export function updateCard(id : number, content : string) {
  const data: UpdateCardReq = { id: id, content: content}
  return request({
    url: '/api/v1/card',
    method: 'put',
    data
  })
}

export type DeleteCardReq = {
  id: number
}

export function deleteCard(id : number) {
  const data: DeleteCardReq = { id: id }
  return request({
    url: '/api/v1/card',
    method: 'delete',
    data
  })
}

export type GetCardReq = {
  id: number
}

export function getCard(id : number) {
  const params: GetCardReq = { id: id }
  return request({
    url: '/api/v1/card',
    method: 'get',
    params
  })
}

export type GetCardsPageReq = {
  page: number
}

export type GetCardsPageResp = {
    nickname: string;
    avatar:   string;
    cards:    Card[];
    count:    number;
}

export type Card = {
    id:         number;
    created_at: string;
    content:    string;
}

export function getCardsPage(page : number) {
  const params: GetCardsPageReq = { page: page }
  return request({
    url: '/api/v1/cards',
    method: 'get',
    params
  })
}


