import request from "@/request/axios"

let home = {
    async rotationList(data) {
        return request.post(
            '/BaseInfo/RotationList',
            data
        );
    }
}

export default home