import { useCommonStore } from "../../stores/commonStore";
import { NavLink } from "react-router-dom"
import { Helmet } from 'react-helmet'
import './navbar.css'

function Navbar() {

    const { changeTheme } = useCommonStore();

    const { theme } = useCommonStore((state) => ({
        theme: state.theme
    }));

    return (
        <div className="container">
            <Helmet htmlAttributes={theme ? { 'data-theme': 'light' } : { 'data-theme': 'dark' }} >
                <title>Riada Go</title>
            </Helmet>
            <nav>
                <ul>
                    <li><strong><NavLink to="/">Riada</NavLink></strong></li>
                </ul>
                <ul>
                    <li><NavLink to="/members" >Members</NavLink></li>
                    <li></li>
                    <li><button className="contrast switcher" onClick={() => changeTheme()} >
                        <i className="bi bi-lightbulb">

                        </i>
                    </button></li>
                </ul>
            </nav>
        </div>
    )
}

export default Navbar;