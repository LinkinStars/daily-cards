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
  created_at: string
}

export function updateCard(id : number, content : string, created_at : string) {
  const data: UpdateCardReq = { id: id, content: content, created_at: created_at}
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
  offset: number
}

export function getCard(id : number) {
  const params: GetCardReq = { id: id }
  return request({
    url: '/api/v1/card',
    method: 'get',
    params
  })
}

export function getCardWithOffset(id : number, offset : number) {
  const params: GetCardReq = { id: id, offset: offset }
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

export type GetCardsStatReq = {
  start_time: string;
}

export type CardStat = {
  checked_days: string[];
}

export function getCardsStat(startTime : string) {
  const params: GetCardsStatReq = { start_time : startTime }
  return request({
    url: '/api/v1/cards/stat',
    method: 'get',
    params
  })
}


