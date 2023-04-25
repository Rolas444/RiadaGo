import React, { useEffect, useRef, useState } from 'react';
import { ControlledInput } from '../../components/inputs/Input';
import { loginRequest } from '../../services/axios/userAxios';
import { useCommonStore } from '../../stores/commonStore';
import { validateEmail } from '../../tools/validators/formValidators';
import './login.css'

const Login = () => {
    const setToken = useCommonStore()

    const  userRef = useRef<HTMLInputElement | null>(null);
    const passRef = useRef<HTMLInputElement | null>(null);
    const rmbRef = useRef<HTMLInputElement | null>(null);
    const [vForm, setVform] = useState(false); 
    var [user,setUser]=useState('');
    var [pass,setPass] = useState('');
    

    useEffect(()=>{})

    const changeUser = () => {
        // setUser(userRef.current?.value ===undefined ? '': userRef.current?.value );
        

        setVform(validateFormLogin(user,pass));
        
    }

    const changePass = () => {
        setPass(passRef.current?.value ===undefined ? '': passRef.current?.value);
        setVform(validateFormLogin(user,pass));
    }

    const validateFormLogin=(user : any, pass: any) => {
        
        if(validateEmail(user) && pass.length>0) return true;
        else return false;
    }

    const submitHandler= async (e: React.FormEvent<HTMLFormElement>)=>{
        e.preventDefault();
        let rmb : any = rmbRef.current?.checked
        if (vForm){
            console.log(user+' '+pass+' '+rmb);
            // const resLogin = await loginRequest(user,pass,rmb)
            // console.log(resLogin);
        }else alert ("username or password invalid")

    }
    
    return (
        <div className='login'>
            <main className="container" data-theme="dark">
                <article className="grid" >
                    <div>
                        <hgroup>
                            <h1>Sign in</h1>
                        </hgroup>
                        <form onSubmit={submitHandler}>
                            {/* <input ref={userRef} onInput={changeUser} type="input" name="username" placeholder="username" autoComplete='off' aria-label="username" required /> */}
                            <ControlledInput type='text' placeholder='username' onUpdate={val=>{setUser(val)}} />
                            <input ref={passRef} onInput={changePass} type="password" name="password" placeholder="Password" aria-label="Password" required />
                            <fieldset>
                                <label htmlFor="remember">
                                    <input ref={rmbRef} type="checkbox" role="switch" id="remember" name="remember" />
                                    Remember me
                                </label>
                            </fieldset>
                            <button type="submit" className="contrast" disabled={!vForm} >Login</button>
                        </form>
                    </div>
                    <div></div>
                </article>
            </main>
        </div>
    )
}
export default Login;