import { useState } from 'react'
import './App.css'
import { Helmet } from 'react-helmet'
import Navbar from './commons/nav/navbar.component'
import { useCommonStore } from './stores/commonStore'

function App() {
  const { theme } = useCommonStore((state)=>({
    theme: state.theme
  }));
  

  return (
    <>
      <Helmet htmlAttributes={theme?{'data-theme':'light'}:{'data-theme':'dark'} } >
        <title>Riada Go</title>
      </Helmet>
      <div className="container">

        <Navbar />
      </div>
    </>

  )
}

export default App
