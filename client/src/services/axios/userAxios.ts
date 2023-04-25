import { userInterface } from "../../interface"
import axios from "./authApi"


export const loginRequest = async (username: string, password: string, remember:boolean)=>{
    axios.post("/user/login",{
        user_name: username,
        password,
        remember,
    })
}

export const registerRequest = async (user: userInterface.createUser)=>
    axios.post("/user/create",user)

export const profileRequest = async ()=>axios.get("/user/profile")