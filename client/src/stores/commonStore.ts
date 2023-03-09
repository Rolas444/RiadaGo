import { create } from "zustand";

export interface CommonState {
    theme:boolean;
    changeTheme:()=>void;
}

export const useCommonStore = create<CommonState>((set,get)=>({
    theme: false,
    changeTheme: () => set((state)=>({...state,theme:!state.theme}),)

}))


