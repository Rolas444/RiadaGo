import { createBrowserRouter } from "react-router-dom";
import Home from "../pages/home/home.page";
import Members from "../pages/members/members.page";
import Login from "../pages/login/login.page";


const router = createBrowserRouter([
    {
        path: '/',
        element:<Home/>,

    },
    {
        path: '/members',
        element: <Members/>
    },
    {
        path: '/login',
        element: <Login/>
    }
])

export default router;