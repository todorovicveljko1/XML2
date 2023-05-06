import Axios from "axios"

const axios = Axios.create({
    baseURL: process.env.API_URL ??  "http://localhost:8000/api/v1",
    headers: {
        'X-Requested-With': 'XMLHttpRequest',
        "Content-Type": "application/json",
    }
})


export default axios
