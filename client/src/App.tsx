import { useState } from 'react'
import './App.css'
import Navbar from './commons/nav/navbar.component'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="container">
      <Navbar/>
    </div>
  )
}

export default App
