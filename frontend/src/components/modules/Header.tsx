import { useNavigate } from "react-router-dom";
import { useState } from 'react';
import { withTokenRequest, requestHeaders } from '../../http';
import logo from "../../assets/img/logo.png";
import IconButton from "@mui/material/IconButton";
import Button from "@mui/material/Button";
import MenuItem from "@mui/material/MenuItem";
import Menu from "@mui/material/Menu";
import AccountCircleIcon from '@mui/icons-material/AccountCircle';

const Header = () => {
    let navigate = useNavigate();
    const [anchorEl, setAnchorEl] = useState(null);
    const isMenuOpen = Boolean(anchorEl);
    const handleProfileMenuOpen = (event: any) => {
        setAnchorEl(event.currentTarget);
      }
    const handleMenuClose = () => {
      setAnchorEl(null);
    };

    const renderMenu = () => {
        return (
            <Menu
              anchorEl={anchorEl}
              anchorOrigin={{ vertical: "top", horizontal: "right" }}
              transformOrigin={{ vertical: "top", horizontal: "right" }}
              open={isMenuOpen}
              onClose={handleMenuClose}
            >
              <MenuItem onClick={() => {handleMenuClose(); SignOutEvent();}}>Sign Out</MenuItem>
            </Menu>
          );
    }

    const SignOutEvent = () => {
      requestHeaders.Authorization = `${localStorage.getItem('token_type')} ${localStorage.getItem('access_token')}`;
      withTokenRequest.post('/sign-out', {
      },{
        headers: requestHeaders
      }).then(() => {
        localStorage.removeItem('access_token');
        localStorage.removeItem('user_id');
        navigate('/');
      })
    };

    const headerButtonStyle = {
        cssFloat: 'right',
        'margin-right': '30px',
        padding: '20px 0'
    }
    let AuthMenu = null;
    if (localStorage.getItem('access_token')) {
      AuthMenu = (
        <>
        <IconButton 
        aria-owns={isMenuOpen ? "material-appbar" : undefined}
        aria-haspopup="true"
        onClick={handleProfileMenuOpen}
        style={headerButtonStyle}
        color="inherit">
            <AccountCircleIcon />
        </IconButton>
        </>
      );
    } else {
      AuthMenu = (
      <Button style={headerButtonStyle} color="inherit" onClick={() => navigate('/sign-in')}>
        SIGN IN
      </Button>
      );
    }

    return (
        <>
            <img src={logo} alt="picture" onClick={() => navigate('/')} style={{cursor: 'pointer', height: '60px', 'marginLeft': '50px'}}/>
            {AuthMenu}
            {renderMenu()}
        </>
    )
}

export default Header;