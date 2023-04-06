import axios from "axios"
import { useCommonStore } from "../../stores/commonStore";

const baseURL = `http://localhost:3003`;

const authApi = axios.create({
    baseURL,
    withCredentials: true,
});

authApi.interceptors.request.use((config)=>{
    const token = useCommonStore.getState().token;
    config.headers = {
        Authorization: "Bearer " + token,
    };
    return config;
});

export default authApi;