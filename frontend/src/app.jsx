import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Layout from './components/Layout';
import Home from './pages/Home';
import Programs from './pages/Programs';
import Grants from './pages/Grants';
import ProgramDetail from './pages/ProgramDetail';
import NotFound from './pages/NotFound';

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/programs" element={<Programs />} />
          <Route path="/programs/:id" element={<ProgramDetail />} />
          <Route path="/grants" element={<Grants />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </Layout>
    </Router>
  );
}

export default App;