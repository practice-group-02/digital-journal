import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Home from "./pages/Home.jsx";
import Grants from "./pages/Grants.jsx";
import Programs from "./pages/Programs.jsx";
import NotFound from "./pages/NotFound.jsx";
import Layout from "./components/Layout.jsx"; 

import "./index.css";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      { path: "/", element: <Home /> },
      { path: "/grants", element: <Grants /> },
      { path: "/programs", element: <Programs /> },
      { path: "*", element: <NotFound /> },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
