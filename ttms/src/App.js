import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
//导入页面组件
import Login from "./views/Login/Login";
import Layout from "./views/Layout/Layout";
import Studio from './views/Studio/Studio';
import Seat from './views/Seat/Seat';
import Play from './views/Play/Play';
import Schedule from './views/Schedule/Schedule';
import Order from './views/Order/Order';
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Login />} />
        <Route path='/layout' element={<Layout />}>
          <Route path='/layout/studio' element={<Studio />} />
          <Route path='/layout/seat' element={<Seat />} />
          <Route path='/layout/play' element={<Play />} />
          <Route path='/layout/schedule' element={<Schedule />} />
          <Route path='/layout/order' element={<Order />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
