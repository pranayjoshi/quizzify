// import React from "react";
import { Link } from "react-router-dom";
import Navbar from "../Layout/navbar";
import { useSelector } from "react-redux";
// import "../App.css";
// import backgroundImage from '../assets/12.jpg'; // Update with the correct path

const HomePage = () => {
//    console.log( useSelector((state: any) => state.token))
  return (
    // <div className="h-screen" style={{ backgroundImage: url(${backgroundImage}), backgroundSize: 'cover' }}>
    <div className="h-screen bg-inherit">
      <Navbar />
      <div className="h-full mx-auto w-8/12 flex flex-col content-center align-center justify-center">
        <h1 className="text-6xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-green-400 to-blue-500 text-center mt-8 mb-4">
          Quizzify: Unleash Your Knowledge!
        </h1>
      <div className="h-full mx-auto w-full flex flex-col content-center align-center justify-center">
        <div className="grid grid-cols-2 gap-8 h-full items-center justify-items-center px-8 py-4">
          <Link to="/create" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 opacity-100 hover:opacity-75">
            <h2 className="text-2xl font-bold mb-2">Create Quiz</h2>
          </Link>
          <Link to="/join" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 opacity-100 hover:opacity-75">
            <h2 className="text-2xl font-bold mb-2">Join Quiz</h2>
          </Link>
        </div>

        <div className="grid grid-cols-3 gap-6 h-full items-start justify-items-start px-8 py-4">
          <Link to="/quick-quiz" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 opacity-100 hover:opacity-75">
            <h2 className="text-2xl font-bold mb-2">Quick Quiz</h2>
          </Link>
          <Link to="/my-quizes" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 opacity-100 hover:opacity-75">
            <h2 className="text-2xl font-bold mb-2">My Quizes</h2>
          </Link>
          <Link to="/results" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 opacity-100 hover:opacity-75">
            <h2 className="text-2xl font-bold mb-2">Results</h2>
          </Link>
        </div>
        </div>

      </div>
    </div>
  );
};

export default HomePage;