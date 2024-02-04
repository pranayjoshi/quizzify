import React from "react";
import { Link } from "react-router-dom";
import Navbar from "../Layout/navbar";
// import backgroundImage from '../assets/12.jpg'; // Update with the correct path

const HomePage = () => {
  return (
    // <div className="h-screen" style={{ backgroundImage: `url(${backgroundImage})`, backgroundSize: 'cover' }}>
    <div className="h-screen bg-inherit">
      <Navbar />
      <div className="h-full w-full flex flex-col">
        <div className="grid grid-cols-2 gap-4 h-full items-end justify-items-center px-8 py-4">
          <Link to="/create" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 hover:opacity-100">
            <h2 className="text-2xl font-bold mb-2">Create Quiz</h2>
          </Link>
          <Link to="/join" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 hover:opacity-100">
            <h2 className="text-2xl font-bold mb-2">Join Quiz</h2>
          </Link>
        </div>
        <div className="grid grid-cols-3 gap-4 h-full items-start justify-items-start px-8 py-4">
          <Link to="/quick-quiz" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 hover:opacity-100">
            <h2 className="text-2xl font-bold mb-2">Quick Quiz</h2>
          </Link>
          <Link to="/notes" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 hover:opacity-100">
            <h2 className="text-2xl font-bold mb-2">Notes</h2>
          </Link>
          <Link to="/results" className="flex items-center justify-center p-6 rounded shadow-md text-center h-64 w-full text-white bg-gradient-to-r from-green-400 to-blue-500 opacity-60 hover:opacity-100">
            <h2 className="text-2xl font-bold mb-2">Results</h2>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default HomePage;