import axios from "axios";
import QS from "qs";
import router from "@/router";
import config from "@/request/config";

const http = {}

const Axios = axios.create({
    baseURL: config.baseUrl,
    timeout: config.timeout,
    headers: config.header
});


//响应拦截
Axios.interceptors.request.use(
    async config => {
        //存在加在本地缓存获取token 加在头部
        config.headers.token = sessionStorage.getItem('token')
        return config;
    },
    error => {
        return Promise.error(error)
    }
)

Axios.interceptors.response.use(
    response => {
        if (response.status === 200){
            return Promise.resolve(response)
        } else {
            return Promise.reject(response)
        }
    },
    error => {
        if (error.response.status) {
            switch (error.response.status){
                case 401:
                    alert(error.message)
                    break;
                case 403:
                    sessionStorage.clear();
                    router.push('/home');
                    break;
                case 404:
                    alert(error.message)
                    break;
                default:
            }
            return Promise.reject(error.response);
        }
    }
);

/**
 *  get 方法
 * @param url string
 * @param params
 * @returns {Promise<unknown>}
 */
http.get = function (url, params = {}) {
    return new Promise((resolve, reject) => {
        Axios.get(url, {
            params: params
        }).then(res => {
            resolve(res.data)
        }).catch(err => {
            reject(err.data)
        })
    })
}

/**
 * post 请求
 * @param url
 * @param params
 * @returns {Promise<unknown>}
 */
http.post = function (url, params) {
    return new Promise((resolve , reject) => {
        Axios.post(url, params)
            .then(res => {
                resolve(res.data)
            }).catch(err => {
                reject(err.data)
            });
    });
}

export default http;

