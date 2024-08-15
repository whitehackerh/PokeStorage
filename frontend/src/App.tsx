import React from 'react';
import logo from './logo.svg';
import './App.css';
import {
  BrowserRouter,
  Routes,
  Route,
  Outlet,
  Link,
  useLocation
} from "react-router-dom";
import Header from "./components/modules/Header";
import HomeBar from "./components/modules/HomeBar";
import { createTheme, ThemeProvider } from '@mui/material/styles';


const theme = createTheme({
  typography: {
    fontFamily: 'Century Gothic, sans-serif',
  },
});

const Top = () => {
  return (
    <h1></h1>
  );
};

const Menu = () => {
  let location = useLocation();
  return (
    <>
      <Header />
      <HomeBar />
      {location.pathname === "/" ? <Top /> : <></>}
      <Outlet />
    </>
  );
};

function App() {
  return (
    <ThemeProvider theme={theme}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Menu />}>
          </Route>
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
