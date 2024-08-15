import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { noTokenRequest } from '../../../http';
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { Link } from "react-router-dom";

const SignIn = () => {
  const navigate = useNavigate();

  const signUpStyle = {
    width: '300px',
    border: '1px solid gray',
    margin: '50px'
  }

  return (
    <>
        <div style={signUpStyle}>
            <label style={{'paddingLeft': '20px'}}>Don't have an account?</label>
            <Link to="/sign-up" style={{'paddingLeft': '20px'}}>Sign Up</Link>
        </div>
    </>
  );
};

export default SignIn;