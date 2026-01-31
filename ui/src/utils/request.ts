import axios, { AxiosRequestConfig } from 'axios'
import router from '@/router/index'
import { showError } from './toast'

interface BaseResponse<T> {
  code: number
  data: T
  message: string
}

let baseURL = String(import.meta.env.VITE_BASE_URL)

if (!baseURL) {
  baseURL = ''
}

// 创建axios实例
const axiosApiInstance = axios.create({
  baseURL,
  timeout: 15000
  // withCredentials: true
})

// 请求拦截器
axiosApiInstance.interceptors.request.use(
  (config) => {
    let token = localStorage.getItem('accessToken')
    if (token) {
      if (config.headers) {
        config.headers.Authorization = token
      }
    }
    return config
  },
  (error) => {
    console.log(error)
    return Promise.reject(error)
  }
)

const request = <T>(config: AxiosRequestConfig): Promise<BaseResponse<T>> => {
  return new Promise((resolve, reject) => {
    axiosApiInstance.request<BaseResponse<T>>(config).then(
      (res) => {
        if (res.data !== undefined) {
          resolve(res.data)
        }
      },
      (err) => {

        // 判断请求超时
        if (err.code === 'ECONNABORTED' && err.message.indexOf('timeout') !== -1) {
          reject(err)
          return
        }
        const errorInfo = err.response

        console.log(errorInfo)

        switch (errorInfo.status) {
          case 401:
            router.push({ name: 'user-login' })
            localStorage.removeItem('accessToken')
            break
          case 400:
            showError(errorInfo.data.message)
            break
          default:
            showError(errorInfo.statusText)
            break
        }
      }
    )
  })
}

export default request
