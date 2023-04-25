import { useState } from "react"

export const ControlledInput = ({ type='text', placeholder='placeholder' , onUpdate = function (e: string):void{}})=>{
    const [value, setValue] = useState('');
    const handleChange = (e : React.ChangeEvent<HTMLInputElement>) => {
        setValue(e.target.value);    
        onUpdate(value);
    };
    return (<><input type={type} onChange={handleChange} placeholder={placeholder} autoComplete='off'/></>);

}