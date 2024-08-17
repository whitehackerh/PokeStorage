import { noTokenRequest } from '../../../http';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";

const SignUp = () => {
    const navigate = useNavigate();
    const [values, setValues] = useState({
        username: '',
        password: '',
        email_address: '',
        name: ''
    });

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const target = e.target;
        const value = target.value;
        const name = target.name;
        setValues({...values, [name]: value});
    }

    function registerUser() {    
        noTokenRequest.post('/sign-up', {
            username: values.username,
            password: values.password,
            email_address: values.email_address,
            name: values.name
        })
        .then((res) => {
            localStorage.setItem('token_type', res.headers.authorization.split(' ')[0]);
            localStorage.setItem('access_token', res.headers.authorization.split(' ')[1]);
            localStorage.setItem('user_id', res.data.data.id);
            navigate('/home');
        })
        .catch((error) => {
            console.log(error);
        });
    }

    const signUpForm = {
        'marginTop': '50px',
        'text-align': 'center'
    };

    return (
        <>
            <div style={signUpForm}>
                <h2>Sign Up</h2>
                <TextField id="outlined-basic" label="Username" variant='outlined' name="username" value={values.username} onChange={handleChange}/><br /><br />
                <TextField id="outlined-basic" type="password" label="Password" variant='outlined' name="password" value={values.password} onChange={handleChange}/><br /><br />
                <TextField id="outlined-basic" label="Email" variant='outlined' name="email_address" value={values.email_address} onChange={handleChange}/><br /><br />
                <TextField id="outlined-basic" label="Name" variant='outlined' name="name" value={values.name} onChange={handleChange}/><br /><br />
                <Button variant="contained" style={{ margin: "10px" }} onClick={registerUser}>Register</Button>
            </div>
        </>
    )
}

export default SignUp;

