import './login.css'

const Login = () => {
    return (
        <div className='login'>
            <main className="container" data-theme="dark">
                <article className="grid" >
                    <div>
                        <hgroup>
                            <h1>Sign in</h1>
                            <h2>start session for use this application</h2>
                        </hgroup>
                        <form>
                            <input type="text" name="login" placeholder="Login" aria-label="Login" required />
                            <input type="password" name="password" placeholder="Password" aria-label="Password" required />
                            <fieldset>
                                <label htmlFor="remember">
                                    <input type="checkbox" role="switch" id="remember" name="remember" />
                                    Remember me
                                </label>
                            </fieldset>
                            <button type="submit" className="contrast" >Login</button>
                        </form>
                    </div>
                    <div></div>
                </article>
            </main>
        </div>
    )
}
export default Login;