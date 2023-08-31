import request from '@/utils/request'

export type GetSiteInfoResp = {
    site_name:   string;
    nickname: string;
    avatar:   string;
    show_qrcode: boolean;
    hide_link_corner: boolean;
}

export function getSiteInfo() {
  return request<GetSiteInfoResp>({
    url: '/api/v1/card/site',
    method: 'get',
  })
}
