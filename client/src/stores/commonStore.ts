import { create } from "zustand";

export interface CommonState {
    theme:boolean;
    user:{
        userId: number;
        username:string;
        rol:string;
        login:boolean;
    }
    changeTheme:()=>void;
}

export const useCommonStore = create<CommonState>((set,get)=>({
    theme: false,
    user: {userId:0,username: "",rol: "",login:true},
    changeTheme: () => set((state)=>({...state,theme:!state.theme}),)

}))


