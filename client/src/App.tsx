import { useState } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom' // Add the Switch import statement
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import HomePage from './pages/HomePage'
import CreatePage from './pages/CreatePage'
import CodeEnterPage from './pages/CodeEnterPage'
import Quiz from './pages/quiz'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
    <div className=" bg-gray-900 p-1 dark:bg-gray-900 ">
      <BrowserRouter>
        <Routes>
          <Route path="/"  element={<HomePage/>}/>
          <Route path="/create" element={<CreatePage/>} />
          <Route path="/join" element={<CodeEnterPage/>} />
          <Route path="/quick-quiz" element={<Quiz/>} />

          {/* <Route path="/about" component={About} /> */}
        </Routes>
      </BrowserRouter>
    </div>
    </>
  )
}

export default App
