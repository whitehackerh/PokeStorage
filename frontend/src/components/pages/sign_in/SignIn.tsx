import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { noTokenRequest } from '../../../http';
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { Link } from "react-router-dom";

const SignIn = () => {
  const navigate = useNavigate();
  const [values, setValues] = useState({
    username: '',
    password: ''
  });

  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    const target = e.target;
    const value = target.value;
    const name = target.name;
    setValues({ ...values, [name]: value });
  }

  function signInOperation() {
    noTokenRequest.post('/sign-in', {
      username: values.username,
      password: values.password
    })
    .then((res) => {
      localStorage.setItem('access_token', res.headers.authorization.split(' ')[1]);
      localStorage.setItem('user_id', res.data.data.id);
      navigate('/home');
    })
    .catch((error) => {
      console.log(error);
    })
  }

  const signInForm: React.CSSProperties = {
    margin: '50px',
    display: 'flex',
    flexDirection: 'column' as 'column',
    alignItems: 'center'
  };

  const textFieldContainer: React.CSSProperties = {
    display: 'flex',
    alignItems: 'center',
    gap: '10px'
  };

  const signUpContainer: React.CSSProperties = {
    display: 'flex',
    justifyContent: 'center'
  };
  
  const signUpStyle: React.CSSProperties = {
    width: '300px',
    border: '1px solid gray',
    margin: '50px'
  };

  return (
    <>
      <div style={signInForm}>
        <div style={textFieldContainer}>
          <TextField id="outlined-basic" label="Username" variant="outlined" name="username" value={values.username} onChange={handleChange} />
          <TextField id="outlined-basic" type="password" label="Password" variant='outlined' name="password" value={values.password} onChange={handleChange}/>
          <Button variant="contained" onClick={signInOperation}>
            Sign In
          </Button>
        </div>
      </div>
      <div style={signUpContainer}>
        <div style={signUpStyle}>
          <label style={{'paddingLeft': '20px'}}>Don't have an account?</label>
          <Link to="/sign-up" style={{'paddingLeft': '20px'}}>Sign Up</Link>
        </div>
      </div>
    </>
  );
};

export default SignIn;
