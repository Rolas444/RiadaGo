import { create } from "zustand";
import { devtools, persist } from "zustand/middleware"
import { userInterface } from "../interface";

export interface CommonState {
    theme: boolean;
    user: {
        userId: number;
        username: string;
        rol: string;
        login: boolean;
    },
    token: string;
    setUser: (nuser: userInterface.userLoged)=>void;
    setToken: (ntoken: string)=>void;
    changeTheme: () => void;
}

export const useCommonStore = create<CommonState>()(
    devtools(
        persist(
            (set) => ({
                theme: false,
                user: { userId: 0, username: "", rol: "", login: false },
                token: "",
                setUser:(nuser)=>set((state)=>({...state,user:nuser})),
                setToken: (ntoken)=>set((state)=>({...state, token: ntoken})),
                changeTheme: () => set((state) => ({ ...state, theme: !state.theme })),
            }),
            { name: 'common-storage', }
        ),
    ))






