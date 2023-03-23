import request from "@/request/axios"

let article = {
    async articleList(data) {
        return request.post(
            '/Article/ArticleList',
            data
        );
    },
    async articleDetail(id) {
        return request.post(
            '/Article/ArticleDetail',
            {articleId: id}
        );
    }
}

// 文章列表
// API.articleList = function (data) {
//     return request.post(
//         '/Article/ArticleList',
//             data
//     );
// }

export default article