
import './App.css'
import { Helmet } from 'react-helmet'
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

      </div>
    </>

  )
}

export default App
