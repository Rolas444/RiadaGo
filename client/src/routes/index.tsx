import { createBrowserRouter } from "react-router-dom";
import Home from "../pages/home/home.page";
import Members from "../pages/members/members.page";

const router = createBrowserRouter([
    {
        path: '/',
        element:<Home/>,

    },
    {
        path: '/members',
        element: <Members/>
    }
])

export default router;