import { useCommonStore } from "../../stores/commonStore";
import {useEffect} from "react"
import { NavLink, useNavigate } from "react-router-dom"
import { Helmet } from 'react-helmet'
import './navbar.css'

function Navbar() {

    const navigate=useNavigate();
    const { changeTheme } = useCommonStore();

    const { theme, user } = useCommonStore((state) => ({
        theme: state.theme, user:state.user
    }));

    useEffect(() => {
        if(user.login===false){
            navigate('/login')
        }
    })

    const ChangeTheme=(e: React.FormEvent<HTMLAnchorElement>)=>{
        e.preventDefault();
        changeTheme();

    }
    

    return (
        <div className="container">
            <Helmet htmlAttributes={theme ? { 'data-theme': 'light' } : { 'data-theme': 'dark' }} >
                <title>Riada Go </title>
            </Helmet>
            <nav>
                <ul>
                    <li><strong><NavLink to="/">Riada</NavLink></strong></li>
                </ul>
                <ul>
                    <li><NavLink to="/members" >Members</NavLink></li>
                    <li><NavLink to="/login"><i className="bi bi-person-fill"></i></NavLink></li>
                    
                    <li><NavLink to="/" onClick={ChangeTheme} className='contrast' ><i className={theme ? "bi bi-brightness-high":"bi bi-moon" }></i></NavLink></li>
                </ul>
            </nav>
        </div>
    )
}

export default Navbar;