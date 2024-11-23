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
import Home from "./components/pages/home/Home";
import SignIn from "./components/pages/sign_in/SignIn";
import SignUp from "./components/pages/sign_up/SignUp";
import Menu from "./components/pages/menu/Menu";
import BredPokemonList from "./components/pages/bred_pokemon_list/BredPokemonList";
import BredPokemonDetail from "./components/pages/bred_pokemon_detail/BredPokemonDetail";
import RegisterPokemon from "./components/pages/register_pokemon/RegisterPokemon";
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

const Index = () => {
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
          <Route path="/" element={<Index />}>
            <Route path="/home" element={<Home />} />
            <Route path="/sign-in" element={<SignIn />} />
            <Route path="/sign-up" element={<SignUp />} />
            <Route path="/menu" element={<Menu />} />
            <Route path="/bred-pokemon-list" element={<BredPokemonList />} />
            <Route path="/bred-pokemon-detail" element={<BredPokemonDetail />} />
            <Route path="/register-pokemon" element={<RegisterPokemon />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
